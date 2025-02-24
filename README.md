
# Docker (broken, need to fix)
* `docker build -t groceries .`
* `docker run -p 57457:57457 groceries`

# Native
* Both means of running the app natively require you to install
  * nodejs and npm
  * go
  * sqlite3

## Built
* run `sh run.sh`

## Dev
* `cd ./spa && yarn start`
* `cd ./api/db && sh generate.sh`
* `cd ./api && go run main.go`
