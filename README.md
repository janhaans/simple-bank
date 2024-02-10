# simple-bank

## SQL script to create tables
Postgresql SQL statements to create `simple bank` tables exported
from dbdiagram.io: [simple bank](https://dbdiagram.io/d/simple-bank-65c391b8ac844320aea89c27)
to  `db/sql/simple-bank.sql`

## golang-migrate
Github Repo: [golang-migrate](https://github.com/golang-migrate/migrate)

**Installation**
`> brew install golang-migrate`

**Create migration with name `init_schema`**
```
> migrate create -ext sql -dir db/migration -seq init_schema

-ext sql                : sql extension
-dir db/migration       : store the migration files in directory db/migration
-seq                    : generate sequential version number for the migration files

```
**Edit db/migration/000001_init_schema.up.sql**
Copy and paste SQL script in `db/sql/simple-bank.sql` to `db/migration/000001_init_schema.up.sql`

**Edit db/migration/000001_init_schema.down.sql**
```
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS accounts;
```

## Run Postgress database
```
> make postgres
```
## Install and connect to Postgress database with Table-Plus
```
> brew install --cask tableplus
```

## Create database simple_bank
```
> make createdb
```

## Drop database simple_bank
```
> make dropdb
```

## Migrate Up
```
> make migrateup
```
Refresh Table-Plus `Cmd + R`

## Migrate down
```
> make migratedown
```
Refresh Table-Plus `Cmd + R`
