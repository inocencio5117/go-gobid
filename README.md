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