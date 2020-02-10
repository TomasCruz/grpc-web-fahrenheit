# grpc-web-fahrenheit
delegates request to grpc-server-fahrenheit

## API
### /health/
displays service health status

#### Return codes:
- OK (200) is returned if service DB is working properly
- FailedDependency (424) is returned otherwise
- InternalServerError (500) is returned in case of general errors.

### /c2f/[:cels]
returnes a body containing celsius/fahrenheit pair

#### input
Path parameter with number of degrees celsius

#### Return codes:
- OK (200) is returned when ran successfully, with the body containing celsius/fahrenheit pair.
- NotAcceptable (406) if path parameter can't be parsed to a float64
- FailedDependency (424) if gRPC dependancy returned error
- InternalServerError (500) is returned in case of general errors.

### /f2c/[:fahr]
returnes a body containing celsius/fahrenheit pair

#### input
Path parameter with number of degrees fahrenheit

#### Return codes:
- OK (200) is returned when ran successfully, with the body containing celsius/fahrenheit pair.
- NotAcceptable (406) if path parameter can't be parsed to a float64
- FailedDependency (424) if gRPC dependancy returned error
- InternalServerError (500) is returned in case of general errors.

## Build
### Prerequisites:
- standard Docker installation
- standard gRPC installation

### building
- run 'source ./env'
- run 'make'

## Run
- run grpc-server-fahrenheit
- 'make run' from terminal

## Unit tests
'make test' from terminal for unit tests

## Integration tests
If not ran already, run steps for build and run.
In another terminal, navigate to directory containing Makefile, then 'source ./env', then 'make integration'
