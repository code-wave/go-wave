# go-wave
> Web project for recruting programming study and team project members
<<<<<<< HEAD
=======
---
>>>>>>> 764382abe57948bc4d13ce5abd2181ad21c064e3

## Envirionment & Prerequisites

| Name | Version | Link |
|:-:|:-:|:-:|
| Docker | 20.10.7 | <https://docs.docker.com/engine/install/> |
| Docker-compose | 1.27.0 | <https://docs.docker.com/compose/install/> |
| Compose file version | 3.8 | <https://docs.docker.com/compose/compose-file/compose-file-v3/> |

---

## How to Use

```bash
git clone https://github.com/atg0831/go-wave.git
cd go-wave
```
<<<<<<< HEAD
### Run Containers
- **Develop Mode**
=======
### Develop mode
- run containers
>>>>>>> 764382abe57948bc4d13ce5abd2181ad21c064e3
```bash
# run dev mode using shell script
./runserver.sh --dev
```
```bash
# Or you can just use docker-compose command
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build
```

<<<<<<< HEAD
- **Production Mode**
=======
- down(stop and remove)containers
```bash
./downserver.sh 
```

### Production mode
- run containers
>>>>>>> 764382abe57948bc4d13ce5abd2181ad21c064e3
```bash
# run prod mode using shell script
./runserver.sh --prod 
```
```bash
# Or you can just use docker-compose command
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build
```

<<<<<<< HEAD
### Down Containers
=======
- down(stop and remove)containers
>>>>>>> 764382abe57948bc4d13ce5abd2181ad21c064e3
```bash
./downserver.sh 
```
```bash
# Or you can just use docker-compose command
docker-compose down
```

## Each services role
<<<<<<< HEAD
- **proxy** 
   - A container for the reverse proxy role. When connecting to the Nginx server, the static files built by the react(frontend) server are displayed. And if you access the /api path, it will route to the api server.
- **api**
   - A container that receives requests from clients and sends responses back to clients.
- **postgres**
   - PostgreSQL Database container
- **pgadmin**
   - A container to easily manage PostgreSQL database with GUI
- **redis**
=======
- proxy 
   - A container for the reverse proxy role. When connecting to the Nginx server, the static files built by the react(frontend) server are displayed. And if you access the /api path, it will route to the api server.
- api
   - A container that receives requests from clients and sends responses back to clients.
- postgres
   - PostgreSQL Database container
- pgadmin
   - A container to easily manage PostgreSQL database with GUI
- redis
>>>>>>> 764382abe57948bc4d13ce5abd2181ad21c064e3
   - A container that is in-memory key-value data structure store and support pub/sub for chat

---

## Deploy Ports

<<<<<<< HEAD
|Container Name |   # Port   |
|:-------------:|:----------:|
|   proxy       |    8081    |
|   api         |    58080   |
|   postgres    |    54320   |
|   pgadmin     |    54330   |
|   redis       |    56379   |
=======
|Container Name |  # Port   |
|:-------------:|:---------:|
|   proxy       |   8081    |
|   api         |   58080   |
|   postgres    |   54320   |
|   pgadmin     |   54330   |
|   redis       |   56379   |
>>>>>>> 764382abe57948bc4d13ce5abd2181ad21c064e3

---

## Contributors' info
  
Taegeon An - <https://github.com/atg0831>  
Jungmin Kim - <https://github.com/PudgeKim>
