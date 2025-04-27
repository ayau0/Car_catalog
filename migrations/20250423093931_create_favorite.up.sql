CREATE TABLE favorites (
                           id SERIAL PRIMARY KEY,
                           user_id INTEGER NOT NULL REFERENCES users(id),
                           car_id INTEGER NOT NULL REFERENCES cars(id),
                           created_at TIMESTAMP DEFAULT current_timestamp
);
