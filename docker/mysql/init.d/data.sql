CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL,
    create_at DATETIME NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (name, email, create_at) VALUES ('TOM','xxxx@mail.co.jp', NOW());

