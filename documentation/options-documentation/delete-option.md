#### #Delete an option

|URL | `/v1/options/{optionid}`  |
|:-:|:-:|
|  Method  | Delete |
|  Header  | Authorization: Bearer $Token |
|  Authorization Level | all |
|Success Response | Code: 200  |

Content:

        "message": "option was deleted successfully",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "option not found",
        "status": 404