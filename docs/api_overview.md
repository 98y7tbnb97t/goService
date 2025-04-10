# API Documentation Overview

This document provides an overview of the available API endpoints, request/response formats, and guidelines on how to interact with the API.

---

## Endpoints

### Tasks

- **GET /tasks**  
  **Description:** Retrieves a list of tasks.  
  **Response:**  
  - **200 OK**: Returns an array of tasks.
  
- **POST /tasks**  
  **Description:** Creates a new task.  
  **Request Body:**  
  A JSON object following the `Task` schema.  
  **Response:**  
  - **201 Created**: Returns the created task.

- **PUT /tasks/{id}**  
  **Description:** Fully updates an existing task by ID.  
  **Request Body:**  
  A JSON object following the `Task` schema.  
  **Response:**  
  - **200 OK**: Returns the updated task.

- **PATCH /tasks/{id}**  
  **Description:** Partially updates an existing task by ID.  
  **Request Body:**  
  A JSON object with the fields to update.  
  **Response:**  
  - **200 OK**: Returns the updated task.

- **DELETE /tasks/{id}**  
  **Description:** Deletes a task by ID (using a soft delete by updating the deleted timestamp).  
  **Response:**  
  - **200 OK** or **204 No Content**: Confirms successful deletion.

### Users

- **GET /users**  
  **Description:** Retrieves a list of users.  
  **Response:**  
  - **200 OK**: Returns an array of users.

- **POST /users**  
  **Description:** Creates a new user.  
  **Request Body:**  
  A JSON object following the `User` schema.  
  **Response:**  
  - **201 Created**: Returns the created user.

- **PATCH /users/{id}**  
  **Description:** Updates fields of a user by ID.  
  **Request Body:**  
  A JSON object with the fields to update.  
  **Response:**  
  - **200 OK**: Returns the updated user.

- **DELETE /users/{id}**  
  **Description:** Deletes a user by ID.  
  **Response:**  
  - **204 No Content**: Indicates successful deletion.

---

## Schemas

### Task Schema
- **id** (integer): Unique identifier of the task.
- **task** (string): Description of the task.
- **is_done** (boolean): Flag indicating whether the task is completed.
- **user_id** (integer): (Optional) ID of the user associated with the task.
- **created_at** (date-time): Timestamp when the task was created.
- **updated_at** (date-time): Timestamp when the task was last updated.
- **deleted_at** (date-time): Timestamp for soft deletion of the task (if applicable).

### User Schema
- **id** (integer): Unique identifier of the user.
- **email** (string): Email address of the user.
- **password** (string): Password of the user (not exposed in API responses).
- **created_at** (date-time): Timestamp when the user was created.
- **updated_at** (date-time): Timestamp when the user was last updated.
- **deleted_at** (date-time): Timestamp for soft deletion of the user (if applicable).
- **tasks** (array of Task): List of tasks associated with the user.

---

## Additional Information

- **Format:** All requests and responses use JSON.
- **Error Handling:** In case of errors, the API responds with appropriate HTTP status codes and a JSON payload detailing the error.
- **Authentication:** If authentication is needed (not described here), refer to the authentication documentation for token generation and usage.
- **Conventions:** The API follows RESTful conventions with predictable endpoints. Endpoints are case-sensitive.

This documentation provides a high-level view of the endpoints and data schemas. For more details on query parameters, headers, and error codes, further documentation may be added as needed.

---

Happy coding!