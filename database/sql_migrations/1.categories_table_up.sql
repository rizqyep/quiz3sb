-- +migrate Up 
-- +migrate StatementBegin

CREATE TABLE categories(
    id SERIAL PRIMARY KEY, 
    name VARCHAR(256),
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd