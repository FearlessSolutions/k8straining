#!/bin/bash

echo 'Running: npm ' $1

docker exec -w /reference-app -it reference-app sh -c "npm $1"
