# Software Engineer / DevOps Challenge

An implementation of the challenge.

**Data all run in memory and no persistence has been implemented.**

## Collector (./collector)

The collector subscribes to mosquitto and fires a function callback on every received message. The Value of this is then appended to an array. The aggregator runs every 5 seconds and saves this into the storage.

The web service can be accessed via port 3000 under /average. If using docker-compose.yml as go to http://localhost:3000/average.

## RandGen (./randgen)

The random generator fires every 100 milliseconds and publishes the value to mosquitto. This in turn then get's received by subscribers such as the collector above.
Values being generated are integers of the value 1 to 10000 (inclusive).

## MQTT (./mqtt)

A library wrapper to ease the use of Paho MQTT for this particular project.

## Development

Just run `./build-n-run.sh` to see the application up and running. Go to http://localhost:3000/average to see the aggregated average value.

Scale up and down by running `docker-compose scale rand-gen=3`

## Deployment

Deployment will be done via Ansible. Questions regarding the machines, services and the infrastructure arise. These have to be answered before proceeding makes sense.
Ansible shines here as it does not require any agents to be installed, rather more just the SSH keys and SSH to run commands remotely.

## Monitoring

Monitoring should be done via Prometheus. Metrices and application health can easily be exported to Prometheus. Grafana suits well as a frontend.

## Todos
  - [ ] ~~Implement broker~~
  - [x] Random generator
  - [x] Collector with aggregator
  - [x] Docker-Compose file
  - [ ] Deployment via Ansible. Knowledge of infrastructure required. Place for storage of artefacts.
  - [ ] Clean up code
  - [ ] Add configuration via env and .env
  - [ ] Extend to expose metrics to Prometheus
