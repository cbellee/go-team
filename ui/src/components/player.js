import React from 'react';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';

class Player extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isFetching: false,
            player: ""
        }
    }

    componentDidMount() {
        this.GetPlayer(this.props.match.params.id);
    }

    GetPlayer(id) {
        fetch(`http://localhost:8080/teamPlayers/${id}`)
            .then(res => res.json())
            .then((data) => {
                this.setState({ player: data })
            })
            .catch(console.log)
    }

    render() {
        return (
            <Container>
                <Card key={this.state.player.ID}>
                    <Card.Header style={{ background: this.state.player.TeamColourHex }}>
                        {this.state.player.TeamName}
                    </Card.Header>
                    <Card.Body>
                        <Card.Title>{this.state.player.FirstName} {this.state.player.LastName}</Card.Title>
                        <Card.Subtitle>{this.state.player.Email}</Card.Subtitle>
                        <Card.Text></Card.Text>
                    </Card.Body>
                    <Card.Footer>
                    </Card.Footer>
                </Card>
            </Container>
        )
    }
}

export default Player
