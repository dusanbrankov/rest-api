# REST API written in Go

Source: <https://www.youtube.com/watch?v=7VLmLOiQ3ck&t=4513s>

## Instructions

### migrate

[migrate](https://github.com/golang-migrate/migrate/tree/v4.17.1) is a Go library and CLI tool for executing database migrations.

**Installation**

```sh
# download deb file
curl -O https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb

# install
sudo dpkg -i migrate.linux-amd64.deb
```

**Usage**

```sh
# the following command creates the respective SQL files in the
# directory specified in 'Makefile'
make migration add-user-table

# write the queries in the newly created SQL files and run:
make migrate-up
```


