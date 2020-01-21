import React from 'react';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import { Link } from 'react-router-dom';

class Players extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isFetching: false,
            teamPlayers: []
        }
    }

    componentDidMount() {
        this.PlayerList();
    }

    PlayerList() {
        fetch('http://localhost:8080/teamPlayers')
            .then(res => res.json())
            .then((data) => {
                this.setState({ teamPlayers: data })
            })
            .catch(console.log)
    }

    render() {
        const players = this.state.teamPlayers.map((player) => (
            <Container>
                <Card key={player.ID}>
                    <Card.Header style={{ background: player.TeamColourHex }}>
                        {player.TeamName}
                    </Card.Header>
                    <Card.Body>
                        <Card.Title><Link to={"/players/" + player.ID}> {player.FirstName} {player.LastName}</Link></Card.Title>
                        <Card.Subtitle>{player.Email}</Card.Subtitle>
                        <Card.Text></Card.Text>
                    </Card.Body>
                    <Card.Footer>
                    </Card.Footer>
                </Card>
            </Container>
        ));

        return (<div>{players}</div>)
    }
}

export default Players
