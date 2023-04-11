#!/bin/bash
git reset --hard origin/master
git pull

docker build -t chat_go_img .
docker rm -f chat_go
docker run -p 8888:8888 --name chat_go -v $(pwd):/app chat_go_img