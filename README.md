# Dcard_2023_backend_intern_homework

Hi, I'm Sun. This is my [Dcard's 2023 backend intern homework](https://drive.google.com/file/d/1XIQnm_p3waggh4PQCKEXaUa030cqtZeq/view).

## Environment

- Go 1.20
- PostgreSQL 15.2

## Installation

1. Set up enviroment configurations

```sh
cp .env.example .env
cp docker-compose.yml.prod.example docker-compose.yml # There are two example files (development and production for different situation)
# edit `.env` and `docker-compose.yml` before continue
```

2. Create and Start containers

```sh
docker-compose up -d
```

Done! the application is running!

## Testing

1. change `.env` file from:

```env
...
DB_DATABASE=postgres
...
```

to:

```env
...
DB_DATABASE=testing
...
```

2. enter the container `go`, and run test

```sh
go test -p 1 -count 1 ./...
```

## APIs

### POST /head

#### request

```json
{
    "key": "head1"
}
```

#### response

http status: 201

### POST /node

#### request

```json
{
    "key": "node1",
    "value": "content",
    "prev": "head1"
}
```

#### response

http status: 201

### GET /head/:key

Get head by key.

#### response

```json
{
    "key": "head1",
    "next": "node1",
    "created_at": "2023-04-05T14:20:59.206051Z",
    "updated_at": "2023-04-05T14:20:59.206051Z"
}
```

### GET /node/:key

Get node by key.

#### response

```json
{
    "key": "node1",
    "value": "some content",
    "next": "node2",
    "created_at": "2023-04-05T14:01:43.666822Z",
    "updated_at": "2023-04-05T14:01:43.666822Z"
}
```

### DELETE /head/:key

#### response

http status: 204

### DELETE /node/:key

#### response

http status: 204

## Why PostgreSQL?

Why use PostgreSQL for this key-value system?

First. Because in the system, the two neighboring nodes are linked together. So we can benefit from the `cascade` feature in relational database. For example, if there are nodes like:

```txt
nodeA -> nodeB -> nodeC
```

When we delete NodeC, we need to update nodeB after deletion successful. But we have linked two system in our database. So we can do it by simply set `cascade` as `set null on delete`.

Second. Because the structure of the table is fixed. We can use the same schema to finish this system.

Base on the reason above, I choose the relational database. And I want to use PostgreSQL instead of MySQL (or other relational db) at this time because it has more feature than MySQL (although not used this time). So I thought it's a good idea to give it a try on a simple project.

## Some TODOs

There are some more features I plan to do, but there is no time QQ

- [ ] Use migration instead of `.sql`
- [ ] Refactor with DI
- [ ] Add a reverse proxy in front of the application
- [ ] More formal API documents
- [ ] Automation testing by GitHub Actions
- [ ] Package Release by GitHub Actions
- [ ] Use gRPC instead of RESTful API
