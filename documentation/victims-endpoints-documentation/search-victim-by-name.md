#### #Search a victim by NAME

|URL | `/v1/victims?victim-name={victim-name}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params  |
    "victim-name" = {victim's name}
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
- You will receive a null value in the field of the date of birth.