# How to deploy the API to Heroku

1. First, you  must create a Heroku account with this link to register:
    [Sign up Heroku](https://signup.heroku.com/)

2. Now must download the Heroku Command Line Interface (CLI) to manage the applications. With the following link, you can download CLI through the operating system you use.

    [Heroku-Cli Download and Install](https://devcenter.heroku.com/articles/heroku-cli#download-and-install)

    :warning: *Note: The Heroku CLI requires **[GIT](https://git-scm.com/downloads)***

3. To verify that the installation was successful, use the **heroku –version** command in the terminal, which should show you a message indicating the version of the Heroku that you installed.

    ![heroku-version](heroku-documentation-img/heroku-version.png)

4. After you install the CLI, run the **heroku login** command  You’ll be prompted to enter any key to go to your web browser to complete login. The CLI will then log you in automatically.

    ![heroku-login](heroku-documentation-img/heroku-login.png)

5. Must clone the application on your computer with the following command **git clone name-project**. As the code has already been uploaded to my personal heroku account, you do not need to add or install any additional dependencies. After the project is cloned, in the terminal you must go to the folder where the project is cloned

    ![git-clone](heroku-documentation-img/git-clone.png)

6. You must create an app in Heroku, which will be the one that will receive the project code, to create it you do it from the same terminal with the following command **heroku create name-project**

    ![create-app-heroku](heroku-documentation-img/create-app-heroku.png)

7. When you create an app, a git remote (called heroku) is also created and associated with your local git repository. You can use the **git remote -v** command to confirm that a remote named heroku has been set for your app

    ![git-remote](heroku-documentation-img/git-remote.png)

8. To deploy in Heroku you need to enter the following command which will send your code to the app you just created in Heroku **git push heroku master**. *Note *: this same command you must enter if there was a change or you added something new to your code locally and want to send it to Heroku.*
     
    ![git-push-heroku](heroku-documentation-img/git-push-heroku.png)
    ![git-push-heroku-2](heroku-documentation-img/git-push-heroku-2.png)

9. If everything went correctly on the Heroku website it should appear that the application was built successfully.

    ![show-build-succesufl](heroku-documentation-img/show-build-succesufl.png)

10. The application is already running in Heroku you can open it in your browser with the command
**heroku open** but you still have some configurations. You can also observe the dyno (container) application is running **heroku ps**

    ![heroku-open-dyno](heroku-documentation-img/heroku-open-dyno.png)

11. The next thing is setting up a remote Postgresql database run: *Hobby-dev* represents the plan that Heroku offers for free. **heroku addons: create heroku-postgresql: hobby-dev**

    ![heroku-postgresql](heroku-documentation-img/heroku-postgresql.png)

12. Go to heroku.com, you should find your newly created app on your dashboard, click on it. After that, click on settings, then click on *Reveal Config Vars*
There is a var named *DATABASE_URL* , this was automatically added to your *.env* file when you created the postgresql database (Note: Heroku automatically replace your local .env when you deploy your app), from this var, you will extract our database connection parameter.

    *postgres://username:password@host/dbName*

    ![heroku-settings](heroku-documentation-img/heroku-settings.png)

    ![var-heroku](heroku-documentation-img/var-heroku.png)

    ![setting-hide-var-postgres](heroku-documentation-img/setting-hide-var-postgres.png)

13. If all this went well, your API should be live now!

    ![postman](heroku-documentation-img/postman.png)

    ![database-user-heroku](heroku-documentation-img/database-user-heroku.png)

