CREATE TABLE decks (
    deck_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    name VARCHAR(255) NOT NULL,
    tags VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
