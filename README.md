# Payment Settler 

Better than your average payment split app

## Project structure
- `domain` directory:
    - `api`:
        - `handlers`: API facing tasks and logic
        - `routes`: setup routes (API endpoints) for the API
    - `config`: load config from `.env` file
    - `models`: Business layer models 
    - `services`: core business logic
- `infrastructure` directory:
    - `database`: define database object (postgres/nosql/mysql etc.)
    - `repositories`: abstract out the database logic here

### Service interaction

```
[HTTP Client]
      ↓
[Router] - Maps HTTP routes to handlers
      ↓
[Handlers] - HTTP-specific tasks
      ↓
[Services] - Business logic and data orchestration
      ↓
[Repositories] - Database access
      ↓
[Database]
```

# Running the program

1. Ensure postgres is up and running
2. Populate a `.env` file in the root of the project:

```sh
DB_HOST=<DATABASE URL>
DB_PORT=<DATABASE SERVER PORT>
DB_USER=<USER>
DB_PASSWORD=<PASSWORD>
DB_NAME=<DATABASE_NAME>
```

3. Run `go run src/cmd/main.go`
