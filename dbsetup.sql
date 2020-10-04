-- 1. Create database
create database golocker;

-- 2. Apply the required schema from https://github.com/corverroos/goku/blob/master/db/schema.sql to the golocker database

-- 3. Create a golocker user
create user 'golocker'@'localhost' identified by '';

-- 4. Apply db privileges for golocker user
grant all privileges on * . * to 'golocker'@'localhost';
