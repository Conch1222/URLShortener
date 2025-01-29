# URLShortener

This project will help you convert long URL to short URL,

and allow user to set the expiration time for short URL.

## Usage

[Link](https://drive.google.com/file/d/1_yZyP1PF0BiUaR1WcTTOGwPIDxRK_yLz/view?usp=sharing)

## Setup

### Use Docker Compose

1. Deploy the project on Docker

`docker-compose up --build`

2. Once done, enter http://localhost:8080/shorten into the webpage to shorten your URL

**NOTE: Run the Docker command in the root project directory**

### Use Kubernetes

1. Apply all Kubernetes setting in k8s folder

`kubectl apply -f ./k8s/`

2. Once done, enter http://localhost:30080/shorten into the webpage to shorten your URL

**NOTE: Run the k8s command in the root project directory**
