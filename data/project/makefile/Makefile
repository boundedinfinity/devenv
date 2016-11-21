makefile_dir	:= $(abspath $(shell pwd))

app_name		:= {{.ProjectName}}

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort
