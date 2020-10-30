#### This API Allow at most 100 request per 1 second.

## Authentication:

Details about JWT:

Expiration Time: 24 hours.

## Authorization:

There are 3 types of authorization levels

1. all
2. admin
3. editor

Errors JWT:

| Error Response | Code: 403  |
|:-:|:-:|       

Content:

        "message": "Missing auth token",
        "status": 403

| Error Response | Code: 403  |
|:-:|:-:|       

Content:

        "message": "Invalid/Malformed auth token",
        "status": 403

| Error Response | Code: 403  |
|:-:|:-:|       

Content:

        "message": "Malformed or Expired authentication token",
        "status": 403

| Error Response | Code: 403  |
|:-:|:-:|       

Content:

        "message": "Token is not valid.",
        "status": 403

| Error Response | Code: 403  |
|:-:|:-:|       

Content:

        "message": "You are not authorized to access.",
        "status": 403
