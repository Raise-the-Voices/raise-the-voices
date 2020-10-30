#### #Read all media from a victim

|URL | `/v1/victimmedias?idvictim={victimID}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params | "idvictim" = $victimID |

| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "medias": []Media-array,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

1. **(If the authorization header with the token is not sent, it will not receive the victim Media with the status of the report in revision.)**
2. **(If the request was made without authorization, the victim files with the tag: "documents" will not be sent in the response.)**