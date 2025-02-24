
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

# Demo
![Kapture 2025-02-24 at 17 46 28](https://github.com/user-attachments/assets/3b6c510e-d9c9-4c0c-aae2-0b18cf9e31b7)
