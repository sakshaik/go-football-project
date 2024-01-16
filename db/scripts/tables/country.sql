CREATE TABLE IF NOT EXISTS country (
    country_id INTEGER PRIMARY KEY AUTOINCREMENT,
    country_name TEXT NOT NULL UNIQUE,
    continent_id INTEGER NOT NULL,
    FOREIGN KEY(continent_id) REFERENCES continent(continent_id)
)