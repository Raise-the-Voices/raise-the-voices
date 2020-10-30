#### #Delete a specific incident

|URL | `/v1/incidents/{idincident}`  |
|:-:|:-:|
|  Method  |Delete|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|Success Response | Code: 200  |

Content:

        "message": "Incident was deleted successfully",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Incident not Found",
        "status": 404