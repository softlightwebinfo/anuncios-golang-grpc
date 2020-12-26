cd ../proto
protoc --go_out=plugins=grpc:. *.proto
#protoc -I=. *.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.

protoc \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:. \
    --js_out=import_style=commonjs:. -I=. *.proto