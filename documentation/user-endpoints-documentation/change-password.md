#### #Change Password from a user

|URL | `/v1/users/change-password`  |
|:-:|:-:|
|  Method  |Post|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "email": "test@gmail.com",
        "old_password": "123456",
        "new_password": "aaaaaa"
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "user Password have been updated",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Password does not match!",
        "status": 400
