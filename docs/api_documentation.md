# Task Management API Documentation

## Endpoints

### GET /tasks
- Retrieves a list of all tasks.
- Response:
  - 200 OK: Returns an array of tasks.

### GET /tasks/:id
- Retrieves the details of a specific task.
- Parameters:
  - id (path): Task ID
- Response:
  - 200 OK: Returns the task details.
  - 404 Not Found: Task not found.

### POST /tasks
- Creates a new task.
- Request body:
  - title (string): Task title
  - description (string): Task description
  - due_date (string): Task due date
  - status (string): Task status
- Response:
  - 201 Created: Returns the created task.
  - 400 Bad Request: Invalid request payload.

### PUT /tasks/:id
- Updates a specific task.
- Parameters:
  - id (path): Task ID
- Request body:
  - title (string): Task title
  - description (string): Task description
  - due_date (string): Task due date
  - status (string): Task status
- Response:
  - 200 OK: Returns the updated task.
  - 400 Bad Request: Invalid request payload.
  - 404 Not Found: Task not found.

### DELETE /tasks/:id
- Deletes a specific task.
- Parameters:
  - id (path): Task ID
- Response:
  - 204 No Content: Task successfully deleted.
  - 404 Not Found: Task not found.
