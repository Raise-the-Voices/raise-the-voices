#### #Update a victim translation

|URL | `/v1/translations/{idtranslation}`  |
|:-:|:-:|
|  Method  |Put|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "language": "es",
        "nationality": "Chileno",
        "health_status": "Normal",
        "health_issues": "Ninguno",
        "languagues_spoken": " Ingles",
        "profession": "Educacdor",
        "about_the_victim": "Joven",
        "additional_information": "Adicional"
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "Victim Translation have been updated",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Victim Translation not found or was deleted",
        "status": 404
