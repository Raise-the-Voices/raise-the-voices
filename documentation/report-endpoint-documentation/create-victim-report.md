#### #Create Victim Report

|URL | `/v1/reports`  |
|:-:|:-:|
|  Method  |Post|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "name": "Jesus Gonzalez",
        "legal_name": "Jesus Rafael Gonzalez Leon",
        "place_of_birth": "Puerto La Cruz",
        "date_of_birth": "1993-03-15T22:42:31+07:00",
        "current_status": "prisoner",
        "country": "cuba",
        "gender": "Male",
        "last_seen_date": "2010-09-15T22:42:31+07:00",
        "last_seen_place": "Barcelona",
        "profile_image_url": "https://pngimage.net/wp-content/uploads/2019/05/taehyung-side-profile-png-.png",
        "Report": {
            "name_of_reporter" : "jesus",
            "email_of_reporter" : "testjesus@gmail.com",
            "discovery" : "test discovery",
            "is_direct_testimony": true
        },
        "VictimTranslation": [
            {
                "language": "en",
                "nationality": "Venezuela",
                "languagues_spoken": "Spanish, English",
				"about_the_victim": "about test"
            },
            {
                "language": "ug",
                "nationality": "ئۇيغۇر تىلى ",
                "languagues_spoken": "ياخشىمۇسىز",
                "additional_information": "ياخشىمۇسىز ياخشىمۇسىز"
            }
        ],
        "Incident": [
            {
                "date_of_incident": "1990-09-22T22:42:31+07:00",
                "location": "France",
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
                ]
            },
            {
                "date_of_incident": "2015-09-22T22:42:31+07:00",
                "location": "spain",
                "IncidentMedia": [
                    {
                        "date_of_media": "2115-09-22T22:42:31+07:00",
                        "mediaurl": "https://eloutput.com/app/uploads-eloutput.com/2019/03/sonic-real-imagen-pelicula.jpg",
                        "tag":"jpg"
                    }
                ]
            }
        ],
        "VictimMedia": [
            {
                "date_of_media": "2115-09-22T22:42:31+07:00",
                "mediaurl": "https://eloutput.com/app/uploads-eloutput.com/2019/03/sonic-real-imagen-pelicula.jpg",
                "tag":"jpg"
            }
        ]
    }
| Success Response | Code: 201  |
|:-:|:-:|

Content:

        "message": "report has been created",
        "status": 201
        "victim": #DataParams

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400
