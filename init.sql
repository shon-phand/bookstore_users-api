CREATE DATABASE postgres;
\c postgres;
CREATE TABLE users (
id SERIAL  PRIMARY KEY NOT NULL UNIQUE,
first_name varchar(50),
last_name varchar(50),
email varchar(50) NOT NULL UNIQUE ,
password varchar(70),
status varchar(50),
date_created varchar(50)
);