# Goal

- Simple Web server cartoons API CRUD, without DB using In-Memory
- no external libraries
- Dockerrize to run anywhere

## Using Docker

- Build the image

```
docker build -t cartoon-api .
```

- Run the container

```
docker run -p 8080:8080 cartoon-api:latest
```

- The api endpoint is `http://localhost:8080/api/v1/cartoons`
