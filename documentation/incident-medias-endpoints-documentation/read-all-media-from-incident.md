#### #Read all media from a incident

|URL | `/v1/incident-medias?idincident={incidentID}`  |
|:-:|:-:|
|  Method  |Get|
|  Query Params | "idincident" = $incidentID |

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