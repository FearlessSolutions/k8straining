## TESTBED

This project is a teastbed for later kubernetes work. instead of building out something annoying as hell, and then trying to force that into a microservice architecture, and then from there trying to put cool stuff on top, I figured "let's build something dumb and easy".

This project implements integer math. It adds integers at /add, it multiplies intigers at /multiply (by making multiple calls to `/add`). It subtracts, it divides. 

It may even, if we work hard enough, do powers. And maybe even have a (really shitty) ui. 

The goal is just to have a lot of services that talk to eachother, and don't actually do anything too complicated. It barely even matters if it's right. 

It's just a testbed to try out some kubernetes crap with. 


## Usage

`docker-compose up --build`

`curl -X POST -G 'http://localhost:9696/add' -d a=3 -d b=36`
`curl -X POST -G 'http://localhost:9696/multiply' -d a=3 -d b=36`