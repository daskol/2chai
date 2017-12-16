--  schema.sql

CREATE TABLE IF NOT EXISTS boards (
    board_id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    abbr VARCHAR(16) NOT NULL UNIQUE,
    name VARCHAR(64) NOT NULL,
    description VARCHAR(256) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CLOCK_TIMESTAMP(),
    updated_at TIMESTAMPTZ DEFAULT CLOCK_TIMESTAMP()
);

CREATE TABLE IF NOT EXISTS threads (
    thread_id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    board_id INTEGER,
    subject VARCHAR(256),
    created_at TIMESTAMPTZ DEFAULT CLOCK_TIMESTAMP(),
    updated_at TIMESTAMPTZ DEFAULT CLOCK_TIMESTAMP(),
    FOREIGN KEY (board_id) REFERENCES boards (board_id)
);

CREATE TABLE IF NOT EXISTS posts (
    post_id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    thread_id INTEGER,
    subject VARCHAR(256),
    comment VARCHAR(16000),
    created_at TIMESTAMPTZ DEFAULT CLOCK_TIMESTAMP(),
    updated_at TIMESTAMPTZ DEFAULT CLOCK_TIMESTAMP(),
    FOREIGN KEY (thread_id) REFERENCES threads (thread_id)
);
