# Task Manager

This project is a task management system with a Go backend and a Vite-based React frontend.

## Initial Step
Unzip or clone the project
>>> There should be two folder `frontend/` and `backend`

## Backend Setup (Go)

### Prerequisites

- **Go**: Make sure you have Go installed on your system. You can download it from the official site: [Go Downloads](https://go.dev/dl/)


### Steps to Set Up the Backend
Initialize Go modules:
```bash
go mod tidy
```

### To run the backend, you can use the following commands for different tasks:

- To **start the API server**:
    ```bash
    go run ./cmd api
    ```
      
- To **list tasks**:
    ```bash
    go run ./cmd list
    ```
      
- To **add a task**:
    ```bash
    go run ./cmd add <task-name> <task-description>
    ```
      
- To **process tasks**:
    ```bash
    go run ./cmd process
    ```

### For Swagger
- Run the API server
- Then head to <backend-url>/swagger/index.html

## Frontend Setup (Vite React)

### Prerequisites

- **Node.js**: Ensure you have Node.js installed. You can download it from [Node.js Downloads](https://nodejs.org/).
- **npm**: It comes with Node.js, so you should have it installed once Node.js is set up.

### Steps to Set Up the Frontend

1. Install dependencies:
    ```bash
    npm install
    ```

2. Add .env file :
    ```bash 
    VITE_API_BASE_URL= <your-backend-url>
    ```

3. Start the Vite development server:
    ```bash
    npm run dev
    ```

The frontend should now be running at [http://localhost:5173](http://localhost:5173).

---


