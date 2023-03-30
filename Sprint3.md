# Sprint 3

Sprint 3 video link: [https://drive.google.com/file/d/1j7ntU4_jUTQY_42-pF5TltjOtmekP7Sv/view?usp=sharing](https://drive.google.com/file/d/1j7ntU4_jUTQY_42-pF5TltjOtmekP7Sv/view?usp=sharing)

## Goals
### Frontend:
- reorganize where each page is located
  - take out login and create account from navigation bar after the user is logged in
- create a home page and make one of the buttons in the nav bar link to it
- cypress unit tests
- create the separate student and tutor pages
  
### Backend:
- search function (subjects, ratings, name)
- sort by highest rating, etc.
- authentication feature (tutors can't access student websites and vice versa)
- logout feature
- add name attribute during registration
- unit tests

 
### Overall Goals: 
- Styling frontend and implementing angular material 
- Set up profiles for tutors and students
- Add logout feature
- Create different views for students and teachers
- Begin implementing search function
- registration: add spot for username to put their whole name
- Unit tests for new functionality


## Work Completed
  - Unit tests for both frontend and backend
  - Frontend: Implemented angular material to improve the design of the website and added a logout button
  - Backend: Created functionality for a user logging out and a search function where the user can search for a particular rating, subject, and person.         Unit tests for new functionality were implemented as well. 

## Unit Tests
### Backend
Tests can be found in `backend/user_handlers_test.go`.
- Test for logging out: `func TestLogout(t *testing.T)`
- Test for searching the database: `func TestSearchDatabase(t *testing.T)`


### Frontend
- Test for _ can be found using Cypress in `frontend/src/app/_/_.component.cy.ts`


## Updated Backend Documentation
### New routes added in Sprint 3:

### POST `/api/logout`
Gets the `session` cookie from the request and removes the corresponding session from the backend, logging the user out.

### GET `/api/user`
Gets the `session` cookie from the request and returns the JSON representation of the user. If the user is a student, the JSON response will be in this format:
```json
{
  "id": 1,
  "username": "foo",
  "password": "bar",
  "first_name": "Foo",
  "last_name": "Bar",
  "is_tutor": false,
  "rating": 10.0,
  "email": "foo@bar.com",
  "phone": "000-000-0000",
  "about": "foo bar, foo bar. Foo foo foo, foo bar bar.",
  "grade": 1
}
```
If the user is a tutor, the JSON response will be in this format:

```json
{
  "id": 1,
  "username": "foo",
  "password": "bar",
  "first_name": "Foo",
  "last_name": "Bar",
  "is_tutor": true,
  "rating": 10.0,
  "subjects":[
    {"name": "math"},
    {"name": "reading"}
  ],
  "email": "foo@bar.com",
  "phone": "000-000-0000",
  "about": "foo bar, foo bar. Foo foo foo, foo bar bar."
}
```

### GET `/api/search`
Searches the database for the queried text and returns a JSON list of matching users.

Parameters:
- `q`: Usernames to look for
- `subject`: Subjects to filter the resulting users by

Example:

**GET** `/api/search?q=foo&subject=math` returns a JSON list of all users with usernames like `foo` who teach the subject `math`. The JSON list will be in the same format as above.

### Previous documentation from Sprint 2:
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
