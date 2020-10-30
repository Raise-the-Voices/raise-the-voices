#### #Update report data

|URL | `/v1/reports/{idreport}`  |
|:-:|:-:|
|  Method  | Put |
|  Header  | Content-Type: application/json |
|  Header  | Authorization: Bearer $Token |
|  Authorization Level | all |
|  Body  |Example:
    {
        "name_of_reporter": "pedro",
        "email_of_reporter": "testjesus@gmail.com",
        "discovery": "test false",
        "is_direct_testimony": false
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "Report have been updated",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "report not found or was deleted",
        "status": 404
