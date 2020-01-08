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
        
    4. On the browser, enter: localhost:3000 in the url to bring up page.
    5. Use mongo shell or Robomongo to verify updates to the database 'test'

Future Additions:
    - Create microservices using Docker or client, server and MongoDB.
    - Include test cases and how to run them.

Author:
Kim Nguyen, December 2019
