-- Active: 1700787304186@@mysql@3306@knownanchor
CREATE TABLE
    decks (
        deck_id INT PRIMARY KEY AUTO_INCREMENT,
        owner_id INT,
        name VARCHAR(255) NOT NULL,
        tags VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (owner_id) REFERENCES users(id)
    );