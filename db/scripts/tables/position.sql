CREATE TABLE IF NOT EXISTS position (
    position_id INTEGER PRIMARY KEY AUTOINCREMENT,
    position_short_name TEXT NOT NULL UNIQUE,
    position_long_name TEXT NOT NULL UNIQUE
)