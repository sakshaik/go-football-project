CREATE TABLE IF NOT EXISTS club (
    club_id INTEGER PRIMARY KEY AUTOINCREMENT,
    club_name TEXT NOT NULL,
    league_id INTEGER NOT NULL,
    UNIQUE(club_name, league_id),
    FOREIGN KEY(league_id) REFERENCES league(league_id)
)