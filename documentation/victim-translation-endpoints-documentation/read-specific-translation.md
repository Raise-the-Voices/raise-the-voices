#### #Read a specific translation

|URL | `/v1/translations?idtranslation={translationID}`  |
|:-:|:-:|
|  Method  |Get|
|  Query Params | "idtranslation" = $translationID|

|Success Response | Code: 200  |
|:-:|:-:|

Content:

        "translation": #translation-data,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Translation not found",
        "status": 404
