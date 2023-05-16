
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

# Linter 
 Linter is a tool for detecting style errors, potential errors, and violations of generally accepted coding conventions in the code you write. Linters play an important role in improving the quality of the code you write and building a more robust code base.

# ![image](https://github.com/hakanemreyesil/Login/assets/59126857/71a7cd38-3b51-4228-9601-ecf089fbccdc) ![image](https://github.com/hakanemreyesil/Login/assets/59126857/82c00f4c-f401-471b-a327-aba54ce0154f) ![image](https://github.com/hakanemreyesil/Login/assets/59126857/4fb5f6e4-8d9b-4a06-a608-d767f5e90c6c)
