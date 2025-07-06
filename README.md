# Todo List Application

This project implements a simple Todo list application using a microservices architecture. The backend is built with Go, leveraging gRPC for inter-service communication and SQLite for data persistence. The frontend is a React application that consumes the backend's REST API.

## Makefile Commands

Update the `Makefile` provides convenient commands to manage backend services.
 
**Navigate to the `todo-service` directory for backend services:**

### Running the Backend Services

1.  **Generate Protobuf Code:** This command compiles `proto/task_service.proto` into `proto/task_service.pb.go` and `proto/task_service_grpc.pb.go`. It's crucial to run this whenever you modify the `.proto` definition.
    ```bash
    make proto
    ```
2.  **Start Storage Service:** Starts the Storage microservice Service
    ```bash
    make run-storage
    ```
3.  **Start API Gateway:**
    ```bash
    make run-api
    ```
4.  **Test API Gateway:** Executes Go tests for the `api-gateway/` module. It includes race detection (`-race`) and generates a code coverage profile (`-coverprofile`) in `reports/api-gateway/cov.out`
    ```bash
    make test-api
    ```
5.  **Test Storage:** Executes Go tests for the `storage-service/` module. It also includes race detection (`-race`) and generates a code coverage profile in `reports/storage-service/cov.out`.
    ```bash
    make test-storage
    ```
6.  **Build:** This command compiles both the `api-gateway` and `storage-service` applications. The compiled binaries will be placed in the `./build/` directory (e.g., `./build/api-gateway`, `./build/storage-service`).
    ```bash
    make build
    ```
7.  **Linting:** Performs static code analysis
    ```bash
    make lint
    ```
8.  **Formating:** Formats Go source code
    ```bash
    make fmt
    ```

### Frontend Setup (React App)

1.  **Navigate to the `web-app` directory:**
    ```bash
    cd web-app
    ```
2.  **Install Dependencies:**
    ```bash
    npm install # or yarn install
    ```
3.  **Start the React Development Server:**
    ```bash
    npm start # or yarn start
    ```
4.  **Unit test:**
    ```bash
    npm run test
    ```
4.  **Linting:** 
    ```bash
    npm run lint
    ```
