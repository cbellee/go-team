import React, { Component } from 'react';
import Players from './components/players';

class App extends Component {
  render() {
    return (
      <Players players={this.state.players}/>
    )
  }

  state = {
    players: []
  };

  componentDidMount() {
    fetch('http://localhost:8080/players')
      .then(res => res.json())
      .then((data) => {
        this.setState({ players: data })
      })
      .catch(console.log)
  }
}

export default App;
