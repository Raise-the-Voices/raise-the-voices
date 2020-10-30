#### #Get user token and response if is valid or not

|URL | `/v1/password/reset/{token}`  |
|:-:|:-:|
|  Method  | Get |
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "the Token is valid",
        "user-email": #user-email
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Error! The Token was not found or not valid",
        "status": 404

| Error Response | Code: 401  |
|:-:|:-:|

Content:

        "message": "the Token is expired",
        "status": 401
