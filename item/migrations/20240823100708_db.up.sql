CREATE TABLE IF NOT EXISTS items(
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    name TEXT,
    type TEXT,
    amount INT
);