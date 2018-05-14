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

## Todos

- [ ] Save database to disk
- [ ] Log file
- [ ] Configuration file
- [ ] Tests
- [ ] Code example