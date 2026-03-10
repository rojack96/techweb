CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS postgis_topology;

CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE users.accounts (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    language CHAR(2) DEFAULT 'en',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE users.profiles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--

CREATE SCHEMA sightings;

CREATE TABLE sightings.animals (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,  -- es: 'gatto', 'cane'
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE sightings.breeds (
    id BIGSERIAL PRIMARY KEY,
    animal_id BIGINT NOT NULL REFERENCES sightings.animals(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,       -- es: 'Persiano', 'Labrador'
    UNIQUE (animal_id, name),         -- stesso nome razza non ripetuto per lo stesso animale
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE sightings.sightings (
    id BIGSERIAL PRIMARY KEY,
    animal_id BIGINT NOT NULL REFERENCES sightings.animals(id),
    breed_id BIGINT REFERENCES sightings.breeds(id),  -- opzionale se razza sconosciuta
    latitude NUMERIC(9,6) NOT NULL,   -- precisione fino a 0.11 m circa
    longitude NUMERIC(9,6) NOT NULL,
    spotted_at TIMESTAMP DEFAULT NOW(), -- quando è stato avvistato
    notes TEXT,                        -- eventuali note dell'avvistatore
    created_at TIMESTAMP DEFAULT NOW()
);