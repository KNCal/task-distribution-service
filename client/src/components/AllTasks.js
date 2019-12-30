import React, { Component } from "react";
import { Card, Header, Icon } from "semantic-ui-react";

class AllTasks extends Component {
  constructor(props) {
    super(props);
    
    this.state = {
      items: []
    }
  }

  render() {
    console.log(this.props);
    return (
      <div>
        <Header className="header" as="h1">
          Manage Tasks
        </Header>
        <br/>
        <div className="row">
          <Card.Group>
           {this.props.tasks.map(item => {
              // if task is assigned
              if (item.status) {
                let color = "green";
                return (
                    <Card key={item._id} color={color} fluid>
                      <Card.Content>
                        <Card.Header textAlign="left">
                          <div style={{ wordWrap: "break-word" }}>Task: {item.task}</div>
                        </Card.Header>
                        <Card.Header textAlign="center">
                          <div style={{ wordWrap: "break-word" }}>Agent: {item.agent}</div>
                        </Card.Header>
                        <Card.Meta textAlign="right">
                          <Icon
                            name="check circle"
                            color={color}
                            onClick={() => this.props.taskComplete(item._id)}
                          />
                          <span style={{ paddingRight: 10 }}>Mark Complete</span>
                        </Card.Meta> 
                      </Card.Content>
                    </Card>
                );
              }
              else {
                let color = "yellow";
                return (
                  <Card key={item._id} color={color} fluid>
                  <Card.Content>
                    <Card.Header textAlign="left">
                      <div style={{ wordWrap: "break-word" }}>Task: {item.task}</div>
                      <Card.Header textAlign="center">
                          <div style={{ wordWrap: "break-word" }}>No Agent Assigned</div>
                      </Card.Header>
                    </Card.Header>
                    <Card.Meta textAlign="right">
                      <Icon
                        name="check circle"
                        color={color}
                        onClick={() => this.props.deleteTask(item._id)}
                      />
                      <span style={{ paddingRight: 10 }}>Delete</span>
                    </Card.Meta> 
                  </Card.Content>
                </Card>
                )
              }
            })
          }
        </Card.Group>
        </div>
       </div>
    )
  }
}

export default AllTasks;