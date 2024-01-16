CREATE TABLE IF NOT EXISTS confederation (
    confederation_id INTEGER PRIMARY KEY AUTOINCREMENT,
    confederation_name TEXT NOT NULL UNIQUE,
    confederation_description TEXT NOT NULL,
    continent_id INTEGER NOT NULL,
    FOREIGN KEY(continent_id) REFERENCES continent(continent_id)
)