-- +migrate Up 
-- +migrate StatementBegin

CREATE TABLE books(
    id SERIAL PRIMARY KEY, 
    category_id INTEGER, 
    title VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL,
    image_url VARCHAR(256) NOT NULL,
    release_year INTEGER NOT NULL,
    price VARCHAR(16) NOT NULL,
    total_page INTEGER NOT NULL, 
    thickness VARCHAR(16) NOT NULL, 
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_category
      FOREIGN KEY(category_id) 
	  REFERENCES categories(id)
)

-- +migrate StatementEnd