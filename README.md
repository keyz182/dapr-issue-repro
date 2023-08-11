# Dapr pluggable binding - problem hunting

Minimal reproduction of my issue starting a pluggable binding. 

Run with 

```
go run main.go
```


Then run the test with:

```
dapr run --app-id batch-sdk --app-port 6004 -G 6005 --resources-path components --log-level debug -- go test

```


Errors with:
```
WARN[0000] Error processing component, daprd process will exit gracefully  app_id=batch-sdk instance=APOLLO scope=dapr.runtime type=log ver=1.11.2
FATA[0000] process component mylog error: incorrect type binding.mylog  app_id=batch-sdk instance=APOLLO scope=dapr.runtime type=log ver=1.11.2
❌  The daprd process exited with error code: exit status 1
```