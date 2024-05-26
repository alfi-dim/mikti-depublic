CREATE TABLE IF NOT EXISTS depublic_transactions (
    id varchar(20) PRIMARY KEY,
    events_id varchar(20) NOT NULL,
    users_id varchar(20) NOT NULL,
    quantity INT NOT NULL,
    date DATE NOT NULL,
    status VARCHAR(255) NOT NULL,
    total_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_events_id FOREIGN KEY (events_id) REFERENCES depublic_events(id),
    CONSTRAINT fk_users_id FOREIGN KEY (users_id) REFERENCES depublic_users(id)
);
