# List all databases in the Postgres container
docker exec -it crypto-postgres-v2 psql -U user -d postgres -c "\l"
# Create the new database for v2
docker exec -it crypto-postgres-v2 createdb -U user crypto_exchange_v2

