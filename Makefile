RUN_FLAGS=
ifneq ($(USE_APP_DEFAULT),)
	RUN_FLAGS+= -use-app-default
endif
ifneq ($(VERBOSE),)
	RUN_FLAGS+= -verbose
endif
GOPATH=$(shell pwd)/gopath

help:
	@echo 'Makefile for GOLANG BigTable sample                                       '
	@echo '                                                                          '
	@echo '   make path            Create the GOPATH dir                             '
	@echo '   make install         Install the GO dependencies                       '
	@echo '   make update_install  Update the installed GO dependencies              '
	@echo '   make list_clusters   Run example for Cluster Admin API                 '
	@echo '   make create_cluster  Cluster Admin API: Create new cluster             '
	@echo '   make delete_cluster  Cluster Admin API: Delete a cluster               '
	@echo '   make list_zones      Cluster Admin API: List zones                     '
	@echo '   make list_tables     Run example for Table Admin API                   '
	@echo '   make create_table    Example for Table Admin API with table creation   '
	@echo '                                                                          '
	@echo 'NOTE: Append USE_APP_DEFAULT=True to the end of your make command to      '
	@echo '      switch from a service account to a user account (via the application'
	@echo '      default credentials).                                               '
	@echo '                                                                          '
	@echo 'NOTE: Append VERBOSE=True to the end of your make command to log more     '
	@echo '      output from your examples.                                          '

path:
	mkdir -p $(GOPATH)

install: path
	GOPATH=$(GOPATH) go get github.com/dhermes/gcloud-golang/bigtable

update_install: path
	GOPATH=$(GOPATH) go get -u github.com/dhermes/gcloud-golang/bigtable

list_clusters: install
	GOPATH=$(GOPATH) go run main_list_clusters.go consts.go helpers.go $(RUN_FLAGS)

create_cluster: install
	GOPATH=$(GOPATH) go run main_create_cluster.go consts.go helpers.go $(RUN_FLAGS)

delete_cluster: install
	GOPATH=$(GOPATH) go run main_delete_cluster.go consts.go helpers.go $(RUN_FLAGS)

list_zones: install
	GOPATH=$(GOPATH) go run main_list_zones.go consts.go helpers.go $(RUN_FLAGS)

list_tables: install
	GOPATH=$(GOPATH) go run main_list_tables.go consts.go helpers.go $(RUN_FLAGS)

create_table: install
	GOPATH=$(GOPATH) go run main_create_table.go consts.go helpers.go $(RUN_FLAGS)

.PHONY: path install update_install list_clusters create_clutser list_zones list_tables create_table
