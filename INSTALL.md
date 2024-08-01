go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/cosmtrek/air@latest
go get github.com/lib/pq
go get github.com/spf13/viper
go get -u github.com/gin-gonic/gin

curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate

------------------
docker exec -it postgres bash
psql -U root contact_db
select * from pg_available_extensions;
CREATE EXTENSION IF NOT EXISTS  "uuid-ossp";

```sh
migrate -path db/migration -database  "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose up

migrate -path db/migration -database "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose down
```

sqlc init
sqlc generate