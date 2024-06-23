CREATE TABLE IF NOT EXISTS depublic_events (
    id varchar(20) PRIMARY KEY,
    administrators_id VARCHAR(20) NOT NULL ,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    date DATE NOT NULL,
    tickets INT NOT NULL,
    tickets_sold INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_administrators_id FOREIGN KEY (administrators_id) REFERENCES depublic_administrators(id)
);
