package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type Reservation struct {
	Base
	CreateDate timestamp.Timestamp
	Meeting    Meeting
}
