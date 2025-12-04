CREATE TABLE transactions (
    id uuid PRIMARY KEY,
    match_id uuid REFERENCES matches(id),
    user_id uuid REFERENCES users(id),
    amount_kobo bigint NOT NULL,
    type text NOT NULL,
    created_at timestamptz DEFAULT now(),
    meta jsonb DEFAULT '{}'
);
CREATE INDEX idx_transactions_user ON transactions(user_id);
