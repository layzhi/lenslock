# lenslock
Test go webapp repo


## SQL Migration tool - goose
Installiation:
```
go install github.com/pressly/goose/v3/cmd/goose@v3
```
pressly/goose [https://pressly.github.io/goose/]

Create Migration File:
```
goose create <name> sql
```

Run Migration File:
```
# The \ is used to break a command into several lines in bash

goose postgres \
  "host=<host> port=<port> user=<username for db> password=<password for user> dbname=<db name> sslmode=disable" \
  status

# Run The Migration 

goose postgres \
  "host=<host> port=<port> user=<username for db> password=<password for user> dbname=<db name> sslmode=disable" \
  up  
```

Undo Migration:
```
goose postgres \
  "host=<host> port=<port> user=<username for db> password=<password for user> dbname=<db name> sslmode=disable" \
  down  
```

Typical Local Goose Workflow:
```
1. Rollback migrations with goose down or reset.
goose down

2. Pull changes that other developers have pushed to our team's repo.
#    This assumes that the origin remote branch is configured to something
#    like GitHub where I can pull code other devs have submitted.
git pull origin <main branch> --rebase

3. Run goose fix to rename my migrations with the correct versions.
goose fix

4. Run the migrations to verify they work.
goose up

5. Test everything
go test ./...

6. Commit and merge everything to the main branch my team uses.
#    This might be via a GitHub pull request, or something else.
```