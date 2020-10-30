#### #Read a report list

|URL | `/v1/reports?state={varState}&offset={varOffset}&limit={varLimit}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params  |
    "state" = 'all' or 'in-review' or 'published'
    "offset" = #numberOffset
    "limit" = #numberlimit
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
