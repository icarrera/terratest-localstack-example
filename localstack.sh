#!/bin/sh

echo "deploying localstack"
docker run -it -p 4566:4566 -p 4571:4571 localstack/localstack