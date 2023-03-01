# Sprint 2

## Goals
### Frontend:
- Fix button
- Make switch for teacher or student
- Connect with back end
  
### Backend:
- Create two different student and tutor classes
- Create a search function to search for users
- Hash passwords as we currently we just take in passwords as a string
 
### Both:
- Connect backend and frontend

## Work Completed
  - Unit tests for both frontend and backend
  - Integrated frontend and backend together
  - Frontend: Created all the different pages the user will have access to 
  - Backend: Created functionality for hashing a password, created a login route, and added a check to see if the username is unique

## Unit Tests
### Backend
Tests can be found in `backend/user_handlers_test.go`.
- Test for creating a user: `func TestNewUserHandler(t *testing.T)`
- Test for deleting a user: `func TestDeleteUserHandler(t *testing.T)`
- Test for updating a user: `func TestUpdateUserHandler(t *testing.T)`
- Test for getting a user: `func TestGetUserHandler(t *testing.T)`
- Test for getting all users: `func TestGetAllUsersHandler(t *testing.T)`
- Test for logging in: `func TestLogin(t *testing.T)`


### Frontend
- Test for login can be found using Cypress in `frontend/src/app/login/login.component.cy.ts`
- Test for register can be found using Cypress in `frontend/src/app/register/register.component.cy.ts`

## Backend API Documentation
Note: All successful requests to the backend should have a `200 OK` response code.
### GET `/api/users`
Returns a list of all users in JSON. Example response format:
```json
[
  {
    "username": "foo",
    "password": "bar",
    "first_name": "John",
    "last_name": "Doe",
    "is_tutor": false,
    "rating": 7.0,
    "subjects": [
      {"name": "Intro to Software Engineering"},
      {"name": "Programming Language Concepts"}
    ],
    "email": "johndoe@example.com",
    "phone": "123-456-7890",
    "about": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
    "grade": 13
  },
]
```

### POST `/api/users`
Adds a new user to the database. The request body should contain JSON formatted like so: 
```json
{
    "username": "myUsername",
    "password": "myPassword",
    "is_tutor": true
}
```

The response body will contain the JSON representation of the user.

### GET `/api/users/{id}`
Returns data for a user with the corresponding `id`. The response body will contain JSON in this format:
```json
{
  "username": "foo",
  "password": "bar",
  "first_name": "John",
  "last_name": "Doe",
  "is_tutor": false,
  "rating": 7.0,
  "subjects": [
    {"name": "Intro to Software Engineering"},
    {"name": "Programming Language Concepts"}
  ],
  "email": "johndoe@example.com",
  "phone": "123-456-7890",
  "about": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
  "grade": 13 
}
```

### PUT `/api/users/{id}`
Updates user data for the user with the corresponding `id`. The request body should contain have the same JSON format as the above GET request, with each parameter being optional (e.g. you may include `rating` and `email`, and exclude all other properties).

The response body will contain the JSON representation of the user, as in the above GET request.

### DELETE `/api/users/{id}`
Deletes the user with the corresponding `id`. The response body will contain the JSON representation of the user that was deleted, as in the above GET request.

### POST `/api/login`
Sends a login request to the backend. The request body should contain JSON in this format:
```json
{
  "username": "foo",
  "password": "bar"
}
```

If the login is successful, the user will receive a cookie with a `session-name` attribute to identify the session, and the response will contain the following body:
```json
{
  "status": 200,
  "message": "Successfully logged in."
}
```

If the login was unsuccessful, the response will contain the following body with the appropriate error and error message:
```json
{
  "status": 401,
  "message": "Username or password was incorrect."
}
```
