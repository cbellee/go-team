import React, { Component } from 'react';
import Players from './components/players';
import Teams from './components/teams';
import { Tabs, Tab } from 'react-bootstrap-tabs';

class App extends Component {
  render() {
    return (
      <Tabs onSelect={(label => console.log(label + ' selected'))}>
        <Tab label="Players">
          <Players players={this.state.players} />
        </Tab>
        <Tab label="Teams">
          <Teams teams={this.state.teams} />
        </Tab>
      </Tabs>
    )
  }

  state = {
    players: [],
    teams: []
  };

  componentDidMount() {
    fetch('http://localhost:8080/players')
      .then(res => res.json())
      .then((data) => {
        this.setState({ players: data })
      })
      .catch(console.log)

    fetch('http://localhost:8080/teams')
      .then(res => res.json())
      .then((data) => {
        this.setState({ teams: data })
      })
      .catch(console.log)
  }
}

export default App;
