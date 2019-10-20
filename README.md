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
curl localhost:8080/nest1
```


## Kubernetes:
for local kubernetes setup, see instructions in `/localk8ssetup`

`kubectl apply -f k8s.yml`

```bash
curl -X POST -G 'http://localhost/add' -d a=3 -d b=36 -H"Content-Length:0"
curl -X POST -G 'http://localhost/multiply' -d a=3 -d b=36 -H"Content-Length:0"
```

### Deploy to GKE

1) Follow the instructions [here](https://cloud.google.com/kubernetes-engine/docs/quickstart) for configuring your shell and deploying / authenticating a cluster
2) Ensure you are using the correct [context](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/) to address the remote cluster (I use [kubectx](https://github.com/ahmetb/kubectx) and [kube_ps1](https://github.com/jonmosco/kube-ps1) to make this easy)
3) Edit the `k8s.yml` file: change the `image:` entries to use the remote registries, and enable `type: NodePort` on the services
4) Deploy the YAML: 
```bash
kubectl apply -f k8s.yml
```