## ✅ Prioritized & Completed

### Backend:

* ✅ Current implementation has an issue with CORS. I have implemented the CORS middleware to fix the issue.

* ✅ Split Monolith into Microservices: Created `api-gateway` (handles HTTP REST endpoints) and `storage-service` (manages data persistence).

* ✅ Implemented REST API endpoints in API Gateway with proper request handling, routing, and response formatting for frontend consumption.

* ✅ Established gRPC communication between gateway and storage service.

* ✅ Defined and generated `.proto` contracts for gRPC.

* ✅ Organized microservice structure and common `proto/` directory.

* ✅ Comprehensive error handling with gRPC status codes and mapped HTTP status codes in API Gateway for consistent client experience.

* ✅ Migrated data persistence from in-memory to SQLite database in the storage service for lightweight, transactional, and persistent storage.

* ✅ Added input validation on API gateway (e.g., for task creation).

* ✅ Added middleware for CORS and Logger 

* ✅ Established basic configuration management via environment variables for flexible deployments.

* ✅ Implemented graceful shutdowns handling SIGINT/SIGTERM signals to ensure clean termination and completion of in-flight requests.

* ✅ Integrated structured logging `(log/slog)` across both services for consistent, key-value based log entries, aiding debugging and monitoring.

* ✅ Developed backend unit tests with mocked gRPC clients and database layers to verify business logic correctness and isolate external dependencies.

* ✅ Updated `Makefile` to automate build, run, and protobuf code generation tasks for both services, including linting and formatting.


### Frontend:

* ✅ Upgrade Vite and React.js to latest stable versions for security and performance.

* ✅ Implement modular, reusable React components with clear separation.

* ✅ Implement tailwindcss library and improve UI/UX including form validation for better user experience.

* ✅ Create API client library and React hooks for consistent backend communication.

* ✅ Implement logic for marking tasks complete/incomplete with real-time UI sync.

* ✅ Configure and added unit tests for React components and API hooks.

---

## 🛠️ Deferred Production Considerations

### ⚙️ Observability & Monitoring

* [ ] Integrate **Prometheus metrics** and **Grafana dashboards**.
* [ ] Configure centralized logging (e.g., **Grafana Loki**, **ELK Stack**).

### 🔐 Security

* [ ] Harden **input validation and sanitization** across services.

### 💃️ Data Management

* [ ] Add **database migration** using `golang-migrate`.
* [ ] Tune **connection pooling** and **timeouts** for DB access.

### 🚀 DevOps & Scalability

* [ ] **Dockerize** both services and deploy on **Kubernetes**.
* [ ] Set up **CI/CD pipelines** (e.g., GitHub Actions).
* [ ] Configure **Kubernetes CPU/memory resource limits**.

### 📈 Performance

* [ ] Benchmark core functions with `go test -bench`.
* [ ] Profile services using `pprof` for hot path optimization.
* [ ] Perform **load testing** with tools like k6 or Artillery.

### 🧑‍💻 Maintainability

* [ ] Enforce **code review practices** and merge guidelines.
* [ ] Introduce **API versioning** (e.g., `/v1/tasks`).
* [ ] Write full **API and architecture documentation**.

 ### ⚠️ Unit / Integration Testing 
 * [ ] Add integration and load tests for backend services (critical for CI/CD).
 * [ ] Add integration and end-to-end tests for frontend.
