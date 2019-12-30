import React, { Component } from "react";
import { Card, Header, Form, Input, Button } from "semantic-ui-react";
import InputCheckbox from "./InputCheckbox";

const priorityLevel = {
  LOW: "Low",
  HIGH: "High"
}

class CreateTask extends Component {
  constructor(props) {
    super(props);    
    
    this.state = {
      task: "",
      items: [],
      priority: priorityLevel.LOW,
      boxSelect : { '1':false, '2': false, '3':false}
    }
    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  onChange = event => {
    switch (event.target.name) {
      case 'task':
        this.setState({ task: event.target.value });
        break;
      case 'priority':
        this.setState({ priority: event.target.value });
        break;  
      default:
        console.log("Unknown event: " + event.target.name);
    }
  };

  onChangeSelect = box_key => {
    console.log("Received: " + box_key);
    this.setState({boxSelect: {...this.state.boxSelect, [box_key]:this.state.boxSelect[box_key]?false:true}})
  };

  onSubmit = (e) => {
    e.preventDefault();
    this.props.createTask(this.state.task, this.state.priority, this.state.boxSelect);
    this.setState({ 
      task: "",
      priority: priorityLevel.LOW,
      boxSelect: { '1':false, '2': false, '3':false}
    });
  };

  render() {
    console.log("This state: " + this.state);
    console.log("This props: " + this.props);    
    return (
      <div>
        <div className="row">
          <Header className="header" as="h1">
            Create Task
          </Header>
        </div>
        <br/>

        <div className="row">
          <Form onSubmit={e => this.onSubmit(e)}>
            <Input
              type="text"
              name="task"
              value={this.state.task}
              required
              onChange={this.onChange}
              fluid
              placeholder="Task Name"
            />
            <br/>
            <label htmlFor="priority-options">Priority:</label>
            <select 
              className="browser-default custom-select" 
              required
              name="priority"
              value={this.state.priority}
              onChange={this.onChange} 
              id="priority-options">
              <option >{priorityLevel.LOW}</option>
              <option >{priorityLevel.HIGH}</option>
            </select>
            <br/>
              <InputCheckbox
                name="skills"
                boxSelect={this.state.boxSelect}
                onChangeSelect={this.onChangeSelect}
              />
            <br/>
            <Button type="submit">Create</Button>
          </Form>

        </div>
        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
    );
  }            
}

export default CreateTask;
