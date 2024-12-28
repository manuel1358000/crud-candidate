
# CRUD Candidates Project

This project is a Go application that allows you to manage candidates. Below are the steps required to run the project in your local environment.

## Prerequisites

- Have [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed.
- Have [Flyway](https://flywaydb.org/) installed.
- Have [Go](https://golang.org/) installed.

## Steps to Run the Project

### 1. Create the `.env` file

Create a `.env` file in the root directory of the project with the following content:

```
DB_USER=crud-candidates
DB_PASSWORD=crud-candidates
DB_NAME=candidates_db
DB_HOST=localhost
DB_PORT=5432
DB_SSLMODE=disable
SECRET_KEY=secret
```

> **Note:** These variables should match the configurations in the `docker-compose.yaml` file.

### 2. Start the PostgreSQL container

Run the following command to start the container with the PostgreSQL instance:

```bash
docker-compose up -d
```

### 3. Set up environment variables for Flyway

Set the following environment variables on your computer:

```bash
export FLYWAY_URL=jdbc:postgresql://localhost:5432/candidates_db
export FLYWAY_USER=crud-candidates
export FLYWAY_PASSWORD=crud-candidates
export FLYWAY_SCHEMAS=public
export FLYWAY_LOCATIONS=filesystem:./sql
```

### 4. Run migrations with Flyway

With the variables set, run the following command to apply the initial migrations:

```bash
flyway migrate
```

### 5. Run the project

Run the following command to start the project:

```bash
go run main.go
```

### 6. Test the API with Postman

You can use the provided Postman collection to make queries and test the project’s functionalities.

---

Done! You can now use the CRUD Candidates project in your local environment.
