CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users(
                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      email TEXT NOT NULL,
                      username TEXT NOT NULL,
                      password TEXT NOT NULL,
                      created_at TIMESTAMPTZ DEFAULT NOW(),
                      archived_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_user
    ON users(email)
    WHERE archived_at IS NULL;

CREATE TABLE todo(
                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                     user_id UUID NOT NULL REFERENCES users(id),
                     title TEXT NOT NULL,
                     description TEXT,
                     is_completed BOOLEAN DEFAULT FALSE,
                     expires_at TIMESTAMPTZ,
                     created_at TIMESTAMPTZ DEFAULT NOW(),
                     archived_at TIMESTAMPTZ
);

CREATE TABLE user_session(
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             user_id UUID NOT NULL REFERENCES users(id),
                             expires_at TIMESTAMPTZ NOT NULL,
                             created_at TIMESTAMPTZ DEFAULT NOW(),
                             archived_at TIMESTAMPTZ
);