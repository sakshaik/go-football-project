CREATE TABLE IF NOT EXISTS league (
    league_id INTEGER PRIMARY KEY AUTOINCREMENT,
    league_name TEXT NOT NULL,
    country_id INTEGER NOT NULL,
    confederation_id INTEGER NOT NULL,
    UNIQUE(league_name, country_id),
    FOREIGN KEY(country_id) REFERENCES country(country_id),
    FOREIGN KEY(confederation_id) REFERENCES confederation(confederation_id)
)