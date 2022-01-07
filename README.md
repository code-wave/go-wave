# go-wave
> Web project for recruting programming study and team project members



## Envirionment

| Name | Version | Link |
|:-:|:-:|:-:|
| Docker | 20.10.7 | <https://docs.docker.com/engine/install/> |
| Docker-compose | 1.27.0 | <https://docs.docker.com/compose/install/> |
| Compose file  | 3.8 | <https://docs.docker.com/compose/compose-file/compose-file-v3/> |
| golang | 1.15.7 | <https://golang.org/dl> |


---

## How to Use

```bash
git clone https://github.com/atg0831/go-wave.git
cd go-wave
```
### Run Containers
- **Local Mode**

```bash
# run dev mode using shell script
./runserver.sh --local
```
```bash
# Or you can just use docker-compose command
docker-compose -f docker-compose.yml -f docker-compose.local.yml up -d --build
```

- **Develop Mode**
```bash
# run dev mode using shell script
./runserver.sh --dev 
```
```bash
# Or you can just use docker-compose command
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build
```

### Down Containers
```bash
./downserver.sh 
```
```bash
# Or you can just use docker-compose command
docker-compose down
```

## Each services role
- **proxy** 
   - A container for the reverse proxy role. When connecting to the Nginx server, the static files built by the react(frontend) server are displayed. And if you access the `/api` path, it will route to the api server.
- **api**
   - A container that receives requests from clients and sends responses back to clients.
- **postgres**
   - PostgreSQL Database container
- **pgadmin**
   - A container to easily manage PostgreSQL database with GUI
- **redis**

   - A container that is in-memory key-value data structure store and support pub/sub for chat

---

## Deploy Ports
|Container Name |   # Port   |
|:-------------:|:----------:|
|   proxy       |    8081    |
|   api         |    58080   |
|   postgres    |    54320   |
|   pgadmin     |    54330   |
|   redis       |    56379   |


---
## Contributors' info
Taegeon An - <https://github.com/atg0831>  
Jungmin Kim - <https://github.com/PudgeKim>
