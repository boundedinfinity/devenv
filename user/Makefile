makefile_dir	:= $(abspath $(shell pwd))

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

test-clean:
	rm -rf $(makefile_dir)/test-output
	go test

test-dirty:
	go test
