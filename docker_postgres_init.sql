CREATE TABLE IF NOT EXISTS users(
    id INT GENERATED ALWAYS AS IDENTITY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    firstName VARCHAR(255),
    lastName VARCHAR(255),
    PRIMARY KEY(id),
    UNIQUE(email)
);


INSERT INTO users(email, password, firstName, lastName)
VALUES ('Jade@email.com','abc123','Jade','Smith');

