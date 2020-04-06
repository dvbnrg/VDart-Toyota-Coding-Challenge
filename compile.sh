#!/bin/bash

protoc --proto_path=proto --go_out=plugins=grpc:proto hello.proto