#### #Add a new victim translation

|URL | `/v1/victims/{idvictim}/translations`  |
|:-:|:-:|
|  Method  |Post|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "language": "es",
        "nationality": "Venezolano",
        "health_status": "Normal",
        "health_issues": "Ninguno",
        "languagues_spoken": " Espanol",
        "profession": "Educador",
        "about_the_victim": "Joven",
        "additional_information": "Test adicional"
    }
| Success Response | Code: 201  |
|:-:|:-:|

Content:

        "message": "victim Translation has been added",
        "status": 201
        "victim_translation": #victim-translation-data

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Victim ID not Exist",
        "status": 404
