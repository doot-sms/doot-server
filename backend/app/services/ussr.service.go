package services

import (
	"context"

	"github.com/doot-sms/doot-server/pkg/db"
)

type USSRCreateRequestParams struct {
	UserId    int32
	SenderId  int32
	Requestor db.UssrRequestor
}

type IUSSRService interface {
	CreateRequest(c context.Context, data USSRCreateRequestParams) (db.UserSenderReq, error)
}
