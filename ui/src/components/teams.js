import React from 'react'

const Teams = ({ teams }) => {
    return (
        <div>
            <center><h1>Team List</h1></center>
            {teams.map((team) => (
                <div class="card" key={team.ID}>
                    <div class="card-body" style={{color: 'white', backgroundColor: team.Colour.Hex}}>
                        <h5 class="card-title">{team.Name}</h5>
                        <div class="card-subtitle mb-2 text-muted">{team.Players.map((player) => <div>{(player.FirstName)} {(player.LastName)}</div>)}</div>
                    </div>
                </div>
            ))}
        </div>
    )
};

export default Teams