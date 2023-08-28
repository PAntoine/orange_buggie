#
# Name  : makefile
# Desc  : Makefile for building the orange buggie voice system
#
# Author: Peter Antoine
# Date  : 08/08/2023
#
#                     Copyright (c) 2023 Peter Antoine
#                            All rights Reserved.
#                      Released Under the MIT Licence

all: build_graph run_model

bin:
	@- mkdir -p bin

build_graph: bin
	@go build -o bin/build_graph cmd/build_graph.go

run_model: bin
	@go build -o bin/run_model cmd/run_model.go

tests:
	@go test ./source/language_model

debug_tests:
	@go test ./source/language_model -v

