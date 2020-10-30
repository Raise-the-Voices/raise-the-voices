#### #Read a specific report

|URL | `/v1/reports?idreport={ParamVar}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params  |
    "idreport" = #idreport
|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "report": report data,
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
