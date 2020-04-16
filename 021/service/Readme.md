# Docker image for application

## Description
A docker image for the web application in web directory + a podspec for a 2 container pod with this application and a postgres db.

## Prerequisites
1. docker
2. golang >= 1.12

## setup

`make dockerv1` - for v1
`make dockerv2` - for v2

`kubectl apply -f webpod.yml`
