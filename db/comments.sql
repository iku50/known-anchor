CREATE TABLE comments (
    comment_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    post_id INT,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (post_id) REFERENCES posts(post_id)
);
