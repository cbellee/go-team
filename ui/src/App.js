import React, { Component } from 'react';
import Players from './components/players';
import Teams from './components/teams';
import Tabs from 'react-bootstrap/Tabs';
import Tab from 'react-bootstrap/Tab';

//import 'jquery/src/jquery'; //for bootstrap.min.js
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.min.js';

class App extends Component {
  render() {
    return (
      <div>
        <h1>Go Team</h1>
      </div>
    );
  }
}

export default App


/*  <Tabs defaultActiveKey="teams" activeKey="teams">
      <Tab title="Players" eventkey="players">
        <Players players={this.state.teamPlayers} teams={this.state.teams} />
      </Tab>
      <Tab title="Teams" eventkey="teams">
        <Teams teams={this.state.teams} />
      </Tab>
  </Tabs> */

       
  /* 
    state = {
      teamPlayers: [],
      teams: []
    };
  
    componentDidMount() {
      fetch('http://localhost:8080/teamPlayers')
        .then(res => res.json())
        .then((data) => {
          this.setState({ teamPlayers: data })
        })
        .catch(console.log)
  
      fetch('http://localhost:8080/teams')
        .then(res => res.json())
        .then((data) => {
          this.setState({ teams: data })
        })
        .catch(console.log)
    }
  } */
