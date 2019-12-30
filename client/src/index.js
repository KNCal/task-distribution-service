import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import axios from "axios";
import App from './App';
import * as serviceWorker from './serviceWorker';

let endpoint = "http://localhost:8080";

function createTask(task, priority, boxSelect) {
    console.log('====> here')
    console.log(boxSelect)
    
    let skillsArr = Object.keys(boxSelect).filter(item => boxSelect[item])

    axios
        .post(
        endpoint + "/api/task",
        JSON.stringify({
          "taskName" : task,
          "prioritySelect" : priority,
          "skillsArr" : skillsArr
        })
        ,
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded"
          }
        }
        )
        .then(res => { getTaskList() })
        };
  
function taskComplete(id) {
    axios
        .put(endpoint + "/api/alltasks/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
        })
        .then(res => {
            console.log(res);
            getTaskList();
        });
}

function deleteTask(id) {
  axios
      .delete(endpoint + "/api/deleteTask/" + id, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      }
      })
      .then(res => {
          console.log(res);
          getTaskList();
      });
}

function getTaskList() {
    axios
        .get(endpoint + "/api/alltasks").then(res => {
            console.log(res)
            if (res.data) { ReactDOM.render(<App createTask={createTask} deleteTask={deleteTask} taskComplete={taskComplete} tasks={res.data}/>, document.getElementById('root')); }
            else { ReactDOM.render(<App createTask={createTask} deleteTask={deleteTask} taskComplete={taskComplete} tasks={[]}/>, document.getElementById('root'));}
        });
};
getTaskList()

serviceWorker.unregister();