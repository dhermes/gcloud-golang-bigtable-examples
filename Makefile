RUN_FLAGS=
ifneq ($(USE_APP_DEFAULT),)
	RUN_FLAGS+= -use-app-default
endif
ifneq ($(VERBOSE),)
	RUN_FLAGS+= -verbose
endif
GOPATH=$(shell pwd)/gopath

help:
	@echo 'Makefile for GOLANG BigTable sample                                             '
	@echo '                                                                                '
	@echo '   make path                     Create the GOPATH dir                          '
	@echo '   make install                  Install the GO dependencies                    '
	@echo '   make list_clusters            Run example for Cluster Admin API              '
	@echo '   make list_zones               Cluster Admin API: List zones                  '
	@echo '   make list_tables              Run example for Table Admin API                '
	@echo '   make list_tables_with_create  Example for Table Admin API with table creation'
	@echo '                                                                                '
	@echo 'NOTE: Append USE_APP_DEFAULT=True to the end of your make command to            '
	@echo '      switch from a service account to a user account (via the application      '
	@echo '      default credentials).                                                     '
	@echo '                                                                                '
	@echo 'NOTE: Append VERBOSE=True to the end of your make command to log more           '
	@echo '      output from your examples.                                                '

path:
	mkdir -p $(GOPATH)

install: path
	GOPATH=$(GOPATH) go get github.com/dhermes/gcloud-golang/bigtable

list_clusters: install
	GOPATH=$(GOPATH) go run main_with_cluster_admin.go consts.go helpers.go cluster_api.go $(RUN_FLAGS)

list_zones: install
	GOPATH=$(GOPATH) go run main_list_zones.go consts.go helpers.go cluster_api.go $(RUN_FLAGS)

list_tables: install
	GOPATH=$(GOPATH) go run main_with_table_admin.go consts.go helpers.go $(RUN_FLAGS)

list_tables_with_create: install
	GOPATH=$(GOPATH) go run main_with_table_admin_and_create.go consts.go helpers.go $(RUN_FLAGS)

.PHONY: path install list_clusters list_tables list_tables_with_create
