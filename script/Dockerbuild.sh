CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out .
docker build -t goboilerplates/micro-rpc .
docker tag  goboilerplates/micro-rpc goboilerplates/micro-rpc:0.0.0