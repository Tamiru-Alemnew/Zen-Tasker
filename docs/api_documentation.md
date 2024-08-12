# Task Management API Documentation

## Endpoints

### GET /tasks
  - **Status:** `200 OK`
  - **Body:**

    ```json
    [
      {
        "id": "uuid",
        "title": "string",
        "description": "string",
        "dueDate": "string ",
        "status": "string",
      }
    ]
    ```

### GET /tasks/:id
- Retrieves the details of a specific task.

  **URL:** `/tasks/{id}`

  **Method:** `GET`

  **Path Parameters:**

  - `id` : The ID of the task to retrieve.

  **Response:**

  - **Status:** `200 OK`
  - **Body:**

      ```json
      {
        "id": "uuid",
        "title": "string",
        "description": "string",
        "dueDate": "string ",
        "status": "string",
      }
      ```
  **Errors:**

  - `400 Bad Request`: Invalid task ID.
  - `404 Not Found`: Task not found.

### POST /tasks
- Creates a new task.

  **URL:** `/tasks`

  **Method:** `POST`

  **Request Body:**

    ```json
    {
      "title": "string",
      "description": "string",
      "dueDate": "string ",
      "status": "string"
    }
    ```

  **Response:**

  - **Status:** `201 Created` : Task successfully created.

  **Errors:**

  - `400 Bad Request`: Invalid request payload.

### PUT /tasks/:id
- Updates a specific task.
  **URL:** `/tasks/{id}`

  **Method:** `PUT`

  **Parameter:**

  - `id`: The ID of the task to update.

  **Request Body:**

    ```json
    {
      "title": "string",
      "description": "string",
      "dueDate": "string ",
      "status": "string"
    }
    ```

  **Response:**

  - **Status:** `200 OK`

  **Errors:**

  - `400 Bad Request`: Invalid task ID.
  - `404 Not Found`: Task not found.


### DELETE /tasks/:id
- Deletes a specific task.
  **URL:** `/tasks/{id}`

  **Method:** `DELETE`

  **Path Parameters:**

  - `id`: The ID of the task to delete.

  **Response:**

  - **Status:** `200 OK`

  **Errors:**

  - `400 Bad Request`: Invalid task ID.
  - `404 Not Found`: Task not found.



#### **Authentication**
- **Sign up**: `POST /register`

  - **Request Body**:
    ```json
    {
      "username": "example",
      "password": "************"
    }
    ```
  - **Response**: `201 Created`

    ```json
    {
      "id": "00000000",
      "username": "example",
      "role": "admin"
    }
    ```

- **Sign in**: `POST /login`

  - **Request Body**:
    ```json
    {
      "username": "example",
      "password": "************"
    }
    ```
  - **Response**: `200 OK`

    ```json
    {
      "id": "00000000",
      "username": "example",
      "role": "admin"
    }
    ```
- **Promote User**: `PATCH /promote/:id`

  - **Path Parameters**:
    - `id`: The ID of the user to promote.

  - **Response**: `200 OK`

    ```json
      {
          "message": "User promoted to admin successfully"
      }
    ```