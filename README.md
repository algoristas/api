# Algoristas API
This repo contains the API for the [Algoristas' dashboard](http://letmethink.mx:3333/).

## Running app on docker
Run the following command for building the image:
```
docker build -t algoristas-api .
```

Followed by:
```
docker run -p 8080:8080 algoristas-api
```

This last command will run the application inside the container and will expose the port 8080. You can test the API by querying localhost at 8080.

```
curl localhost:8080/v1/results
curl -i localhost:8080/v1/problems/sets
curl -i localhost:8080/v1/standings
```
