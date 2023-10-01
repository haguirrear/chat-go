migrateup:
	migrate  -path ./db/migrations/ -database ${DATABASE_URL}?sslmode=disable -verbose up
migratedown:
	migrate  -path ./db/migrations/ -database ${DATABASE_URL}?sslmode=disable -verbose up


.PHONY: migrateup migratedown