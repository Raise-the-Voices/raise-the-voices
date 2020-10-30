#### #Reset Password

|URL | `/v1/password/reset`  |
|:-:|:-:|
|  Method  | Post |
|  Header  | Content-Type: application/json |
|  Body  |Example:
    {
        "email": "testemail@hotmail.com",
        "password": "123456"
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "User Password have been modified",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Error! changing password.",
        "status": 404
