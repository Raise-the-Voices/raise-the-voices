#### #Delete a specific victim media

|URL | `/v1/victimmedias/{idvictimmedia}`  |
|:-:|:-:|
|  Method  |Delete|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|Success Response | Code: 200  |

Content:

        "message": "Media was deleted successfully",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Media not Found",
        "status": 404