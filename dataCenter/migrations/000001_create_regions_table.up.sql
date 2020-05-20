CREATE TABLE IF NOT EXISTS regions (  
    isocode VARCHAR(15) UNIQUE,
    region VARCHAR(200) UNIQUE,
    infected INTEGER,
    recovered INTEGER,
    deceased INTEGER,
    country VARCHAR(100)
)
