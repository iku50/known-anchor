CREATE TABLE cards (
    card_id INT PRIMARY KEY AUTO_INCREMENT,
    deck_id INT,
    owner_id INT,
    front VARCHAR(255) NOT NULL,
    back VARCHAR(255) NOT NULL, 
    FOREIGN KEY (deck_id) REFERENCES decks(deck_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
