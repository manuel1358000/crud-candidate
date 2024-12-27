CREATE TABLE candidates (
    id SERIAL PRIMARY KEY,              
    name VARCHAR(255) NOT NULL,         
    email VARCHAR(255) UNIQUE NOT NULL,
    gender VARCHAR(50),
    salary DECIMAL(10, 2)
);