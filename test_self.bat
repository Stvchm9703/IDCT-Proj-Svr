protoc ^
    -I proto ^
    -I vendor/github.com/grpc-ecosystem/grpc-gateway/ ^
    -I vendor/github.com/gogo/googleapis/ ^
    -I vendor/ ^
    --gogo_out=plugins=grpc,^
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,^
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,^
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:^
C:/Users/"Steven Chm"/Go/src/ ^
    --grpc-gateway_out=allow_patch_feature=false,^
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,^
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,^
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:^
C:/Users/"Steven Chm"/Go/src/ ^
    --swagger_out=third_party/OpenAPI/ ^
    --govalidators_out=gogoimport=true,^
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,^
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,^
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:^
C:/Users/"Steven Chm"/Go/src ^
    proto/GameCtl.proto



REM sed -i.bak "s/empty.Empty/types.Empty/g" proto/GameCtl.pb.gw.go && rm proto/GameCtl.pb.gw.go.bak
statik -m -f -src third_party/OpenAPI/