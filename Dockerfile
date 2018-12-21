# build stage
FROM golang AS build-env
ADD . /src
RUN cd /src && GOOS=linux go build -o vashisht-api-linux

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/vashisht-api-linux /app/
ENTRYPOINT ./vashisht-api-linux