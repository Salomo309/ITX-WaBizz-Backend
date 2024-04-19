CREATE TABLE IF NOT EXISTS Users (
    email VARCHAR(255) NOT NULL PRIMARY KEY,
    is_active BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS Chatroom (
    chatroom_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_phone VARCHAR(50) NOT NULL,
    customer_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS Chat (
    chat_id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255),
    chatroom_id INT NOT NULL,
    timendate DATETIME NOT NULL,
    isRead ENUM('0','1', 'null'),
    statusRead ENUM('sent', 'delivered', 'read', 'null'),
    content TEXT NOT NULL,
    messageType ENUM('text','photo','video'),
    FOREIGN KEY (chatroom_id) REFERENCES Chatroom(chatroom_id),
    FOREIGN KEY (email) REFERENCES Users(email)
);

INSERT INTO Chatroom VALUES (NULL,'08141341134','Budiono');
INSERT INTO Chatroom VALUES (NULL,'08983517193','Vivi');
INSERT INTO Chatroom VALUES (NULL,'08183920340','Aldo');
INSERT INTO Users VALUES ('asep@gmail.com', True);
INSERT INTO Users VALUES ('bakti@gmail.com', False);
INSERT INTO Chat VALUES (NULL,'asep@gmail.com',1,'2024-03-07 13:00:00', 'null', 'read', 'Sudah diterima', 'text');
INSERT INTO Chat VALUES (NULL,NULL,1,'2024-03-07 13:01:00', '0', 'null', 'Ok', 'text');
INSERT INTO Chat VALUES (NULL,NULL,1,'2024-03-07 13:01:00', '0', 'null', 'Terima kasih', 'text');
INSERT INTO Chat VALUES (NULL,NULL,2,'2024-03-07 14:23:00', '1', 'null', 'Bisa kirimkan resinya?', 'text');
INSERT INTO Chat VALUES (NULL,'bakti@gmail.com',2,'2024-03-07 14:24:00', 'null', 'delivered', '/photo', 'photo');
INSERT INTO Chat VALUES (NULL,NULL,3,'2024-03-01 15:01:00', '1', 'null', 'Halo', 'text');