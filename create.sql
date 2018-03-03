CREATE TABLE purchase (
	id serial PRIMARY KEY,
	price numeric NOT NULL DEFAULT 0,
	name varchar(255) NOT NULL
);