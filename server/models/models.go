package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task     string             `json:"taskName"`
	Priority string             `json:"prioritySelect"`
	Skills   []string           `json:"skillsArr"`
	Agent    string             `json:"agent"`
	Status   bool               `json:"status"`
}

type Assignment struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AgentID      string             `json:"agentid"`
	TaskID       string             `json:"taskid"`
	TaskPriority string             `json:"taskpriority"`
	Busy         bool               `json:"busy"`
	TimeAssigned time.Time          `json:"timeassigned"`
}
