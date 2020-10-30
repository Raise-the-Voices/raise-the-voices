#### #Read all translation from victim

|URL | `/v1/translations?idvictim={victimID}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params | "idvictim" = $victimID |

| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "translations": []Translations-array,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

**(If the authorization header with the token is not sent, it will not receive all translation with the status of the report in revision.)**