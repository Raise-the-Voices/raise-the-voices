#### #Read all users

|URL | `/v1/users`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token |
|  Authorization Level | admin |

|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "users": []users-Array,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 403  |
|:-:|:-:|

Content:

        "message": "You are not authorized to access.",
        "status": 403
