# CEN3031

## Project Name
Find A Tutor

## Project Description
  
Our project will essentially create a website for a tutoring service. It will be a system to help tutors and students connect. The website will be accessible by both the tutor and student, which means once the user logs in they will see different screens. The screen shown depends on whether student or tutor was chosen when the account was created. The student will be able to search for subjects and find tutors who they can contact. The tutor will be able to add their contact information and subjects they are comfortable with. There will also be a section where students can rate tutors they have interacted with. Overall, our project "Find A Tutor" will help students and tutors connect and communicate easily on this new platform. 

## Members
**Front-end**: Selina Lin & Maya Jenkins

**Back-end**: Manasi Patel & Mai Tran

# Installation
First, make sure that you have [Node.js](https://nodejs.org/en/), the [Angular CLI tool](https://angular.io/cli), and [Go](https://go.dev/) installed.

Then, clone the repository with this command in the command line:
```
git clone https://github.com/mayajenk/CEN3031.git
```
Navigate to the folder containing the project.
```
cd CEN3031
```
## Installing dependencies
Navigate to the `frontend` folder and run `npm install`:
```
cd frontend
npm install
```

# Usage
## Running only the frontend
If you just want to run the frontend, navigate to the frontend folder and run `ng serve`.
```
cd frontend
ng serve
```
The app will be live at `http://localhost:4200`

## Unit Testing for the frontend
Install cyress, you can use the following command in the terminal:
```
npm install cypress --save-dev
```
Open cypress in the terminal using:
```
npx cypress open
```
Click on "Component Testing" in the cypress window, and select angular for the frontend framework.
Make sure you have all of the correct dependencies installed and continue.

## Running the frontend and the backend
If you want to run both the frontend and the backend and see real-time updates of the frontend, first navigate to the frontend folder and run `ng serve`.
```
cd frontend
ng serve
```
Then, navigate to the backend and run go with the debug tag.
```
cd ../backend
go run -tags debug .
```
Now, you can see the frontend at `http://localhost:8080` and also access the API at `http://localhost:8080` because the backend automatically forwards requests to the frontend.

## Building the frontend and running the backend
If you want to only work on the backend and have a static frontend, first navigate to the frontend folder and run `ng build`.
```
cd frontend
ng build
```
Then open a new terminal and navigate to the backend and run without tags.
```
cd ../backend
go run .
```
The server will now be running at http://localhost:8080/.
