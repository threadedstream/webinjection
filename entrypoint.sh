cd migrations/
goose postgres "host=postgres-db port=5432 user=postgres dbname=postgres sslmode=disable" up
cd ..
./main