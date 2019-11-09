// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: GameCtl.proto

package proto

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/gogo/googleapis/google/api"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *RoomListRequest) Validate() error {
	return nil
}
func (this *RoomListResponse) Validate() error {
	for _, item := range this.Result {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Result", err)
			}
		}
	}
	return nil
}
func (this *RoomCreateRequest) Validate() error {
	return nil
}
func (this *RoomRequest) Validate() error {
	return nil
}
func (this *Room) Validate() error {
	for _, item := range this.CellStatus {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CellStatus", err)
			}
		}
	}
	return nil
}
func (this *CellStatus) Validate() error {
	return nil
}
func (this *CreateCredReq) Validate() error {
	return nil
}
func (this *Cred) Validate() error {
	return nil
}
