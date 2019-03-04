.PHONY: all
all:
	@if $(MAKE) test; then \
		echo "\033[1;32mSuccessful native run! uname: $$(uname -a)\033[0m"; \
	else \
		echo "\033[1;31mFailed native run! uname: $$(uname -a)\033[0m"; \
		exit 1; \
	fi
	@if $(MAKE) test-docker; then \
		echo "\033[1;32mSuccessful docker run! uname: $$(uname -a)\033[0m"; \
	else \
		echo "\033[1;31mFailed docker run! uname: $$(uname -a)\033[0m"; \
		exit 1; \
	fi

.PHONY: test
test: test-loader test-cloader

.PHONY: test-docker
test-docker:
	docker run --rm -it \
		--name golang-plugin \
		-v $$PWD:/src \
		golang:1.11-stretch \
			bash -c 'set -ex; cd /src; make clean test'

.PHONY: pow
pow:
	go build -buildmode=plugin -o ./out/pow ./pow/main.go

.PHONY: loader
loader: out pow
	go build -o out/loader-main ./loader/loader.go

.PHONY: test-loader
test-loader: loader
	./out/loader-main


.PHONY: cloader
cloader: out pow
	go build -buildmode=c-shared -o out/cloader.so ./cloader/loader.go
	gcc -o out/cloader-main ./cloader/main.c -ldl

.PHONY: test-cloader
test-cloader: cloader
	./out/cloader-main


out:
	mkdir out

.PHONY: clean
clean:
	rm -rf out
