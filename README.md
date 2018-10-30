creating digital ocean server with docker-machines
```
apt update
apt-get update

snap install docker
apt-get install virtualbox
docker-machine create --driver virtualbox myvm1
```
Not the correct way to run docker-machines on digital ocean, missing vboxmanage

Create docker-machines using the digital ocean api --- following stevens guide
```
wget https://github.com/docker/machine/releases/download/v0.15.0/docker-machine-$(uname -s)-$(uname -m)
 
mv docker-machine-Linux-x86_64 docker-machine
chmod +x docker-machine 
mv docker-machine /usr/local/bin
docker-machine create --driver digitalocean --digitalocean-access-token=<token> <name>

docker swarm init --advertise-addr <IP>
```
token= aefc74993c1843b7ab2761ecbb2ddb41f4cb20e9880ad17ab93eb40759389d50
could not determine meaning of error: switched to containers

deploying goChat in a container
```
docker container run golang go get -v github.com/ClayGale/331goChat/...
docker container commit $(docker container ls -lq) test1
docker container run test1 331goChat
```
panic: open ui/templates/layout.html: no such file or directory

goroutine 1 [running]:
html/template.Must(0x0, 0x7fe6f4f8f078, 0xc20803b530, 0x1)
    /usr/lib/golang/src/html/template/template.go:304 +0x50



