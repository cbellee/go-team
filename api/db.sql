USE goteamdb

CREATE TABLE IF NOT EXISTS colours (
    id INT AUTO_INCREMENT PRIMARY KEY,
    rgb CHAR(6) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO colours(rgb, name) 
VALUES
    ('FFFFFF', 'white'),
    ('000000', 'black'),
    ('FF0000', 'red'),
    ('00FF00', 'green'),
    ('0000FF', 'blue');

CREATE TABLE IF NOT EXISTS teams (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    colour_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO teams (name, colour_id) 
VALUES
    ('the white team', 1),
    ('the green team', 4),
    ('the blue team', 5),
    ('the black team', 2);

CREATE TABLE IF NOT EXISTS players (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    team_id INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO players (email, first_name, last_name, team_id) 
VALUES
    ('cbellee@microsoft.com', 'chris', 'bellee', 1),
    ('daved@home.net', 'david', 'double', 1),
    ('fgreene@gmail.com', 'fred', 'greene', 2),
    ('sam.smith@googlemail.com.au', 'sam', 'smith', 3);

/* CREATE TABLE IF NOT EXISTS team_members (
    team_member_id INT AUTO_INCREMENT PRIMARY KEY,
    player_id INT NOT NULL,
    team_id INT NOT NULL,
    CONSTRAINT fk_player
    FOREIGN KEY (player_id)
        REFERENCES players(player_id)
            ON UPDATE CASCADE 
            ON DELETE RESTRICT,
    CONSTRAINT fk_team
    FOREIGN KEY (team_id)
        REFERENCES teams(team_id)
            ON UPDATE CASCADE 
            ON DELETE RESTRICT 
); */
