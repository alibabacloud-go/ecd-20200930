// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTaskApprovalCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferTaskApprovalCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferTaskApprovalCallbackResponse
	GetStatusCode() *int32
	SetBody(v *TransferTaskApprovalCallbackResponseBody) *TransferTaskApprovalCallbackResponse
	GetBody() *TransferTaskApprovalCallbackResponseBody
}

type TransferTaskApprovalCallbackResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferTaskApprovalCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferTaskApprovalCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferTaskApprovalCallbackResponse) GoString() string {
	return s.String()
}

func (s *TransferTaskApprovalCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferTaskApprovalCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferTaskApprovalCallbackResponse) GetBody() *TransferTaskApprovalCallbackResponseBody {
	return s.Body
}

func (s *TransferTaskApprovalCallbackResponse) SetHeaders(v map[string]*string) *TransferTaskApprovalCallbackResponse {
	s.Headers = v
	return s
}

func (s *TransferTaskApprovalCallbackResponse) SetStatusCode(v int32) *TransferTaskApprovalCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferTaskApprovalCallbackResponse) SetBody(v *TransferTaskApprovalCallbackResponseBody) *TransferTaskApprovalCallbackResponse {
	s.Body = v
	return s
}

func (s *TransferTaskApprovalCallbackResponse) Validate() error {
	return dara.Validate(s)
}
