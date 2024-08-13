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
 
  ### postman Documentation
  https://documenter.getpostman.com/view/32082781/2sA3s6E9iu
