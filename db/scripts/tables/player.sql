CREATE TABLE IF NOT EXISTS player (
    player_id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INTEGER NOT NULL,
    city_id INTEGER NOT NULL,
    FOREIGN KEY(city_id) REFERENCES city(city_id)
)