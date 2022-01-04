# timescale-filla
## all filla no killa

A small package/script to fill a timescaledb database with random data.
An alternative to [TSBS](https://github.com/timescale/tsbs).

Docker container to run timescale-filla against:

```docker run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb:latest-pg14```