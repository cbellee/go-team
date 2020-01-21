import React from 'react';
import Card from 'react-bootstrap/Card';
import Container from 'react-bootstrap/Container';
import { Route, Link, Switch, BrowserRouter as Router } from 'react-router-dom'

class Teams extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isFetching: false,
            teams: []
        }
    }

    componentDidMount() {
        this.TeamList();
    }

    TeamList() {
        fetch('http://localhost:8080/teams')
            .then(res => res.json())
            .then((data) => {
                this.setState({ teams: data })
            })
            .catch(console.log)
    }

    render() {
        const teams = this.state.teams.map((team) => (
            <Container>
                <Card key={team.ID}>
                    <Card.Header style={{ background: team.Colour.Hex }}>
                        {team.Name}
                    </Card.Header>
                    <Card.Title></Card.Title>
                        <Card.Subtitle>Players</Card.Subtitle>
                    <Card.Body>
                        <ul>
                            {
                                team.Players.map((player) => (
                                    <li>
                                        <Link key={player.ID} to={"/players/" + player.ID}>
                                            {player.FirstName}
                                        </Link>
                                    </li>
                                ))
                            }
                        </ul>
                    </Card.Body>
                    <Card.Footer>
                    </Card.Footer>
                </Card>
            </Container>
        ));

        return (<div>{teams}</div>)
    }
}

export default Teams

/* const Teams = ({ teams }) => {
    return (
        <Container>
            {
                teams.map((team) => (
                    <Card bg={team.TeamColourHex} key={team.ID}>
                        <Card.Body>
                            <Card.Text>
                                {team.ID}
                                {team.Name}
                            </Card.Text>
                        </Card.Body>
                        <Card.Footer>
                        </Card.Footer>
                    </Card>
                ))
            }
        </Container>
    )
};

state = {
    teams: []
  };

  componentDidMount() {
    fetch('http://localhost:8080/teams')
      .then(res => res.json())
      .then((data) => {
        this.setState({ teams: data })
      })
      .catch(console.log)
  }; */
