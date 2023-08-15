DROP TABLE IF EXISTS users;

CREATE TABLE users (
    email VARCHAR(320) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (email)
);
