CREATE TABLE IF NOT EXISTS photos  (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    -- fichier
    filename TEXT NOT NULL,
    filepath TEXT NOT NULL,

    -- info générale
    date_taken TEXT,
    -- width INTEGER,
    -- height INTEGER,

    -- appareil
    camera_make TEXT,
    camera_model TEXT,
    -- focal_length REAL,
    -- aperture REAL,
    -- exposure_time TEXT,
    -- iso INTEGER,

    -- localisation
    latitude REAL,
    longitude REAL,

    -- album (catégorie perso)
    -- album TEXT,

    -- date d'import
    imported_at TEXT DEFAULT CURRENT_TIMESTAMP
);