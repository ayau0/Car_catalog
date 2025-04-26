CREATE TABLE brands (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(100) NOT NULL,
                        country VARCHAR(100),
                        description TEXT,
                        logo_url VARCHAR(255),
                        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
