#### #Add a new incident

|URL | `/v1/victims/{idvictim}/incidents`  |
|:-:|:-:|
|  Method  |Post|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "date_of_incident": "2015-09-22T22:42:31+07:00",
        "location": "spain",
        "is_disappearance": true,
        "IncidentTranslation": [
            {
                "language": "en",
                "narrative_of_incident": "Narrrative test"
            },
            {
                "language": "ch",
                "narrative_of_incident": "书、杂志等中区别于图片的"
            }
        ],
        "IncidentMedia": [
            {
                "date_of_media": "2115-09-22T22:42:31+07:00",
                "mediaurl": "https://eloutput.com/app/uploads-eloutput.com/2019/03/sonic-real-imagen-pelicula.jpg",
                "tag": "jpg"
            }
        ]
    }
| Success Response | Code: 201  |
|:-:|:-:|

Content:

        "message": "Incident has been added has been added",
        "status": 201
        "Incident": #Incident-data

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Victim ID not Exist",
        "status": 404