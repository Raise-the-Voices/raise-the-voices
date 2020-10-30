#### #Read a specific user

|URL | `/v1/users?iduser={userID}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token |
|  Authorization Level | admin |
|  Query Params  |
    "iduser" = #iduser
|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "user": user-data,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "User not found",
        "status": 404

| Error Response | Code: 403  |
|:-:|:-:|

Content:

        "message": "You are not authorized to access.",
        "status": 403
