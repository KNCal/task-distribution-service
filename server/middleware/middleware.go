package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"../models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//********************
//** Initializations
//********************

// DB connection string
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "test"

// Collection Task List
const collTaskList = "tasklist"

// Collection Agent Assignment List
const collAgentAssignList = "assignlist"

// Collection object/instance
var collectionTaskList *mongo.Collection
var collectionAgentAssignList *mongo.Collection

// Array of agents
type Agents struct {
	Agents []Agent `json:"agents"`
}

// Agent has ID and skills
type Agent struct {
	AgentID string   `json:"agentID"`
	Skills  []string `json:"skills"`
}

var agents Agents

// create connection with mongo db
func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collectionTaskList = client.Database(dbName).Collection(collTaskList)
	collectionAgentAssignList = client.Database(dbName).Collection(collAgentAssignList)

	fmt.Println("collectionTaskList instance created!")

	// Read agents list from file
	jsonFile, e := ioutil.ReadFile("./agents.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Println("Successfully opened agents.json")

	// map to local struct
	error := json.Unmarshal([]byte(jsonFile), &agents)
	if error != nil {
		panic(error)
	}
}

//********************
//** Request Handlers
//********************

// Get all tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)
}

// Create task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("rBody", r.Body)

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(task, r.Body)

	taskID := insertOneTask(task)
	agentAssigned := getAgent(task)
	if agentAssigned != "0" {
		updateAssignment(taskID, task, agentAssigned)
		updateTask(taskID, agentAssigned)
		fmt.Println("AgentAssigned: ", agentAssigned)
	} else {
		fmt.Println("Error - No agents available")
	}
	task.Agent = agentAssigned
	json.NewEncoder(w).Encode(task)
}

// Task complete
func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//********************
//** CRUD functions
//********************

// Get all tasks
func getAllTasks() []primitive.M {
	cur, err := collectionTaskList.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

// Insert task
func insertOneTask(task models.Task) string {
	task.Agent = ""
	task.Status = false
	insertResult, err := collectionTaskList.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record into Task List", insertResult.InsertedID)

	if old, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return old.Hex()
	}
	return "0"
}

// Update task
func updateTask(task string, agent string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}

	// task status true means it's running/assigned an agent
	update := bson.M{"$set": bson.M{
		"status": true,
		"agent":  agent}}
	result, err := collectionTaskList.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

// Update Assignment Table with agent status busy, timeAssigned, task id
func updateAssignment(taskid string, task models.Task, agent string) {
	fmt.Println("Agent inside updateAssignment: ", agent)
	id, _ := primitive.ObjectIDFromHex(taskid)
	filter := bson.M{"agentid": agent}
	update := bson.M{"$set": bson.M{
		"busy":         true,
		"taskpriority": task.Priority,
		"timeassigned": time.Now(),
		"taskid":       id}}
	result, err := collectionAgentAssignList.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

// task complete method, update task's status to true
func taskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collectionTaskList.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

// Delete task
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collectionTaskList.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document", d.DeletedCount)
}

//********************
//** Utility functions
//********************

// check priority of tasks agents are working on
func checkLowPriority(aRecords []models.Assignment) bool {
	for i := 0; i < len(aRecords); i++ {
		if aRecords[i].TaskPriority != "Low" {
			return false
		}
	}
	return true
}

func findMatch(aRecords []models.Assignment) string {
	var agent string
	var latest time.Time
	agent = ""

	for _, record := range aRecords {
		// loop until agent is assigned
		if agent != "" {
			break
		}
		// if agent is not busy, return agent
		if !record.Busy {
			agent = record.AgentID
		}
	}

	if agent == "" {
		lowPriority := checkLowPriority(aRecords)
		if lowPriority {
			latest = aRecords[0].TimeAssigned
			for j := 1; j < len(aRecords); j++ {
				fmt.Println("Checking for most recently assigned agent...")
				if aRecords[j].TimeAssigned.After(latest) {
					latest = aRecords[j].TimeAssigned
					agent = aRecords[j].AgentID
				}
			}
			fmt.Println("Latest: ", latest)
			fmt.Println("Agent: ", agent)
		} else {
			agent = "0"
		}
	}
	return (agent)
}

func subset(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value] += 1
	}
	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}
	return true
}

// Get assignment records
func getAgentsRecords(matchedAgents []string) []models.Assignment {
	var results []models.Assignment
	result := models.Assignment{}
	for _, agent := range matchedAgents {
		filter := bson.D{primitive.E{Key: "agentid", Value: agent}}
		cur, err := collectionAgentAssignList.Find(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}
		for cur.Next(context.Background()) {
			e := cur.Decode(&result)
			if e != nil {
				log.Fatal(e)
			}
			results = append(results, result)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		cur.Close(context.Background())
	}
	return results
}

func getAgent(task models.Task) string {
	var matchedAgents []string
	var agent string

	for i := 0; i < len(agents.Agents); i++ {
		isSubset := subset(task.Skills, agents.Agents[i].Skills)
		if isSubset {
			matchedAgents = append(matchedAgents, agents.Agents[i].AgentID)
		}
	}
	agent = agents.Agents[0].AgentID
	agentsRecords := getAgentsRecords(matchedAgents)
	agent = findMatch(agentsRecords)
	fmt.Println("Agent Assigned: ", agent)
	return agent
}
