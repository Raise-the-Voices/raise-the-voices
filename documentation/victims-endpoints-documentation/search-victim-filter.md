#### #Search victim by Filter

|URL | `/v1/victims?victim-name={victim-name}&country={country}&status={status}&report-state={report-state}&sort={sort_by order}`  |
|:-:|:-:|
|  Method  |Get|
|  Header  | Authorization: Bearer $Token|
|  Authorization Level | all |
|  Query Params  |
    "victim-name" = {victim-name}
    "country" = {country}
    "status" = {status}
    "report-state" = {report-state} (published, in-review or all)
    "sort" = {sort_by order} (Example: created_at asc) (sort_by: some field of the victim) (order: desc or asc)
|Success Response(Example) | Code: 200  |
|:-:|:-:|

Content:

    {
        "status": 200,
        "victim": [
            {
                "ID": 1,
                "CreatedAt": "2020-04-16T09:49:12.31873-04:00",
                "UpdatedAt": "2020-04-16T09:49:12.31873-04:00",
                "DeletedAt": null,
                "name": "jesus gonzalez",
                "legal_name": "Jesus Rafael Gonzalez Leon",
                "aliases": "",
                "place_of_birth": "Puerto La Cruz",
                "date_of_birth": "1993-03-15T11:42:31-04:00",
                "current_status": "released",
                "country": "Venezuela",
                "last_seen_date": "2010-09-15T11:12:31-04:30",
                "last_seen_place": "Barcelona",
                "profile_image_url": "https://pngimage.net/wp-content/uploads/2019/05/taehyung-side-profile-png-.png",
                "Report": {
                    "ID": 0,
                    "CreatedAt": "0001-01-01T00:00:00Z",
                    "UpdatedAt": "0001-01-01T00:00:00Z",
                    "DeletedAt": null,
                    "victimid": 0,
                    "state": ""
                },
                "Incident": null,
                "VictimMedia": null,
                "VictimTranslation": null
            }
        ]
    }

|Success Response (no victim was found) | Code: 200  |
|:-:|:-:|

Content:

    {
        "status": 200,
        "victim": []
    }

| Error Response | Code: 400  |
|:-:|:-:|

Content:

        "message": "Params Error",
        "status": 400

| Error Response | Code: 404  |
|:-:|:-:|

Content:

        "message": "Some column in the database does not exist",
        "status": 404

##### *1-)NOTE*: all 5 filters must be in the URL, no matter if they are empty

##### *2-)NOTE*: If the authorization header with the token is not sent:

- You will receive a null value in the field of the date of birth.
