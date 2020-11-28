package enums

import (
	"encoding/json"
	"errors"
)

type EnumCode int

const (
	GetErrorOneResult  EnumCode = 1001
	GetErrorFullResult EnumCode = 1002
)

func (lt *EnumCode) UnmarshalJSON(b []byte) error {
	var s int
	_ = json.Unmarshal(b, &s)
	leaveType := EnumCode(s)
	switch leaveType {
	case GetErrorFullResult, GetErrorOneResult:
		*lt = leaveType
		return nil
	}
	return errors.New("Invalid enum code type")
}

func (lt EnumCode) IsValid() error {
	switch lt {
	case GetErrorFullResult, GetErrorOneResult:
		return nil
	}
	return errors.New("Invalid enum code type")
}
