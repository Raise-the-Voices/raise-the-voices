#### #Delete a specific victim

|URL | `/v1/victims/{idvictim}`  |
|:-:|:-:|
|  Method  |Delete|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|Success Response | Code: 200  |

Content:

        "message": "the victim was deleted successfully",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "victim not found",
        "status": 404

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Error! this is the last Administrator user please create an additional administrator user.",
        "status": 404

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Error! to delete user.",
        "status": 404
