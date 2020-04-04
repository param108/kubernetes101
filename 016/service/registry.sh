docker run -d --restart=always -p "5000:5000" --name "kind-registry" registry:2
# print the ip address of the registry
docker inspect kind-registry -f '{{.NetworkSettings.IPAddress}}'
