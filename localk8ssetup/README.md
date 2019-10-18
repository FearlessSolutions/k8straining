### Local Cluster Options:

1) [docker desktop kubernetes](https://www.docker.com/products/kubernetes) (this is the method supported in this repo)
2) [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/)
3) [k3s](https://k3s.io/)

## Setup:
1) follow the steps for whichever you choose for installing and starting your cluster
2) ensure you have access to the kubectl context with `kubectl config get-contexts` 
3) ensure nodes are running with `kubectl get nodes`
4) [follow instructions here](https://kubernetes.github.io/ingress-nginx/deploy/_) to deploy ingress controller to your cluster

## Project Installation:
1) make sure you've build the containers `docker-compose build`
2) `kubectl apply -f appk8s.yml`
