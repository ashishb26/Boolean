# Boolean
## Features implemented
- This api provides the user the ability to perform CRUD operation on boolean values.
- A MYSQL database is used to perform the aforementioned operations.
- Redis is further used to execute a distributed mutex locking mechanism

### Authorization
- The api uses a jwt based authentication. 
- The endpoint for authentication is `/user/login`and the default credentials are: `username: root , password: password`

### Endpoints
- **Add New Boolean** --> **POST** `/api/`. 

  The expected JSON input format is 
  ```sh
  "value": true/false
  "key" : "Example key"
  ```
- **Get Boolean** --> **GET** `/api/:id`. 
  Here id refers to the unique id given to the boolean as it is stored in the database
  
- **Update Boolean** --> **PATCH** `/api/:id`. 
  This endpoint can be used to update either the value, key or both 
  
- **Delete Boolean** --> **DELETE** `/api/:id`
