#### #Add a new option

|URL | `/v1/options`  |
|:-:|:-:|
|  Method  | Post |
|  Header  | Content-Type: application/json |
|  Header  | Authorization: Bearer $Token |
|  Authorization Level | all |
|  Body  |Example:
    {
        "title": "imprisoned",
        "group": "current_status"
    }

| Success Response | Code: 201  |
|:-:|:-:|

Content:

        "message": "option has been added",
        "status": 201
        "option": #option

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400