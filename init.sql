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

INSERT INTO Chatroom VALUES (NULL,'08141341134','Budiono');
INSERT INTO Chatroom VALUES (NULL,'08983517193','Vivi');
INSERT INTO Users VALUES (NULL, '1', 'asep@gmail.com', 'Asep', '/asep', 0);
INSERT INTO Users VALUES (NULL, '2', 'bakti@gmail.com', 'Bakti', '/bakti', 0);
INSERT INTO Chat VALUES (NULL,1,1,'2024-03-07 13:00:00', NULL, 'read', 'Sudah diterima', 'text');
INSERT INTO Chat VALUES (NULL,NULL,1,'2024-03-07 13:01:00', 0, NULL, 'Ok', 'text');
INSERT INTO Chat VALUES (NULL,NULL,2,'2024-03-07 14:23:00', 1, NULL, 'Bisa kirimkan resinya?', 'text');
INSERT INTO Chat VALUES (NULL,2,2,'2024-03-07 14:24:00', NULL, 'Delivered', '/photo', 'photo');
