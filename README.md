# Boolean as a service
## Setup Instructions:
- Packages required to run the app: 
  ```go
  go get github.com/gin-gonic/gin
  go get github.com/dgrijalva/jwt-go
  go get github.com/go-sql-driver/mysql
  go get github.com/jinzhu/gorm
  go get github.com/rs/xid
  ```
- Database :
  By default the app uses a local MySQL database with user credentials as shown below and intialised in the **db.go** file in the package **dbConfig**
   ```go
   var dbUserName = "root"
   var dbPassword = ""
   var dbName = "booldb"
   ```
   The same can be modified if any changes to the credentials or the database are necessary

## Instructions to run the API:
The API server is configured to listen to the port **8080**
#### Authentication:
- It is required for the user to authenticate him/herself to make any API calls. The API itself performs token based authentication with the help of cookies. The cookies expire after 5 minutes
- Authentication needs to be performed using the below endpoint by passing the username and password in the JSON format
  ```sh
  /login
  ```
- By default it accepts the following credentials 
  ```sh
  username: "root"
  password: "password"
- Authenticated users can further add new users using the endpoint
  ```sh
  /adduser
  ```
  and supply the new username and password in JSON format

#### Endpoints:
-  ```sh
   POST "/" 
   ```
   Adds a new boolean to the database. Input JSON format expected is:
   ```sh
   "value": true/false (If not supplied default is false)
   "label": "Sample label"  (Optional)
   ```
- ```sh
  GET "/:id"  
  ```
  Extracts a boolean record (if it exists) from the database whose id matches the given input id
- ```sh
  PATCH "/:id"
  ```
  Updates the boolean record (if it exists) from the database whose id matches the given input id, using the information supplied in the JSON format:
  ```sh
  "value":   (Optional)
  "label":   (Optional)
  ```
- ```sh
  DELETE "/:id"
  ```
  Deletes a boolean record (if it exists) from the database whose id matches the given input id

- Also note that sync.RWMutex has been used to synchronize database access.
  
# Project Structure
The project has three packages (apart from main)
- **models**  
The models.go file in the models package consists of the definitions of various structs to be used to bind to user JSON input and/or define database schemas
- **controller**
The controller package consists of logic to handle various endpoints and authentication
- **dbConfig**
This package is used to handle database connection and configuration
