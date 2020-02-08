
	# Generate gogo, gRPC-Gateway, swagger, go-validators output.
	#
	# -I declares import folders, in order of importance
	# This is how proto resolves the protofile imports.
	# It will check for the protofile relative to each of these
	# folders and use the first one it finds.
	#
	# --gogo_out generates GoGo Protobuf output with gRPC plugin enabled.
	# --grpc-gateway_out generates gRPC-Gateway output.
	# --swagger_out generates an OpenAPI 2.0 specification for our gRPC-Gateway endpoints.
	# --govalidators_out generates Go validation files for our messages types, if specified.
	#
	# The lines starting with Mgoogle/... are proto import replacements,
	# which cause the generated file to import the specified packages
	# instead of the go_package's declared by the imported protof files.
	#
	# $$GOPATH/src is the output directory. It is relative to the GOPATH/src directory
	# since we've specified a go_package option relative to that directory.
	#
	# proto/example.proto is the location of the protofile we use.
echo:
	echo $(CURDIR)
	ls -a $$GOPATH/src/
	cp $$GOPATH/src/GameCtl.pb.go $(CURDIR)/proto/
	cp $$GOPATH/src/GameCtl.validator.pb.go $(CURDIR)/proto/
	rm $$GOPATH/src/GameCtl.pb.go
	rm $$GOPATH/src/GameCtl.validator.pb.go
generate:
	protoc \
		-I proto \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/ \
		-I vendor/github.com/gogo/googleapis/ \
		-I vendor/ \
		--gogo_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		--grpc-gateway_out=allow_patch_feature=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		--swagger_out=third_party/OpenAPI/ \
		--govalidators_out=gogoimport=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		proto/GameCtl.proto
	# gvm issue :  move the genrated file to current directory
	mv $(CURDIR)/vendor/GameCtl.pb.go $(CURDIR)/proto/
	mv $(CURDIR)/vendor/GameCtl.validator.pb.go $(CURDIR)/proto/
	mv $(CURDIR)/vendor/GameCtl.pb.gw.go $(CURDIR)/proto/
	## Workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/229.
	sed -i.bak "s/empty.Empty/types.Empty/g" proto/GameCtl.pb.gw.go && rm proto/GameCtl.pb.gw.go.bak

	## Generate static assets for OpenAPI UI
	statik -m -f -src third_party/OpenAPI/

install:
	go get \
		github.com/gogo/protobuf/protoc-gen-gogo \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		github.com/rakyll/statik

build: 
	go build -o bin/roomstatus_server.exe main.try.go



generate_vcred:
	protoc \
		-I proto/ \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/ \
		-I vendor/github.com/gogo/googleapis/ \
		-I vendor/ \
		--go_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		--grpc-gateway_out=allow_patch_feature=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		--swagger_out=third_party/OpenAPI/ \
		--govalidators_out=gogoimport=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		proto/cred.proto
	# gvm issue :  move the genrated file to current directory
	mv $(CURDIR)/vendor/cred.pb.go $(CURDIR)/proto/
	mv $(CURDIR)/vendor/cred.validator.pb.go $(CURDIR)/proto/
	mv $(CURDIR)/vendor/cred.pb.gw.go $(CURDIR)/proto/
	## Workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/229.
	sed -i.bak "s/empty.Empty/types.Empty/g" proto/cred.pb.gw.go && rm proto/cred.pb.gw.go.bak

	# ## Generate static assets for OpenAPI UI
	# statik -m -f -src third_party/OpenAPI/
