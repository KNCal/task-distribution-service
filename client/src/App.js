import React from "react";
import "./App.css";
import { Segment, Divider, Grid } from "semantic-ui-react";
import CreateTask from "./components/CreateTask";
import AllTasks from "./components/AllTasks";

function App(props) {
  console.log('app.js rendering')
  return (
    <div>
      <Segment>
        <Grid columns={2} relaxed='very'>
          <Grid.Column>
            <CreateTask createTask={props.createTask}/>
          </Grid.Column>
          <Grid.Column>
            <AllTasks deleteTask={props.deleteTask} taskComplete={props.taskComplete} tasks={props.tasks}/>
          </Grid.Column>
        </Grid>
        <Divider vertical>OR</Divider>
      </Segment>
    </div>
  );
}

export default App;

