CREATE TABLE cards (
    card_id INT PRIMARY KEY AUTO_INCREMENT,
    deck_id INT,
    user_id INT,
    front VARCHAR(255) NOT NULL,
    back VARCHAR(255) NOT NULL,
    status ENUM('new', 'review', 'mastered') DEFAULT 'new',
    last_review TIMESTAMP,
    next_review TIMESTAMP,
    interval INT,
    repetitions INT,
    ease FLOAT,
    FOREIGN KEY (deck_id) REFERENCES decks(deck_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
