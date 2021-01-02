from golang:latest
copy . /app
WORKDIR /app
run go build
expose 7010
cmd ["/app/ruz"]