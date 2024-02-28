CREATE TABLE IF NOT EXISTS Chatroom (
    chatroom_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_phone VARCHAR(50) NOT NULL,
    customer_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS Chat (
    chat_id INT AUTO_INCREMENT PRIMARY KEY,
    chatroom_id INT NOT NULL,
    timendate DATETIME NOT NULL,
    isRead BOOLEAN NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY (chatroom_id) REFERENCES Chatroom(chatroom_id)
);

CREATE TABLE IF NOT EXISTS User (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    admin_name VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS Reply (
    reply_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    chatroom_id INT NOT NULL,
    timendate DATETIME NOT NULL,
    content TEXT NOT NULL,
    statusRead ENUM('sent', 'delivered', 'read'),
    FOREIGN KEY (chatroom_id) REFERENCES Chatroom(chatroom_id),
    FOREIGN KEY (user_id) REFERENCES User(user_id)
);