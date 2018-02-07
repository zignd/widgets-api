CREATE TABLE public.widgets (
	id serial PRIMARY KEY NOT NULL,
	"name" varchar(128) NOT NULL,
	color_id int4 NOT NULL,
	price numeric NOT NULL,
	melts bool NOT NULL,
	inventory int4 NOT NULL,
	created_at timestamptz NOT NULL,
	FOREIGN KEY (color_id) REFERENCES colors(id)
);