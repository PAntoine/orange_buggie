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

all: build_graph

bin:
	@- mkdir -p bin

build_graph: bin
	@go build -o bin/build_graph cmd/build_graph.go
