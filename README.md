# Go Micro Boilerplate

**This is the structure of the boilerplate:**

**`./proto`** contains protocol buffers models of each service where it exists.
Your services use protocol buffer messages as the main communication objects.


**`./scripts`** contains bash scripts. You may see `generate.sh` scripts there. 
It's needed to generate GoLang models of proto messages.
To generate proto files, execute the following command from the root of the project 
```bash
$ ./scripts/generate.sh .
```

    
**`./services`** contains the logic of all of the services.

    
**`./utils`** contains common packages used across the services.


**`./vendor`** (not exists yet) contains dependencies. 
Execute the following from the root of the project to setup `vendor`:
```bash
$ export GO111MODULE=on
$ go mod download
$ go mod vendor
```
    

**`./docker-compose.yaml`** contains configuration of services to run inside docker containers.
Useful commands:
- `docker-compose up -d --build` - restart all services.
- `docker-compose up -d --build <service-name>` - restart a certain service.
- `docker-compose down` - destroy all services. All data will be removed.
- `docker-compose stop` - stop all services.
    

**`./go.mod`** and **`./go.sum`** needed to manage dependencies using [go modules][gomodules].


**`./.env`** contains environment variables that are passing to the services 
when we run them using docker-compose (see `docker-compose.yaml`).


**`./README.md`** recommended to describe your project.


