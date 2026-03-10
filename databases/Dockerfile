FROM postgres:16-bookworm

RUN apt-get update && apt-get install -y \
    postgis \
    postgresql-16-postgis-3 \
    postgresql-16-postgis-3-scripts \
    gdal-bin \
    && rm -rf /var/lib/apt/lists/*

COPY init-postgis.sql /docker-entrypoint-initdb.d/01-postgis.sql
COPY import-regions.sh /docker-entrypoint-initdb.d/02-import-regions.sh
COPY data/regions.shp /data/
COPY data/regions.shx /data/
COPY data/regions.dbf /data/
COPY data/regions.prj /data/