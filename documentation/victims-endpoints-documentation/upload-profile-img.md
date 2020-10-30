#### #Upload a victim's profile photo

|URL | `/v1/victims/profile-img/{idvictim}`  |
|:-:|:-:|
|  Method  | Post |
|  Body    | formdata |
|  Type    | File |
|  KEY : myfile  | VALUE: The image |

| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "the upload was successful",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

        "message": "Victim not found",
        "status": 404

| Error Response | Code: 400  |
|:-:|:-:|

        "message": "Params Error",
        "status": 400

| Error Response | Code: 500  |
|:-:|:-:|

        "message": "Error! body form-part exceeds the allowed limit",
        "status": 500

| Error Response | Code: 500  |
|:-:|:-:|

        "message": "Error! formatting the file",
        "status": 500

| Error Response | Code: 500  |
|:-:|:-:|

        "message": "Error! The Object exceeds the allowed limit",
        "status": 500

| Error Response | Code: 500  |
|:-:|:-:|

        "message": "Error! Reading the entire object",
        "status": 500

| Error Response | Code: 500  |
|:-:|:-:|

        "message": "Error! It is not a permitted format",
        "status": 500

| Error Response | Code: 500  |
|:-:|:-:|

        "message": "Error! on AWS",
        "status": 500

Notes* :

| allowed formats | "image/jpeg" , "image/jpg", "image/png"  |
|:-:|:-:|

| maximum file size allowed | 32MB  |
|:-:|:-:| 

Example* :
    ![postman-example-upload](imgs/postman-example-upload.png)