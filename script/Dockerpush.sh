echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push goboilerplates/micro-rpc
docker push goboilerplates/micro-rpc:0.0.0