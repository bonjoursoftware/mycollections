# mycollections

A simple Go web application using Gin and pure-Go SQLite, demonstrating user registration, login, and a personalized homepage.

## Features
- User Registration and Login
- Secure password hashing using bcrypt
- Session management using Gin sessions (cookie store)
- Database interactions using pure-Go SQLite (`modernc.org/sqlite`), ideal for easy cross-compilation
- Personalized homepage

## Requirements
- Go 1.20+ (or newer)

## Configuration (Environment Variables)
The application strictly requires certain environment variables to start.

- `SESSION_SECRET` **(Required)**: A secure string used to encrypt session cookies. The application will immediately crash if this is not provided.
- `PORT` (Optional): The port to run the web server on (defaults to `8080`).
- `DB_PATH` (Optional): The file path for the SQLite database (defaults to `app.db`).

## Usage

You can use the provided `Makefile` to build, run, and test the application.

### Run Locally
```bash
make run
```
The application will start on `http://localhost:8080`.

### Testing
To run all tests:
```bash
make test
```

To run tests with coverage:
```bash
make test-coverage
```

### Building
To build for your current architecture:
```bash
make build
```

To build for Linux AMD64:
```bash
make build-amd64
```

To build for Linux ARM64 (e.g., Raspberry Pi):
```bash
make build-arm64
```

### Clean
To clean up generated binaries and coverage files:
```bash
make clean
```
