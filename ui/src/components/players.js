import React from 'react'

const join = (x, y, primary, foreign, sel) => {
    const idx = x.reduce((idx, row) =>
        idx.set(row[primary], row),
        new Map);

    return y.map(row =>
        sel(idx.get(row[foreign]),
            row));
};

const result = join(this.players, this.teams, "teamID", "ID", ({name}, {name}, {name}));
console.table(result);

const Players = ({ players }) => {
    return (
        <div>
            <center><h1>Player List</h1></center>
            {players.map((player) => (

                <div class="card" key={player.ID}>
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