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

## Note

Currently only GET and POST calls will work as we couldn't find a way to have query params using the net/http package. Hence we will try using gorilla tools in another branch.

### Please check `"using-gorilla-library" branch` for more a complete REST API.

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
