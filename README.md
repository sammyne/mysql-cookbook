# mysql-cookbook

## Prerequisite

- Docker
- MySQL-8.0.16

## Scripts

### mysqld.sh

- Run a MySQL server as daemon with
  - Name `mysqld`
  - Root password as `hello`
  - Data stored in `datadir` of current working directory
  - The `playground` directory of current working directory is synchronized with `/playground` in the container
  - Listening on port `3306`

### mysql.sh

- Connect to the container `mysqld` and execute the `mysql` command |

## Progress

- [x] Chapter01 - Using the mysql Client Program
- [ ] Chapter02 - Writing MySQL-Based Programs

# References

- [mysql-Docker Official Images](https://hub.docker.com/_/mysql?tab=description)
- [Connector/C++](https://dev.mysql.com/downloads/connector/cpp/)
