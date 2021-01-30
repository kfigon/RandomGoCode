# DB
run:
```
docker run --rm --name pg-docker -e POSTGRES_PASSWORD=mypass -e POSTGRES_USER=myuser -d -p 5432:5432 postgres
```

connect in shell:
```
docker exec -it pg-docker bash
psql -U myuser

create database mydb
\c mydb

CREATE TABLE person(
    id SERIAL PRIMARY KEY,
    creation_date DATE NOT NULL,
    name VARCHAR(10) NOT NULL
);

INSERT INTO person (creation_date, name ) VALUES ('2021-01-30', 'Foo');
INSERT INTO person (creation_date, name ) VALUES ('2021-01-29', 'Bar');
```

connect in intelliJ:
```
jdbc:postgresql://localhost:5432/postgres
```


# GO
```
go mod init hello
go get github.com/jackc/pgx/v4
go get github.com/lib/pq
```
connection string:
` postgresql://localhost:5432/mydb?user=myuser&password=mypass`
