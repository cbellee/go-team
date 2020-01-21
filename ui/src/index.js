import React from 'react';
import ReactDOM from 'react-dom';
import { Route, Link, Switch, BrowserRouter as Router } from 'react-router-dom'
import App from './App';
import Players from './components/players';
import Player from './components/player';
import Teams from './components/teams';
import NotFound from './components/notfound';
import './index.css';

const routing = (
  <Router>
    <div>
      <ul>
        <li><Link to="/">Home</Link></li>
        <li><Link to="/players">Players</Link></li>
        <li><Link to="/teams">Teams</Link></li>
      </ul>
      <Switch>
        <Route exact path="/" component={App} />
        <Route exact path="/players" component={Players} />
        <Route path="/players/:id" component={Player} />
        <Route path="/teams" component={Teams} />
        <Route component={NotFound} />
      </Switch>
    </div>
  </Router>
)

ReactDOM.render(
  routing,
  document.getElementById('root')
);
