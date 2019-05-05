#!/bin/bash -x

docker rm -f $(docker ps -aq)
export LOCAL_IP=$(python -c "import socket; print socket.gethostbyname(socket.gethostname())")

docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from test
