CREATE TABLE coffe_plus.refresh_tokens (
    token_hash       TEXT        NOT NULL PRIMARY KEY,
    user_id     UUID        NOT NULL REFERENCES coffe_plus.users(id) ON DELETE CASCADE,
    device_name TEXT        NOT NULL,
    ip_address  TEXT,
    expires_at  TIMESTAMPTZ NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (user_id,device_name)
)