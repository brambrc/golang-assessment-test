CREATE TABLE records (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         marks INTEGER[] NOT NULL,
                         createdAt TIMESTAMPTZ NOT NULL DEFAULT NOW()
);