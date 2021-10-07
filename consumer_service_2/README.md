
## External packages used:

* [viper](github.com/spf13/viper)
* [gorilla/websocket](github.com/gorilla/websocket)


## Running the application

Easiest ways are either with `make run` or with docker `docker-compose up -d`

## Run tests

Some unit tests need RabbitMQ running in the background. You can easily start a container in Docker with the following command:

```
docker run -d --hostname my-rabbit -p 5672:5672 -p 15672:15672 --name some-rabbit rabbitmq:management
```

Running locally is as easy as `make test`.