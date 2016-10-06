# User Server

User server is a microservice to store user accounts and perform simple authentication.

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Start a mysql database

4. Download and start the service

	```shell
	go get github.com/suicidegang/user-srv
	user-srv --database_url="root:root@tcp(192.168.99.100:3306)/user"
	```

## The API
User server implements the following RPC Methods

Account
- Create


### Account.Create
```shell
$ micro query sg.micro.srv.user Account.Create '{"user":{"username": "asim", "email": "asim@example.com"}, "password": "password1"}'
{}
```
