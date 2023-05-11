
# Login with Golang and MAMP (MySQL)
This is a simple web application that allows users to log in using a username and password. The login information is checked against a MySQL database hosted in MAMP.

The app uses the httprouter library to handle GET and POST requests, and the html/template library to render HTML templates.

# Installation and Usage
To run this app, you need to have Go and MAMP installed on your machine.

Clone this repository to your local machine.
Start MAMP and make sure the MySQL server is running.
Create a new MySQL database and import the users.sql file included in the repository to create a users table.
Edit the config.json file to match your MySQL server settings (database name, username, and password).
Open a terminal window and navigate to the cloned repository folder.
Run the command go run main.go to start the app.
Open a web browser and go to http://localhost:8080 to access the login page.

# Login
On the login page, enter your username and password and click the "Log In" button. The app will check the entered information against the users table in the MySQL database. If the information is correct, you will be redirected to a welcome page that displays your username. If the information is incorrect, you will see an error message.
# ![image](https://user-images.githubusercontent.com/59126857/235627313-35ec8ea6-8623-4f01-adb0-628de7e35625.png ) ![image](https://user-images.githubusercontent.com/59126857/235627369-e17e3e01-ab7a-4a54-8456-cae7326a4212.png ) ![image](https://user-images.githubusercontent.com/59126857/235627661-b8e1c785-1bac-40e9-a352-2861d1c3ba93.png )

# ![afterlogin](https://github.com/hakanemreyesil/Login/assets/59126857/30c0eceb-8c76-40c0-8bf2-0b7872bb2539) ![login](https://github.com/hakanemreyesil/Login/assets/59126857/36cdcd8f-0492-43d0-8501-ad78fc6eeb69) ![Error](https://github.com/hakanemreyesil/Login/assets/59126857/8aa90c70-d38c-4be1-b025-71ab60bb91f9)
