import React from 'react'

const Players = ({ players }) => {
    return (
        <div>
            <center><h1>Player List</h1></center>
            {players.map((player) => (
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">yy{player.Email}</h5>
                            <h6 class="card-subtitle mb-2 text-muted">{player.FirstName} {player.LastName}</h6>
                            <h6 class="card-subtitle mb-2 text-muted"></h6>
                        </div>
                    </div>
                ))}
        </div>
    )
};

export default Players