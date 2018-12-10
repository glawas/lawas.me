#!/bin/bash
docker run -it -p 8686:8686 \
-v /home/glawas/projects/golang/lawas.me/app:/go/src/github.com/glawas/lawas.me/app \
-v /home/glawas/projects/golang/lawas.me/public:/go/src/github.com/glawas/lawas.me/public \
go-lawas
