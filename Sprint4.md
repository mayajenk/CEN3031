# Sprint 4

**Video link: [https://drive.google.com/file/d/1wQHMYxLe8jfF7rbHSZe4wJlsDjFqqkvE/view?usp=sharing](https://drive.google.com/file/d/1wQHMYxLe8jfF7rbHSZe4wJlsDjFqqkvE/view?usp=sharing)**

## Goals
- Finish the project
- Implement search -- display profiles when searching for tutors
- **Frontend**:
  - Add editing of contact info/subjects
  - Make search display profiles
    - Make new component for search profile
    - Add an add to favorites button
  - Add stuff to the home page
  - Display message if incorrect login/username already taken
- **Backend**:
  - Add a way to connect students and tutors
  - Add prices/rates to tutors
  - Update search to sort by ratings/prices
  - Make mock database
- Both:
  -  Add review system
  -  If time allows, add profile pictures
  -  Fix bugs

## Work Completed
- Unit tests were completed for both frontend and backend
- **Frontend:** 
  - Updated the home page 
  - Refomatted both profile views for student and tutor
  - Added buttons to edit contact and personal information 
  - Added upload profile picture button 
  - Created a "My Subjects" section in the tutor profile to add/save subjects
  - Created a "Students" section in the tutor profile
  - Created a "Reviews" section for both tutor and students in the profile page
  - Student can add a tutor after searching which creates a connection 
  - Student and tutor can give each other reviews after connection is created
  - Created the search page that displays tutor cards
  - Added functionality to display tutors and redirect to tutors on student profile page 
  
- **Backend:**
  - Added functionality for profile pictures 
  - Calculate ratings based on reviews 
  - Added prices and rates attribute to the tutors
  - Added functionality to create a connection between a student and tutor 

## Unit Tests
### Backend
Tests can be found in `backend/connections_test.go`.
- Test for adding a connection: `func TestAddConnection(t *testing.T)`
- Test for deleting a connection: `func TestDeleteConnection(t *testing.T) `

Tests can be found in `backend/reviews_test.go`.
- Test for adding a review: `func TestAddReview(t *testing.T)`
- Test for deleting a review: `func TestDeleteReview(t *testing.T)`

Tests can be found in `backend/search_test.go`.
- Test for searching the database: `func TestSearchDatabase(t *testing.T)`

Tests can be found in `backend/subjects_test.go`.
- Test for adding a subject: `func TestAddSubjectHandler(t *testing.T)`
- Test for updating a subject: `func TestUpdateSubjectsHandler(t *testing.T)`

Tests can be found in `backend/users_test.go`.
- Test for getting a profile picture: `func TestGetProfilePicture(t *testing.T)`
- Test for updating a profile picture: `func TestUploadProfilePicture(t *testing.T)`

### Frontend
- Cypress tests
  - All tests can be found in `frontend/cypress/e2e/spec.cy.ts`
  - Check Links: 
    - Tests visiting website and clicking login and register. 
  - Login Form: 
    - Tests that students and tutors can log in with valid credentials. 
  - Register Form:
    - Tests that students and tutors can register unless the username is not unique. 
   - Home Page:
     - Tests that the home page is fully working. 
   - Search Page:
     - Tests that the search page is fully working including all functionality. 
   - Search Profile Page:
     - Tests that the search profile pages is fully working. 
    - Tutor Profile Page:
      - Tests that the tutor profile pages is fully working including all functionality. 
   - Student Profile Page:
     - Tests that the search profile pages is fully working including all functionality. 
  - Logout:
    - Tests that the logout component mounts properly.

- Unit tests
  - A unit test for each component and service can be found in its corresponding `.spec` file. Run `ng test` in the `frontend` folder to see the results. 

## Updated Backend Documentation
### POST `/api/connection`
Adds a connection between the specified users. The request body should should contain JSON formatted like so:
```json
{
  "user_1": 1,
  "user_2": 2,
}
```
where `user_1` and `user_2` are the IDs of the corresponding users.

### DELETE `/api/connection`
Removes the connection between the specified users. The request body should contain JSON formatted in the same way as above.
The response body will contain the users between whom the connection was deleted.

### POST `/api/review`
Adds a review to the specified user. The request body should contain JSON formatted like so:
```json
{
  "reviewer_id": 1,
  "reviewee_id": 2,
  "review_text": "The review text goes here",
  "rating": 0.0
}
```
where the reviewer is the user that made the review, and the reviewee is the user being reviewed.

### DELETE `/api/review/{id}`
Deletes the review with the specified `id`. 

### POST `/api/users/{id}/subjects`
Adds a subject to the user with the corresponding `id`. The JSON in the request body should be formatted like so:
```json
{
  "name": "Math"
}
```

### PUT `/api/users/{id}/subjects`
Replaces all of the subjects from the user with the matching `id` with those in the request body. The JSON in the request body should be a list of subjects formatted like so:
```json
[
  {"name": "Math"},
  {"name": "Reading"}
]
```

### DELETE `/api/users/{id}/subjects`
Deletes all of the subjects from the user with the matching `id`.

### GET `/api/search?subject={subject_name}`
Returns a list of tutors who teach `subject_name`.

## Previous Backend Documentation

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
