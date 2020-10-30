#### #Modify data of a user

|URL | `/v1/users/{iduser}`  |
|:-:|:-:|
|  Method  |Put|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "name": "jesus3",
        "email": "test-5@gmail.com",
        "phone": "04161234569",
        "user_role": "admin"
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "user have been updated",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "User not found",
        "status": 404
