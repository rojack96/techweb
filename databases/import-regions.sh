#!/bin/bash
set -e

echo "Importing regions shapefile into schema 'geofence'..."

# 1️⃣ Crea schema geofence se non esiste
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "
CREATE SCHEMA IF NOT EXISTS geofence;
"

# 2️⃣ Import shapefile nella tabella 'regions' dentro schema 'geofence' con SRID corretto (32632)
shp2pgsql -I -s 32632 /data/regions.shp geofence.regions \
  | psql -U "$POSTGRES_USER" -d "$POSTGRES_DB"

echo "Creating geom_wgs84 column (lat/lon) in geofence.regions..."

# 3️⃣ Aggiungi colonna geom_wgs84 in lat/lon
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "
ALTER TABLE geofence.regions
ADD COLUMN geom_wgs84 geometry(MULTIPOLYGON, 4326);
"

# 4️⃣ Trasforma le geometrie da 32632 a 4326
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "
UPDATE geofence.regions
SET geom_wgs84 = ST_Transform(geom, 4326);
"

# 5️⃣ Crea indice GIST sulla nuova colonna
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "
CREATE INDEX idx_regions_geom_wgs84 ON geofence.regions USING GIST(geom_wgs84);
"

echo "Done! You can now query geofence.regions.geom_wgs84 with GPS points (lat/lon)."