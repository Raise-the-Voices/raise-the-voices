# EndPoint List

1. ### _Users Endpoints_

    * [Create a new user](user-endpoints-documentation/create-new-user.md) `Post /v1/users`
    * [Signin User](user-endpoints-documentation/signin.md) `Post /v1/signin`
    * [Refresh Token](user-endpoints-documentation/refresh-token.md) `Post /v1/refresh`   
    * [Change Password from a user](user-endpoints-documentation/change-password.md) `Post /v1/users/change-password`
    * [Modify data of a user](user-endpoints-documentation/modify-user-data.md) `Put /v1/users/{iduser}`
    * [Read a specific user](user-endpoints-documentation/read-specific-user.md) `Get /v1/users?iduser={userID}`
    * [Read all users](user-endpoints-documentation/read-all-users.md) `Get /v1/users`
    * [Search by name to users](user-endpoints-documentation/search-by-name-users.md) `Get /v1/users?name={ParamVar}`
    * [Delete a specific user](user-endpoints-documentation/delete-specific-user.md) `Delete /v1/users/{iduser}`
    * [Forgot your password](user-endpoints-documentation/forgot-your-password.md) `Post /v1/password/email`
    * [Get user token and response if is valid or not](user-endpoints-documentation/token-validation.md) `Get /v1/password/reset/{token}`
    * [Reset Password](user-endpoints-documentation/reset-password.md) `Post /v1/password/reset`

    :warning: *Note:

    To activate the password recovery function, add an email and your password.
    link: [Func send()](./../models/user.go)
    in the send () function

    To send emails with your gmail you must disable Control access to less secure apps
    **[GMAIL](https://support.google.com/accounts/answer/6010255)**

    And add the page path of the website developed to capture the token
    add it in the function:
    [Func GenerateEmailToResetPassword()](./../models/user.go)

    The password recovery token has a duration of **1 hour**.

2. ### _Reports Endpoints_

    * [Create a victim report](report-endpoint-documentation/create-victim-report.md) `Post /v1/reports`
    * [Modify report state](report-endpoint-documentation/modify-report-state.md) `Patch /v1/reports/{idreport}`
    * [Update data from a report](report-endpoint-documentation/update-report-data.md) `Put /v1/reports/{idreport}`
    * [Read a report list](report-endpoint-documentation/read-report-list.md) `Get /v1/reports?state={varState}&offset={varOffset}&limit={varLimit}`
    * [Read a specific report](report-endpoint-documentation/read-specific-report.md) `Get /v1/reports?idreport={ParamVar}`
    * [Read all reports](report-endpoint-documentation/read-all-reports.md) `Get /v1/reports`
    * [Delete a specific report](report-endpoint-documentation/delete-report.md) `Delete /v1/reports/{idreport}`

3. ### _Victims Endpoints_

    * [Read a specific victim](victims-endpoints-documentation/read-specific-victim.md) `Get /v1/victims?idvictim={ParamVar}`
    * [Read all victims](victims-endpoints-documentation/read-all-victims.md) `Get /v1/victims`
    * [Read all victims by report state](victims-endpoints-documentation/read-all-victims-by-report-state.md) `Get /v1/victims?report-state={report-state}`
    * [Search a victim by NAME](victims-endpoints-documentation/search-victim-by-name.md) `Get /v1/victims?victim-name={ParamVar}`
    * [Search victim by Filter](victims-endpoints-documentation/search-victim-filter.md) `Get /v1/victims?victim-name={victim-name}&country={country}&status={status}&report-state={report-state}&sort={sort_by order}`
    * [Update data from a victim](victims-endpoints-documentation/update-victim-data.md) `Put /v1/victims/{idvictim}`
    * [Delete a specific victim](victims-endpoints-documentation/delete-victim.md) `Delete /v1/victims/{idvictim}`
    * [Get Base64 string PDF file of a ID-victims array](victims-endpoints-documentation/pdf-victim.md) `Get /v1/victims/pdf?idvictim={victim-ID}&idvictim={victim-ID}`   
    * [Upload a victim's profile photo according to the victim's ID in S3 (AWS)](victims-endpoints-documentation/upload-profile-img.md) `Post /v1/victims/profile-img/{idvictim}`
    * [Delete the victim's profile photo according to the victim's ID in S3 (AWS).](victims-endpoints-documentation/delete-profile-img.md) `Delete /v1/victims/profile-img/{idvictim}`

4. ### _Victim_Translations Endpoints_

    * [Add a new victim translation](victim-translation-endpoints-documentation/add-new-translation.md) `Post /v1/victims/{idvictim}/victim-translations`
    * [Update a victim translation](victim-translation-endpoints-documentation/update-translation.md) `Put /v1/victim-translations/{idvictim-translation}`
    * [Read a specific translation](victim-translation-endpoints-documentation/read-specific-translation.md) `Get /v1/victim-translations?idvictim-translation={translationID}`
    * [Read all translations from a victim](victim-translation-endpoints-documentation/read-all-translation-from-victim.md) `Get /v1/victim-translations?idvictim={victimID}`
    * [Delete a specific translation](victim-translation-endpoints-documentation/delete-translation.md) `Delete /v1/victim-translations/{idvictim-translation}`

5. ### _Victim_Medias Endpoints_

    * [Add a new victim media](victim-medias-endpoints-documentation/add-new-victim-media.md) `Post /v1/victims/{idvictim}/victimmedias`
    * [Read a specific victim media](victim-medias-endpoints-documentation/read-specific-victim-media.md) `Get /v1/victimmedias?idvictimmedia={victim-mediaID}`
    * [Read all media from a victim](victim-medias-endpoints-documentation/read-all-media-from-victim.md) `Get /v1/victimmedias?idvictim={victimID}`
    * [Delete a specific victim media](victim-medias-endpoints-documentation/delete-victim-media.md) `Delete /v1/victimmedias/{idvictimmedia}`
    * [Upload Victim Media File in S3 (AWS)](victim-medias-endpoints-documentation/upload-victim-media.md) `Post /v1/victimmedias/upload`

6. ### _Incidents Endpoints_

    * [Add a new incident](incident-endpoints-documentation/add-new-incident.md) `Post /v1/victims/{idvictim}/incidents`
    * [Update a specific incident](incident-endpoints-documentation/update-incident.md) `Put /v1/incidents/{idincident}`
    * [Read all incidents from a victim](incident-endpoints-documentation/read-all-incident-victim.md) `Get /v1/incidents?idvictim={victimID}`
    * [Delete a specific incident](incident-endpoints-documentation/delete-incident-by-id.md) `Delete /v1/incidents/{idincident}`

7. ### _Incidents_Translations Endpoints_

    * [Add a new incident translation](incident-translation-endpoints-documentation/add-incident-translation.md) `Post /v1/incidents/{idincident}/incident-translations`
    * [Read a specific incident translation](incident-translation-endpoints-documentation/read-specific-incident-translation.md) `Get /v1/incident-translations?idincident-translation={idincident-translationID}`
    * [Read all translations from a incident](incident-translation-endpoints-documentation/read-all-translation-from-incident.md) `Get /v1/incident-translations?idincident={incidentID}`
    * [Update a incident translation](incident-translation-endpoints-documentation/update-incident-translation.md) `Put /v1/victim-translations/{idvictim-translation}`
    * [Delete a specific incident translation](incident-translation-endpoints-documentation/delete-incident-translation.md) `Delete /v1/incident-translations/{idincident-translation}`

8. ### _Incident_Medias Endpoints_

    * [Add a new incident media](incident-medias-endpoints-documentation/add-incident-media.md) `Post /v1/incidents/{idincident}/incident-medias`
    * [Read a specific incident media](incident-medias-endpoints-documentation/read-specific-incident-media.md) `Get /v1/incident-medias?idincident-media={incident-mediaID}`
    * [Read all media from a incident](incident-medias-endpoints-documentation/read-all-media-from-incident.md) `Get /v1/incident-medias?idincident={incidentID}`
    * [Delete a specific incident media](incident-medias-endpoints-documentation/delete-incident-media.md) `Delete /v1/incident-medias/{idincident-media}`
    * [Upload Incident Media File in S3 (AWS)](incident-medias-endpoints-documentation/upload-incident-media.md) `Post /v1/incident-medias/upload`

9. ### _Options Endpoints_

    * [Add a new option](options-documentation/add-option.md) `Post /v1/options`
    * [Get all options](options-documentation/get-all-options.md) `Get /v1/options`
    * [Delete an option](options-documentation/delete-option.md) `Delete /v1/options/{optionid}`