#!/bin/bash

docker run --rm --name mysqld -p 3306:3306 -v ${PWD}/datadir:/var/lib/mysql -v ${PWD}/playground:/playground -e MYSQL_ROOT_PASSWORD=hello -d mysql:8.0.16