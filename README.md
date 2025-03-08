# Simple Bank Project Setup

## Database Setup

### 1. Start PostgreSQL in Docker
```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine
```
- `--name postgres`: Names the container "postgres" for easy reference
- `-p 5432:5432`: Maps container's PostgreSQL port to host machine
- `-e POSTGRES_USER=root`: Sets database superuser to "root"
- `-e POSTGRES_PASSWORD=secret`: Sets password for the superuser
- `-d`: Runs container in detached mode (background)
- `postgres:17-alpine`: Uses lightweight Alpine-based PostgreSQL 17 image

### 2. Verify Container is Running
```bash
docker ps | grep postgres
```
This command lists all running containers and filters to show only the PostgreSQL container.

> **Note**: For learning purposes, you can stop and remove the container with:
> ```bash
> docker stop postgres
> docker rm postgres
> ```

### 3. Connect to PostgreSQL
```bash
docker exec -it postgres psql -U root
```
- `docker exec`: Executes a command inside the running container
- `-it`: Provides interactive terminal
- `postgres`: Target container name
- `psql -U root`: Runs PostgreSQL command-line client as root user

## Migration Setup

### 1. Install Migration Tool
```bash
brew install golang-migrate
```

Verify installation:
```bash
migrate --version
```

### 2. Create Migration Directory
```bash
mkdir -p db/migration
```
- `mkdir`: Creates a directory
- `-p`: Creates parent directories if they don't exist (creates `db` if it doesn't exist)

### 3. Create Initial Migration Files
```bash
migrate create -ext sql -dir db/migration -seq init_schema
```
- `-ext sql`: Sets file extension to SQL
- `-dir db/migration`: Specifies directory for migration files
- `-seq`: Uses sequential version numbers
- `init_schema`: Name of the migration

This creates two files:
- `NNNNNN_init_schema.up.sql`: For applying the migration
- `NNNNNN_init_schema.down.sql`: For reverting the migration

### 4. Create Database
```bash
docker exec -it postgres createdb --username=root --owner=root simple_bank
```
Creates a new database named "simple_bank" owned by the "root" user.

### 5. Verify Database Creation
```bash
docker exec -it postgres psql -U root
```

Then in the PostgreSQL shell:
```sql
\list
```

Or connect directly to the database:
```bash
docker exec -it postgres psql -U root simple_bank
```

## Running Migrations

### Apply Database Migrations
```bash
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```
- `-path`: Specifies the directory containing migration files
- `-database`: Connection string for the database
- `-verbose`: Shows detailed output
- `up`: Applies all pending migrations

### Rollback Migrations (if needed)
```bash
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
```

TODO
- install sqlc brew install sqlc
- sqlc init sqlc.yaml is added
- sqlc generate
package db
error parsing queries: path error: stat /Users/murtaza.sajjad/golang_projects/simple_bank/db/query: no such file or directory

so create db/query folder
and create account account.sql file

now run 	sqlc generate
it will create 3 files
models
db.go
account sql go

- then run go mod init github.com/sajjadmurtaza/simple_bank
- go mod tidy