/*
* Change the host from 'localhost' to '%' to adapt for docker container.
* As explained by https://github.com/docker-library/mysql/issues/230
*/
CREATE USER 'cbuser'@'%' IDENTIFIED BY 'cbpass';
GRANT ALL ON cookbook.* TO 'cbuser'@'%';