#### #Delete a specific report

|URL | `/v1/reports/{idreport}`  |
|:-:|:-:|
|  Method  |Delete|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|Success Response | Code: 200  |

Content:

        "message": "the report was deleted successfully",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Report not found",
        "status": 404
