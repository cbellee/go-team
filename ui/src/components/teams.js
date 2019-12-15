import React from 'react'

const Teams = ({ teams }) => {
    return (
        <div>
            <center><h1>Team List</h1></center>
            {teams.map((team) => (
<<<<<<< HEAD
                <div class="card" key={team.ID}>
                    <div class="card-body" style={{color: 'white', backgroundColor: team.Colour.Hex}}>
                        <h5 class="card-title">{team.Name}</h5>
                        <div class="card-subtitle mb-2 text-muted">{team.Players.map((player) => <div>{(player.FirstName)} {(player.LastName)}</div>)}</div>
=======
                <div class="card">
                    <div class="card-body">
                        <h5 class="card-title">{team.Name}</h5>
                        <h6 class="card-subtitle mb-2 text-muted">{team.colourName}</h6>
                        <h6 class="card-subtitle mb-2 text-muted"></h6>
>>>>>>> 87ce71de57f6fac2dadffc2996d73276cb9da3a4
                    </div>
                </div>
            ))}
        </div>
    )
};

export default Teams