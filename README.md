# blazedb

BlazeDB is a simple key value storage system.

It uses a tcp server to communicate with clients with the following available commands:

- `set key value`
- `get key`
- `update key value`
- `delete key`

There is a client embedded within the blazedb binary but you can also use telnet to communicate with the server:

```
$ telnet localhost 9876
Trying ::1...
Connected to localhost.
Escape character is '^]'.
set x 1
success
get x
1
```

## Starting the server

To start the blazedb server simply do

`blazedb server -c <path_to_the_config_file>`

## Configuration

The sample configuration can be found in server/config.template.toml and contains the following:

```
# Default blazedb configuration

port = 9876
max_queue_size = 100
log_file = "/tmp/blazedb.log"
save_directory = "/tmp"
```


## Todos

- [ ] Save database to disk
- [ ] Log file
- [x] Configuration file
- [ ] Tests
- [ ] Code example