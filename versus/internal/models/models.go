package models

import "github.com/gofrs/uuid"

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Rating   int       `db:"rating"`
}

type Wallet struct {
	UserID    uuid.UUID `db:"user_id"`
	Balance   int64     `db:"balance_kobo"`
	Locked    int64     `db:"locked_kobo"`
	UpdatedAt string    `db:"updated_at"`
}

type MatchStatus string

const (
	StatusPending   MatchStatus = "pending"
	StatusMatched   MatchStatus = "matched"
	StatusConfirmed MatchStatus = "confirmed"
	StatusStarted   MatchStatus = "started"
	StatusFinished  MatchStatus = "finished"
	StatusCancelled MatchStatus = "cancelled"
)

type Match struct {
	ID           uuid.UUID   `db:"id"`
	StakeSlices  int         `db:"stake_slices"`
	StakeKobo    int64       `db:"stake_kobo"`
	Player1      uuid.UUID   `db:"player1"`
	Player2      *uuid.UUID  `db:"player2"`
	Player1Ready bool        `db:"player1_ready"`
	Player2Ready bool        `db:"player2_ready"`
	Winner       *uuid.UUID  `db:"winner"`
	Status       MatchStatus `db:"status"`
	CreatedAt    string      `db:"created_at"`
}
