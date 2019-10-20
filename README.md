## TESTBED

This project is a teastbed for kubernetes training. instead of building out something annoying as hell, and then trying to force that into a microservice architecture, and then from there trying to put cool stuff on top, I figured "let's build something dumb and easy".

This project implements integer math. It adds integers at /add, it multiplies intigers at /multiply (by making multiple calls to `/add`). It subtracts, it divides. 

It may even, if we work hard enough, do powers. And maybe even have a (really shitty) ui. 

The goal is just to have a lot of services that talk to eachother, and don't actually do anything too complicated. It barely even matters if it's right. 

It's just a testbed to try out some kubernetes crap with. 




# Usage

## Docker

`docker-compose up --build`

```bash
curl -X POST -G 'http://localhost:8080/add' -d a=3 -d b=36
curl -X POST -G 'http://localhost:8080/multiply' -d a=3 -d b=36
```


## Kubernetes:
for local kubernetes setup, see instructions in `/localk8ssetup`

`kubectl apply -f appk8s.yml`

```bash
curl -X POST -G 'http://localhost/add' -d a=3 -d b=36 -H"Content-Length:0"
curl -X POST -G 'http://localhost/multiply' -d a=3 -d b=36 -H"Content-Length:0"
```
