#!/bin/bash

cd ./test-network

./network.sh up
./network.sh createChannel -s couchdb
./network.sh deployCC -ccn examcret -ccp ../contract/ -ccl go

