#### #Refresh Token

|URL | `/v1/refresh`  |
|:-:|:-:|
|  Method  | Post |
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |

| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "token":   tokenString,
        "message": "Refreshed token",
        "Expires": expirationTime,

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Still have more than 30 seconds for the token time to expire",
        "status": 400