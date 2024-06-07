
## Go Auth

This is a small dating CRUD app built using Go and MySQL.

### Implementation details and assumptions:

Storage:

- MySQL is used as the storage for the application.
- The current design and implementation won't scale to millions of users, but it shows the basic concepts of how to implement a user management system like a dating app using Go and an ORM on top of MySQL.

Security:

- I assumed that the fråont-end will be responsible for utilising the endpoints securely.
- The password is hashed using bcrypt before storing it in the database.

Adding Location to users and distanceFromMe:

- To keep the project simple and achievable in a short time frame, I have used the random package to generate a random location for each user.
- When the user calls the discover endpoint, a call is made to the database, and then a random location is generated for each user. Then all the users are sorted by location ascending and returned back to the request.
- The location number is generated in a way that is always higher then the calling user so it is easy to sort and serve in a response.


#### Endpoints testing:

I have included a postman collection in the root directory of the project. You can import the collection and test the endpoints.

#### Prerequisites:

- Docker
- Docker-compose

#### Installation:

1 - Clone the repository:

    git clone https://github.com/af-md/go-auth

2 - Navigate to the project directory:

    cd go-auth


3 - Run the docker script to run the application with Docker:

    ./docker_run_dev.sh
This script will build the Docker image and run the application in a Docker container.

4 - Test the application by using the postman collection in the root directory of the project. Below there is a guide that litsts the baisc endpoints that are exposed.

#### Exposed endpoints that are exposed and their descriptions

##### Create user:

Endpoint:

    POST: /api/user/create

Successful Response:

```json

{
    "result": {
        "ID": 24,
        "Email": "ali.al-mansour@gmail.com",
        "Password": "$2a$10$7kfL1N0AuCq.4X5tKS.a/uiPDH6NTrpoKbgKHbuxmnYO7OVY68ikq",
        "Name": "Ali Al-Mansour",
        "Gender": "Female",
        "Age": 21,
        "Location": 45
    }
} 
 ```

##### Login:

Endpoint:

    POST: /api/login

Successful Response:

```json
{
    "message": "User is logged in, make sure to save and attach the token to the next request as a Bearer token to show that the user is logged in. Without the token the request will be rejected.",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhbmlhLnNhbGVoQGV4YW1wbGUuY29tIn0.33MT2ubjNDpKJ0ZXF-x4_AjMNDo-6R5tWqYZqVYxIwo"
}
 ```

##### Discover:

    GET: /api/discover?age=$age&gender=$gender (query params must be provided)

Successful Response:

```json
{
    "results": [
        {
            "user": {
                "ID": 2,
                "Email": "omar.farooq@hotmail.com",
                "Password": "$2a$10$GA5MSnywJFA.4rY53dfyMe/YzkDA5l/Sweby43C9zxx6sVgPdaiDS",
                "Name": "Omar Farooq",
                "Gender": "Male",
                "Age": 66,
                "Location": 155
            },
            "distanceFromMe": 155
        },
        {
            "user": {
                "ID": 11,
                "Email": "ibrahim.khan@hotmail.com",
                "Password": "$2a$10$akeIpz99ApeYD9lJtrund.jH8uqEM7AJGj.QVsONwkqn2dLLoLGae",
                "Name": "Ibrahim Khan",
                "Gender": "Male",
                "Age": 41,
                "Location": 337
            },
            "distanceFromMe": 337
        },
        {
            "user": {
                "ID": 19,
                "Email": "maryam.al-hakim@outlook.com",
                "Password": "$2a$10$jfB/0Gb9v44bM.HAJWW7NujOeEyR43ZenbhTsiRNFZTj3hXrssJrS",
                "Name": "Maryam Al-Hakim",
                "Gender": "Male",
                "Age": 44,
                "Location": 554
            },
            "distanceFromMe": 554
        }
    ]
}
 ```

##### Swipe:

Endpoint:

    POST: /api/swipe

Successful Response:

```json
{
    "result": {
        "matched": true,
        "matchID": 3
    }
}
```