// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSnatEntryResponseBody
	GetRequestId() *string
}

type DeleteSnatEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSnatEntryResponseBody) SetRequestId(v string) *DeleteSnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
