import React from 'react'

<<<<<<< HEAD
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

=======
>>>>>>> 87ce71de57f6fac2dadffc2996d73276cb9da3a4
const Players = ({ players }) => {
    return (
        <div>
            <center><h1>Player List</h1></center>
            {players.map((player) => (
<<<<<<< HEAD

                <div class="card" key={player.ID}>
                    <div class="card-body">
                        <h5 class="card-title">{player.FirstName} {player.LastName}</h5>
                        <h6 class="card-subtitle mb-2 text-muted">{player.Email}</h6>
                    </div>
                </div>
            ))}
=======
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">{player.FirstName} {player.LastName}</h5>
                            <h6 class="card-subtitle mb-2 text-muted">{player.Email}</h6>
                        </div>
                    </div>
                ))}
>>>>>>> 87ce71de57f6fac2dadffc2996d73276cb9da3a4
        </div>
    )
};

export default Players