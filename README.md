# go-wave
Web project for recruting programming study and team project members
---

## Envirionment & Prerequisites

| Name | Version | Link |
|:-:|:-:|:-:|
| Docker | 20.10.7 | <https://docs.docker.com/engine/install/> |
| Docker-compose | 1.27.0 | <https://docs.docker.com/compose/install/> |
| Compose file version | 3.8 | <https://docs.docker.com/compose/compose-file/compose-file-v3/> |


## How to Use

```bash
git clone https://github.com/atg0831/go-wave.git
cd go-wave
```
### Develop mode
If you want to run for develop mode, try this command
1. run containers(you can use script or docker-compose command)
```bash
./runserver.sh --dev 
```

or

```bash
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build
```

2. down(stop and remove)containers
```bash
./downserver.sh 
```

### Production mode
If you want to run for prodcution mode, try this command
1. run containers(you can use script or docker-compose command)
```bash
./runserver.sh --prod 
```

or

```bash
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build
```

2. down(stop and remove)containers
```bash
./downserver.sh 
```

## Each services role
1. proxy 
   - A container for the reverse proxy role. When connecting to the Nginx server, the static files built by the react(frontend) server are displayed. And if you access the /api path, it will route to the api server.
2. api
   - A container that receives requests from clients and sends responses back to clients.
3. postgres
   - PostgreSQL Database container
4. pgadmin
   - A container to easily manage PostgreSQL database with GUI
5. redis
   - A container that is in-memory key-value data structure store and support pub/sub for chat


## Deploy Ports

| Name     | Port # |
|:--------:|:------:|
| proxy    | 8081   |
| api      | 58080  |
| postgres | 54320  |
| pgadmin  | 54330  |
| redis    | 56379  |


## Contributors' info
  
Taegeon An - <https://github.com/atg0831>
Jungmin Kim - <https://github.com/PudgeKim>