## Simple GO API 

<!-- ABOUT THE PROJECT -->

![image](https://user-images.githubusercontent.com/7758970/204043976-feed56b2-eca0-4de0-9b06-1ec42848bfe0.png)

I've deployed this project to my personal VPS, and deployed to `single-node kubernetes`

You can access via http://103.134.154.18:32012


#### Built With

* Go (Mux, JWT)
* Docker
* Kubernetes
* MongoDB


<!-- GETTING STARTED -->
## Getting Started


### Prerequisites

This project is configured to be able to containerized using docker (Dockerfile). And deployed in single-node kubernetes cluster. (Deployment.yaml)

Here is what i mean by single-node cluster, go checkout my story on medium here:

https://reinhardjsilalahi.medium.com/beginners-guide-simple-hello-kubernetes-all-in-one-on-a-single-vps-fcfdfee9edfc

Below is the commands to deploy the changes to the pod

```sh
git pull
```

```sh
docker build -t simple-api .
```

```sh
kubectl rollout restart deployment/simple-api
```


<!-- USAGE EXAMPLES -->
## Usage
You can see the example usages from `swagger.yml ` and open it on https://editor.swagger.io

Our you can import the postman request collection, `simple-api.postman_collection.json`
