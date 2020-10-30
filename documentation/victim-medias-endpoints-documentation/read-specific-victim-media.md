#### #Read a specific victim media

|URL | `/v1/victimmedias?idvictimmedia={victim-mediaID}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params | "idvictimmedia" = $victim-mediaID|

|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "media": #media-data,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "media not found",
        "status": 404

**(If the request was made without authorization, the victim files with the tag: "documents" will not be sent in the response.)**