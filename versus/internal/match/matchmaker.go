package match

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/chessclub2205-dev/versus-service/internal/models"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Matchmaker struct {
	rdb *redis.Client
	db  *sqlx.DB
}

func NewMatchmaker(rdb *redis.Client, db *sqlx.DB) *Matchmaker {
	return &Matchmaker{rdb: rdb, db: db}
}

type Request struct {
	RequestID string    `json:"request_id"`
	UserID    uuid.UUID `json:"user_id"`
	Rating    int       `json:"rating"`
	Stake     int       `json:"stake_slices"`
	Timestamp int64     `json:"ts"`
}

// queue key format: versus:stake:<slices>
func (m *Matchmaker) Enqueue(ctx context.Context, req Request) error {
	key := fmt.Sprintf("versus:stake:%d", req.Stake)
	b, _ := json.Marshal(req)
	// use RPUSH so older entries match earlier
	return m.rdb.RPush(ctx, key, string(b)).Err()
}

// TryMatch polls queue and pairs two players if possible. For brevity this is a simple pop pair.
// Production: use sorted sets and rating buckets + locks.
func (m *Matchmaker) TryMatch(ctx context.Context, stake int) (*models.Match, error) {
	key := fmt.Sprintf("versus:stake:%d", stake)
	// block pop two items safely: BLPOP ... then BLPOP again or use transaction
	// simplified: pop left twice (non-blocking)
	res1, err := m.rdb.LPop(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	res2, err := m.rdb.LPop(ctx, key).Result()
	if err != nil {
		// push back res1 if second pop failed
		m.rdb.LPush(ctx, key, res1)
		return nil, err
	}
	var r1, r2 Request
	_ = json.Unmarshal([]byte(res1), &r1)
	_ = json.Unmarshal([]byte(res2), &r2)

	// create DB match record
	matchID := uuid.Must(uuid.NewV4())
	stakeKobo := int64(stake) * 6250 // example: one slice = 6250 kobo? (See doc for conversion)
	q := `INSERT INTO matches (id, stake_slices, stake_kobo, player1, player2, status, created_at)
          VALUES ($1,$2,$3,$4,$5,'matched',NOW())`
	if _, err := m.db.ExecContext(ctx, q, matchID, stake, stakeKobo, r1.UserID, r2.UserID); err != nil {
		// in case of db error push requests back
		m.rdb.LPush(ctx, key, res1)
		m.rdb.LPush(ctx, key, res2)
		return nil, err
	}

	// return lightweight match object
	var mm models.Match
	mm.ID = matchID
	mm.StakeSlices = stake
	mm.StakeKobo = stakeKobo
	mm.Player1 = r1.UserID
	mm.Player2 = &r2.UserID
	mm.Status = models.StatusMatched
	mm.CreatedAt = time.Now().Format(time.RFC3339)
	return &mm, nil
}
