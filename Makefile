ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

GO_ARGS:=--go_opt=paths=source_relative --go_out=plugins=grpc:codegen/go

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
		protoc -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
			--proto_path=proto $(GO_ARGS) proto/$$a/*.proto \
			--grpc-gateway_out=logtostderr=true:$(GOPATH)/src \
            --swagger_out=logtostderr=true,json_names_for_fields=true:codegen/swagger; \
		echo "Done $$a."; \
	done;
	@echo "Done!"
