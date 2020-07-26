ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

all: ./proto/*/
	@echo "Clean codegen..."; \
	rm -rf codegen; \
	mkdir codegen; \
	mkdir codegen/go; \
	mkdir codegen/swagger; \
	echo "Done."; \
	echo "Processing proto files..."; \
	for dir in $^ ; do \
		a=`echo $(basename $${dir}) | cut -d/ -f2` ; \
		echo "Processing $$a..."; \
		mkdir $(ROOT_DIR)/codegen/go/$$a; \
		echo "  Generating go for $$a..."; \
		protoc -I$(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
			-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6 \
			-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
			--proto_path=proto --go_opt=paths=source_relative --go_out=plugins=grpc:codegen/go proto/$$a/*.proto \
			--grpc-gateway_out=logtostderr=true:$(GOPATH)/src \
            --swagger_out=logtostderr=true,json_names_for_fields=true:codegen/swagger \
            --validate_out="lang=go:$(GOPATH)/src"; \
		echo "Done $$a."; \
	done;
	@echo "Done!"
