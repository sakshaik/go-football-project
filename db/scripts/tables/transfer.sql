CREATE TABLE IF NOT EXISTS transfer_player (
    transfer_id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    from_club_id INTEGER NOT NULL,
    to_club_id INTEGER NOT NULL,
    fee FLOAT not null,
    currency_code TEXT NOT null,
    status TEXT not null,
    FOREIGN KEY(player_id) REFERENCES player(player_id),
    FOREIGN KEY(from_club_id) REFERENCES club(club_id),
    FOREIGN KEY(to_club_id) REFERENCES club(club_id)
)