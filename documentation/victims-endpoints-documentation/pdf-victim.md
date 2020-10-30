#### #Get Base64 string PDF file of a ID-victims array

|URL | `/v1/victims/pdf?idvictim={victim-ID}&idvictim={victim-ID}`  |
|:-:|:-:|
|  Method  |Get|

|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "PDF": #Base64-String,
        "message": "the pdf file from victim array was created successfully",
        "status": 200

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Error when creating the PDF",
        "status": 404

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

### Note: If you want to check that the string in the base64 format of the PDF you can use the following page which converts a base 64 string to a PDF file.

[converter BASE64 to PDF](https://base64.guru/converter/decode/pdf) link.

