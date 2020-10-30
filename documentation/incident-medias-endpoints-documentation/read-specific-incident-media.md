#### #Read a specific incident media

|URL | `/v1/incident-medias?idincident-media={incident-mediaID}`  |
|:-:|:-:|
|  Method  |Get|
|  Query Params | "idincident-media" = $incident-mediaID|

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
