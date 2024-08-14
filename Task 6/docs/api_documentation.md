# API Endpoints Documentation

## Postman Documentation

- https://documenter.getpostman.com/view/31532211/2sA3s4nAx4

## Setup Environmental Variables and Make Connection with Database

- Create .env file
    - add the following environment-specific variables on new lines in the form of NAME=VALUE
          - CONNECTION_STRING = mongodb://0.0.0.0:27017
          - PORT = 8080 //
          - SECRET_KEY = //
          - DB_NAME = //

## Running the App 

- go run .


## Endpoints

### Get All Tasks

- **Endpoint:** `GET /tasks`
- **Description:** Retrieves a list of all tasks.

#### Middleware:
- **JWTValidation:** Validates the JWT token in the request.
- **RoleAuth("ADMIN"):** Authorizes the request for users with the "ADMIN" and "USER" role.

#### Response:
- **200 OK:** Successful response, returns the list of tasks.

---

### Get Task by ID

- **Endpoint:** `GET /tasks/:id`
- **Description:** Retrieves a specific task by its ID.

#### Middleware:
- **JWTValidation:** Validates the JWT token in the request.
- **RoleAuth("ADMIN"):** Authorizes the request for users with the "ADMIN" and "USER" role.

#### Parameters:
- `id` (required): The ID of the task to retrieve.

#### Response:
- **200 OK:** Successful response, returns the requested task.

---

### Create Task

- **Endpoint:** `POST /tasks`
- **Description:** Creates a new task.

#### Middleware:
- **JWTValidation:** Validates the JWT token in the request.
- **RoleAuth("ADMIN"):** Authorizes the request for users with the "ADMIN" role.

#### Request Body:
- The data for the new task.

#### Response:
- **201 Created:** Successful response, returns the created task.

---

### Update Task

- **Endpoint:** `PUT /tasks/:id`
- **Description:** Updates an existing task.

#### Middleware:
- **JWTValidation:** Validates the JWT token in the request.
- **RoleAuth("ADMIN"):** Authorizes the request for users with the "ADMIN" role.

#### Parameters:
- `id` (required): The ID of the task to update.

#### Request Body:
- The updated data for the task.

#### Response:
- **200 OK:** Successful response, returns the updated task.

---

### Update Specific Field

- **Endpoint:** `PATCH /tasks/:id`
- **Description:** Updates a specific field of an existing task.

#### Middleware:
- **JWTValidation:** Validates the JWT token in the request.
- **RoleAuth("ADMIN"):** Authorizes the request for users with the "ADMIN" role.

#### Parameters:
- `id` (required): The ID of the task to update.

#### Request Body:
- The field(s) to update and their new value(s).

#### Response:
- **200 OK:** Successful response, returns the updated task.

---

### Delete Task

- **Endpoint:** `DELETE /tasks/:id`
- **Description:** Deletes an existing task.

#### Middleware:
- **JWTValidation:** Validates the JWT token in the request.
- **RoleAuth("ADMIN"):** Authorizes the request for users with the "ADMIN" role.

#### Parameters:
- `id` (required): The ID of the task to delete.

#### Response:
- **204 No Content:** Successful response, the task has been deleted.

---

### Register User

- **Endpoint:** `POST /user/register`
- **Description:** Registers a new user.

#### Request Body:
- The user registration data.

#### Response:
- **201 Created:** Successful response, returns the registered user.

---

### Login User

- **Endpoint:** `POST /user/login`
- **Description:** Authenticates a user and returns a JWT token.

#### Request Body:
- The user login credentials.

#### Response:
- **200 OK:** Successful response, returns the JWT token.

---

### Get Users

- **Endpoint:** `POST /users`
- **Description:** Retrives all user. Can only be used by ADMINs

#### Request Body:
- 

#### Response:
- **200 OK:** Successful response, returns all users

---

### Promote Users

- **Endpoint:** `POST /user/promote_user/:id`
- **Description:** Promotes a USER to ADMIN. Can only be used by ADMINs

#### Request Body:
- 

#### Response:
- **200 OK:** Successful response, returns the updated user profile
