# API Endpoints Documentation

## Task Management Endpoints

### Get All Tasks

- **Endpoint:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieve a list of all tasks.
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
    ```json
    [
        {
            "id": 1,
            "title": "Task 1",
            "description": "Description of Task 1",
            "dueDate": "2024-08-01",
            "status": "In Progress"
        },
        
    ]
    ```

### Get Task by ID

- **Endpoint:** `/tasks/:id`
- **Method:** `GET`
- **Description:** Retrieve a task by its ID.
- **Parameters:**
  - `id` (path): The ID of the task to retrieve.
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
    ```json
    {
        "id": 1,
        "title": "Task 1",
        "description": "Description for Task 1",
        "dueDate": "2024-08-01",
        "status": "Pending"
    }
    ```
  - **Status Code:** `404 Not Found`
  - **Body:**
    ```json
    {
        "error": "task not found"
    }
    ```

### Add a New Task

- **Endpoint:** `/tasks`
- **Method:** `POST`
- **Description:** Create a new task.
- **Request Body:**
  ```json
  {
      "title": "New Task",
      "description": "Description for the new task",
      "dueDate": "2024-08-01",
      "status": "Pending"
  }
  ```
- **Response:**
  - **Status Code:** `201 OK`
  - **Body:**
  ```json
  {
    "id": 2,
    "title": "New Task",
    "description": "Description for the new task",
    "dueDate": "2024-08-01",
    "status": "Pending"
  }
  ```
  - **Status Code:** `404 Not Found`
  - **Body:**
  ```json
  {
    "error": "error message"
  }
  ```
### Update Task

- **Endpoint:** `/tasks/:id`
- **Method:** `PATCH`
- **Description:** Update an existing task.
- **Parameters:**
  - `id` (path): The ID of the task to update.
- **Request Body:**
  ```json
  {
    "title": "Updated Task",
    "description": "Updated description",
    "dueDate": "2024-08-02",
    "status": "Completed"
  }
  ```
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
  ```json
  {
    "id": 1,
    "title": "Updated Task",
    "description": "Updated description",
    "dueDate": "2024-08-02",
    "status": "Completed"
  }
  ```
  - **Status Code:** `404 Not Found`
  - **Body:**
  ```json
  {
    "error": "error message"
  }
  ```

### Delete Task

- **Endpoint:** `/tasks/:id`
- **Method:** `DELETE`
- **Description:** Delete a Task
- **Parameters:**
  - `id` (path): The ID of the task to delete.
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```
  - **Status Code:** `404 Not Found`
  - **Body:**
  ```json
  {
    "error": "error message"
  }
  ```




  