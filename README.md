# Go Quiz
A simple Golang backend for making a quiz and a few frontends

## Helpers
Import database from file to container:
```sh
cat backup.sql | docker exec -i <container_id> /usr/bin/mysql -u root --password=qwerty whatcanido
```

As `container_id` put mariadb container ID. You can get it by using `docker ps` command with enabled db.

## How to launch
```sh
## Generate .env
 cd scripts;
 yarn;
 yarn generate-dotenv;

# Launching DB + PHPMyAdmin
 cd server;
 docker-compose up;

# Launching Rest API
 ./air
## or
 air

# Launching PWA
 cd client;
 yarn dev;
```
