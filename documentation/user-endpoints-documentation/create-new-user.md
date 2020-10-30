#### #Create a new user

|URL | `/v1/users`  |
|:-:|:-:|
|  Method  | Post |
|  Header  | Content-Type: application/json |
|  Header  | Authorization: Bearer $Token |
|  Authorization Level | admin 
|  Body  |Example:
    {
        "name": "jesus",
        "password": "123456",
        "email": "test-1@gmail.com",
        "phone": "04161234567",
        "user_role": "admin"
    }
| Success Response | Code: 201  |
|:-:|:-:|

Content:

        "message": "User has been created",
        "status": 201
        "user": #UserData

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 403  |
|:-:|:-:|

Content:

        "message": "You are not authorized to access.",
        "status": 403
