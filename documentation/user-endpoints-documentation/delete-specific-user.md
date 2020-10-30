#### #Delete a specific user

|URL | `/v1/users/{iduser}`  |
|:-:|:-:|
|  Method  |Delete|
|  Header  | Authorization: Bearer $Token |
|  Authorization Level | admin |
|Success Response | Code: 200  |

Content:

        "message": "the user was deleted successfully",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "user Not Found",
        "status": 404

| Error Response | Code: 403  |
|:-:|:-:|

Content:

        "message": "You are not authorized to access.",
        "status": 403
