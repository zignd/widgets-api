# widgets-api

## Pre-requisites

The application requires a PostgreSQL database, you can set up one easily using a Docker container. The command below, runs a container configured for the application.

    $ docker run --name widgetapi -e POSTGRES_DB=widgetapi -e POSTGRES_USER=widgetapi -e POSTGRES_PASSWORD=ze4v43recg1fg32 -p 5432:5432 -d postgres:10-alpine

We also need to create the database and initialize it with some data and for that we're going to use a tool specialized for migrations. We're going to use the go tools to retrieve and compile the tool and its dependency, therefore, use the following commands.

    $ go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
    $ go build -tags 'postgres' -o $GOPATH/bin/migrate github.com/mattes/migrate/cli
    
And now let's retrieve the code in this repository, it has the migrations we're going to execute against the database.

    $ go get -u github.com/zignd/widgets-api

And then we can execute the migrations.

    $ cd $GOPATH/src/github.com/zignd/widgets-api
    $ migrate -database "postgres://widgetapi:ze4v43recg1fg32@localhost:5432/widgetapi?sslmode=disable" -path ./migrations up

## Endpoints

`POST /auth`

`GET /users`

`GET /users/:id`

`POST /users`

`GET /widgets`

`GET /widgets/:id`

`POST /widgets`

`PUT /widgets/:id`