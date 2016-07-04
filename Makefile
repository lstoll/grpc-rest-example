.PHONY: proto gogo-protobuf bundle

proto:	tmp/protoc/protoc gogo-protobuf bundle
	mkdir -p go/server/releaser
	mkdir -p ruby/lib
	tmp/protoc/protoc --proto_path=vendor:. --gogo_out=plugins=grpc:go/server/releaser service.proto
	cd ruby && bundle exec ../tmp/protoc/protoc --ruby_out=lib --proto_path=../vendor:.. --grpc_out=lib --plugin=protoc-gen-grpc=$(shell which grpc_tools_ruby_protoc_plugin.rb) ../service.proto

tmp/protoc/protoc:
# ahahahahahahahahaha
	mkdir -p tmp/protoc && cd tmp/protoc && curl -sL https://github.com/google/protobuf/releases/download/v3.0.0-beta-3/protoc-3.0.0-beta-3-osx-x86_64.zip | python -c "import zipfile,sys,StringIO;zipfile.ZipFile(StringIO.StringIO(sys.stdin.read())).extractall(sys.argv[1] if len(sys.argv) == 2 else '.')" && chmod +x protoc

gogo-protobuf:
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	go get -u github.com/gogo/protobuf/gogoproto

bundle:
	cd ruby && bundle