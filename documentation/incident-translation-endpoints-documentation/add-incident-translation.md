#### #Add a new incident translation

|URL | `/v1/incidents/{idincident}/translations`  |
|:-:|:-:|
|  Method  |Post|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "language": "es",
        "narrative_of_incident": "hola eventos narrativa"
    }
| Success Response | Code: 201  |
|:-:|:-:|

Content:

        "message": "incident Translation has been added",
        "status": 201
        "incident_translation": #incident-translation-data

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Incident ID not Exist",
        "status": 404
