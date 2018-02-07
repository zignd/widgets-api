CREATE TABLE public.users (
	id serial PRIMARY KEY,
	"name" varchar(128) NOT NULL,
	avatar varchar(256) NOT NULL,
	username varchar(16) UNIQUE NOT NULL,
	"password" varchar(32) NOT NULL,
	"admin" bool NOT NULL,
	created_at timestamptz NOT NULL
);