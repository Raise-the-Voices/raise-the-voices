#### #Add a new victim media

|URL | `/v1/victims/{idvictim}/victimmedias`  |
|:-:|:-:|
|  Method  |Post|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "date_of_media": "1990-09-22T22:42:31+07:00",
        "mediaurl": "path_from_media",
        "tag":"media_format"
    }

| Success Response | Code: 201  |
|:-:|:-:|

Content:

        "message": "victim media has been added",
        "status": 201
        "victim_media": #victim_media-data

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Victim ID not Exist",
        "status": 404
