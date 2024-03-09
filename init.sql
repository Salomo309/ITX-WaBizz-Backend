CREATE TABLE IF NOT EXISTS Chatroom (
    chatroom_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_phone VARCHAR(50) NOT NULL,
    customer_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS Users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    google_id VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    picture VARCHAR(1024) NOT NULL,
    admin TINYINT(1) NOT NULL
);

CREATE TABLE IF NOT EXISTS Chat (
    chat_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    chatroom_id INT NOT NULL,
    timendate DATETIME NOT NULL,
    isRead TINYINT(1),
    statusRead ENUM('sent', 'delivered', 'read'),
    content TEXT NOT NULL,
    messageType ENUM('text','photo','video'),
    FOREIGN KEY (chatroom_id) REFERENCES Chatroom(chatroom_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);