CREATE TABLE IF NOT EXISTS club (
    club_id INTEGER PRIMARY KEY AUTOINCREMENT,
    club_name TEXT NOT NULL,
    league_id INTEGER NOT NULL,
    FOREIGN KEY(league_id) REFERENCES league(league_id)
)