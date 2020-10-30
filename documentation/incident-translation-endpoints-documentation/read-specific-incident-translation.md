#### #Read a specific incident translation

|URL | `/v1/incident-translations?idincident-translation={idincident-translationID}`  |
|:-:|:-:|
|  Method  |Get|
|  Query Params | "idincident-translation" = $idincident-translationID|

|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "translation": #translation-data,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Incident Translation not found",
        "status": 404
