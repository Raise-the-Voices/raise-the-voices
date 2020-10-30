#### #Update data from a victim

|URL | `/v1/victims/{idvictim}`  |
|:-:|:-:|
|  Method  |Put|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "name": "Jesus 4334",
        "legal_name": "Jesus Gonzalez Leon",
        "Aliases": "tato",
        "place_of_birth": "Puerto La Cruz",
        "date_of_birth": "1993-09-22T22:42:31+07:00",
        "current_status": "released",
        "country": "Venezuela",
        "gender": "Male",
        "last_seen_date": "2002-09-22T22:42:31+07:00",
        "last_seen_place": "Lecherias",
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "victim have been updated",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Victim not found",
        "status": 404

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Error! this is the last Administrator user please create an additional administrator user.",
        "status": 404

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Error! updating user",
        "status": 404


