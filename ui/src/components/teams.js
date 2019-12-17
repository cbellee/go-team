import React from 'react'
import { Card, CardBlock, CardFooter, CardTitle, CardText } from 'react-bootstrap'

const Teams = ({ teams }) => {
    return (
        <Card>
            <CardBlock>
                <Card.Title>Team List</Card.Title>
                <CardText>
                    {teams.map((team) => (
                        <div class="card" key={team.ID}>
                            <div class="card-body" style={{ color: 'white', backgroundColor: team.Colour.Hex }}>
                                <h5 class="card-title">{team.Name}</h5>
                                <div class="card-subtitle mb-2 text-muted">{team.Players.map((player) => <div>{(player.FirstName)} {(player.LastName)}</div>)}</div>
                            </div>
                        </div>
                    ))}
                </CardText>
            </CardBlock>
        </Card>
    )
};

export default Teams