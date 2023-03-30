# Sprint 3

Sprint 3 video link: [https://drive.google.com/file/d/1j7ntU4_jUTQY_42-pF5TltjOtmekP7Sv/view?usp=sharing](https://drive.google.com/file/d/1j7ntU4_jUTQY_42-pF5TltjOtmekP7Sv/view?usp=sharing)

## Goals
### Frontend:
- Reorganize where each page is located
  - Take out login and create account from navigation bar after the user is logged in
- Create a home page and make one of the buttons in the nav bar link to it
- Add Cypress and unit tests
- Create the separate student and tutor pages
  
### Backend:
- Create a search function (subjects, ratings, name)
  - Sort by highest rating, etc.
- Add authentication feature (tutors can't access student websites and vice versa)
- Add logout feature
- Create unit tests for each new function

 
### Overall Goals: 
- Styling frontend and implementing Angular Material 
- Set up profiles for tutors and students
- Add logout feature
- Create different views for students and teachers
- Begin implementing search function
- Registration: Add spot for username to put their whole name
- Unit tests for new functionality


## Work Completed
- Unit tests were completed for both frontend and backend
- **Frontend:** 
  - Implemented Angular Material to improve the design of the website.
  - Added a logout button.
  - Added authentication, preventing logged out users from accessing the `/profile` and `/search` routes.
  - Implemented different Tutor and Student views of the website.
  - Changed profile view so that it displayed the corresponding user's information when logged in.
- **Backend:**
  - Created functionality for a user logging out
  - Created a search function where the user can search for a particular subject and user.
  - Added a function that gets the user based on the cookies in the request.
  - Revised the API response message so that the browser could properly render it.

## Unit Tests
### Backend
Tests can be found in `backend/user_handlers_test.go`.
- Test for logging out: `func TestLogout(t *testing.T)`
- Test for searching the database: `func TestSearchDatabase(t *testing.T)`


### Frontend
- Cypress tests
  - Login: `frontend/src/app/login/login.component.cy.ts`
    - Tests that you can pass in input to username/password.  
  - Logout: `frontend/src/app/logout/logout.component.cy.ts`
    - Tests that the logout component mounts properly.
  - `frontend/src/app/register/register.component.cy.ts`
    - Tests that you can pass in input to username/password/is_tutor. 
  - `frontend/src/app/search/search.component.cy.ts`
    - Tests that the search component renders properly.
- Unit tests
  - A unit test for each component and service can be found in its corresponding `.spec` file. Run `ng test` in the `frontend` folder to see the results. 


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
