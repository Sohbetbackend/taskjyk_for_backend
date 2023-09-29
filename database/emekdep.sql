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
    address     VARCHAR(50),
    PRIMARY KEY (`id`)
);
