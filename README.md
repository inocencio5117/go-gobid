# Gobid

Gobid is a API for an auction house where users can place bids on items. The API allows users to create accounts, list items for sale, place bids, and view their bidding history.


## Requirements

- Go 1.16 or later
- Air v1.52.0 or later
- Tern v2.3.0 or later
- sqlc v1.28.0 or later
- Docker and Docker Compose (optional)

## Running the Application

To install Gobid, follow these steps:
1. Clone the repository

2. Download the dependencies using Go modules:

```
$ go mod download
```

3. Use Air to run the application:

```
$ air --build.cmd "go build -o ./bin/api ./cmd/api" --build.bin "./bin/api"
```

## Creating and running migrations

Go to "gobid/internal/store/pgstore/migrations"

Then, inside the folder, create a new migration using the following tern command:
```
$ tern create <migration_name>
```

Next you can add the SQL queries in the newly created file.

Finally, you can go back to the project root folder and run the migrations:
```
$ go run ./cmd/terndotenv/
```

## Generating queries with sqlc

To generate queries using sqlc, follow these steps:
1. Navigate to the "gobid/internal/store/pgstore" directory.
2. Add your SQL query in a `.sql` file within this directory.
3. Run the following command to generate the Go code:
```
$ sqlc generate -f ./internal/store/pgstore/sqlc.yaml
```