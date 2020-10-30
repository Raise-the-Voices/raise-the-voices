#### #Read all victims by report state

|URL | `/v1/victims?report-state={report-state}`  |
|:-:|:-:|
|  Method  |Get|
|  Query Params  |
    "report-state" = {report-state} (published or in-review)
|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "victim": victim-list,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400
