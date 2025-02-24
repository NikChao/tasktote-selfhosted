(cd ./spa && npm i && npm run build)
(cd ./api/db && sh generate.sh)
(cd ./api && go build -o dist/bootstrap . && ./dist/bootstrap)
