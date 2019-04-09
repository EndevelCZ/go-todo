# REPOSITORY TO LEARN FROM
[google bookshelf](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/getting-started/bookshelf)

[hex-example](https://github.com/Holmes89/hex-example)

# go microservices frameworks
[micro](https://micro.mu/)

[go kit](https://gokit.io/)

```sh
# start datastore server
go get -u cloud.google.com/go/datastore
gcloud components install cloud-datastore-emulator
gcloud beta emulators datastore start

# export variable for datastore localhost
$(gcloud beta emulators datastore env-init)

# run server
go run main.go

# add todo
curl -X POST http://localhost/todos -H 'Content-Type: application/json' -d '{ "text": "second todo" } '
# list todos
curl http://localhost/todos -H 'Content-Type: application/json'
```


# grpc debugging
```sh
cat pb/add.json | prototool grpc pb \
--address 0.0.0.0:5000 \
--method pb.Todos/AddTodo \
--stdin
```


# DEBUGGER
1) install DELVE!!
[vscode debugger DELVE!!](https://github.com/go-delve/delve)
[vscode doc doc](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code)
```sh
```

### vscode debug package
```json
{
    "name": "Launch Package Working",
    "type": "go",
    "request": "launch",
    "mode": "debug",
    "program": "${fileDirname}",
    "env": {"DATASTORE_EMULATOR_HOST":"0.0.0.0:8081"},
},
```

# remote DEBUGGER
1)
```json
{
    "name": "Connect to server",
    "type": "go",
    "request": "launch",
    "mode": "remote",
    "remotePath": "${workspaceFolder}",
    "port": 2345,
    "host": "127.0.0.1",
    "program": "${workspaceFolder}",
    "env": {"DATASTORE_EMULATOR_HOST":"0.0.0.0:8081"},
    "args": [],
},
```
2) run binary
```sh
./server
```

3) GET PID of running process
```sh
lsof -i:5000
# or
ps aux | grep server
```

4) run dlv
```sh
dlv attach 58400 --headless --listen=:2345 --log --api-version=2
```

5) run vscode remote debugger and create breakpoints