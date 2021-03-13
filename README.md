# bookstore_users-api

- Language : Golang
- Framework Used - Gin
- Database : postgres
- Authentication : JWT
- Running on port: 8000

#Rest Endpoints:
- GET /ping               : ping the webserver(jwt authentication NOT reuqired)

- POST /login             : login with username and password (jwt authentication NOT reuqired)

- GET /users/:user_id     : fetch user details with id (jwt authentication reuqired)
- POST /users/:user_id    : create a new user (jwt authentication NOT reuqired)
- PUT /users/:user_id     : update user id (jwt authentication reuqired)
- PATCH /users/:user_id   : update specific attribute of user  (jwt authentication reuqired)
- DELETE /users/:user_id  : delete user (jwt authentication reuqired)
- GET /users              : get/serach/retrieve users with status (e.g /users?status=active) (jwt authentication reuqired)


- postgres details
 (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "shon1234"
	dbname   = "postgres"
)

table name users 

CREATE TABLE users (
id SERIAL  PRIMARY KEY NOT NULL UNIQUE,
first_name varchar(50),
last_name varchar(50),
email varchar(50) NOT NULL UNIQUE ,
password varchar(70),
status varchar(50),
date_created varchar(50)
);
