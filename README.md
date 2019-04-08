# REPOSITORY TO LEARN FROM
[google bookshelf](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/getting-started/bookshelf)
[hex-example](https://github.com/Holmes89/hex-example)

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