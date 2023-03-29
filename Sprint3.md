# Sprint 3

Sprint 3 video link: 

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
- Test for logging out: `TestLogout(t *testing.T)`
- Test for searching the database: `func TestSearchDatabase(t *testing.T)`


### Frontend
- Test for _ can be found using Cypress in `frontend/src/app/_/_.component.cy.ts`

### POST `/api/logout`




