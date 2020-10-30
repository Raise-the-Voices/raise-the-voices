#### #Read all reports

|URL | `/v1/reports`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |

|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "reports": []Report-Array,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400
