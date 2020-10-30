#### #Update a specific incident

|URL | `/v1/incidents/{incident}`  |
|:-:|:-:|
|  Method  |Put|
|  Header  |Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Body  |Example:
    {
        "date_of_incident": "2020-09-22T22:42:31+07:00",
        "location": "Venezuela",
        "is_disappearance": true,
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "Incident have been updated",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Incident not found or was deleted",
        "status": 404
