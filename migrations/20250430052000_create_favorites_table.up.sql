CREATE TABLE IF NOT EXISTS favorites (
                                         id SERIAL PRIMARY KEY,
                                         user_id INTEGER NOT NULL,
                                         car_id INTEGER NOT NULL,
                                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                                         FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);
