# API Endpoints
## GET `/users`
Returns a list of all users in JSON.

## POST `/users`
Adds a new user to the database. JSON format should follow: 
```json
{
    "username": "myUsername",
    "password": "myPassword",
    "isTutor": true
}
```

## GET `/users/{id}`
Returns data for a user with the corresponding `id`. 

## PUT `/users/{id}`
Updates user data for the user with the corresponding `id`. Follows same JSON format as above.

## DELETE `/users/{id}`
Deletes the user with the corresponding `id`.