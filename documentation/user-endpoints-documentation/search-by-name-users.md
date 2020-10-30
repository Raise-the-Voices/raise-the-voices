#### #Search by name to users

|URL | `/v1/users?name={ParamVar}`  |
|:-:|:-:|
|  Method  | Get |
|  Query Params  |
    "name" = {name of user}
|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "users": users data,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "users not found",
        "status": 404
