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

# Port on which to listen for the blazedb server
# port = 9876

# Max queue size for the blazedb server commands
# max_queue_size = 100

# The logfile
# log_file = "/tmp/blazedb.log"

# Where to save data on disk
# save_file = "/tmp/blaze.dump"

# How often we need to save data on disk
# save_period = "1m"

# Debug mode (0 or 1)
# debug = 0
```


## Todos

- [x] Save database to disk
- [x] Log file
- [x] Configuration file
- [x] Add a go client
- [ ] Tests
- [ ] Documentation
- [ ] Code example
- [ ] Add a parameter to dump the state every N write commands
- [ ] Add SSL support