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

## SQLC Setup

### 1. Install SQLC
```bash
brew install sqlc
```

### 2. Initialize SQLC Configuration
```bash
sqlc init
```
This creates a `sqlc.yaml` configuration file in your project.

### 3. Create Query Directory
```bash
mkdir -p db/query
```

### 4. Create SQL Query Files
Create SQL query files in the `db/query` directory. For example:
```bash
touch db/query/account.sql
```

### 5. Generate Go Code from SQL
```bash
sqlc generate
```
This generates three types of files:
- Models: Go structs that map to database tables
- `db.go`: Database interface
- Query implementations (e.g., `account.sql.go`)

## Go Module Setup

### 1. Initialize Go Module
```bash
go mod init github.com/sajjadmurtaza/simple_bank
```

### 2. Download Dependencies
```bash
go mod tidy
```

TODO
Testing
everyunit test case should start with Test prefix 