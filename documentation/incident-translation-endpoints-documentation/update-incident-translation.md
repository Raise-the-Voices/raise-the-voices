#### #Update a incident translation

|URL | `/v1/victim-translations/{idvictim-translation}`  |
|:-:|:-:|
|  Method  |Put|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "language": "es",
        "narrative_of_incident": "hola eventos narrativa"
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "Incident Translation have been updated",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Incident Translation not found or was deleted",
        "status": 404
