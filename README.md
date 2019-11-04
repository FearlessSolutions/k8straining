## TESTBED

This project is a teastbed for kubernetes training. instead of building out something annoying as hell, and then trying to force that into a microservice architecture, and then from there trying to put cool stuff on top, I figured "let's build something dumb and easy".


This project consists of 4 microservices: add, multiply, nest1, and nest2
- Add: takes 2 integers, adds them together, and then returns the result
- Multiply: takes 2 integers, checks which is smaller, calls Add enough times with the larger to get the correct result, decides positive or negative, and returns the result
- Nest1: Gets Nest2, then returns "Nest2 responded: [nest2's response]"
- Nest2: Returns the contents of an environment variable

The goal is just to have a lot of services that talk to eachother, and don't actually do anything too complicated. It barely even matters if it's right. This can then be used for testing kubernetes assumtions and systems, as well as a testbed for new systems (Helm, service mesh, virtualkubelt). 

If you want to use this to test any of those additional complexities, please create a fork of this repo. 


# Usage

## Docker

`docker-compose up --build`

```bash
curl -X POST -G 'http://localhost:8080/add' -d a=3 -d b=36
curl -X POST -G 'http://localhost:8080/multiply' -d a=3 -d b=36
curl localhost:8080/nest1
```


## Kubernetes:
for local kubernetes setup, see instructions in `/localk8ssetup`

`kubectl apply -f k8s.yml`

```bash
curl -X POST -G 'http://localhost/add' -d a=3 -d b=36 -H"Content-Length:0"
curl -X POST -G 'http://localhost/multiply' -d a=3 -d b=36 -H"Content-Length:0"
curl http://localhost/nest1
```

### Deploy to GKE

1) Follow the instructions [here](https://cloud.google.com/kubernetes-engine/docs/quickstart) for configuring your shell and deploying / authenticating a cluster
2) Ensure you are using the correct [context](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/) to address the remote cluster (I use [kubectx](https://github.com/ahmetb/kubectx) and [kube_ps1](https://github.com/jonmosco/kube-ps1) to make this easy)
3) Edit the `k8s.yml` file: change the `image:` entries to use the remote registries.
4) Deploy the YAML: 
```bash
kubectl apply -f k8s.yml
```

## Network Diagram:
```

                                 +----------------+
                                 |                |
                          +------>    Frontend    |
                          |      |                |
                          |      +----------------+
                          |
                          |      +----------------+
                          |      |                |
                          +----->+     Adder      |
                          |      |                |
                          |      +----------------+
                          |              ^
                          |              |
                          |              |
       +------------+     |      +----------------+
       |            |     |      |                |
       |   Client   +----------->+   Multiplier   |
       |            |     |      |                |
       +------------+     |      +----------------+
                          |
                          |      +----------------+      +----------------+
                          |      |                |      |                |
                          +----->+     Nest1      +----->+     Nest2      |
                                 |                |      |                |
                                 +----------------+      +----------------+
```