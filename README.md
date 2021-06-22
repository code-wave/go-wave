# go-wave
---

## How to Use

```bash
git clone https://github.com/atg0831/booking_study.git
cd booking_study
```

If you want run for development, try this command.  

1. run all containers
```bash
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build    
```
2. run specific containers
```bash
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build ${service_name}
```

---
