# go-mentorship - Phase 2

This is a self created go mentorship program where I go from beginner to intermediate stage. I follow all the best practices and ask gemini to review my code throughout the program. (Check phase1_README for beginning of the project)

## The 6-Week Roadmap

* **Week 1: Enterprise Architecture & HTTP Testing.** Focus: Migrating from a monolithic `cmd/` to the standard `internal/` layout. Introduction to `net/http/httptest` for testing handlers without starting a server. *Assignment: The Great Migration & Health Check Test.*
* **Week 2: Interfaces & Database Mocking.** Focus: Dependency injection at scale. Creating a `Store` interface so we can test HTTP handlers using a mock in-memory database instead of the real SQLite file. *Assignment: 100% Handler Test Coverage.*
* **Week 3: Containerization (DevOps Foundation).** Focus: Writing `Dockerfile` and `docker-compose.yml` optimized for WSL2 environments. Creating a `Makefile` to standardize build/run/test commands. *Assignment: Dockerized Task Manager.*
* **Week 4: CI/CD & Production Polish.** Focus: GitHub Actions. Automating tests and linting on every push. Upgrading `log` to the structured `log/slog` package. *Assignment: The Deployment Pipeline.*
* **Week 5: Event-Driven Architecture (WebSockets).** Focus: Upgrading HTTP connections. Implementing the `gorilla/websocket` package to allow bidirectional communication. *Assignment: The WebSocket Handshake & Client Hub.*
* **Week 6: Real-Time Systems.** Focus: Concurrency at scale. Broadcasting state changes (Task creations/updates) to all connected clients instantly. *Assignment: Real-Time Task Board.*


## Current Status

* **Current Week:** 2
* **Current Task:** Interfaces & Decoupling - Decouple handlers.go from the concrete implementation of the SQLite database using an interface.
