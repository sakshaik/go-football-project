CREATE TABLE IF NOT EXISTS city (
    city_id INTEGER PRIMARY KEY AUTOINCREMENT,
    city_name TEXT NOT NULL,
    country_id INTEGER NOT NULL,
    UNIQUE(city_name, country_id),
    FOREIGN KEY(country_id) REFERENCES country(country_id)
)