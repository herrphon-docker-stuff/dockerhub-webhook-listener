#!/bin/bash

sudo docker pull nathanleclaire/octoblog:latest
docker stop blog
docker rm blog
docker run --name blog -d -p 8000:80 nathanleclaire/octoblog