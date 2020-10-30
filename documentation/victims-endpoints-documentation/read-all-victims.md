#### #Read all reports

|URL | `/v1/victims`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |

|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "victims": []victims-Array,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

Note: If the authorization header with the token is not sent:
- You will receive a null value in the field of the date of birth