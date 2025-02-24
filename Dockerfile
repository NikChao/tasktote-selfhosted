FROM node:18-alpine as spa-builder

# Compile SPA
WORKDIR /app
COPY ./spa ./spa
RUN cd ./spa && npm i && npm run build

# Compile backend
FROM golang:1.21-alpine as api-builder
WORKDIR /app
RUN apk add --no-cache sqlite sqlite-dev
COPY ./api ./api
COPY --from=spa-builder /app/spa/dist ./spa/dist
RUN cd ./api/db && sh generate.sh
RUN cd ./api && sh build.sh

# Run webserver
FROM alpine:3.18
WORKDIR /app
RUN apk add --no-cache sqlite sqlite-dev
COPY --from=api-builder /app/api/dist/bootstrap ./api/dist/bootstrap
COPY --from=api-builder /app/api/db/groceries.db ./api/db/groceries.db
COPY --from=spa-builder /app/spa/dist ./spa/dist
RUN chmod +x ./api/dist/bootstrap
EXPOSE 57457


WORKDIR /app/api
CMD ["./dist/bootstrap"]
