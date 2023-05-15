
# Login with Golang, MAMP (MySQL), and JWT Authentication
This is a simple web application that allows users to log in using a username and password. The login information is checked against a MySQL database hosted in MAMP. Upon successful login, a JSON Web Token (JWT) is generated and returned to the frontend for secure authentication.

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

# ![image](https://github.com/hakanemreyesil/Login/assets/59126857/8aa1845c-ddf8-4dc7-83b1-66b9292931c4) ![image](https://github.com/hakanemreyesil/Login/assets/59126857/19d2d83d-ea6d-4ae4-91dc-f3723a2a25b2) ![image](https://github.com/hakanemreyesil/Login/assets/59126857/6e70ee09-9521-4a7c-a10c-885b41bff594)
