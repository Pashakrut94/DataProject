CREATE TABLE IF NOT EXISTS regions (
    isocode VARCHAR(10) UNIQUE,
    region VARCHAR(30) UNIQUE,
    infected INTEGER,
    recovered INTEGER,
    deceased INTEGER,
    country VARCHAR(30)
)