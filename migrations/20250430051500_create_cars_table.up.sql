CREATE TABLE cars (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(100) NOT NULL,
                      brand_id INTEGER NOT NULL REFERENCES brands(id) ON DELETE CASCADE,
                      year INT,
                      price NUMERIC(10, 2),
                      image_url VARCHAR(255),
                      created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                      user_id INTEGER NOT NULL
);
