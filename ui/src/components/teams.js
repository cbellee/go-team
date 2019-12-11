import React from 'react'

const Teams = ({ teams }) => {
    return (
        <div>
            <center><h1>Team List</h1></center>
            {teams.map((team) => (
                <div class="card">
                    <div class="card-body">
                        <h5 class="card-title">{team.Name}</h5>
                        <h6 class="card-subtitle mb-2 text-muted">{team.colourName}</h6>
                        <h6 class="card-subtitle mb-2 text-muted"></h6>
                    </div>
                </div>
            ))}
        </div>
    )
};

export default Teams