#!/bin/sh
echo "---------------------> build image name forum"
docker build -t forum .
docker images

echo "---------------------> build container and run name forumcon"
docker container run -p 8080:8080 --detach --name forumcon forum
docker ps -a

echo "---------------------> clean up unused images."
docker image prune -a -f