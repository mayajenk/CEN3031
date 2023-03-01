# API Endpoints
## GET `/api/users`
Returns a list of all users in JSON.

## POST `/api/users`
Adds a new user to the database. JSON format should follow: 
```json
{
    "username": "myUsername",
    "password": "myPassword",
    "is_tutor": true
}
```

## GET `/api/users/{id}`
Returns data for a user with the corresponding `id`. 

## PUT `/api/users/{id}`
Updates user data for the user with the corresponding `id`. Follows same JSON format as above.

## DELETE `/api/users/{id}`
Deletes the user with the corresponding `id`.