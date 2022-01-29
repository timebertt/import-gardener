.PHONY: modules
modules:
	go mod tidy
	# vendoring to see what is really pulled in by go modules
	go mod vendor
