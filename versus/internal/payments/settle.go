package payments

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/chessclub2205-dev/versus-service/internal/models"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type Manager struct {
	db *sqlx.DB
}

func NewManager(db *sqlx.DB) *Manager {
	return &Manager{db: db}
}

// SettleMatch is idempotent: it returns nil if already settled.
func (m *Manager) SettleMatch(ctx context.Context, matchID uuid.UUID, winnerID uuid.UUID) error {
	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var mm models.Match
	// lock the row
	q := `SELECT * FROM matches WHERE id = $1 FOR UPDATE`
	if err := tx.GetContext(ctx, &mm, q, matchID); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("match not found")
		}
		return err
	}
	if mm.Status == models.StatusFinished {
		// already settled
		return nil
	}
	if mm.Status != models.StatusStarted && mm.Status != models.StatusMatched && mm.Status != models.StatusConfirmed {
		return fmt.Errorf("invalid match status for settlement: %s", mm.Status)
	}
	if mm.Player2 == nil {
		return errors.New("cannot settle match without player2")
	}

	stake := mm.StakeKobo
	wAmt, lAmt, commission, _, err := Payouts(stake)
	if err != nil {
		return err
	}

	var loser uuid.UUID
	if mm.Player1 == winnerID {
		loser = *mm.Player2
	} else if mm.Player2 != nil && *mm.Player2 == winnerID {
		loser = mm.Player1
	} else {
		return errors.New("winner not part of match")
	}

	// update winner wallet
	if _, err := tx.ExecContext(ctx,
		`UPDATE wallets SET locked_kobo = locked_kobo - $1, balance_kobo = balance_kobo + $2 WHERE user_id = $3`,
		stake, wAmt, winnerID); err != nil {
		return err
	}
	// update loser wallet
	if _, err := tx.ExecContext(ctx,
		`UPDATE wallets SET locked_kobo = locked_kobo - $1, balance_kobo = balance_kobo + $2 WHERE user_id = $3`,
		stake, lAmt, loser); err != nil {
		return err
	}

	// credit app commission to app_wallet (single row table)
	if _, err := tx.ExecContext(ctx,
		`UPDATE app_wallet SET balance_kobo = balance_kobo + $1`, commission); err != nil {
		return err
	}

	// record transactions
	// (simplified: insert two transactions + commission row)
	_, err = tx.ExecContext(ctx, `INSERT INTO transactions (id, match_id, user_id, amount_kobo, type) VALUES ($1,$2,$3,$4,$5)`,
		uuid.Must(uuid.NewV4()), matchID, winnerID, wAmt, "payout")
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO transactions (id, match_id, user_id, amount_kobo, type) VALUES ($1,$2,$3,$4,$5)`,
		uuid.Must(uuid.NewV4()), matchID, loser, lAmt, "payout")
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO transactions (id, match_id, user_id, amount_kobo, type) VALUES ($1,$2,$3,$4,$5)`,
		uuid.Must(uuid.NewV4()), matchID, nil, commission, "commission")
	if err != nil {
		return err
	}

	// update match status
	if _, err := tx.ExecContext(ctx,
		`UPDATE matches SET status = $1, winner = $2 WHERE id = $3`, models.StatusFinished, winnerID, matchID); err != nil {
		return err
	}
	return tx.Commit()
}
