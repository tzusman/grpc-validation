SHELL=/bin/bash -o pipefail

PROTOCOLS=$(wildcard ./protocols/*/*.proto)
GOPBS=$(PROTOCOLS:.proto=.pb.go)

.PHONY: protocols
protocols: $(GOPBS)
%.pb.go: %.proto
	protoc -I $(dir $^) $^ --go_out=plugins=grpc:$(dir $@)
	# go run inject/main.go 