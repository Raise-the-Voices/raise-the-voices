#### #Read all translations from a incident

|URL | `/v1/incident-translations?idincident={incidentID}`  |
|:-:|:-:|
|  Method  |Get|
|  Query Params | "idincident" = $incidentID |

| Success Response | Code: 200  |
|:-:|:-:|

Content:

        "translations": []Translations-array,
        "status": 200

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400