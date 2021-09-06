# Cart Service
<!--- [![CircleCI](https://circleci.com/gh/Ralphbaer/hash.svg?style=svg&circle-token=37cdd1fd1e89719b206adf5ae503bc83073b9c3a)](https://circleci.com/gh/Ralphbaer/hash/?branch=master) -->

The badge above shows the status of the last pipeline. If everything goes well, the badge will be "green" indicating success or "red" (failed) indicating some error in the build.

This repo contains the source code of the Cart service.

## Architecture

<!--- ![alt text](./hexagonal-macro.png "Title") -->

## Requirements

| Name | Version | Notes | Mandatory
|------|---------|---------|---------|
| [golang](https://golang.org/dl/) | >= go1.15.14 | Main programming language | true
| [docker](https://www.docker.com/) | n/a | Used to handle core mock images and start local service | true
| [aws-cli](https://aws.amazon.com/pt/cli/) | v2 | Used to create all AWS Enviroment (Just in case you want to know) | false
| [sh/bash] | depending on OS. Anyway, you should be able do execute any .sh file | Used to lint checks, test processes and some console interface customizations | true
| [make](https://www.gnu.org/software/make/) | depending on OS. Anyway, you should be able do execute make commands to run the project, tests and localenvironment | n/a | true

## Providers

| Name | Version | Notes
|------|---------|---------|
| [aws](https://aws.amazon.com/pt/) | n/a | All the infraestructure are on AWS

# Usage

### Start Local
Inside /cart, you can run:
```bash
make run           # Start service on port 3000 without test data (empty database)
make run-test      # Start service on port 3000 with test data (json](files/pdvs.json))
```

# Testing

```bash
make test                 # Run all unit tests
```
## Documentation

Visit [this link](http://hash-microservices-elb-6682139.us-east-2.elb.amazonaws.com/cart/docs#overview) for API documentation. If you want to access the docs locally, just change the host in the url to localhost:3000. Something like: http://localhost:3000/cart/docs

# Deploy
The service uses a Circle pipeline that is triggered when a code is pushed to master branch.
The master has no PR approval criteria to facilitate the tests for the Hash team.
For testing deploy purposes, you guys can change anything in the code or just change the signal.id file inside /cart, commit and push the change to master and wait the circle-ci to deploy the code. :)

All provisioning on the AWS Cloud has already been done and its running.

CircleCi is also working. Maybe you could create an account to access the project there through this link https://app.circleci.com/pipelines/github/Ralphbaer/hash?invite=true.

If you guys cant get the access, no problem, I can talk calmly and give all the details of both AWS provisioning (ec2, load balancer, system manager, secrets, etc.) and the circleCI pipeline.