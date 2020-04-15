[gomicro]: https://github.com/micro/go-micro
[nats]: https://nats.io/
[docker]: https://www.docker.com/
[dockercompose]: https://docs.docker.com/compose/

# Go Micro Boilerplate

This is the simple, clear, and useful architecture/structure of the project based on 
GoLang programming language and micro-services architecture.

[Go Micro][gomicro] is used as a framework that allows you to setup and manage your 
micro-services really easily. Read more about this framework on their official GitHub page.

We use [NATS][nats] as a discovery service, message broker service, and transport channel.

All services are [dockerized][docker] and can be adjusted by [docker-compose][dockercompose] configuration.
Check Dockerfiles and docker-compose.yaml file in the root of the project.

**This is the structure of the boilerplate:**

**`./proto`** contains protocol buffers models of each service where it exists.
Your services use protocol buffer messages as the main communication objects.


**`./scripts`** contains bash scripts. You may see `generate.sh` scripts there. 
It's needed to generate GoLang models of proto messages.
To generate proto files, execute the following command from the root of the project 
```bash
$ ./scripts/generate.sh .
```
Also, to build the base image for services, you may use `build-base-image.sh` script:
```bash
$ ./scripts/build-base-image.sh
```

    
**`./services`** contains the logic of all of the services.
Once you wand to create a new service, you may create a new directory with `<service-name>-svc` format.
Inside of the created directory you can put the logic of your service like it's done with `account-svc`.

    
**`./utils`** contains common packages used across the services.


**`./vendor`** (not exists yet) contains dependencies. 
Execute the following from the root of the project to setup `vendor`:
```bash
$ export GO111MODULE=on
$ go mod download
$ go mod vendor
```
    

**`./docker-compose.yaml`** contains configuration of services to run inside docker containers.
If you want to add a new service there, you can use the way how it's done with `account-svc`.
Useful commands:
- `docker-compose up -d --build` - restart all services.
- `docker-compose up -d --build <service-name>` - restart a certain service.
- `docker-compose down` - destroy all services. All data will be removed.
- `docker-compose stop` - stop all services.
    

**`./go.mod`** and **`./go.sum`** needed to manage dependencies using [go modules][gomodules].


**`./.env`** contains environment variables that are passing to the services 
when we run them using docker-compose (see `docker-compose.yaml`).


**`./README.md`** recommended to describe your project.


