
## An Example of a Task Distrubtion Service

This is an example of a work distribution service. It receives a task with the skill(s) required by it, an agent can be assigned to it based on the following conditions:

•	The agent must possess at least all the skills required by the task
•	An agent cannot be assigned a task if they’re already working on a task of equal or higher priority.
•	The system will always prefer an agent that is not assigned any task to an agent already assigned to a task.
•	If all agents are currently working on a lower priority task, the system will pick the agent that started working on his/her current task the most recently.
•	If no agent is able to take the task, the service should return an error.

This project was built with go, ReactJS and MongoDB.

### Getting Started

Requirements: MongoDB, Go compiler and JS compiler

    For running the project:
    1. Install MongoDB: https://docs.mongodb.com/manual/installation/

    2. Install Go: https://golang.org/doc/install

    3. Install npm: https://www.npmjs.com/package/npm

    For viewing the data in MongoDB:
    You can use Mongo shell that comes with MongoDB download or use RoboMongo GUI: https://robomongo.org/download


### To Run the Project
    1. Start the database server. At command line, type:
        $ mongod
    
    2. Start the server. At root of project, type:
        $ cd server
        $ go run main.go 
    
    3. Start the web app. At the root of project, type:
        $ cd client
        $ npm install  (this installs the app's dependencies)
        $ npm start

Note:
    At the moment, when you start the server ($go run main.go) the first time, it will create a collection in the "test" database. 

    As a work around, for subsequent starts the server, comment out the following lines in main.go:

        - "./setupdb"
        - setupdb.SetUpTable()

    before starting the server again.

    If you accidentally start the server more than once without commenting, use Robomongo or Mongo shell to drop the table and start again. 

Future Additions:
    - Put in a check if collection exists before creating it.
    - Create microservices using Docker or client, server and MongoDB.
    - Include test cases and how to run them.

Author:
Kim Nguyen, December 2019
