#!/bin/bash
go build
shopt -s extglob
strip tproxy
echo "Image: urtho/algotelemetry"
docker build . -t urtho/algotelemetry:latest 
docker push -a urtho/algotelemetry
