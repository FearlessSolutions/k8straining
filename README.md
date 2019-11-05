<!-- Add project logo here -->
<p align="center"><img src="https://fearless.tech/static/brand/logo_lockup.png"></p>

<h1 align="center">Kubernetes Training</h1>

<h2 align="center">Testbed</h2>

<!-- The badges below are included as examples. Include any relevant badges as desired (e.g., PR badge, license badge, open source badge, status badge) -->

<p align="center">
  <a href="http://makeapullrequest.com" target="blank"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square"></a>
  <a href="https://github.com/ellerbrock/open-source-badges/" target="blank"><img src="https://badges.frapsoft.com/os/v1/open-source.svg?v=103"></a>
 </p>

<!-- Be sure to update TOC to match final headings -->
<p align="center">
  <a href=#background>Background</a> • 
  <a href=#usage>Usage</a> • 
  <a href=#docker>Docker</a> • 
  <a href=#kubernetes>Kubernetes</a> • 
  <a href=#deploy-to-gke>Deploy to GKE</a> •  
  <a href=#network-diagram>Network Diagram</a> • 
  <a href=#contributors>Contributors</a>
</p>

## Background
This testbed was created as an "Idiot's Guide" to learning Kubernetes. Kubernetes training can be unnecessarily complicated when it involves building an application, trying to force it into a microservices architecture, and then trying to put features on top of that architecture. 

Instead, this testbed allows for the creation of multiple services that talk to each other and only perform simple functions. This testbed can therefore be used to test Kubernetes assumptions and systems without too much added complexity. 

This testbed can also be used for new systems (e.g., Helm, service mesh, virtualkubelt). If you want to test any of these additional complexities, please create a fork of this repo.

This testbed includes the following four microservices:
1. **Add:** Takes 2 integers, adds them together, and returns the result.
2. **Multiply:** Takes 2 integers, checks which is smaller, calls Add enough times with the larger integer to get the correct result, decides positive or negative, and returns the result.
3. **Nest1:** Gets Nest2, then returns "Nest2 responded: [Nest2's response]."
4. **Nest2:** Returns the contents of an environment variable.

## Usage

### Docker

`docker-compose up --build`

```bash
curl -X POST -G 'http://localhost:8080/add' -d a=3 -d b=36
curl -X POST -G 'http://localhost:8080/multiply' -d a=3 -d b=36
curl localhost:8080/nest1
```

### Kubernetes
For local kubernetes setup, see instructions in `/localk8ssetup`

`kubectl apply -f k8s.yml`

```bash
curl -X POST -G 'http://localhost/add' -d a=3 -d b=36 -H"Content-Length:0"
curl -X POST -G 'http://localhost/multiply' -d a=3 -d b=36 -H"Content-Length:0"
curl http://localhost/nest1
```

#### Deploy to GKE

1. Follow the instructions [here](https://cloud.google.com/kubernetes-engine/docs/quickstart) for configuring your shell and deploying / authenticating a cluster.
2. Ensure you are using the correct [context](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/) to address the remote cluster (I use [kubectx](https://github.com/ahmetb/kubectx) and [kube_ps1](https://github.com/jonmosco/kube-ps1) to make this easy).
3. Edit the `k8s.yml` file: change the `image:` entries to use the remote registries.
4. Deploy the YAML as follows:
```bash
kubectl apply -f k8s.yml
```

### Network Diagram
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

## Contributors

<a href="https://github.com/mkantzerFearless" target="blank"><img src="https://avatars0.githubusercontent.com/u/39307507?s=60&v=4" width="100px;"/></a>
