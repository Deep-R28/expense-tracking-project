-- User Database

CREATE DATABASE user_db;

USE user_db;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fullname VARCHAR(100),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100),
    mobile VARCHAR(20),
    password VARCHAR(255) NOT NULL,
    gender ENUM('male', 'female', 'other'),
    dob DATE,
    currency VARCHAR(10),
    location VARCHAR(100),
    terms BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
