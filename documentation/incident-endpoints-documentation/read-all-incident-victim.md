#### #Read all incidents from a victim

|URL | `/v1/incidents?idvictim={victimID}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params | "idvictim" = $victimID |

| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "incidents": []incidents-array,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

**(If the authorization header with the token is not sent, it will not receive the incidents with the status of the report in revision.)**