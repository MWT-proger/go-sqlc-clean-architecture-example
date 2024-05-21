# go-sqlc-clean-architecture-example
This is an example of how we can use the SQLC package and still follow the principles of a clean architecture.

## Project Deployment

1. Clone the repository to any suitable directory on your computer.

```bash
git clone https://github.com/MWT-proger/go-sqlc-clean-architecture-example.git
```


2. Copy the template of the environment variables file.

```bash
  cp deployments/.env.example deployments/.env
```

3. Specify the correct environment variables in the newly created file [.env](deployments/.env).

*The following variables are available*
```bash
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=testDB
POSTGRES_PORT=5432
```
4. Run the Postgres database using the following command.

```bash
  docker compose -f deployments/docker-compose.yaml --env-file deployments/.env up -d
```

5. Start the server.
```
go run ./cmd/example -a "localhost:8000" -d "user=postgres password=postgres host=localhost port=5432 dbname=testDB sslmode=disable"
```
