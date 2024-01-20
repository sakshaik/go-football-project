CREATE TABLE IF NOT EXISTS player_ext_ref (
    player_ext_ref_id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    club_id INTEGER NOT NULL,
    UNIQUE (player_id, club_id)
    FOREIGN KEY(player_id) REFERENCES player(player_id),
    FOREIGN KEY(club_id) REFERENCES club(club_id)
)