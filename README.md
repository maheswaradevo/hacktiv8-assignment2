# Hacktiv8 Assignment 2 - Order API

## How To Run

### Migrations

- First you need to install the golang-migrate to do database migrations

MacOS

```bash
brew install golang-migrate
```

Windows (use scoop)

```bash
scoop install migrate
```

To run a migrations

```bash
migrate -source file://./db/migrations -database "mysql://root:@tcp(localhost:3306)/order_by" up
```

To rollback a migrations

```bash
migrate -source file://./db/migrations -database "mysql://root:@tcp(localhost:3306)/order_by" down
```

### Run the Code

To run the code, we use Makefile.

```bash
make run
```

### Important

To use .env file, delete the _.example_ in .env file
