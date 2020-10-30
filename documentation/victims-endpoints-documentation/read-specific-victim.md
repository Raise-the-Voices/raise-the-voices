#### #Read a specific victim

|URL | `/v1/victims?idvictim={ParamVar}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params  |
    "idvictim" = #idvictim
|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "victim": victim data,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "victim not found",
        "status": 404

Note: If the authorization header with the token is not sent:
- you will not receive the victim with the status of the report in revision.
- If the case is published you will receive a null value in the field of the date of birth.