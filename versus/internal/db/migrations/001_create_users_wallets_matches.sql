CREATE TABLE users (
    id uuid PRIMARY KEY,
    username text UNIQUE,
    rating int NOT NULL,
    created_at timestamptz DEFAULT now()
);

CREATE TABLE wallets (
    user_id uuid PRIMARY KEY REFERENCES users(id),
    balance_kobo bigint NOT NULL DEFAULT 0,
    locked_kobo bigint NOT NULL DEFAULT 0,
    updated_at timestamptz DEFAULT now()
);

CREATE TYPE match_status AS ENUM ('pending','matched','confirmed','started','finished','cancelled');

CREATE TABLE matches (
    id uuid PRIMARY KEY,
    stake_slices int NOT NULL,
    stake_kobo bigint NOT NULL,
    player1 uuid NOT NULL REFERENCES users(id),
    player2 uuid REFERENCES users(id),
    player1_ready boolean DEFAULT false,
    player2_ready boolean DEFAULT false,
    winner uuid REFERENCES users(id),
    status match_status DEFAULT 'pending',
    created_at timestamptz DEFAULT now()
);

-- small app_wallet single-row table to collect commission
CREATE TABLE app_wallet (
    id serial PRIMARY KEY,
    balance_kobo bigint NOT NULL DEFAULT 0
);
INSERT INTO app_wallet (balance_kobo) VALUES (0);
