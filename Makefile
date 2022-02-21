SHELL := /bin/bash
BASEDIR = $(shell pwd)

.PHONY: osx
osx:
	@bash ./scripts/osx
	
clean: 
	@rm -fr build