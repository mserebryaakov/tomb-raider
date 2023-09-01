CREATE TABLE namespace (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE application (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    namespace_id INTEGER NOT NULL REFERENCES namespace(id),
    code TEXT NOT NULL,
    data JSONB
);