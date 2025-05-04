CREATE TABLE saved_cars (
                            id SERIAL PRIMARY KEY,
                            user_id INTEGER NOT NULL,
                            car_id INTEGER NOT NULL REFERENCES cars(id) ON DELETE CASCADE,
                            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
