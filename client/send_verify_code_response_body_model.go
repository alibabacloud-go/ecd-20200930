// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerifyCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendVerifyCodeResponseBody
	GetRequestId() *string
}

type SendVerifyCodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerifyCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendVerifyCodeResponseBody) SetRequestId(v string) *SendVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendVerifyCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
