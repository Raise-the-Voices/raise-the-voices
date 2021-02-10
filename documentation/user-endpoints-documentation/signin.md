#### #Signin User

|URL | `v1/signin`  |
|:-:|:-:|
|  Method  | Post |
|  Header  | Content-Type: application/json |
|  Body  | Example:
    {
        "email": "exampleraisethevoices@gmail.com",
        "password": "fX2Weey=y?*Z98&N"
    }
| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "expires": expirationTime,
        "message": "Logged In",
        "role": user-role,
        "status" : 200,
        "name": user-name,
        "token":   tokenString,
       
Example:

        {
                "expires": "2020-07-17T08:57:41.2560226-04:00",
                "message": "Logged In",
                "role": "admin",
                "status": 200,
                "name": "Jesus Gonzalez",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVyYWlzZXRoZXZvaWNlc0BnbWFpbC5jb20iLCJ1c2VyX3JvbGUiOiJhZG1pbiIsImV4cCI6MTU5NDk5MDY2MX0.qkc_weWIUgrT7awKFdHal1n-sw5Kl4wLRBgAag3gvpI"
        }

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "invalid request",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|       

Content:

        "message": "Invalid login credentials. Please try again",
        "status": 404

| Error Response | Code: 404  |
|:-:|:-:|       

Content:

        "message": "Email address not found",
        "status": 404