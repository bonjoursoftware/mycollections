# Go Web Application Implementation Plan

## Objective
Create a simple, robust Go web application using Gin and SQLite. The app will support user registration, login, and a personalized homepage. It will have comprehensive unit tests, high coverage, and support cross-compilation for both `amd64` and `arm64` (Raspberry Pi) architectures.

## Architecture & Stack
- **Web Framework:** Gin (`github.com/gin-gonic/gin`)
- **UI:** Server-side rendered HTML using Go's standard `html/template` (no JS frameworks).
- **Database:** SQLite using `modernc.org/sqlite`. We are explicitly choosing this pure-Go SQLite driver to allow trivial cross-compilation for ARM64 without needing CGO cross-compilers.
- **Authentication:** `golang.org/x/crypto/bcrypt` for secure password hashing.
- **Session Management:** `github.com/gin-contrib/sessions` using an encrypted Cookie store.
- **Testing:** `net/http/httptest` for handlers, and an in-memory SQLite database (`file::memory:`) for models.

## Project Structure
```text
/workspace/
├── Makefile             # Build, test, and cross-compilation targets
├── README.md            # Documentation
├── main.go              # Entry point and Gin router setup
├── db/
│   └── db.go            # SQLite connection and initialization
├── models/
│   └── user.go          # User struct, DB queries, password verification
├── handlers/
│   ├── auth.go          # Registration, login, and logout handlers
│   └── home.go          # Homepage handler
├── middleware/
│   └── auth.go          # Route protection middleware
└── templates/
    ├── layout.html      # Base HTML layout
    ├── index.html       # Personalized homepage
    ├── login.html       # Login form
    └── register.html    # Registration form
```

## Implementation Phases

### Phase 1: Core Setup & Makefile
- Initialize Go module (`go mod init app`).
- Update `Makefile` with targets: `run`, `build-amd64`, `build-arm64`, `test`, `test-coverage`, `clean`.
- Add `README.md` with instructions on how to use the Makefile and run the app.

### Phase 2: Database & Models
- Implement `db.go` to initialize a connection to `modernc.org/sqlite` and auto-create the `users` table.
- Implement `models/user.go` to handle `CreateUser` and `AuthenticateUser`.
- Write unit tests for models using an in-memory SQLite database.

### Phase 3: Web Layer & Templates
- Create the HTML templates (`layout.html`, `login.html`, `register.html`, `index.html`) using standard UI components.
- Implement `handlers/auth.go` for `GET/POST /register`, `GET/POST /login`, and `POST /logout`.
- Implement `handlers/home.go` for the `GET /` authenticated homepage.
- Implement session-based authentication middleware in `middleware/auth.go`.

### Phase 4: Main Application & Integration
- Wire up Gin router, session middleware, template rendering, and database initialization in `main.go`.
- Ensure all static assets (if any) or templates are embedded or correctly referenced.

### Phase 5: Comprehensive Testing
- Write `httptest` driven unit tests for all handlers (happy path & error scenarios like duplicate emails, bad passwords, etc.).
- Ensure test coverage target is maximized across handlers, models, and middleware.

## Verification & Migration
- Run `make test-coverage` to ensure tests pass and coverage is high.
- Run `make build-amd64` and `make build-arm64` to verify successful cross-compilation.
- Run `make run` to manually test the web flow.
