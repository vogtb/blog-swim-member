## To bring up the vagrant boxes
```
IP=192.168.50.6 NAME=node01 vagrant up
```

## To build main.go
```
GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -i -ldflags "-X main.build=master" -o ./blog-swim-member.linux.amd64
```
