# Dcard_2023_backend_intern_homework

Hi, I'm Sun. This is my [Dcard's 2023 backend intern homework](https://drive.google.com/file/d/1XIQnm_p3waggh4PQCKEXaUa030cqtZeq/view).

## Environment

- Go 1.20
- PostgreSQL 15.2

## Installation

1. Set up enviroment configurations

```sh
cp .env.example .env
cp docker-compose.yml.example docker-compose.yml
# edit `.env` and `docker-compose.yml` before continue
```

2. Create and Start containers

```sh
docker-compose up -d
```

## Execution

1. Enter the container

```sh
docker exec -it go sh
```

2. Run the application

```sh
go run .
```
