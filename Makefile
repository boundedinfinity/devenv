makefile_dir	:= $(abspath $(shell pwd))

app_name		:= avatar

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

project-bootstrap:
	brew install glide
	make go-bootstrap

project-clean:
	make go-clean

clean:
	make go-clean

go-bootstrap:
	go get -u github.com/jteeuwen/go-bindata/...
	glide install

go-data-create:
	cd $(makefile_dir)/data && go-bindata -pkg data ./...

go-data-clean:
	rm -rf $(makefile_dir)/data/bindata_assetfs.go
	rm -rf $(makefile_dir)/data/bindata.go

go-clean:
	go clean

go-scrub:
	rm -rf $(makefile_dir)/vendor
	make go-clean
	make go-data-clean

glide-clean:
	rm -rf $(makefile_dir)/vendor
	rm -rf $(makefile_dir)/glide.lock

go-build:
	make go-data-create
	go build
