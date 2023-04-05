CREATE DATABASE testing;
GRANT ALL PRIVILEGES ON DATABASE testing TO postgres;

\c testing;

CREATE TABLE IF NOT EXISTS nodes (
   key VARCHAR(255) PRIMARY KEY,
   value TEXT,
   next VARCHAR(255),
   created_at TIMESTAMP NOT NULL DEFAULT Now(),
   updated_at TIMESTAMP NOT NULL DEFAULT Now(),
   
   FOREIGN KEY (next) REFERENCES nodes (key) ON DELETE SET NULL
);

CREATE INDEX node_key ON nodes (key);