# General overview



![overview](https://i.imgur.com/5HMsCQx.png)



## Publisher Service

Golang app that publishes events to a queue. Events are generated via messages that arrive by users connected via websocket.

Events published have the following structure:

```
{"event_type":"TECH_TASK_NAME","data":"some message"}
```

and the content-type of `application/json`.

### External packages used:

* [viper](github.com/spf13/viper)
* [gorilla/websocket](github.com/gorilla/websocket)


### Running the application

Depends on a running instance of RabbitMQ. The easiest way is to start a docker container with `docker run -d --hostname my-rabbit -p 5672:5672 -p 15672:15672 --name some-rabbit rabbitmq:management`

Easiest ways are either with `make run` or by building the docker image with `docker build -t publisher_service .` and
running it with `docker run publisher -p 8200:8200`

### Run tests

Some unit tests need RabbitMQ running in the background:

```
docker run -d --hostname my-rabbit -p 5672:5672 -p 15672:15672 --name some-rabbit rabbitmq:management
```

Running locally is as easy as `make test`.




## Consumer Service 

Golang app that subscribe to a queue in order to consume events from it. Users can register their interested in receiving a 
broadcast of each event that is consumed.

Events consumed must have the type "TECH_TASK_NAME" (by default) or another as long as it is configured previously to running
the app. Data inside events must also respect the following structure:

```
{"event_type":"TECH_TASK_NAME","data":"some message"}
```

and events must have the content-type of `application/json`.

### External packages used:

* [viper](github.com/spf13/viper)
* [gorilla/websocket](github.com/gorilla/websocket)


### Running the application

Depends on a running instance of RabbitMQ. The easiest way is to start a docker container with `docker run -d --hostname my-rabbit -p 5672:5672 -p 15672:15672 --name some-rabbit rabbitmq:management`  

Easiest ways are either with `make run` or by building the docker image with `docker build -t consumer_service .` and
running it with `docker run consumer_service -p 8208:8208`

### Run tests

Some unit tests need RabbitMQ running in the background:

```
docker run -d --hostname my-rabbit -p 5672:5672 -p 15672:15672 --name some-rabbit rabbitmq:management
```

Running locally is as easy as `make test`.
