# Sprint 4

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




## Unit Tests
### Backend
Tests can be found in `backend/connections_test.go`.
- Test for adding a connection: `func TestAddConnection(t *testing.T)`
- Test for deleting a connection: `func TestDeleteConnection(t *testing.T) `


### Frontend


## Backend Documentation
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
