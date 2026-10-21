CREATE TABLE links (
    id           INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id      INT NOT NULL REFERENCES users(id),
    original_url TEXT NOT NULL,
    slug         VARCHAR(32) NOT NULL UNIQUE,
    delete_at    TIMESTAMPTZ,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);