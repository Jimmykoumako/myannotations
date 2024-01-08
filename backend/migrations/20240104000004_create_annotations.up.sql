CREATE TABLE annotations (id SERIAL PRIMARY KEY, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, deleted_at TIMESTAMP, content TEXT NOT NULL, type VARCHAR(255) NOT NULL, timestamp INTEGER, color VARCHAR(255), text_id INTEGER REFERENCES texts(id));
