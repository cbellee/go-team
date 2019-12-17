import React from 'react';
import { Card, CardBlock, CardFooter, CardTitle, CardText } from 'react-bootstrap';

const Players = ({ players }) => {
    return (
        <Card>
            <CardBlock>
                <CardTitle>
                    Player List
                </CardTitle>
                <CardText>
                    {players.map((player) => (
                        player.FirstName
                        //<div class="card" key={player.ID}>
                        //<div class="card-body" style={{ backgroundColor: player.TeamColourHex }}>
                        //<h5 class="card-title">{player.FirstName} {player.LastName}</h5>
                        //</div>
                        //</div>
                    ))}
                </CardText>
            </CardBlock>
            <CardFooter>
            </CardFooter>
        </Card>
    )
};

export default Players