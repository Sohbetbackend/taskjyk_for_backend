DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id          INT AUTO_INCREMENT NOT NULL,
    first_name  VARCHAR(50) NOT NULL,
    last_name   VARCHAR(50) NOT NULL,
    middle_name VARCHAR(50) NOT NULL,
    user_name   VARCHAR(50) NOT NULL,
    phone       INT,
    email       VARCHAR(60),
    birthday    VARCHAR(50),
    gender      INT NOT NULL,
    address     VARCHAR(50),
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS classroom;
CREATE TABLE classroom (
    id          INT AUTO_INCREMENT NOT NULL,
    name        VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    PRIMARY KEY (`id`)
);
