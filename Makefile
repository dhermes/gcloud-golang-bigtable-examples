ifeq ($(USE_APP_DEFAULT),)
	RUN_FLAGS=""
else
	RUN_FLAGS="-use-app-default"
endif
GOPATH=$(shell pwd)/gopath

help:
	@echo 'Makefile for GOLANG BigTable sample                                           '
	@echo '                                                                              '
	@echo '   make path                   Create the GOPATH dir                          '
	@echo '   make install                Install the GO dependencies                    '
	@echo '   make run_cluster            Run example for Cluster Admin API              '
	@echo '   make run_table              Run example for Table Admin API                '
	@echo '   make run_table_with_create  Example for Table Admin API with table creation'

path:
	mkdir -p $(GOPATH)

install: path
	GOPATH=$(GOPATH) go get google.golang.org/cloud/bigtable

run_cluster: install
	GOPATH=$(GOPATH) go run main_with_cluster_admin.go consts.go helpers.go $(RUN_FLAGS)

run_table: install
	GOPATH=$(GOPATH) go run main_with_table_admin.go consts.go helpers.go $(RUN_FLAGS)

run_table_with_create: install
	GOPATH=$(GOPATH) go run main_with_table_admin_and_create.go consts.go helpers.go $(RUN_FLAGS)

.PHONY: path install run_cluster run_table run_table_with_create
