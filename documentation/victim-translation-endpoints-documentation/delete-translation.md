#### #Delete a specific translation

|URL | `/v1/translations/{idtranslation}`  |
|:-:|:-:|
|  Method  |Delete|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|Success Response | Code: 200  |

Content:

        "message": "translation was deleted successfully",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Translation not Found",
        "status": 404