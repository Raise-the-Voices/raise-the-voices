#### #Modify Report State

|URL | `/v1/reports/{idreport}`  |
|:-:|:-:|
|  Method  | Patch |
|  Header  | Content-Type: application/json|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
| Success Response | Code: 200  |

Content:

        "message": "report state modified",
        "status": 200

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "report Not Found",
        "status": 404
