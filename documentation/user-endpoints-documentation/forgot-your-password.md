#### #Forgot your password

|URL | `/v1/password/email`  |
|:-:|:-:|
|  Method  | Post |
|  Header  | Content-Type: application/json |
|  Body  |Example:
    {
        "email": "jesuseltato@hotmail.com"
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "message": "the email was sent successfully.",
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Email address is required",
        "status": 404

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Error to send email",
        "status": 404
