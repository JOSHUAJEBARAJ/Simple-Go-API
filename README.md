# Simple-Go-API

### Lesson -1 
- Create the mux, create the handle func to handle the request and response, and create the server to listen and serve the request.

## Lesson -2

- We are going to the decouple the application into three tiers 
 - Presentation 
 - Logic Tier (Server)
 - Data Tier (Database)

 - Internal is the special folder that is used to store the internal packages that are not exported to the outside world.

 - Create a new folder called `internal` and create new folder called `todo` and `transport` inside the internal folder.

 - Next create the service and initialize the constructor . The reason why we are doing it is because we don't want to expose the variables inside the struct to the outside world.

 - Next create the functions to get and create the todo.

 - Next create the transport layer to handle the request and response.


 ### Database 


 ```
 CREATE TABLE todo_items(
id SERIAL PRIMARY KEY,
task VARCHAR(255) NOT NULL,
status VARCHAR(255) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

- We are going to use the third party library because it allows the connection pooling
