import React from 'react'

const Players = ({ players }) => {
    return (
        <div>
            <center><h1>Player List</h1></center>
            {players.map((player) => (
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">{player.FirstName} {player.LastName}</h5>
                            <h6 class="card-subtitle mb-2 text-muted">{player.Email}</h6>
                        </div>
                    </div>
                ))}
        </div>
    )
};

export default Players