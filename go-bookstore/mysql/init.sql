-- Create a new user
CREATE USER 'user'@'%' IDENTIFIED BY 'mypassword';

-- Grant all privileges to the user on the mydatabase database
GRANT ALL PRIVILEGES ON *.* TO 'user'@'%' WITH GRANT OPTION;

-- Create a new database
CREATE DATABASE IF NOT EXISTS bookstore_db;

-- Use the newly created database
USE bookstore_db;