# Goals

- Create a Simple Web server:- Cartoons API
- Do all CRUD operations, without DB only using In-Memory storage
- Dockerrize to run anywhere

## How to run the Cartoons API server Using Docker

- Build the image

```
docker build -t cartoon-api .
```

- Run the container

```
docker run -p 8080:8080 cartoon-api:latest
```

- The api endpoint is `http://localhost:8080/api/v1/cartoons`

## API Docs

### GET request

`http://localhost:8080/api/v1/cartoons`

### POST request

`http://localhost:8080/api/v1/cartoons`

id field is not currently unique

Body

```
{
		"id": 6,
    "title": "X Men",
    "ratings": 8
}
```

### PUT request

`http://localhost:8080/api/v1/cartoons/${id}`

Body

```
{
		"id": 5,
    "title": "Tintin",
    "ratings": 8
}
```

### Delete request

`http://localhost:8080/api/v1/cartoons/${id}`
