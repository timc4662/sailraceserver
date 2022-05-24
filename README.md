# sailrace server

A gRPC server which works with Sailrace Dinghy Race timer redundancy mode.

See https://sailrace.app/

## Running the server

    go run main.go

## Connecting the app

Enter settings > redundancy and switch on.
Click the "Scan other app to connect" button, but rather than scanning, enter the IP address of the server running this code. Note: Ensure that your device and the PC are on the same network. The icon should then go green in the app.

## What does it do

At present, it just writes the race timings to the console. 

Later it may write CSV files to disk in realtime for backup and ease of import into a results system.

## Service definition

The gRPC service definition can be found in protos/sailrace.proto. This can be generated into a gRPC server in most languages if needed.
    