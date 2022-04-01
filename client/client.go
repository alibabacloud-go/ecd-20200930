// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateOfficeSiteRequest struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ActivateOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *ActivateOfficeSiteRequest) SetOfficeSiteId(v string) *ActivateOfficeSiteRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ActivateOfficeSiteRequest) SetRegionId(v string) *ActivateOfficeSiteRequest {
	s.RegionId = &v
	return s
}

type ActivateOfficeSiteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateOfficeSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateOfficeSiteResponseBody) SetRequestId(v string) *ActivateOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

type ActivateOfficeSiteResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ActivateOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivateOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *ActivateOfficeSiteResponse) SetHeaders(v map[string]*string) *ActivateOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *ActivateOfficeSiteResponse) SetBody(v *ActivateOfficeSiteResponseBody) *ActivateOfficeSiteResponse {
	s.Body = v
	return s
}

type AddUserToDesktopGroupRequest struct {
	ClientToken     *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DesktopGroupId  *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	EndUserIds      []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddUserToDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupRequest) SetClientToken(v string) *AddUserToDesktopGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetDesktopGroupId(v string) *AddUserToDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetDesktopGroupIds(v []*string) *AddUserToDesktopGroupRequest {
	s.DesktopGroupIds = v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetEndUserIds(v []*string) *AddUserToDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetRegionId(v string) *AddUserToDesktopGroupRequest {
	s.RegionId = &v
	return s
}

type AddUserToDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserToDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupResponseBody) SetRequestId(v string) *AddUserToDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddUserToDesktopGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUserToDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserToDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupResponse) SetHeaders(v map[string]*string) *AddUserToDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUserToDesktopGroupResponse) SetBody(v *AddUserToDesktopGroupResponseBody) *AddUserToDesktopGroupResponse {
	s.Body = v
	return s
}

type ApproveFotaUpdateRequest struct {
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApproveFotaUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateRequest) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateRequest) SetAppVersion(v string) *ApproveFotaUpdateRequest {
	s.AppVersion = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetDesktopId(v string) *ApproveFotaUpdateRequest {
	s.DesktopId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetRegionId(v string) *ApproveFotaUpdateRequest {
	s.RegionId = &v
	return s
}

type ApproveFotaUpdateResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveFotaUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateResponseBody) SetRequestId(v string) *ApproveFotaUpdateResponseBody {
	s.RequestId = &v
	return s
}

type ApproveFotaUpdateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApproveFotaUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveFotaUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateResponse) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateResponse) SetHeaders(v map[string]*string) *ApproveFotaUpdateResponse {
	s.Headers = v
	return s
}

func (s *ApproveFotaUpdateResponse) SetBody(v *ApproveFotaUpdateResponseBody) *ApproveFotaUpdateResponse {
	s.Body = v
	return s
}

type AttachCenRequest struct {
	CenId        *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId   *int64  `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VerifyCode   *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s AttachCenRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachCenRequest) GoString() string {
	return s.String()
}

func (s *AttachCenRequest) SetCenId(v string) *AttachCenRequest {
	s.CenId = &v
	return s
}

func (s *AttachCenRequest) SetCenOwnerId(v int64) *AttachCenRequest {
	s.CenOwnerId = &v
	return s
}

func (s *AttachCenRequest) SetOfficeSiteId(v string) *AttachCenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *AttachCenRequest) SetRegionId(v string) *AttachCenRequest {
	s.RegionId = &v
	return s
}

func (s *AttachCenRequest) SetVerifyCode(v string) *AttachCenRequest {
	s.VerifyCode = &v
	return s
}

type AttachCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachCenResponseBody) GoString() string {
	return s.String()
}

func (s *AttachCenResponseBody) SetRequestId(v string) *AttachCenResponseBody {
	s.RequestId = &v
	return s
}

type AttachCenResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachCenResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachCenResponse) GoString() string {
	return s.String()
}

func (s *AttachCenResponse) SetHeaders(v map[string]*string) *AttachCenResponse {
	s.Headers = v
	return s
}

func (s *AttachCenResponse) SetBody(v *AttachCenResponseBody) *AttachCenResponse {
	s.Body = v
	return s
}

type CheckUserTagsRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags     []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CheckUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTagsRequest) GoString() string {
	return s.String()
}

func (s *CheckUserTagsRequest) SetRegionId(v string) *CheckUserTagsRequest {
	s.RegionId = &v
	return s
}

func (s *CheckUserTagsRequest) SetTags(v []*string) *CheckUserTagsRequest {
	s.Tags = v
	return s
}

type CheckUserTagsResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckUserTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserTagsResponseBody) SetData(v bool) *CheckUserTagsResponseBody {
	s.Data = &v
	return s
}

func (s *CheckUserTagsResponseBody) SetRequestId(v string) *CheckUserTagsResponseBody {
	s.RequestId = &v
	return s
}

type CheckUserTagsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserTagsResponse) GoString() string {
	return s.String()
}

func (s *CheckUserTagsResponse) SetHeaders(v map[string]*string) *CheckUserTagsResponse {
	s.Headers = v
	return s
}

func (s *CheckUserTagsResponse) SetBody(v *CheckUserTagsResponseBody) *CheckUserTagsResponse {
	s.Body = v
	return s
}

type ClonePolicyGroupRequest struct {
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ClonePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ClonePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupRequest) SetName(v string) *ClonePolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *ClonePolicyGroupRequest) SetPolicyGroupId(v string) *ClonePolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ClonePolicyGroupRequest) SetRegionId(v string) *ClonePolicyGroupRequest {
	s.RegionId = &v
	return s
}

type ClonePolicyGroupResponseBody struct {
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClonePolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClonePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupResponseBody) SetPolicyGroupId(v string) *ClonePolicyGroupResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *ClonePolicyGroupResponseBody) SetRequestId(v string) *ClonePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type ClonePolicyGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClonePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClonePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ClonePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupResponse) SetHeaders(v map[string]*string) *ClonePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ClonePolicyGroupResponse) SetBody(v *ClonePolicyGroupResponseBody) *ClonePolicyGroupResponse {
	s.Body = v
	return s
}

type ConfigADConnectorTrustRequest struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TrustKey     *string `json:"TrustKey,omitempty" xml:"TrustKey,omitempty"`
}

func (s ConfigADConnectorTrustRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigADConnectorTrustRequest) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorTrustRequest) SetOfficeSiteId(v string) *ConfigADConnectorTrustRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ConfigADConnectorTrustRequest) SetRegionId(v string) *ConfigADConnectorTrustRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigADConnectorTrustRequest) SetTrustKey(v string) *ConfigADConnectorTrustRequest {
	s.TrustKey = &v
	return s
}

type ConfigADConnectorTrustResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigADConnectorTrustResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigADConnectorTrustResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorTrustResponseBody) SetRequestId(v string) *ConfigADConnectorTrustResponseBody {
	s.RequestId = &v
	return s
}

type ConfigADConnectorTrustResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigADConnectorTrustResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigADConnectorTrustResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigADConnectorTrustResponse) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorTrustResponse) SetHeaders(v map[string]*string) *ConfigADConnectorTrustResponse {
	s.Headers = v
	return s
}

func (s *ConfigADConnectorTrustResponse) SetBody(v *ConfigADConnectorTrustResponseBody) *ConfigADConnectorTrustResponse {
	s.Body = v
	return s
}

type ConfigADConnectorUserRequest struct {
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DomainUserName *string `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	OUName         *string `json:"OUName,omitempty" xml:"OUName,omitempty"`
	OfficeSiteId   *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigADConnectorUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigADConnectorUserRequest) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorUserRequest) SetDomainPassword(v string) *ConfigADConnectorUserRequest {
	s.DomainPassword = &v
	return s
}

func (s *ConfigADConnectorUserRequest) SetDomainUserName(v string) *ConfigADConnectorUserRequest {
	s.DomainUserName = &v
	return s
}

func (s *ConfigADConnectorUserRequest) SetOUName(v string) *ConfigADConnectorUserRequest {
	s.OUName = &v
	return s
}

func (s *ConfigADConnectorUserRequest) SetOfficeSiteId(v string) *ConfigADConnectorUserRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ConfigADConnectorUserRequest) SetRegionId(v string) *ConfigADConnectorUserRequest {
	s.RegionId = &v
	return s
}

type ConfigADConnectorUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigADConnectorUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigADConnectorUserResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorUserResponseBody) SetRequestId(v string) *ConfigADConnectorUserResponseBody {
	s.RequestId = &v
	return s
}

type ConfigADConnectorUserResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigADConnectorUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigADConnectorUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigADConnectorUserResponse) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorUserResponse) SetHeaders(v map[string]*string) *ConfigADConnectorUserResponse {
	s.Headers = v
	return s
}

func (s *ConfigADConnectorUserResponse) SetBody(v *ConfigADConnectorUserResponseBody) *ConfigADConnectorUserResponse {
	s.Body = v
	return s
}

type CreateADConnectorDirectoryRequest struct {
	DesktopAccessType   *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	EnableAdminAccess   *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Specification       *int64    `json:"Specification,omitempty" xml:"Specification,omitempty"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	VSwitchId           []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s CreateADConnectorDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryRequest) SetDesktopAccessType(v string) *CreateADConnectorDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDirectoryName(v string) *CreateADConnectorDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDnsAddress(v []*string) *CreateADConnectorDirectoryRequest {
	s.DnsAddress = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainName(v string) *CreateADConnectorDirectoryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainPassword(v string) *CreateADConnectorDirectoryRequest {
	s.DomainPassword = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainUserName(v string) *CreateADConnectorDirectoryRequest {
	s.DomainUserName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetEnableAdminAccess(v bool) *CreateADConnectorDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetMfaEnabled(v bool) *CreateADConnectorDirectoryRequest {
	s.MfaEnabled = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetRegionId(v string) *CreateADConnectorDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSpecification(v int64) *CreateADConnectorDirectoryRequest {
	s.Specification = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSubDomainDnsAddress(v []*string) *CreateADConnectorDirectoryRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSubDomainName(v string) *CreateADConnectorDirectoryRequest {
	s.SubDomainName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetVSwitchId(v []*string) *CreateADConnectorDirectoryRequest {
	s.VSwitchId = v
	return s
}

type CreateADConnectorDirectoryResponseBody struct {
	AdConnectors  []*CreateADConnectorDirectoryResponseBodyAdConnectors `json:"AdConnectors,omitempty" xml:"AdConnectors,omitempty" type:"Repeated"`
	DirectoryId   *string                                               `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrustPassword *string                                               `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
}

func (s CreateADConnectorDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponseBody) SetAdConnectors(v []*CreateADConnectorDirectoryResponseBodyAdConnectors) *CreateADConnectorDirectoryResponseBody {
	s.AdConnectors = v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetDirectoryId(v string) *CreateADConnectorDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetRequestId(v string) *CreateADConnectorDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetTrustPassword(v string) *CreateADConnectorDirectoryResponseBody {
	s.TrustPassword = &v
	return s
}

type CreateADConnectorDirectoryResponseBodyAdConnectors struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
}

func (s CreateADConnectorDirectoryResponseBodyAdConnectors) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponseBodyAdConnectors) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponseBodyAdConnectors) SetAddress(v string) *CreateADConnectorDirectoryResponseBodyAdConnectors {
	s.Address = &v
	return s
}

type CreateADConnectorDirectoryResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateADConnectorDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateADConnectorDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponse) SetHeaders(v map[string]*string) *CreateADConnectorDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetBody(v *CreateADConnectorDirectoryResponseBody) *CreateADConnectorDirectoryResponse {
	s.Body = v
	return s
}

type CreateADConnectorOfficeSiteRequest struct {
	AdHostname           *string   `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	Bandwidth            *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CenId                *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId           *int64    `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	CidrBlock            *string   `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	DesktopAccessType    *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DnsAddress           []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DomainName           *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainPassword       *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DomainUserName       *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	EnableAdminAccess    *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	MfaEnabled           *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	OfficeSiteName       *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	ProtocolType         *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Specification        *int64    `json:"Specification,omitempty" xml:"Specification,omitempty"`
	SubDomainDnsAddress  []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName        *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	VerifyCode           *string   `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s CreateADConnectorOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteRequest) SetAdHostname(v string) *CreateADConnectorOfficeSiteRequest {
	s.AdHostname = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetBandwidth(v int32) *CreateADConnectorOfficeSiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCenId(v string) *CreateADConnectorOfficeSiteRequest {
	s.CenId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCenOwnerId(v int64) *CreateADConnectorOfficeSiteRequest {
	s.CenOwnerId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCidrBlock(v string) *CreateADConnectorOfficeSiteRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDesktopAccessType(v string) *CreateADConnectorOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.DnsAddress = v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainName(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainPassword(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainPassword = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainUserName(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainUserName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetEnableAdminAccess(v bool) *CreateADConnectorOfficeSiteRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetEnableInternetAccess(v bool) *CreateADConnectorOfficeSiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetMfaEnabled(v bool) *CreateADConnectorOfficeSiteRequest {
	s.MfaEnabled = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetOfficeSiteName(v string) *CreateADConnectorOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetProtocolType(v string) *CreateADConnectorOfficeSiteRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetRegionId(v string) *CreateADConnectorOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSpecification(v int64) *CreateADConnectorOfficeSiteRequest {
	s.Specification = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSubDomainDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSubDomainName(v string) *CreateADConnectorOfficeSiteRequest {
	s.SubDomainName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetVerifyCode(v string) *CreateADConnectorOfficeSiteRequest {
	s.VerifyCode = &v
	return s
}

type CreateADConnectorOfficeSiteResponseBody struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateADConnectorOfficeSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteResponseBody) SetOfficeSiteId(v string) *CreateADConnectorOfficeSiteResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteResponseBody) SetRequestId(v string) *CreateADConnectorOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

type CreateADConnectorOfficeSiteResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateADConnectorOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateADConnectorOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteResponse) SetHeaders(v map[string]*string) *CreateADConnectorOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *CreateADConnectorOfficeSiteResponse) SetBody(v *CreateADConnectorOfficeSiteResponseBody) *CreateADConnectorOfficeSiteResponse {
	s.Body = v
	return s
}

type CreateBandwidthResourcePackagesRequest struct {
	Amount      *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AutoPay     *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	PackageSize *int32  `json:"PackageSize,omitempty" xml:"PackageSize,omitempty"`
	Period      *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit  *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBandwidthResourcePackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBandwidthResourcePackagesRequest) GoString() string {
	return s.String()
}

func (s *CreateBandwidthResourcePackagesRequest) SetAmount(v int32) *CreateBandwidthResourcePackagesRequest {
	s.Amount = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetAutoPay(v bool) *CreateBandwidthResourcePackagesRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetPackageSize(v int32) *CreateBandwidthResourcePackagesRequest {
	s.PackageSize = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetPeriod(v int32) *CreateBandwidthResourcePackagesRequest {
	s.Period = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetPeriodUnit(v string) *CreateBandwidthResourcePackagesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetPromotionId(v string) *CreateBandwidthResourcePackagesRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetRegionId(v string) *CreateBandwidthResourcePackagesRequest {
	s.RegionId = &v
	return s
}

type CreateBandwidthResourcePackagesResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBandwidthResourcePackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBandwidthResourcePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBandwidthResourcePackagesResponseBody) SetOrderId(v int64) *CreateBandwidthResourcePackagesResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateBandwidthResourcePackagesResponseBody) SetRequestId(v string) *CreateBandwidthResourcePackagesResponseBody {
	s.RequestId = &v
	return s
}

type CreateBandwidthResourcePackagesResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBandwidthResourcePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBandwidthResourcePackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBandwidthResourcePackagesResponse) GoString() string {
	return s.String()
}

func (s *CreateBandwidthResourcePackagesResponse) SetHeaders(v map[string]*string) *CreateBandwidthResourcePackagesResponse {
	s.Headers = v
	return s
}

func (s *CreateBandwidthResourcePackagesResponse) SetBody(v *CreateBandwidthResourcePackagesResponseBody) *CreateBandwidthResourcePackagesResponse {
	s.Body = v
	return s
}

type CreateBundleRequest struct {
	BundleName               *string  `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Description              *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopType              *string  `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	ImageId                  *string  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Language                 *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	RegionId                 *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RootDiskPerformanceLevel *string  `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	RootDiskSizeGib          *int32   `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskPerformanceLevel *string  `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	UserDiskSizeGib          []*int32 `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty" type:"Repeated"`
}

func (s CreateBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleRequest) GoString() string {
	return s.String()
}

func (s *CreateBundleRequest) SetBundleName(v string) *CreateBundleRequest {
	s.BundleName = &v
	return s
}

func (s *CreateBundleRequest) SetDescription(v string) *CreateBundleRequest {
	s.Description = &v
	return s
}

func (s *CreateBundleRequest) SetDesktopType(v string) *CreateBundleRequest {
	s.DesktopType = &v
	return s
}

func (s *CreateBundleRequest) SetImageId(v string) *CreateBundleRequest {
	s.ImageId = &v
	return s
}

func (s *CreateBundleRequest) SetLanguage(v string) *CreateBundleRequest {
	s.Language = &v
	return s
}

func (s *CreateBundleRequest) SetRegionId(v string) *CreateBundleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBundleRequest) SetRootDiskPerformanceLevel(v string) *CreateBundleRequest {
	s.RootDiskPerformanceLevel = &v
	return s
}

func (s *CreateBundleRequest) SetRootDiskSizeGib(v int32) *CreateBundleRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *CreateBundleRequest) SetUserDiskPerformanceLevel(v string) *CreateBundleRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *CreateBundleRequest) SetUserDiskSizeGib(v []*int32) *CreateBundleRequest {
	s.UserDiskSizeGib = v
	return s
}

type CreateBundleResponseBody struct {
	BundleId  *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBundleResponseBody) SetBundleId(v string) *CreateBundleResponseBody {
	s.BundleId = &v
	return s
}

func (s *CreateBundleResponseBody) SetRequestId(v string) *CreateBundleResponseBody {
	s.RequestId = &v
	return s
}

type CreateBundleResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleResponse) GoString() string {
	return s.String()
}

func (s *CreateBundleResponse) SetHeaders(v map[string]*string) *CreateBundleResponse {
	s.Headers = v
	return s
}

func (s *CreateBundleResponse) SetBody(v *CreateBundleResponseBody) *CreateBundleResponse {
	s.Body = v
	return s
}

type CreateDesktopGroupRequest struct {
	AllowAutoSetup          *int32    `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	AllowBufferCount        *int32    `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	AutoPay                 *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BindAmount              *int64    `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	BundleId                *string   `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	ChargeType              *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken             *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Comments                *string   `json:"Comments,omitempty" xml:"Comments,omitempty"`
	DefaultInitDesktopCount *int32    `json:"DefaultInitDesktopCount,omitempty" xml:"DefaultInitDesktopCount,omitempty"`
	DesktopGroupName        *string   `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DirectoryId             *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserIds              []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	KeepDuration            *int64    `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	LoadPolicy              *int64    `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	MaxDesktopsCount        *int32    `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	MinDesktopsCount        *int32    `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	OfficeSiteId            *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OwnType                 *int32    `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	Period                  *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PolicyGroupId           *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResetType               *int64    `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	ScaleStrategyId         *string   `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	VpcId                   *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupRequest) SetAllowAutoSetup(v int32) *CreateDesktopGroupRequest {
	s.AllowAutoSetup = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAllowBufferCount(v int32) *CreateDesktopGroupRequest {
	s.AllowBufferCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAutoPay(v bool) *CreateDesktopGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetBindAmount(v int64) *CreateDesktopGroupRequest {
	s.BindAmount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetBundleId(v string) *CreateDesktopGroupRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetChargeType(v string) *CreateDesktopGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetClientToken(v string) *CreateDesktopGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetComments(v string) *CreateDesktopGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDefaultInitDesktopCount(v int32) *CreateDesktopGroupRequest {
	s.DefaultInitDesktopCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDesktopGroupName(v string) *CreateDesktopGroupRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDirectoryId(v string) *CreateDesktopGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetEndUserIds(v []*string) *CreateDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *CreateDesktopGroupRequest) SetKeepDuration(v int64) *CreateDesktopGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetLoadPolicy(v int64) *CreateDesktopGroupRequest {
	s.LoadPolicy = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMaxDesktopsCount(v int32) *CreateDesktopGroupRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMinDesktopsCount(v int32) *CreateDesktopGroupRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetOfficeSiteId(v string) *CreateDesktopGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetOwnType(v int32) *CreateDesktopGroupRequest {
	s.OwnType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPeriod(v int32) *CreateDesktopGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPeriodUnit(v string) *CreateDesktopGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPolicyGroupId(v string) *CreateDesktopGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetRegionId(v string) *CreateDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetResetType(v int64) *CreateDesktopGroupRequest {
	s.ResetType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetScaleStrategyId(v string) *CreateDesktopGroupRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetVpcId(v string) *CreateDesktopGroupRequest {
	s.VpcId = &v
	return s
}

type CreateDesktopGroupResponseBody struct {
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	OrderIds       []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupResponseBody) SetDesktopGroupId(v string) *CreateDesktopGroupResponseBody {
	s.DesktopGroupId = &v
	return s
}

func (s *CreateDesktopGroupResponseBody) SetOrderIds(v []*string) *CreateDesktopGroupResponseBody {
	s.OrderIds = v
	return s
}

func (s *CreateDesktopGroupResponseBody) SetRequestId(v string) *CreateDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateDesktopGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupResponse) SetHeaders(v map[string]*string) *CreateDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopGroupResponse) SetBody(v *CreateDesktopGroupResponseBody) *CreateDesktopGroupResponse {
	s.Body = v
	return s
}

type CreateDesktopsRequest struct {
	Amount            *int32                      `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AutoPay           *bool                       `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew         *bool                       `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BundleId          *string                     `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	ChargeType        *string                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DesktopName       *string                     `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopNameSuffix *bool                       `json:"DesktopNameSuffix,omitempty" xml:"DesktopNameSuffix,omitempty"`
	DirectoryId       *string                     `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId         []*string                   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	GroupId           *string                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Hostname          *string                     `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	OfficeSiteId      *string                     `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Period            *int32                      `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit        *string                     `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PolicyGroupId     *string                     `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PromotionId       *string                     `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId          *string                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag               []*CreateDesktopsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UserAssignMode    *string                     `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
	UserName          *string                     `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VpcId             *string                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequest) SetAmount(v int32) *CreateDesktopsRequest {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsRequest) SetAutoPay(v bool) *CreateDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopsRequest) SetAutoRenew(v bool) *CreateDesktopsRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDesktopsRequest) SetBundleId(v string) *CreateDesktopsRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsRequest) SetChargeType(v string) *CreateDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopName(v string) *CreateDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopNameSuffix(v bool) *CreateDesktopsRequest {
	s.DesktopNameSuffix = &v
	return s
}

func (s *CreateDesktopsRequest) SetDirectoryId(v string) *CreateDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopsRequest) SetEndUserId(v []*string) *CreateDesktopsRequest {
	s.EndUserId = v
	return s
}

func (s *CreateDesktopsRequest) SetGroupId(v string) *CreateDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetHostname(v string) *CreateDesktopsRequest {
	s.Hostname = &v
	return s
}

func (s *CreateDesktopsRequest) SetOfficeSiteId(v string) *CreateDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriod(v int32) *CreateDesktopsRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriodUnit(v string) *CreateDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopsRequest) SetPolicyGroupId(v string) *CreateDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetPromotionId(v string) *CreateDesktopsRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetRegionId(v string) *CreateDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetTag(v []*CreateDesktopsRequestTag) *CreateDesktopsRequest {
	s.Tag = v
	return s
}

func (s *CreateDesktopsRequest) SetUserAssignMode(v string) *CreateDesktopsRequest {
	s.UserAssignMode = &v
	return s
}

func (s *CreateDesktopsRequest) SetUserName(v string) *CreateDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *CreateDesktopsRequest) SetVpcId(v string) *CreateDesktopsRequest {
	s.VpcId = &v
	return s
}

type CreateDesktopsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDesktopsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestTag) SetKey(v string) *CreateDesktopsRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDesktopsRequestTag) SetValue(v string) *CreateDesktopsRequestTag {
	s.Value = &v
	return s
}

type CreateDesktopsResponseBody struct {
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	OrderId   *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponseBody) SetDesktopId(v []*string) *CreateDesktopsResponseBody {
	s.DesktopId = v
	return s
}

func (s *CreateDesktopsResponseBody) SetOrderId(v string) *CreateDesktopsResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDesktopsResponseBody) SetRequestId(v string) *CreateDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type CreateDesktopsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponse) SetHeaders(v map[string]*string) *CreateDesktopsResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopsResponse) SetBody(v *CreateDesktopsResponseBody) *CreateDesktopsResponse {
	s.Body = v
	return s
}

type CreateImageRequest struct {
	Description       *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId         *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	ImageName         *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageResourceType *string   `json:"ImageResourceType,omitempty" xml:"ImageResourceType,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId        *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotIds       []*string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty" type:"Repeated"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetDesktopId(v string) *CreateImageRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetImageResourceType(v string) *CreateImageRequest {
	s.ImageResourceType = &v
	return s
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotId(v string) *CreateImageRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotIds(v []*string) *CreateImageRequest {
	s.SnapshotIds = v
	return s
}

type CreateImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetHeaders(v map[string]*string) *CreateImageResponse {
	s.Headers = v
	return s
}

func (s *CreateImageResponse) SetBody(v *CreateImageResponseBody) *CreateImageResponse {
	s.Body = v
	return s
}

type CreateNASFileSystemRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateNASFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNASFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemRequest) SetDescription(v string) *CreateNASFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetName(v string) *CreateNASFileSystemRequest {
	s.Name = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetOfficeSiteId(v string) *CreateNASFileSystemRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetRegionId(v string) *CreateNASFileSystemRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetStorageType(v string) *CreateNASFileSystemRequest {
	s.StorageType = &v
	return s
}

type CreateNASFileSystemResponseBody struct {
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemName    *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNASFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNASFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemResponseBody) SetFileSystemId(v string) *CreateNASFileSystemResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetFileSystemName(v string) *CreateNASFileSystemResponseBody {
	s.FileSystemName = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetMountTargetDomain(v string) *CreateNASFileSystemResponseBody {
	s.MountTargetDomain = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetOfficeSiteId(v string) *CreateNASFileSystemResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetRequestId(v string) *CreateNASFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type CreateNASFileSystemResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNASFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNASFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNASFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemResponse) SetHeaders(v map[string]*string) *CreateNASFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CreateNASFileSystemResponse) SetBody(v *CreateNASFileSystemResponseBody) *CreateNASFileSystemResponse {
	s.Body = v
	return s
}

type CreateNetworkPackageRequest struct {
	AutoPay            *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew          *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OfficeSiteId       *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId        *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateNetworkPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageRequest) SetAutoPay(v bool) *CreateNetworkPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetAutoRenew(v bool) *CreateNetworkPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetBandwidth(v int32) *CreateNetworkPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetInternetChargeType(v string) *CreateNetworkPackageRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetOfficeSiteId(v string) *CreateNetworkPackageRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriod(v int32) *CreateNetworkPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriodUnit(v string) *CreateNetworkPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPromotionId(v string) *CreateNetworkPackageRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetRegionId(v string) *CreateNetworkPackageRequest {
	s.RegionId = &v
	return s
}

type CreateNetworkPackageResponseBody struct {
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageResponseBody) SetNetworkPackageId(v string) *CreateNetworkPackageResponseBody {
	s.NetworkPackageId = &v
	return s
}

func (s *CreateNetworkPackageResponseBody) SetOrderId(v string) *CreateNetworkPackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateNetworkPackageResponseBody) SetRequestId(v string) *CreateNetworkPackageResponseBody {
	s.RequestId = &v
	return s
}

type CreateNetworkPackageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNetworkPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNetworkPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageResponse) SetHeaders(v map[string]*string) *CreateNetworkPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkPackageResponse) SetBody(v *CreateNetworkPackageResponseBody) *CreateNetworkPackageResponse {
	s.Body = v
	return s
}

type CreateOrderForHardwareRequest struct {
	Amount          *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AutoPay         *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	CityCode        *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	CityName        *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	ContactName     *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CountryCode     *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	CountryName     *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	DetailAddress   *string `json:"DetailAddress,omitempty" xml:"DetailAddress,omitempty"`
	DistrictCode    *string `json:"DistrictCode,omitempty" xml:"DistrictCode,omitempty"`
	DistrictName    *string `json:"DistrictName,omitempty" xml:"DistrictName,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	HardwareType    *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	HardwareVersion *string `json:"HardwareVersion,omitempty" xml:"HardwareVersion,omitempty"`
	MobilePhone     *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	Phone           *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	PromotionId     *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	ProvCode        *string `json:"ProvCode,omitempty" xml:"ProvCode,omitempty"`
	ProvName        *string `json:"ProvName,omitempty" xml:"ProvName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StreetCode      *string `json:"StreetCode,omitempty" xml:"StreetCode,omitempty"`
	StreetName      *string `json:"StreetName,omitempty" xml:"StreetName,omitempty"`
	ZipCode         *string `json:"ZipCode,omitempty" xml:"ZipCode,omitempty"`
}

func (s CreateOrderForHardwareRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForHardwareRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderForHardwareRequest) SetAmount(v int32) *CreateOrderForHardwareRequest {
	s.Amount = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetAutoPay(v bool) *CreateOrderForHardwareRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetCityCode(v string) *CreateOrderForHardwareRequest {
	s.CityCode = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetCityName(v string) *CreateOrderForHardwareRequest {
	s.CityName = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetContactName(v string) *CreateOrderForHardwareRequest {
	s.ContactName = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetCountryCode(v string) *CreateOrderForHardwareRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetCountryName(v string) *CreateOrderForHardwareRequest {
	s.CountryName = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetDetailAddress(v string) *CreateOrderForHardwareRequest {
	s.DetailAddress = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetDistrictCode(v string) *CreateOrderForHardwareRequest {
	s.DistrictCode = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetDistrictName(v string) *CreateOrderForHardwareRequest {
	s.DistrictName = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetEmail(v string) *CreateOrderForHardwareRequest {
	s.Email = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetHardwareType(v string) *CreateOrderForHardwareRequest {
	s.HardwareType = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetHardwareVersion(v string) *CreateOrderForHardwareRequest {
	s.HardwareVersion = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetMobilePhone(v string) *CreateOrderForHardwareRequest {
	s.MobilePhone = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetPhone(v string) *CreateOrderForHardwareRequest {
	s.Phone = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetPromotionId(v string) *CreateOrderForHardwareRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetProvCode(v string) *CreateOrderForHardwareRequest {
	s.ProvCode = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetProvName(v string) *CreateOrderForHardwareRequest {
	s.ProvName = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetRegionId(v string) *CreateOrderForHardwareRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetStreetCode(v string) *CreateOrderForHardwareRequest {
	s.StreetCode = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetStreetName(v string) *CreateOrderForHardwareRequest {
	s.StreetName = &v
	return s
}

func (s *CreateOrderForHardwareRequest) SetZipCode(v string) *CreateOrderForHardwareRequest {
	s.ZipCode = &v
	return s
}

type CreateOrderForHardwareResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrderForHardwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForHardwareResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderForHardwareResponseBody) SetOrderId(v int64) *CreateOrderForHardwareResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateOrderForHardwareResponseBody) SetRequestId(v string) *CreateOrderForHardwareResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrderForHardwareResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrderForHardwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderForHardwareResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderForHardwareResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderForHardwareResponse) SetHeaders(v map[string]*string) *CreateOrderForHardwareResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderForHardwareResponse) SetBody(v *CreateOrderForHardwareResponseBody) *CreateOrderForHardwareResponse {
	s.Body = v
	return s
}

type CreatePolicyGroupRequest struct {
	AuthorizeAccessPolicyRule   []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule   `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	AuthorizeSecurityPolicyRule []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	ClientType                  []*CreatePolicyGroupRequestClientType                  `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	Clipboard                   *string                                                `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	DomainList                  *string                                                `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	GpuAcceleration             *string                                                `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	Html5Access                 *string                                                `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	Html5FileTransfer           *string                                                `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	LocalDrive                  *string                                                `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	Name                        *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PreemptLogin                *string                                                `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	PreemptLoginUser            []*string                                              `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	PrinterRedirection          *string                                                `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	RegionId                    *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UsbRedirect                 *string                                                `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	UsbSupplyRedirectRule       []*CreatePolicyGroupRequestUsbSupplyRedirectRule       `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	VisualQuality               *string                                                `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	Watermark                   *string                                                `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	WatermarkCustomText         *string                                                `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency       *string                                                `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	WatermarkType               *string                                                `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s CreatePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequest) SetAuthorizeAccessPolicyRule(v []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule) *CreatePolicyGroupRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetAuthorizeSecurityPolicyRule(v []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) *CreatePolicyGroupRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetClientType(v []*CreatePolicyGroupRequestClientType) *CreatePolicyGroupRequest {
	s.ClientType = v
	return s
}

func (s *CreatePolicyGroupRequest) SetClipboard(v string) *CreatePolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetDomainList(v string) *CreatePolicyGroupRequest {
	s.DomainList = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetGpuAcceleration(v string) *CreatePolicyGroupRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5Access(v string) *CreatePolicyGroupRequest {
	s.Html5Access = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5FileTransfer(v string) *CreatePolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLocalDrive(v string) *CreatePolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetName(v string) *CreatePolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPreemptLogin(v string) *CreatePolicyGroupRequest {
	s.PreemptLogin = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPreemptLoginUser(v []*string) *CreatePolicyGroupRequest {
	s.PreemptLoginUser = v
	return s
}

func (s *CreatePolicyGroupRequest) SetPrinterRedirection(v string) *CreatePolicyGroupRequest {
	s.PrinterRedirection = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRegionId(v string) *CreatePolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetUsbRedirect(v string) *CreatePolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetUsbSupplyRedirectRule(v []*CreatePolicyGroupRequestUsbSupplyRedirectRule) *CreatePolicyGroupRequest {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetVisualQuality(v string) *CreatePolicyGroupRequest {
	s.VisualQuality = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermark(v string) *CreatePolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkCustomText(v string) *CreatePolicyGroupRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkTransparency(v string) *CreatePolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkType(v string) *CreatePolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

type CreatePolicyGroupRequestAuthorizeAccessPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreatePolicyGroupRequestAuthorizeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) SetDescription(v string) *CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

type CreatePolicyGroupRequestAuthorizeSecurityPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetType(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

type CreatePolicyGroupRequestClientType struct {
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreatePolicyGroupRequestClientType) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestClientType) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestClientType) SetClientType(v string) *CreatePolicyGroupRequestClientType {
	s.ClientType = &v
	return s
}

func (s *CreatePolicyGroupRequestClientType) SetStatus(v string) *CreatePolicyGroupRequestClientType {
	s.Status = &v
	return s
}

type CreatePolicyGroupRequestUsbSupplyRedirectRule struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceClass     *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	DeviceSubclass  *string `json:"DeviceSubclass,omitempty" xml:"DeviceSubclass,omitempty"`
	ProductId       *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	UsbRedirectType *int64  `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	UsbRuleType     *int64  `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	VendorId        *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s CreatePolicyGroupRequestUsbSupplyRedirectRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetDescription(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetDeviceClass(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.DeviceClass = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetDeviceSubclass(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.DeviceSubclass = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetProductId(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetUsbRedirectType(v int64) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetUsbRuleType(v int64) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetVendorId(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

type CreatePolicyGroupResponseBody struct {
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponseBody) SetPolicyGroupId(v string) *CreatePolicyGroupResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *CreatePolicyGroupResponseBody) SetRequestId(v string) *CreatePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponse) SetHeaders(v map[string]*string) *CreatePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyGroupResponse) SetBody(v *CreatePolicyGroupResponseBody) *CreatePolicyGroupResponse {
	s.Body = v
	return s
}

type CreateRAMDirectoryRequest struct {
	DesktopAccessType    *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DirectoryName        *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	EnableAdminAccess    *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId            []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s CreateRAMDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryRequest) SetDesktopAccessType(v string) *CreateRAMDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetDirectoryName(v string) *CreateRAMDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableAdminAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableInternetAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetRegionId(v string) *CreateRAMDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetVSwitchId(v []*string) *CreateRAMDirectoryRequest {
	s.VSwitchId = v
	return s
}

type CreateRAMDirectoryResponseBody struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRAMDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryResponseBody) SetDirectoryId(v string) *CreateRAMDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateRAMDirectoryResponseBody) SetRequestId(v string) *CreateRAMDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateRAMDirectoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRAMDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRAMDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryResponse) SetHeaders(v map[string]*string) *CreateRAMDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateRAMDirectoryResponse) SetBody(v *CreateRAMDirectoryResponseBody) *CreateRAMDirectoryResponse {
	s.Body = v
	return s
}

type CreateSimpleOfficeSiteRequest struct {
	Bandwidth            *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CenId                *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId           *int64    `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	CidrBlock            *string   `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CloudBoxOfficeSite   *bool     `json:"CloudBoxOfficeSite,omitempty" xml:"CloudBoxOfficeSite,omitempty"`
	DesktopAccessType    *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	EnableAdminAccess    *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	NeedVerifyZeroDevice *bool     `json:"NeedVerifyZeroDevice,omitempty" xml:"NeedVerifyZeroDevice,omitempty"`
	OfficeSiteName       *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId            []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	VerifyCode           *string   `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s CreateSimpleOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSimpleOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteRequest) SetBandwidth(v int32) *CreateSimpleOfficeSiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCenId(v string) *CreateSimpleOfficeSiteRequest {
	s.CenId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCenOwnerId(v int64) *CreateSimpleOfficeSiteRequest {
	s.CenOwnerId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCidrBlock(v string) *CreateSimpleOfficeSiteRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCloudBoxOfficeSite(v bool) *CreateSimpleOfficeSiteRequest {
	s.CloudBoxOfficeSite = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetDesktopAccessType(v string) *CreateSimpleOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEnableAdminAccess(v bool) *CreateSimpleOfficeSiteRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEnableInternetAccess(v bool) *CreateSimpleOfficeSiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetNeedVerifyZeroDevice(v bool) *CreateSimpleOfficeSiteRequest {
	s.NeedVerifyZeroDevice = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetOfficeSiteName(v string) *CreateSimpleOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetRegionId(v string) *CreateSimpleOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetVSwitchId(v []*string) *CreateSimpleOfficeSiteRequest {
	s.VSwitchId = v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetVerifyCode(v string) *CreateSimpleOfficeSiteRequest {
	s.VerifyCode = &v
	return s
}

type CreateSimpleOfficeSiteResponseBody struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSimpleOfficeSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSimpleOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteResponseBody) SetOfficeSiteId(v string) *CreateSimpleOfficeSiteResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateSimpleOfficeSiteResponseBody) SetRequestId(v string) *CreateSimpleOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

type CreateSimpleOfficeSiteResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSimpleOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSimpleOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSimpleOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteResponse) SetHeaders(v map[string]*string) *CreateSimpleOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *CreateSimpleOfficeSiteResponse) SetBody(v *CreateSimpleOfficeSiteResponseBody) *CreateSimpleOfficeSiteResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetDesktopId(v string) *CreateSnapshotRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetSourceDiskType(v string) *CreateSnapshotRequest {
	s.SourceDiskType = &v
	return s
}

type CreateSnapshotResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type CreateUserTagsRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags     []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserTagsRequest) GoString() string {
	return s.String()
}

func (s *CreateUserTagsRequest) SetRegionId(v string) *CreateUserTagsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUserTagsRequest) SetTags(v []*string) *CreateUserTagsRequest {
	s.Tags = v
	return s
}

type CreateUserTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserTagsResponseBody) SetRequestId(v string) *CreateUserTagsResponseBody {
	s.RequestId = &v
	return s
}

type CreateUserTagsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserTagsResponse) GoString() string {
	return s.String()
}

func (s *CreateUserTagsResponse) SetHeaders(v map[string]*string) *CreateUserTagsResponse {
	s.Headers = v
	return s
}

func (s *CreateUserTagsResponse) SetBody(v *CreateUserTagsResponseBody) *CreateUserTagsResponse {
	s.Body = v
	return s
}

type DeleteBundlesRequest struct {
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesRequest) GoString() string {
	return s.String()
}

func (s *DeleteBundlesRequest) SetBundleId(v []*string) *DeleteBundlesRequest {
	s.BundleId = v
	return s
}

func (s *DeleteBundlesRequest) SetRegionId(v string) *DeleteBundlesRequest {
	s.RegionId = &v
	return s
}

type DeleteBundlesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBundlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBundlesResponseBody) SetRequestId(v string) *DeleteBundlesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBundlesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBundlesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesResponse) GoString() string {
	return s.String()
}

func (s *DeleteBundlesResponse) SetHeaders(v map[string]*string) *DeleteBundlesResponse {
	s.Headers = v
	return s
}

func (s *DeleteBundlesResponse) SetBody(v *DeleteBundlesResponseBody) *DeleteBundlesResponse {
	s.Body = v
	return s
}

type DeleteDesktopGroupRequest struct {
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupRequest) SetDesktopGroupId(v string) *DeleteDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DeleteDesktopGroupRequest) SetRegionId(v string) *DeleteDesktopGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupResponseBody) SetRequestId(v string) *DeleteDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDesktopGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupResponse) SetHeaders(v map[string]*string) *DeleteDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDesktopGroupResponse) SetBody(v *DeleteDesktopGroupResponseBody) *DeleteDesktopGroupResponse {
	s.Body = v
	return s
}

type DeleteDesktopsRequest struct {
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsRequest) SetDesktopId(v []*string) *DeleteDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DeleteDesktopsRequest) SetRegionId(v string) *DeleteDesktopsRequest {
	s.RegionId = &v
	return s
}

type DeleteDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsResponseBody) SetRequestId(v string) *DeleteDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDesktopsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsResponse) SetHeaders(v map[string]*string) *DeleteDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDesktopsResponse) SetBody(v *DeleteDesktopsResponseBody) *DeleteDesktopsResponse {
	s.Body = v
	return s
}

type DeleteDirectoriesRequest struct {
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesRequest) SetDirectoryId(v []*string) *DeleteDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DeleteDirectoriesRequest) SetRegionId(v string) *DeleteDirectoriesRequest {
	s.RegionId = &v
	return s
}

type DeleteDirectoriesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponseBody) SetRequestId(v string) *DeleteDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDirectoriesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponse) SetHeaders(v map[string]*string) *DeleteDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectoriesResponse) SetBody(v *DeleteDirectoriesResponseBody) *DeleteDirectoriesResponse {
	s.Body = v
	return s
}

type DeleteImagesRequest struct {
	ImageId  []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesRequest) SetImageId(v []*string) *DeleteImagesRequest {
	s.ImageId = v
	return s
}

func (s *DeleteImagesRequest) SetRegionId(v string) *DeleteImagesRequest {
	s.RegionId = &v
	return s
}

type DeleteImagesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBody) SetRequestId(v string) *DeleteImagesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImagesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponse) SetHeaders(v map[string]*string) *DeleteImagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteImagesResponse) SetBody(v *DeleteImagesResponseBody) *DeleteImagesResponse {
	s.Body = v
	return s
}

type DeleteNASFileSystemsRequest struct {
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNASFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNASFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsRequest) SetFileSystemId(v []*string) *DeleteNASFileSystemsRequest {
	s.FileSystemId = v
	return s
}

func (s *DeleteNASFileSystemsRequest) SetRegionId(v string) *DeleteNASFileSystemsRequest {
	s.RegionId = &v
	return s
}

type DeleteNASFileSystemsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNASFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNASFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsResponseBody) SetRequestId(v string) *DeleteNASFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNASFileSystemsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNASFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNASFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNASFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsResponse) SetHeaders(v map[string]*string) *DeleteNASFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DeleteNASFileSystemsResponse) SetBody(v *DeleteNASFileSystemsResponseBody) *DeleteNASFileSystemsResponse {
	s.Body = v
	return s
}

type DeleteNetworkPackagesRequest struct {
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DeleteNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *DeleteNetworkPackagesRequest) SetRegionId(v string) *DeleteNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

type DeleteNetworkPackagesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesResponseBody) SetRequestId(v string) *DeleteNetworkPackagesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNetworkPackagesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNetworkPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNetworkPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesResponse) SetHeaders(v map[string]*string) *DeleteNetworkPackagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkPackagesResponse) SetBody(v *DeleteNetworkPackagesResponseBody) *DeleteNetworkPackagesResponse {
	s.Body = v
	return s
}

type DeleteOfficeSitesRequest struct {
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteOfficeSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesRequest) SetOfficeSiteId(v []*string) *DeleteOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

func (s *DeleteOfficeSitesRequest) SetRegionId(v string) *DeleteOfficeSitesRequest {
	s.RegionId = &v
	return s
}

type DeleteOfficeSitesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOfficeSitesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesResponseBody) SetRequestId(v string) *DeleteOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOfficeSitesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOfficeSitesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOfficeSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesResponse) SetHeaders(v map[string]*string) *DeleteOfficeSitesResponse {
	s.Headers = v
	return s
}

func (s *DeleteOfficeSitesResponse) SetBody(v *DeleteOfficeSitesResponseBody) *DeleteOfficeSitesResponse {
	s.Body = v
	return s
}

type DeletePolicyGroupsRequest struct {
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DeletePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

func (s *DeletePolicyGroupsRequest) SetRegionId(v string) *DeletePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

type DeletePolicyGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsResponseBody) SetRequestId(v string) *DeletePolicyGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePolicyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsResponse) SetHeaders(v map[string]*string) *DeletePolicyGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyGroupsResponse) SetBody(v *DeletePolicyGroupsResponseBody) *DeletePolicyGroupsResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId []*string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" type:"Repeated"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v []*string) *DeleteSnapshotRequest {
	s.SnapshotId = v
	return s
}

type DeleteSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DeleteUserTagsRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags     []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DeleteUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserTagsRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserTagsRequest) SetRegionId(v string) *DeleteUserTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserTagsRequest) SetTags(v []*string) *DeleteUserTagsRequest {
	s.Tags = v
	return s
}

type DeleteUserTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserTagsResponseBody) SetRequestId(v string) *DeleteUserTagsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserTagsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserTagsResponse) SetHeaders(v map[string]*string) *DeleteUserTagsResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserTagsResponse) SetBody(v *DeleteUserTagsResponseBody) *DeleteUserTagsResponse {
	s.Body = v
	return s
}

type DeleteVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) SetRegionId(v string) *DeleteVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type DeleteVirtualMFADeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponseBody) SetRequestId(v string) *DeleteVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVirtualMFADeviceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *DeleteVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) SetBody(v *DeleteVirtualMFADeviceResponseBody) *DeleteVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type DescribeAlarmEventStackInfoRequest struct {
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EventName  *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
}

func (s DescribeAlarmEventStackInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventStackInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoRequest) SetDesktopId(v string) *DescribeAlarmEventStackInfoRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetEventName(v string) *DescribeAlarmEventStackInfoRequest {
	s.EventName = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetLang(v string) *DescribeAlarmEventStackInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetRegionId(v string) *DescribeAlarmEventStackInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetUniqueInfo(v string) *DescribeAlarmEventStackInfoRequest {
	s.UniqueInfo = &v
	return s
}

type DescribeAlarmEventStackInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackInfo *string `json:"StackInfo,omitempty" xml:"StackInfo,omitempty"`
}

func (s DescribeAlarmEventStackInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventStackInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoResponseBody) SetRequestId(v string) *DescribeAlarmEventStackInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoResponseBody) SetStackInfo(v string) *DescribeAlarmEventStackInfoResponseBody {
	s.StackInfo = &v
	return s
}

type DescribeAlarmEventStackInfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlarmEventStackInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlarmEventStackInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventStackInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoResponse) SetHeaders(v map[string]*string) *DescribeAlarmEventStackInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmEventStackInfoResponse) SetBody(v *DescribeAlarmEventStackInfoResponseBody) *DescribeAlarmEventStackInfoResponse {
	s.Body = v
	return s
}

type DescribeBundlesRequest struct {
	BundleId          []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
	BundleType        *string   `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	CheckStock        *bool     `json:"CheckStock,omitempty" xml:"CheckStock,omitempty"`
	CpuCount          *int32    `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	DesktopTypeFamily *string   `json:"DesktopTypeFamily,omitempty" xml:"DesktopTypeFamily,omitempty"`
	FromDesktopGroup  *bool     `json:"FromDesktopGroup,omitempty" xml:"FromDesktopGroup,omitempty"`
	GpuCount          *float32  `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	MaxResults        *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	MemorySize        *int32    `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	NextToken         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProtocolType      *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBundlesRequest) SetBundleId(v []*string) *DescribeBundlesRequest {
	s.BundleId = v
	return s
}

func (s *DescribeBundlesRequest) SetBundleType(v string) *DescribeBundlesRequest {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesRequest) SetCheckStock(v bool) *DescribeBundlesRequest {
	s.CheckStock = &v
	return s
}

func (s *DescribeBundlesRequest) SetCpuCount(v int32) *DescribeBundlesRequest {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesRequest) SetDesktopTypeFamily(v string) *DescribeBundlesRequest {
	s.DesktopTypeFamily = &v
	return s
}

func (s *DescribeBundlesRequest) SetFromDesktopGroup(v bool) *DescribeBundlesRequest {
	s.FromDesktopGroup = &v
	return s
}

func (s *DescribeBundlesRequest) SetGpuCount(v float32) *DescribeBundlesRequest {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesRequest) SetMaxResults(v int32) *DescribeBundlesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeBundlesRequest) SetMemorySize(v int32) *DescribeBundlesRequest {
	s.MemorySize = &v
	return s
}

func (s *DescribeBundlesRequest) SetNextToken(v string) *DescribeBundlesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesRequest) SetProtocolType(v string) *DescribeBundlesRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeBundlesRequest) SetRegionId(v string) *DescribeBundlesRequest {
	s.RegionId = &v
	return s
}

type DescribeBundlesResponseBody struct {
	Bundles   []*DescribeBundlesResponseBodyBundles `json:"Bundles,omitempty" xml:"Bundles,omitempty" type:"Repeated"`
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBundlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBody) SetBundles(v []*DescribeBundlesResponseBodyBundles) *DescribeBundlesResponseBody {
	s.Bundles = v
	return s
}

func (s *DescribeBundlesResponseBody) SetNextToken(v string) *DescribeBundlesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesResponseBody) SetRequestId(v string) *DescribeBundlesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBundlesResponseBodyBundles struct {
	BundleId             *string                                                 `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	BundleName           *string                                                 `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	BundleType           *string                                                 `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	CreationTime         *string                                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description          *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopType          *string                                                 `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	DesktopTypeAttribute *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute `json:"DesktopTypeAttribute,omitempty" xml:"DesktopTypeAttribute,omitempty" type:"Struct"`
	DesktopTypeFamily    *string                                                 `json:"DesktopTypeFamily,omitempty" xml:"DesktopTypeFamily,omitempty"`
	Disks                []*DescribeBundlesResponseBodyBundlesDisks              `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	ImageId              *string                                                 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName            *string                                                 `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Language             *string                                                 `json:"Language,omitempty" xml:"Language,omitempty"`
	OsType               *string                                                 `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ProtocolType         *string                                                 `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	SessionType          *string                                                 `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	StockState           *string                                                 `json:"StockState,omitempty" xml:"StockState,omitempty"`
}

func (s DescribeBundlesResponseBodyBundles) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundles) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleId(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleId = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleName(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleName = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleType(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetCreationTime(v string) *DescribeBundlesResponseBodyBundles {
	s.CreationTime = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDescription(v string) *DescribeBundlesResponseBodyBundles {
	s.Description = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopType(v string) *DescribeBundlesResponseBodyBundles {
	s.DesktopType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopTypeAttribute(v *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) *DescribeBundlesResponseBodyBundles {
	s.DesktopTypeAttribute = v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopTypeFamily(v string) *DescribeBundlesResponseBodyBundles {
	s.DesktopTypeFamily = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDisks(v []*DescribeBundlesResponseBodyBundlesDisks) *DescribeBundlesResponseBodyBundles {
	s.Disks = v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetImageId(v string) *DescribeBundlesResponseBodyBundles {
	s.ImageId = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetImageName(v string) *DescribeBundlesResponseBodyBundles {
	s.ImageName = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetLanguage(v string) *DescribeBundlesResponseBodyBundles {
	s.Language = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetOsType(v string) *DescribeBundlesResponseBodyBundles {
	s.OsType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetProtocolType(v string) *DescribeBundlesResponseBodyBundles {
	s.ProtocolType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetSessionType(v string) *DescribeBundlesResponseBodyBundles {
	s.SessionType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetStockState(v string) *DescribeBundlesResponseBodyBundles {
	s.StockState = &v
	return s
}

type DescribeBundlesResponseBodyBundlesDesktopTypeAttribute struct {
	CpuCount   *int32   `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	GpuCount   *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuSpec    *string  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	MemorySize *int32   `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetCpuCount(v int32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetGpuCount(v float32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetGpuSpec(v string) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.GpuSpec = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetMemorySize(v int32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.MemorySize = &v
	return s
}

type DescribeBundlesResponseBodyBundlesDisks struct {
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	DiskSize             *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType             *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeBundlesResponseBodyBundlesDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundlesDisks) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskPerformanceLevel(v string) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskSize(v int32) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskType(v string) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskType = &v
	return s
}

type DescribeBundlesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBundlesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponse) SetHeaders(v map[string]*string) *DescribeBundlesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBundlesResponse) SetBody(v *DescribeBundlesResponseBody) *DescribeBundlesResponse {
	s.Body = v
	return s
}

type DescribeCensRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCensRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) SetPageNumber(v int32) *DescribeCensRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensRequest) SetPageSize(v int32) *DescribeCensRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCensRequest) SetRegionId(v string) *DescribeCensRequest {
	s.RegionId = &v
	return s
}

type DescribeCensResponseBody struct {
	Cens       []*DescribeCensResponseBodyCens `json:"Cens,omitempty" xml:"Cens,omitempty" type:"Repeated"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBody) SetCens(v []*DescribeCensResponseBodyCens) *DescribeCensResponseBody {
	s.Cens = v
	return s
}

func (s *DescribeCensResponseBody) SetPageNumber(v int32) *DescribeCensResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensResponseBody) SetPageSize(v int32) *DescribeCensResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCensResponseBody) SetRequestId(v string) *DescribeCensResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCensResponseBody) SetTotalCount(v int32) *DescribeCensResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCensResponseBodyCens struct {
	CenId           *string                                   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CreationTime    *string                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description     *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Ipv6Level       *string                                   `json:"Ipv6Level,omitempty" xml:"Ipv6Level,omitempty"`
	Name            *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PackageIds      []*DescribeCensResponseBodyCensPackageIds `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	ProtectionLevel *string                                   `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	Status          *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            []*DescribeCensResponseBodyCensTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCens) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCens) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCens) SetCenId(v string) *DescribeCensResponseBodyCens {
	s.CenId = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetCreationTime(v string) *DescribeCensResponseBodyCens {
	s.CreationTime = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetDescription(v string) *DescribeCensResponseBodyCens {
	s.Description = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetIpv6Level(v string) *DescribeCensResponseBodyCens {
	s.Ipv6Level = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetName(v string) *DescribeCensResponseBodyCens {
	s.Name = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetPackageIds(v []*DescribeCensResponseBodyCensPackageIds) *DescribeCensResponseBodyCens {
	s.PackageIds = v
	return s
}

func (s *DescribeCensResponseBodyCens) SetProtectionLevel(v string) *DescribeCensResponseBodyCens {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetStatus(v string) *DescribeCensResponseBodyCens {
	s.Status = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetTags(v []*DescribeCensResponseBodyCensTags) *DescribeCensResponseBodyCens {
	s.Tags = v
	return s
}

type DescribeCensResponseBodyCensPackageIds struct {
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
}

func (s DescribeCensResponseBodyCensPackageIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensPackageIds) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensPackageIds) SetPackageId(v string) *DescribeCensResponseBodyCensPackageIds {
	s.PackageId = &v
	return s
}

type DescribeCensResponseBodyCensTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCensResponseBodyCensTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensTags) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensTags) SetKey(v string) *DescribeCensResponseBodyCensTags {
	s.Key = &v
	return s
}

func (s *DescribeCensResponseBodyCensTags) SetValue(v string) *DescribeCensResponseBodyCensTags {
	s.Value = &v
	return s
}

type DescribeCensResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCensResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCensResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponse) GoString() string {
	return s.String()
}

func (s *DescribeCensResponse) SetHeaders(v map[string]*string) *DescribeCensResponse {
	s.Headers = v
	return s
}

func (s *DescribeCensResponse) SetBody(v *DescribeCensResponseBody) *DescribeCensResponse {
	s.Body = v
	return s
}

type DescribeClientEventsRequest struct {
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopIp      *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	DesktopName    *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DirectoryId    *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EventType      *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId   *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClientEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsRequest) SetDesktopId(v string) *DescribeClientEventsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopIp(v string) *DescribeClientEventsRequest {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopName(v string) *DescribeClientEventsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDirectoryId(v string) *DescribeClientEventsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndTime(v string) *DescribeClientEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndUserId(v string) *DescribeClientEventsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEventType(v string) *DescribeClientEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsRequest) SetMaxResults(v int32) *DescribeClientEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeClientEventsRequest) SetNextToken(v string) *DescribeClientEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeClientEventsRequest) SetOfficeSiteId(v string) *DescribeClientEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetOfficeSiteName(v string) *DescribeClientEventsRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeClientEventsRequest) SetRegionId(v string) *DescribeClientEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetStartTime(v string) *DescribeClientEventsRequest {
	s.StartTime = &v
	return s
}

type DescribeClientEventsResponseBody struct {
	Events    []*DescribeClientEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	NextToken *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseBody) SetEvents(v []*DescribeClientEventsResponseBodyEvents) *DescribeClientEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeClientEventsResponseBody) SetNextToken(v string) *DescribeClientEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeClientEventsResponseBody) SetRequestId(v string) *DescribeClientEventsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClientEventsResponseBodyEvents struct {
	AliUid           *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BytesReceived    *string `json:"BytesReceived,omitempty" xml:"BytesReceived,omitempty"`
	BytesSend        *string `json:"BytesSend,omitempty" xml:"BytesSend,omitempty"`
	ClientIp         *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS         *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion    *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	DesktopGroupId   *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopIp        *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	DesktopName      *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryType    *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EventId          *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventTime        *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	EventType        *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	OfficeSiteId     *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName   *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	OfficeSiteType   *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClientEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseBodyEvents) SetAliUid(v string) *DescribeClientEventsResponseBodyEvents {
	s.AliUid = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetBytesReceived(v string) *DescribeClientEventsResponseBodyEvents {
	s.BytesReceived = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetBytesSend(v string) *DescribeClientEventsResponseBodyEvents {
	s.BytesSend = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientIp(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientIp = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientOS(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientOS = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientVersion(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopGroupId(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopGroupName(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopId(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopIp(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopName(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopName = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDirectoryId(v string) *DescribeClientEventsResponseBodyEvents {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDirectoryType(v string) *DescribeClientEventsResponseBodyEvents {
	s.DirectoryType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEndUserId(v string) *DescribeClientEventsResponseBodyEvents {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventId(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventTime(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventTime = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventType(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetOfficeSiteId(v string) *DescribeClientEventsResponseBodyEvents {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetOfficeSiteName(v string) *DescribeClientEventsResponseBodyEvents {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetOfficeSiteType(v string) *DescribeClientEventsResponseBodyEvents {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetRegionId(v string) *DescribeClientEventsResponseBodyEvents {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetStatus(v string) *DescribeClientEventsResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeClientEventsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClientEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClientEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponse) SetHeaders(v map[string]*string) *DescribeClientEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientEventsResponse) SetBody(v *DescribeClientEventsResponseBody) *DescribeClientEventsResponse {
	s.Body = v
	return s
}

type DescribeDesktopGroupsRequest struct {
	DesktopGroupId     *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopGroupName   *string   `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	EndUserIds         []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	ExcludedEndUserIds []*string `json:"ExcludedEndUserIds,omitempty" xml:"ExcludedEndUserIds,omitempty" type:"Repeated"`
	MaxResults         *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId       *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OwnType            *int64    `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	Period             *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PolicyGroupId      *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDesktopGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsRequest) SetDesktopGroupId(v string) *DescribeDesktopGroupsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetDesktopGroupName(v string) *DescribeDesktopGroupsRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetEndUserIds(v []*string) *DescribeDesktopGroupsRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetExcludedEndUserIds(v []*string) *DescribeDesktopGroupsRequest {
	s.ExcludedEndUserIds = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetMaxResults(v int32) *DescribeDesktopGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetNextToken(v string) *DescribeDesktopGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetOfficeSiteId(v string) *DescribeDesktopGroupsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetOwnType(v int64) *DescribeDesktopGroupsRequest {
	s.OwnType = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetPeriod(v int32) *DescribeDesktopGroupsRequest {
	s.Period = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetPeriodUnit(v string) *DescribeDesktopGroupsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetPolicyGroupId(v string) *DescribeDesktopGroupsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetRegionId(v string) *DescribeDesktopGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetStatus(v int32) *DescribeDesktopGroupsRequest {
	s.Status = &v
	return s
}

type DescribeDesktopGroupsResponseBody struct {
	DesktopGroups []*DescribeDesktopGroupsResponseBodyDesktopGroups `json:"DesktopGroups,omitempty" xml:"DesktopGroups,omitempty" type:"Repeated"`
	NextToken     *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponseBody) SetDesktopGroups(v []*DescribeDesktopGroupsResponseBodyDesktopGroups) *DescribeDesktopGroupsResponseBody {
	s.DesktopGroups = v
	return s
}

func (s *DescribeDesktopGroupsResponseBody) SetNextToken(v string) *DescribeDesktopGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBody) SetRequestId(v string) *DescribeDesktopGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDesktopGroupsResponseBodyDesktopGroups struct {
	BindAmount         *int64   `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	Comments           *string  `json:"Comments,omitempty" xml:"Comments,omitempty"`
	Cpu                *int32   `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreateTime         *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator            *string  `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DataDiskCategory   *string  `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskSize       *string  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopGroupId     *string  `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopGroupName   *string  `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DirectoryId        *string  `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryType      *string  `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	EndUserCount       *int32   `json:"EndUserCount,omitempty" xml:"EndUserCount,omitempty"`
	ExpiredTime        *string  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GpuCount           *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuSpec            *string  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	ImageId            *string  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	KeepDuration       *int64   `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	LoadPolicy         *int64   `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	MaxDesktopsCount   *int32   `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	Memory             *int64   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	MinDesktopsCount   *int32   `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	OfficeSiteId       *string  `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName     *string  `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	OfficeSiteType     *string  `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	OwnBundleId        *string  `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	OwnBundleName      *string  `json:"OwnBundleName,omitempty" xml:"OwnBundleName,omitempty"`
	OwnType            *int64   `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	PayType            *string  `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PolicyGroupId      *string  `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PolicyGroupName    *string  `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	ResetType          *int64   `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	Status             *int32   `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemDiskCategory *string  `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskSize     *int32   `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopGroupsResponseBodyDesktopGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopGroupsResponseBodyDesktopGroups) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetBindAmount(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.BindAmount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetComments(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Comments = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetCpu(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetCreateTime(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetCreator(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Creator = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDataDiskCategory(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDataDiskSize(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDesktopGroupId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDesktopGroupName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDirectoryId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDirectoryType(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetEndUserCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.EndUserCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetExpiredTime(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetGpuCount(v float32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetGpuSpec(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetImageId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetKeepDuration(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.KeepDuration = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetLoadPolicy(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.LoadPolicy = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetMaxDesktopsCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.MaxDesktopsCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetMemory(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetMinDesktopsCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.MinDesktopsCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOfficeSiteId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOfficeSiteName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOfficeSiteType(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOwnBundleId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OwnBundleId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOwnBundleName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OwnBundleName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOwnType(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OwnType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetPayType(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.PayType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetPolicyGroupId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetPolicyGroupName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.PolicyGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetResetType(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ResetType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetStatus(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Status = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetSystemDiskCategory(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetSystemDiskSize(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.SystemDiskSize = &v
	return s
}

type DescribeDesktopGroupsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponse) SetHeaders(v map[string]*string) *DescribeDesktopGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopGroupsResponse) SetBody(v *DescribeDesktopGroupsResponseBody) *DescribeDesktopGroupsResponse {
	s.Body = v
	return s
}

type DescribeDesktopIdsByVulNamesRequest struct {
	Necessity    *string   `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type         *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	VulName      []*string `json:"VulName,omitempty" xml:"VulName,omitempty" type:"Repeated"`
}

func (s DescribeDesktopIdsByVulNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetNecessity(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetOfficeSiteId(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetRegionId(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetType(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetVulName(v []*string) *DescribeDesktopIdsByVulNamesRequest {
	s.VulName = v
	return s
}

type DescribeDesktopIdsByVulNamesResponseBody struct {
	DesktopItems []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems `json:"DesktopItems,omitempty" xml:"DesktopItems,omitempty" type:"Repeated"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopIdsByVulNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) SetDesktopItems(v []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) *DescribeDesktopIdsByVulNamesResponseBody {
	s.DesktopItems = v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) SetRequestId(v string) *DescribeDesktopIdsByVulNamesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDesktopIdsByVulNamesResponseBodyDesktopItems struct {
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
}

func (s DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) SetDesktopId(v string) *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) SetDesktopName(v string) *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems {
	s.DesktopName = &v
	return s
}

type DescribeDesktopIdsByVulNamesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopIdsByVulNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopIdsByVulNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetHeaders(v map[string]*string) *DescribeDesktopIdsByVulNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetBody(v *DescribeDesktopIdsByVulNamesResponseBody) *DescribeDesktopIdsByVulNamesResponse {
	s.Body = v
	return s
}

type DescribeDesktopTypesRequest struct {
	CpuCount           *int32   `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	DesktopIdForModify *string  `json:"DesktopIdForModify,omitempty" xml:"DesktopIdForModify,omitempty"`
	DesktopTypeId      *string  `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	GpuCount           *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	InstanceTypeFamily *string  `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	MemorySize         *int32   `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	OrderType          *string  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	RegionId           *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDesktopTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesRequest) SetCpuCount(v int32) *DescribeDesktopTypesRequest {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopIdForModify(v string) *DescribeDesktopTypesRequest {
	s.DesktopIdForModify = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopTypeId(v string) *DescribeDesktopTypesRequest {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetGpuCount(v float32) *DescribeDesktopTypesRequest {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetInstanceTypeFamily(v string) *DescribeDesktopTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetMemorySize(v int32) *DescribeDesktopTypesRequest {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetOrderType(v string) *DescribeDesktopTypesRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetRegionId(v string) *DescribeDesktopTypesRequest {
	s.RegionId = &v
	return s
}

type DescribeDesktopTypesResponseBody struct {
	DesktopTypes []*DescribeDesktopTypesResponseBodyDesktopTypes `json:"DesktopTypes,omitempty" xml:"DesktopTypes,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBody) SetDesktopTypes(v []*DescribeDesktopTypesResponseBodyDesktopTypes) *DescribeDesktopTypesResponseBody {
	s.DesktopTypes = v
	return s
}

func (s *DescribeDesktopTypesResponseBody) SetRequestId(v string) *DescribeDesktopTypesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDesktopTypesResponseBodyDesktopTypes struct {
	AllowDiskSize      []*DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize `json:"AllowDiskSize,omitempty" xml:"AllowDiskSize,omitempty" type:"Repeated"`
	CpuCount           *string                                                      `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	DataDiskSize       *string                                                      `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopTypeId      *string                                                      `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	DesktopTypeStatus  *string                                                      `json:"DesktopTypeStatus,omitempty" xml:"DesktopTypeStatus,omitempty"`
	GpuCount           *float32                                                     `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuSpec            *string                                                      `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	InstanceTypeFamily *string                                                      `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	MemorySize         *string                                                      `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	SystemDiskSize     *string                                                      `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopTypesResponseBodyDesktopTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseBodyDesktopTypes) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetAllowDiskSize(v []*DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.AllowDiskSize = v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetCpuCount(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDataDiskSize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDesktopTypeId(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDesktopTypeStatus(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.DesktopTypeStatus = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetGpuCount(v float32) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetGpuSpec(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetInstanceTypeFamily(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetMemorySize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetSystemDiskSize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.SystemDiskSize = &v
	return s
}

type DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize struct {
	DataDiskSize        *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DefaultDataDiskSize *int32 `json:"DefaultDataDiskSize,omitempty" xml:"DefaultDataDiskSize,omitempty"`
	SystemDiskSize      *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) SetDataDiskSize(v int32) *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) SetDefaultDataDiskSize(v int32) *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize {
	s.DefaultDataDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) SetSystemDiskSize(v int32) *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize {
	s.SystemDiskSize = &v
	return s
}

type DescribeDesktopTypesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponse) SetHeaders(v map[string]*string) *DescribeDesktopTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopTypesResponse) SetBody(v *DescribeDesktopTypesResponseBody) *DescribeDesktopTypesResponse {
	s.Body = v
	return s
}

type DescribeDesktopsRequest struct {
	ChargeType         *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DesktopId          []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	DesktopName        *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus      *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DirectoryId        *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId          []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	ExcludedEndUserId  []*string `json:"ExcludedEndUserId,omitempty" xml:"ExcludedEndUserId,omitempty" type:"Repeated"`
	ExpiredTime        *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FilterDesktopGroup *bool     `json:"FilterDesktopGroup,omitempty" xml:"FilterDesktopGroup,omitempty"`
	GroupId            *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ManagementFlag     *string   `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	MaxResults         *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId       *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName     *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	PolicyGroupId      *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	ProtocolType       *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	QueryFotaUpdate    *bool     `json:"QueryFotaUpdate,omitempty" xml:"QueryFotaUpdate,omitempty"`
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserName           *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsRequest) SetChargeType(v string) *DescribeDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopId(v []*string) *DescribeDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopName(v string) *DescribeDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopStatus(v string) *DescribeDesktopsRequest {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDirectoryId(v string) *DescribeDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetEndUserId(v []*string) *DescribeDesktopsRequest {
	s.EndUserId = v
	return s
}

func (s *DescribeDesktopsRequest) SetExcludedEndUserId(v []*string) *DescribeDesktopsRequest {
	s.ExcludedEndUserId = v
	return s
}

func (s *DescribeDesktopsRequest) SetExpiredTime(v string) *DescribeDesktopsRequest {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsRequest) SetFilterDesktopGroup(v bool) *DescribeDesktopsRequest {
	s.FilterDesktopGroup = &v
	return s
}

func (s *DescribeDesktopsRequest) SetGroupId(v string) *DescribeDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetManagementFlag(v string) *DescribeDesktopsRequest {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsRequest) SetMaxResults(v int32) *DescribeDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsRequest) SetNextToken(v string) *DescribeDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOfficeSiteId(v string) *DescribeDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOfficeSiteName(v string) *DescribeDesktopsRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetPolicyGroupId(v string) *DescribeDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetProtocolType(v string) *DescribeDesktopsRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetQueryFotaUpdate(v bool) *DescribeDesktopsRequest {
	s.QueryFotaUpdate = &v
	return s
}

func (s *DescribeDesktopsRequest) SetRegionId(v string) *DescribeDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetUserName(v string) *DescribeDesktopsRequest {
	s.UserName = &v
	return s
}

type DescribeDesktopsResponseBody struct {
	Desktops   []*DescribeDesktopsResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBody) SetDesktops(v []*DescribeDesktopsResponseBodyDesktops) *DescribeDesktopsResponseBody {
	s.Desktops = v
	return s
}

func (s *DescribeDesktopsResponseBody) SetNextToken(v string) *DescribeDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetRequestId(v string) *DescribeDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetTotalCount(v int32) *DescribeDesktopsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDesktopsResponseBodyDesktops struct {
	BundleId           *string                                         `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	BundleName         *string                                         `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	ChargeType         *string                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ConnectionStatus   *string                                         `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	Cpu                *int32                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime       *string                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataDiskCategory   *string                                         `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskSize       *string                                         `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopGroupId     *string                                         `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopId          *string                                         `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName        *string                                         `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus      *string                                         `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DesktopType        *string                                         `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	DirectoryId        *string                                         `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryType      *string                                         `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	Disks              []*DescribeDesktopsResponseBodyDesktopsDisks    `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	DowngradeQuota     *int64                                          `json:"DowngradeQuota,omitempty" xml:"DowngradeQuota,omitempty"`
	DowngradedTimes    *int64                                          `json:"DowngradedTimes,omitempty" xml:"DowngradedTimes,omitempty"`
	EndUserIds         []*string                                       `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	ExpiredTime        *string                                         `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FotaUpdate         *DescribeDesktopsResponseBodyDesktopsFotaUpdate `json:"FotaUpdate,omitempty" xml:"FotaUpdate,omitempty" type:"Struct"`
	GpuCategory        *int64                                          `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	GpuCount           *float32                                        `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuDriverVersion   *string                                         `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	GpuSpec            *string                                         `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	HostName           *string                                         `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId            *string                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ManagementFlag     *string                                         `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	Memory             *int64                                          `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetworkInterfaceId *string                                         `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	NetworkInterfaceIp *string                                         `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	OfficeSiteId       *string                                         `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName     *string                                         `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	OfficeSiteType     *string                                         `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	OsType             *string                                         `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PolicyGroupId      *string                                         `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PolicyGroupName    *string                                         `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	Progress           *string                                         `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ProtocolType       *string                                         `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Sessions           []*DescribeDesktopsResponseBodyDesktopsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	StartTime          *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SystemDiskCategory *string                                         `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskSize     *int32                                          `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	Tags               []*DescribeDesktopsResponseBodyDesktopsTags     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneType           *string                                         `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktops) SetBundleId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.BundleId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetBundleName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.BundleName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetChargeType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetConnectionStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCpu(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCreationTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDataDiskCategory(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDataDiskSize(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopGroupId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDirectoryId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDirectoryType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDisks(v []*DescribeDesktopsResponseBodyDesktopsDisks) *DescribeDesktopsResponseBodyDesktops {
	s.Disks = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDowngradeQuota(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.DowngradeQuota = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDowngradedTimes(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.DowngradedTimes = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetEndUserIds(v []*string) *DescribeDesktopsResponseBodyDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetExpiredTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetFotaUpdate(v *DescribeDesktopsResponseBodyDesktopsFotaUpdate) *DescribeDesktopsResponseBodyDesktops {
	s.FotaUpdate = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuCategory(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.GpuCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuCount(v float32) *DescribeDesktopsResponseBodyDesktops {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuDriverVersion(v string) *DescribeDesktopsResponseBodyDesktops {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuSpec(v string) *DescribeDesktopsResponseBodyDesktops {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetHostName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.HostName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetImageId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetManagementFlag(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetMemory(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetNetworkInterfaceId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetNetworkInterfaceIp(v string) *DescribeDesktopsResponseBodyDesktops {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOsType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetProgress(v string) *DescribeDesktopsResponseBodyDesktops {
	s.Progress = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetProtocolType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSessions(v []*DescribeDesktopsResponseBodyDesktopsSessions) *DescribeDesktopsResponseBodyDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetStartTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSystemDiskCategory(v string) *DescribeDesktopsResponseBodyDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSystemDiskSize(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetTags(v []*DescribeDesktopsResponseBodyDesktopsTags) *DescribeDesktopsResponseBodyDesktops {
	s.Tags = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetZoneType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ZoneType = &v
	return s
}

type DescribeDesktopsResponseBodyDesktopsDisks struct {
	DiskId           *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize         *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType         *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskId(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskSize(v int32) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskType(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetPerformanceLevel(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.PerformanceLevel = &v
	return s
}

type DescribeDesktopsResponseBodyDesktopsFotaUpdate struct {
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	NewAppVersion     *string `json:"NewAppVersion,omitempty" xml:"NewAppVersion,omitempty"`
	NewTaskUid        *string `json:"NewTaskUid,omitempty" xml:"NewTaskUid,omitempty"`
	ReleaseNote       *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Size              *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsFotaUpdate) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsFotaUpdate) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetCurrentAppVersion(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetNewAppVersion(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.NewAppVersion = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetNewTaskUid(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.NewTaskUid = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNote(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetSize(v int64) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.Size = &v
	return s
}

type DescribeDesktopsResponseBodyDesktopsSessions struct {
	EndUserId         *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty"`
	ExternalUserName  *string `json:"ExternalUserName,omitempty" xml:"ExternalUserName,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsSessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) SetEndUserId(v string) *DescribeDesktopsResponseBodyDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) SetEstablishmentTime(v string) *DescribeDesktopsResponseBodyDesktopsSessions {
	s.EstablishmentTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) SetExternalUserName(v string) *DescribeDesktopsResponseBodyDesktopsSessions {
	s.ExternalUserName = &v
	return s
}

type DescribeDesktopsResponseBodyDesktopsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsTags) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) SetKey(v string) *DescribeDesktopsResponseBodyDesktopsTags {
	s.Key = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) SetValue(v string) *DescribeDesktopsResponseBodyDesktopsTags {
	s.Value = &v
	return s
}

type DescribeDesktopsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponse) SetHeaders(v map[string]*string) *DescribeDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopsResponse) SetBody(v *DescribeDesktopsResponseBody) *DescribeDesktopsResponse {
	s.Body = v
	return s
}

type DescribeDesktopsInGroupRequest struct {
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PayType        *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDesktopsInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupRequest) SetDesktopGroupId(v string) *DescribeDesktopsInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetMaxResults(v int32) *DescribeDesktopsInGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetNextToken(v string) *DescribeDesktopsInGroupRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetPayType(v string) *DescribeDesktopsInGroupRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetRegionId(v string) *DescribeDesktopsInGroupRequest {
	s.RegionId = &v
	return s
}

type DescribeDesktopsInGroupResponseBody struct {
	NextToken                   *string                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OnlinePrePaidDesktopsCount  *int32                                                 `json:"OnlinePrePaidDesktopsCount,omitempty" xml:"OnlinePrePaidDesktopsCount,omitempty"`
	PaidDesktops                []*DescribeDesktopsInGroupResponseBodyPaidDesktops     `json:"PaidDesktops,omitempty" xml:"PaidDesktops,omitempty" type:"Repeated"`
	PaidDesktopsCount           *int32                                                 `json:"PaidDesktopsCount,omitempty" xml:"PaidDesktopsCount,omitempty"`
	PostPaidDesktops            []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops `json:"PostPaidDesktops,omitempty" xml:"PostPaidDesktops,omitempty" type:"Repeated"`
	PostPaidDesktopsCount       *int32                                                 `json:"PostPaidDesktopsCount,omitempty" xml:"PostPaidDesktopsCount,omitempty"`
	PostPaidDesktopsTotalAmount *int32                                                 `json:"PostPaidDesktopsTotalAmount,omitempty" xml:"PostPaidDesktopsTotalAmount,omitempty"`
	RequestId                   *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunningPrePaidDesktopsCount *int32                                                 `json:"RunningPrePaidDesktopsCount,omitempty" xml:"RunningPrePaidDesktopsCount,omitempty"`
	StopedPrePaidDesktopsCount  *int32                                                 `json:"StopedPrePaidDesktopsCount,omitempty" xml:"StopedPrePaidDesktopsCount,omitempty"`
}

func (s DescribeDesktopsInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBody) SetNextToken(v string) *DescribeDesktopsInGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetOnlinePrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.OnlinePrePaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPaidDesktops(v []*DescribeDesktopsInGroupResponseBodyPaidDesktops) *DescribeDesktopsInGroupResponseBody {
	s.PaidDesktops = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktops(v []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktops = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktopsTotalAmount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktopsTotalAmount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetRequestId(v string) *DescribeDesktopsInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetRunningPrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.RunningPrePaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetStopedPrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.StopedPrePaidDesktopsCount = &v
	return s
}

type DescribeDesktopsInGroupResponseBodyPaidDesktops struct {
	ConnectionStatus *string   `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	DesktopId        *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName      *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus    *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DiskType         *string   `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	EndUserId        *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserIds       []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	EndUserName      *string   `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
	EndUserNames     []*string `json:"EndUserNames,omitempty" xml:"EndUserNames,omitempty" type:"Repeated"`
	GpuDriverVersion *string   `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	ImageId          *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName        *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ManagementFlag   *string   `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	OsType           *string   `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ResetTime        *string   `json:"ResetTime,omitempty" xml:"ResetTime,omitempty"`
	SystemDiskSize   *int32    `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopsInGroupResponseBodyPaidDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBodyPaidDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetConnectionStatus(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopId(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopName(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopStatus(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDiskType(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DiskType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserId(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserIds(v []*string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserName(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserNames(v []*string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserNames = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetGpuDriverVersion(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetImageId(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetImageName(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ImageName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetManagementFlag(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetOsType(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetResetTime(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ResetTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetSystemDiskSize(v int32) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.SystemDiskSize = &v
	return s
}

type DescribeDesktopsInGroupResponseBodyPostPaidDesktops struct {
	ConnectionStatus *string   `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	CreateDuration   *string   `json:"CreateDuration,omitempty" xml:"CreateDuration,omitempty"`
	CreateTime       *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DesktopId        *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName      *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus    *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DiskType         *string   `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	EndUserId        *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserIds       []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	EndUserName      *string   `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
	EndUserNames     []*string `json:"EndUserNames,omitempty" xml:"EndUserNames,omitempty" type:"Repeated"`
	GpuDriverVersion *string   `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	ImageId          *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName        *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ManagementFlag   *string   `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	OsType           *string   `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ReleaseTime      *string   `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	ResetTime        *string   `json:"ResetTime,omitempty" xml:"ResetTime,omitempty"`
	SystemDiskSize   *int32    `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopsInGroupResponseBodyPostPaidDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetConnectionStatus(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetCreateDuration(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.CreateDuration = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetCreateTime(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.CreateTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopId(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopName(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopStatus(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDiskType(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DiskType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserId(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserIds(v []*string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserName(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserNames(v []*string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserNames = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetGpuDriverVersion(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetImageId(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetImageName(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ImageName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetManagementFlag(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetOsType(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetReleaseTime(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetResetTime(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ResetTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetSystemDiskSize(v int32) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.SystemDiskSize = &v
	return s
}

type DescribeDesktopsInGroupResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopsInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopsInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponse) SetHeaders(v map[string]*string) *DescribeDesktopsInGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetBody(v *DescribeDesktopsInGroupResponseBody) *DescribeDesktopsInGroupResponse {
	s.Body = v
	return s
}

type DescribeDirectoriesRequest struct {
	DirectoryId     []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	DirectoryStatus *string   `json:"DirectoryStatus,omitempty" xml:"DirectoryStatus,omitempty"`
	DirectoryType   *string   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	MaxResults      *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryStatus(v string) *DescribeDirectoriesRequest {
	s.DirectoryStatus = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryType(v string) *DescribeDirectoriesRequest {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetMaxResults(v int32) *DescribeDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNextToken(v string) *DescribeDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetStatus(v string) *DescribeDirectoriesRequest {
	s.Status = &v
	return s
}

type DescribeDirectoriesResponseBody struct {
	AdHostname  *string                                       `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) SetAdHostname(v string) *DescribeDirectoriesResponseBody {
	s.AdHostname = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetNextToken(v string) *DescribeDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectories struct {
	ADConnectors             []*DescribeDirectoriesResponseBodyDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" type:"Repeated"`
	CreationTime             *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CustomSecurityGroupId    *string                                                   `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty"`
	DesktopAccessType        *string                                                   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DesktopVpcEndpoint       *string                                                   `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	DirectoryId              *string                                                   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryType            *string                                                   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	DnsAddress               []*string                                                 `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DnsUserName              *string                                                   `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty"`
	DomainName               *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainPassword           *string                                                   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DomainUserName           *string                                                   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	EnableAdminAccess        *bool                                                     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	EnableCrossDesktopAccess *bool                                                     `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
	EnableInternetAccess     *bool                                                     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	FileSystemIds            []*string                                                 `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Repeated"`
	Logs                     []*DescribeDirectoriesResponseBodyDirectoriesLogs         `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	MfaEnabled               *bool                                                     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	Name                     *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedVerifyLoginRisk      *bool                                                     `json:"NeedVerifyLoginRisk,omitempty" xml:"NeedVerifyLoginRisk,omitempty"`
	OuName                   *string                                                   `json:"OuName,omitempty" xml:"OuName,omitempty"`
	SsoEnabled               *bool                                                     `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	Status                   *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	SubDnsAddress            []*string                                                 `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" type:"Repeated"`
	SubDomainName            *string                                                   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	TrustPassword            *string                                                   `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
	VSwitchIds               []*string                                                 `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId                    *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetADConnectors(v []*DescribeDirectoriesResponseBodyDirectoriesADConnectors) *DescribeDirectoriesResponseBodyDirectories {
	s.ADConnectors = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCreationTime(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CreationTime = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCustomSecurityGroupId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopVpcEndpoint(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsAddress(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableAdminAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableCrossDesktopAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableInternetAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetFileSystemIds(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.FileSystemIds = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetLogs(v []*DescribeDirectoriesResponseBodyDirectoriesLogs) *DescribeDirectoriesResponseBodyDirectories {
	s.Logs = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetMfaEnabled(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetNeedVerifyLoginRisk(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.NeedVerifyLoginRisk = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetOuName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.OuName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSsoEnabled(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetStatus(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Status = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSubDnsAddress(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSubDomainName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.SubDomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetTrustPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.TrustPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetVSwitchIds(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetVpcId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.VpcId = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectoriesADConnectors struct {
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty"`
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	Specification      *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	TrustKey           *string `json:"TrustKey,omitempty" xml:"TrustKey,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectoriesADConnectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectoriesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetADConnectorAddress(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetConnectorStatus(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetNetworkInterfaceId(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetSpecification(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.Specification = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetTrustKey(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.TrustKey = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetVSwitchId(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.VSwitchId = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectoriesLogs struct {
	Level     *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Step      *string `json:"Step,omitempty" xml:"Step,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectoriesLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectoriesLogs) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetLevel(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Level = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetMessage(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Message = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetStep(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Step = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetTimeStamp(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.TimeStamp = &v
	return s
}

type DescribeDirectoriesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) SetHeaders(v map[string]*string) *DescribeDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoriesResponse) SetBody(v *DescribeDirectoriesResponseBody) *DescribeDirectoriesResponse {
	s.Body = v
	return s
}

type DescribeFlowMetricRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MetricType   *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Period       *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFlowMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricRequest) SetEndTime(v string) *DescribeFlowMetricRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetInstanceId(v string) *DescribeFlowMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetInstanceType(v string) *DescribeFlowMetricRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetMetricType(v string) *DescribeFlowMetricRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetPeriod(v int32) *DescribeFlowMetricRequest {
	s.Period = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetRegionId(v string) *DescribeFlowMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetStartTime(v string) *DescribeFlowMetricRequest {
	s.StartTime = &v
	return s
}

type DescribeFlowMetricResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricResponseBody) SetData(v string) *DescribeFlowMetricResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeFlowMetricResponseBody) SetRequestId(v string) *DescribeFlowMetricResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFlowMetricResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricResponse) SetHeaders(v map[string]*string) *DescribeFlowMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowMetricResponse) SetBody(v *DescribeFlowMetricResponseBody) *DescribeFlowMetricResponse {
	s.Body = v
	return s
}

type DescribeFotaPendingDesktopsRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskUid    *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s DescribeFotaPendingDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFotaPendingDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsRequest) SetMaxResults(v int32) *DescribeFotaPendingDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetNextToken(v string) *DescribeFotaPendingDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetRegionId(v string) *DescribeFotaPendingDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetTaskUid(v string) *DescribeFotaPendingDesktopsRequest {
	s.TaskUid = &v
	return s
}

type DescribeFotaPendingDesktopsResponseBody struct {
	FotaPendingDesktops []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops `json:"FotaPendingDesktops,omitempty" xml:"FotaPendingDesktops,omitempty" type:"Repeated"`
	NextToken           *string                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFotaPendingDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFotaPendingDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsResponseBody) SetFotaPendingDesktops(v []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) *DescribeFotaPendingDesktopsResponseBody {
	s.FotaPendingDesktops = v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBody) SetNextToken(v string) *DescribeFotaPendingDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBody) SetRequestId(v string) *DescribeFotaPendingDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops struct {
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	DesktopId         *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName       *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	FotaProject       *string `json:"FotaProject,omitempty" xml:"FotaProject,omitempty"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetCurrentAppVersion(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetDesktopId(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetDesktopName(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetFotaProject(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.FotaProject = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetOfficeSiteId(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.OfficeSiteId = &v
	return s
}

type DescribeFotaPendingDesktopsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFotaPendingDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFotaPendingDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFotaPendingDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsResponse) SetHeaders(v map[string]*string) *DescribeFotaPendingDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFotaPendingDesktopsResponse) SetBody(v *DescribeFotaPendingDesktopsResponseBody) *DescribeFotaPendingDesktopsResponse {
	s.Body = v
	return s
}

type DescribeFotaTasksRequest struct {
	FotaStatus *string   `json:"FotaStatus,omitempty" xml:"FotaStatus,omitempty"`
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskUid    []*string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty" type:"Repeated"`
	UserStatus *string   `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeFotaTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFotaTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeFotaTasksRequest) SetFotaStatus(v string) *DescribeFotaTasksRequest {
	s.FotaStatus = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetMaxResults(v int32) *DescribeFotaTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetNextToken(v string) *DescribeFotaTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetRegionId(v string) *DescribeFotaTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetTaskUid(v []*string) *DescribeFotaTasksRequest {
	s.TaskUid = v
	return s
}

func (s *DescribeFotaTasksRequest) SetUserStatus(v string) *DescribeFotaTasksRequest {
	s.UserStatus = &v
	return s
}

type DescribeFotaTasksResponseBody struct {
	FotaTasks []*DescribeFotaTasksResponseBodyFotaTasks `json:"FotaTasks,omitempty" xml:"FotaTasks,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFotaTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFotaTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFotaTasksResponseBody) SetFotaTasks(v []*DescribeFotaTasksResponseBodyFotaTasks) *DescribeFotaTasksResponseBody {
	s.FotaTasks = v
	return s
}

func (s *DescribeFotaTasksResponseBody) SetRequestId(v string) *DescribeFotaTasksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFotaTasksResponseBodyFotaTasks struct {
	AppVersion          *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	FotaProject         *string `json:"FotaProject,omitempty" xml:"FotaProject,omitempty"`
	PendingDesktopCount *int32  `json:"PendingDesktopCount,omitempty" xml:"PendingDesktopCount,omitempty"`
	PublishTime         *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	ReleaseNote         *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Size                *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskUid             *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s DescribeFotaTasksResponseBodyFotaTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeFotaTasksResponseBodyFotaTasks) GoString() string {
	return s.String()
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetAppVersion(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.AppVersion = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetFotaProject(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.FotaProject = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetPendingDesktopCount(v int32) *DescribeFotaTasksResponseBodyFotaTasks {
	s.PendingDesktopCount = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetPublishTime(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.PublishTime = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetReleaseNote(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetSize(v int32) *DescribeFotaTasksResponseBodyFotaTasks {
	s.Size = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetStatus(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.Status = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetTaskUid(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.TaskUid = &v
	return s
}

type DescribeFotaTasksResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFotaTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFotaTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFotaTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeFotaTasksResponse) SetHeaders(v map[string]*string) *DescribeFotaTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeFotaTasksResponse) SetBody(v *DescribeFotaTasksResponseBody) *DescribeFotaTasksResponse {
	s.Body = v
	return s
}

type DescribeFrontVulPatchListRequest struct {
	OperateType *string                                    `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	RegionId    *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type        *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	VulInfo     []*DescribeFrontVulPatchListRequestVulInfo `json:"VulInfo,omitempty" xml:"VulInfo,omitempty" type:"Repeated"`
}

func (s DescribeFrontVulPatchListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListRequest) SetOperateType(v string) *DescribeFrontVulPatchListRequest {
	s.OperateType = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetRegionId(v string) *DescribeFrontVulPatchListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetType(v string) *DescribeFrontVulPatchListRequest {
	s.Type = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetVulInfo(v []*DescribeFrontVulPatchListRequestVulInfo) *DescribeFrontVulPatchListRequest {
	s.VulInfo = v
	return s
}

type DescribeFrontVulPatchListRequestVulInfo struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeFrontVulPatchListRequestVulInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListRequestVulInfo) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetDesktopId(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.DesktopId = &v
	return s
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetName(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.Name = &v
	return s
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetTag(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.Tag = &v
	return s
}

type DescribeFrontVulPatchListResponseBody struct {
	FrontPatchList []*DescribeFrontVulPatchListResponseBodyFrontPatchList `json:"FrontPatchList,omitempty" xml:"FrontPatchList,omitempty" type:"Repeated"`
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFrontVulPatchListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBody) SetFrontPatchList(v []*DescribeFrontVulPatchListResponseBodyFrontPatchList) *DescribeFrontVulPatchListResponseBody {
	s.FrontPatchList = v
	return s
}

func (s *DescribeFrontVulPatchListResponseBody) SetRequestId(v string) *DescribeFrontVulPatchListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFrontVulPatchListResponseBodyFrontPatchList struct {
	DesktopId *string                                                         `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	PatchList []*DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList `json:"PatchList,omitempty" xml:"PatchList,omitempty" type:"Repeated"`
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchList) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) SetDesktopId(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchList {
	s.DesktopId = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) SetPatchList(v []*DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) *DescribeFrontVulPatchListResponseBodyFrontPatchList {
	s.PatchList = v
	return s
}

type DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) SetAliasName(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList {
	s.AliasName = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) SetName(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList {
	s.Name = &v
	return s
}

type DescribeFrontVulPatchListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFrontVulPatchListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFrontVulPatchListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponse) SetHeaders(v map[string]*string) *DescribeFrontVulPatchListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFrontVulPatchListResponse) SetBody(v *DescribeFrontVulPatchListResponseBody) *DescribeFrontVulPatchListResponse {
	s.Body = v
	return s
}

type DescribeGroupedVulRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Dealed       *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Necessity    *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeGroupedVulRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulRequest) SetCurrentPage(v int32) *DescribeGroupedVulRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetDealed(v string) *DescribeGroupedVulRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetLang(v string) *DescribeGroupedVulRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetNecessity(v string) *DescribeGroupedVulRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetOfficeSiteId(v string) *DescribeGroupedVulRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetPageSize(v int32) *DescribeGroupedVulRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetRegionId(v string) *DescribeGroupedVulRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetType(v string) *DescribeGroupedVulRequest {
	s.Type = &v
	return s
}

type DescribeGroupedVulResponseBody struct {
	CurrentPage     *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	GroupedVulItems []*DescribeGroupedVulResponseBodyGroupedVulItems `json:"GroupedVulItems,omitempty" xml:"GroupedVulItems,omitempty" type:"Repeated"`
	PageSize        *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupedVulResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseBody) SetCurrentPage(v int32) *DescribeGroupedVulResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetGroupedVulItems(v []*DescribeGroupedVulResponseBodyGroupedVulItems) *DescribeGroupedVulResponseBody {
	s.GroupedVulItems = v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetPageSize(v int32) *DescribeGroupedVulResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetRequestId(v string) *DescribeGroupedVulResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetTotalCount(v int32) *DescribeGroupedVulResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGroupedVulResponseBodyGroupedVulItems struct {
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	AsapCount    *int32  `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	GmtLast      *string `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
	HandledCount *int32  `json:"HandledCount,omitempty" xml:"HandledCount,omitempty"`
	LaterCount   *int32  `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NntfCount    *int32  `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeGroupedVulResponseBodyGroupedVulItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponseBodyGroupedVulItems) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetAliasName(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.AliasName = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetAsapCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.AsapCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetGmtLast(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.GmtLast = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetHandledCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.HandledCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetLaterCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.LaterCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetName(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Name = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetNntfCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.NntfCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetTags(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Tags = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetType(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Type = &v
	return s
}

type DescribeGroupedVulResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupedVulResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupedVulResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponse) SetHeaders(v map[string]*string) *DescribeGroupedVulResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedVulResponse) SetBody(v *DescribeGroupedVulResponseBody) *DescribeGroupedVulResponse {
	s.Body = v
	return s
}

type DescribeHardwareTerminalsRequest struct {
	HardwareType    *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	HardwareVersion *string `json:"HardwareVersion,omitempty" xml:"HardwareVersion,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeHardwareTerminalsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHardwareTerminalsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHardwareTerminalsRequest) SetHardwareType(v string) *DescribeHardwareTerminalsRequest {
	s.HardwareType = &v
	return s
}

func (s *DescribeHardwareTerminalsRequest) SetHardwareVersion(v string) *DescribeHardwareTerminalsRequest {
	s.HardwareVersion = &v
	return s
}

func (s *DescribeHardwareTerminalsRequest) SetRegionId(v string) *DescribeHardwareTerminalsRequest {
	s.RegionId = &v
	return s
}

type DescribeHardwareTerminalsResponseBody struct {
	HardwareTerminals []*DescribeHardwareTerminalsResponseBodyHardwareTerminals `json:"HardwareTerminals,omitempty" xml:"HardwareTerminals,omitempty" type:"Repeated"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHardwareTerminalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHardwareTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHardwareTerminalsResponseBody) SetHardwareTerminals(v []*DescribeHardwareTerminalsResponseBodyHardwareTerminals) *DescribeHardwareTerminalsResponseBody {
	s.HardwareTerminals = v
	return s
}

func (s *DescribeHardwareTerminalsResponseBody) SetRequestId(v string) *DescribeHardwareTerminalsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHardwareTerminalsResponseBodyHardwareTerminals struct {
	HardwareType        *string                                                                      `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	HardwareTypeDetails []*DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails `json:"HardwareTypeDetails,omitempty" xml:"HardwareTypeDetails,omitempty" type:"Repeated"`
}

func (s DescribeHardwareTerminalsResponseBodyHardwareTerminals) String() string {
	return tea.Prettify(s)
}

func (s DescribeHardwareTerminalsResponseBodyHardwareTerminals) GoString() string {
	return s.String()
}

func (s *DescribeHardwareTerminalsResponseBodyHardwareTerminals) SetHardwareType(v string) *DescribeHardwareTerminalsResponseBodyHardwareTerminals {
	s.HardwareType = &v
	return s
}

func (s *DescribeHardwareTerminalsResponseBodyHardwareTerminals) SetHardwareTypeDetails(v []*DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails) *DescribeHardwareTerminalsResponseBodyHardwareTerminals {
	s.HardwareTypeDetails = v
	return s
}

type DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HardwareVersion *int32  `json:"HardwareVersion,omitempty" xml:"HardwareVersion,omitempty"`
	StockState      *string `json:"StockState,omitempty" xml:"StockState,omitempty"`
}

func (s DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails) GoString() string {
	return s.String()
}

func (s *DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails) SetDescription(v string) *DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails {
	s.Description = &v
	return s
}

func (s *DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails) SetHardwareVersion(v int32) *DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails {
	s.HardwareVersion = &v
	return s
}

func (s *DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails) SetStockState(v string) *DescribeHardwareTerminalsResponseBodyHardwareTerminalsHardwareTypeDetails {
	s.StockState = &v
	return s
}

type DescribeHardwareTerminalsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHardwareTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHardwareTerminalsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHardwareTerminalsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHardwareTerminalsResponse) SetHeaders(v map[string]*string) *DescribeHardwareTerminalsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHardwareTerminalsResponse) SetBody(v *DescribeHardwareTerminalsResponseBody) *DescribeHardwareTerminalsResponse {
	s.Body = v
	return s
}

type DescribeImagesRequest struct {
	DesktopInstanceType *string   `json:"DesktopInstanceType,omitempty" xml:"DesktopInstanceType,omitempty"`
	GpuCategory         *bool     `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	GpuDriverVersion    *string   `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	ImageId             []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	ImageStatus         *string   `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty"`
	ImageType           *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	LanguageType        *string   `json:"LanguageType,omitempty" xml:"LanguageType,omitempty"`
	MaxResults          *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OsType              *string   `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ProtocolType        *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) SetDesktopInstanceType(v string) *DescribeImagesRequest {
	s.DesktopInstanceType = &v
	return s
}

func (s *DescribeImagesRequest) SetGpuCategory(v bool) *DescribeImagesRequest {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesRequest) SetGpuDriverVersion(v string) *DescribeImagesRequest {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v []*string) *DescribeImagesRequest {
	s.ImageId = v
	return s
}

func (s *DescribeImagesRequest) SetImageStatus(v string) *DescribeImagesRequest {
	s.ImageStatus = &v
	return s
}

func (s *DescribeImagesRequest) SetImageType(v string) *DescribeImagesRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesRequest) SetLanguageType(v string) *DescribeImagesRequest {
	s.LanguageType = &v
	return s
}

func (s *DescribeImagesRequest) SetMaxResults(v int32) *DescribeImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagesRequest) SetNextToken(v string) *DescribeImagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesRequest) SetOsType(v string) *DescribeImagesRequest {
	s.OsType = &v
	return s
}

func (s *DescribeImagesRequest) SetProtocolType(v string) *DescribeImagesRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeImagesRequest) SetRegionId(v string) *DescribeImagesRequest {
	s.RegionId = &v
	return s
}

type DescribeImagesResponseBody struct {
	Images    []*DescribeImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	NextToken *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBody) SetImages(v []*DescribeImagesResponseBodyImages) *DescribeImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeImagesResponseBody) SetNextToken(v string) *DescribeImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesResponseBody) SetRequestId(v string) *DescribeImagesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImagesResponseBodyImages struct {
	AppVersion         *string   `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CreationTime       *string   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataDiskSize       *int32    `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Description        *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	GpuCategory        *bool     `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	GpuDriverVersion   *string   `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	ImageId            *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType          *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Name               *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OsType             *string   `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Progress           *string   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ProtocolType       *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Size               *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	Status             *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedLanguages []*string `json:"SupportedLanguages,omitempty" xml:"SupportedLanguages,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImages) SetAppVersion(v string) *DescribeImagesResponseBodyImages {
	s.AppVersion = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetCreationTime(v string) *DescribeImagesResponseBodyImages {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetDataDiskSize(v int32) *DescribeImagesResponseBodyImages {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetDescription(v string) *DescribeImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetGpuCategory(v bool) *DescribeImagesResponseBodyImages {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetGpuDriverVersion(v string) *DescribeImagesResponseBodyImages {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageId(v string) *DescribeImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageType(v string) *DescribeImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetName(v string) *DescribeImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetOsType(v string) *DescribeImagesResponseBodyImages {
	s.OsType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProgress(v string) *DescribeImagesResponseBodyImages {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProtocolType(v string) *DescribeImagesResponseBodyImages {
	s.ProtocolType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetSize(v int32) *DescribeImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetStatus(v string) *DescribeImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetSupportedLanguages(v []*string) *DescribeImagesResponseBodyImages {
	s.SupportedLanguages = v
	return s
}

type DescribeImagesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetHeaders(v map[string]*string) *DescribeImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagesResponse) SetBody(v *DescribeImagesResponseBody) *DescribeImagesResponse {
	s.Body = v
	return s
}

type DescribeInvocationsRequest struct {
	CommandType     *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	IncludeOutput   *bool   `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	InvokeId        *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	InvokeStatus    *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetCommandType(v string) *DescribeInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetDesktopId(v string) *DescribeInvocationsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetMaxResults(v int32) *DescribeInvocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNextToken(v string) *DescribeInvocationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

type DescribeInvocationsResponseBody struct {
	Invocations []*DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Repeated"`
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v []*DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetNextToken(v string) *DescribeInvocationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInvocationsResponseBodyInvocations struct {
	CommandContent   *string                                                     `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandType      *string                                                     `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	CreationTime     *string                                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	InvocationStatus *string                                                     `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	InvokeDesktops   []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops `json:"InvokeDesktops,omitempty" xml:"InvokeDesktops,omitempty" type:"Repeated"`
	InvokeId         *string                                                     `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandType(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeDesktops(v []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeDesktops = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeId = &v
	return s
}

type DescribeInvocationsResponseBodyInvocationsInvokeDesktops struct {
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Dropped          *int32  `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo        *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ExitCode         *int64  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	Output           *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Repeats          *int32  `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StopTime         *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvokeDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetDesktopId(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetDropped(v int32) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetErrorCode(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetErrorInfo(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetExitCode(v int64) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetFinishTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetOutput(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetRepeats(v int32) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetStartTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetStopTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetUpdateTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.UpdateTime = &v
	return s
}

type DescribeInvocationsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetHeaders(v map[string]*string) *DescribeInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationsResponse) SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse {
	s.Body = v
	return s
}

type DescribeModificationPriceRequest struct {
	Bandwidth                *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InstanceId               *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType             *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PromotionId              *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType             *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	RootDiskSizeGib          *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	UserDiskSizeGib          *int32  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s DescribeModificationPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceRequest) SetBandwidth(v int32) *DescribeModificationPriceRequest {
	s.Bandwidth = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetInstanceId(v string) *DescribeModificationPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetInstanceType(v string) *DescribeModificationPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetPromotionId(v string) *DescribeModificationPriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetRegionId(v string) *DescribeModificationPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetResourceType(v string) *DescribeModificationPriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetRootDiskPerformanceLevel(v string) *DescribeModificationPriceRequest {
	s.RootDiskPerformanceLevel = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetRootDiskSizeGib(v int32) *DescribeModificationPriceRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetUserDiskPerformanceLevel(v string) *DescribeModificationPriceRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetUserDiskSizeGib(v int32) *DescribeModificationPriceRequest {
	s.UserDiskSizeGib = &v
	return s
}

type DescribeModificationPriceResponseBody struct {
	PriceInfo *DescribeModificationPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeModificationPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBody) SetPriceInfo(v *DescribeModificationPriceResponseBodyPriceInfo) *DescribeModificationPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeModificationPriceResponseBody) SetRequestId(v string) *DescribeModificationPriceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeModificationPriceResponseBodyPriceInfo struct {
	Price *DescribeModificationPriceResponseBodyPriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules []*DescribeModificationPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeModificationPriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) SetPrice(v *DescribeModificationPriceResponseBodyPriceInfoPrice) *DescribeModificationPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) SetRules(v []*DescribeModificationPriceResponseBodyPriceInfoRules) *DescribeModificationPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

type DescribeModificationPriceResponseBodyPriceInfoPrice struct {
	Currency      *string                                                          `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32                                                         `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32                                                         `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*DescribeModificationPriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	TradePrice    *float32                                                         `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribeModificationPriceResponseBodyPriceInfoPricePromotions) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

type DescribeModificationPriceResponseBodyPriceInfoPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

type DescribeModificationPriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribeModificationPriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribeModificationPriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type DescribeModificationPriceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeModificationPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeModificationPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponse) SetHeaders(v map[string]*string) *DescribeModificationPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeModificationPriceResponse) SetBody(v *DescribeModificationPriceResponseBody) *DescribeModificationPriceResponse {
	s.Body = v
	return s
}

type DescribeNASFileSystemsRequest struct {
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
	MaxResults   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNASFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsRequest) SetFileSystemId(v []*string) *DescribeNASFileSystemsRequest {
	s.FileSystemId = v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetMaxResults(v int32) *DescribeNASFileSystemsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetNextToken(v string) *DescribeNASFileSystemsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetOfficeSiteId(v string) *DescribeNASFileSystemsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetRegionId(v string) *DescribeNASFileSystemsRequest {
	s.RegionId = &v
	return s
}

type DescribeNASFileSystemsResponseBody struct {
	FileSystems []*DescribeNASFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
	NextToken   *string                                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNASFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBody) SetFileSystems(v []*DescribeNASFileSystemsResponseBodyFileSystems) *DescribeNASFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeNASFileSystemsResponseBody) SetNextToken(v string) *DescribeNASFileSystemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBody) SetRequestId(v string) *DescribeNASFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNASFileSystemsResponseBodyFileSystems struct {
	Capacity          *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemName    *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	FileSystemStatus  *string `json:"FileSystemStatus,omitempty" xml:"FileSystemStatus,omitempty"`
	FileSystemType    *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredSize       *int64  `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	MountTargetStatus *string `json:"MountTargetStatus,omitempty" xml:"MountTargetStatus,omitempty"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName    *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageType       *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SupportAcl        *bool   `json:"SupportAcl,omitempty" xml:"SupportAcl,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNASFileSystemsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetCapacity(v int64) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.Capacity = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetCreateTime(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.CreateTime = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetDescription(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.Description = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemName(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemStatus(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemStatus = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemType(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemType = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMeteredSize(v int64) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMountTargetDomain(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMountTargetStatus(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MountTargetStatus = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetOfficeSiteId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetOfficeSiteName(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetRegionId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.RegionId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetStorageType(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.StorageType = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetSupportAcl(v bool) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.SupportAcl = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetZoneId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.ZoneId = &v
	return s
}

type DescribeNASFileSystemsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNASFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNASFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponse) SetHeaders(v map[string]*string) *DescribeNASFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNASFileSystemsResponse) SetBody(v *DescribeNASFileSystemsResponseBody) *DescribeNASFileSystemsResponse {
	s.Body = v
	return s
}

type DescribeNetworkPackagesRequest struct {
	InternetChargeType *string   `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	MaxResults         *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NetworkPackageId   []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	NextToken          *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNetworkPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesRequest) SetInternetChargeType(v string) *DescribeNetworkPackagesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetMaxResults(v int32) *DescribeNetworkPackagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DescribeNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNextToken(v string) *DescribeNetworkPackagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetRegionId(v string) *DescribeNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

type DescribeNetworkPackagesResponseBody struct {
	NetworkPackages []*DescribeNetworkPackagesResponseBodyNetworkPackages `json:"NetworkPackages,omitempty" xml:"NetworkPackages,omitempty" type:"Repeated"`
	NextToken       *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponseBody) SetNetworkPackages(v []*DescribeNetworkPackagesResponseBodyNetworkPackages) *DescribeNetworkPackagesResponseBody {
	s.NetworkPackages = v
	return s
}

func (s *DescribeNetworkPackagesResponseBody) SetNextToken(v string) *DescribeNetworkPackagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBody) SetRequestId(v string) *DescribeNetworkPackagesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNetworkPackagesResponseBodyNetworkPackages struct {
	Bandwidth            *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CreateTime           *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EipAddresses         []*string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Repeated"`
	ExpiredTime          *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InternetChargeType   *string   `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	NetworkPackageId     *string   `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	NetworkPackageStatus *string   `json:"NetworkPackageStatus,omitempty" xml:"NetworkPackageStatus,omitempty"`
	OfficeSiteId         *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName       *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
}

func (s DescribeNetworkPackagesResponseBodyNetworkPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesResponseBodyNetworkPackages) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetBandwidth(v int32) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.Bandwidth = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetCreateTime(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.CreateTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetEipAddresses(v []*string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.EipAddresses = v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetExpiredTime(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetInternetChargeType(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetNetworkPackageId(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetNetworkPackageStatus(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.NetworkPackageStatus = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetOfficeSiteId(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetOfficeSiteName(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.OfficeSiteName = &v
	return s
}

type DescribeNetworkPackagesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNetworkPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNetworkPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponse) SetHeaders(v map[string]*string) *DescribeNetworkPackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkPackagesResponse) SetBody(v *DescribeNetworkPackagesResponseBody) *DescribeNetworkPackagesResponse {
	s.Body = v
	return s
}

type DescribeOfficeSitesRequest struct {
	MaxResults     *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId   []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	OfficeSiteType *string   `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeOfficeSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesRequest) SetMaxResults(v int32) *DescribeOfficeSitesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetNextToken(v string) *DescribeOfficeSitesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteType(v string) *DescribeOfficeSitesRequest {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetRegionId(v string) *DescribeOfficeSitesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetStatus(v string) *DescribeOfficeSitesRequest {
	s.Status = &v
	return s
}

type DescribeOfficeSitesResponseBody struct {
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSites []*DescribeOfficeSitesResponseBodyOfficeSites `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOfficeSitesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBody) SetNextToken(v string) *DescribeOfficeSitesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetOfficeSites(v []*DescribeOfficeSitesResponseBodyOfficeSites) *DescribeOfficeSitesResponseBody {
	s.OfficeSites = v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetRequestId(v string) *DescribeOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOfficeSitesResponseBodyOfficeSites struct {
	ADConnectors             []*DescribeOfficeSitesResponseBodyOfficeSitesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" type:"Repeated"`
	AdHostname               *string                                                   `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	Bandwidth                *int32                                                    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CenId                    *string                                                   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CidrBlock                *string                                                   `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CloudBoxOfficeSite       *bool                                                     `json:"CloudBoxOfficeSite,omitempty" xml:"CloudBoxOfficeSite,omitempty"`
	CreationTime             *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CustomSecurityGroupId    *string                                                   `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty"`
	DesktopAccessType        *string                                                   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DesktopCount             *int64                                                    `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	DesktopVpcEndpoint       *string                                                   `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	DnsAddress               []*string                                                 `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DnsUserName              *string                                                   `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty"`
	DomainName               *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainPassword           *string                                                   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DomainUserName           *string                                                   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	EnableAdminAccess        *bool                                                     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	EnableCrossDesktopAccess *bool                                                     `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
	EnableInternetAccess     *bool                                                     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	FileSystemIds            []*string                                                 `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Repeated"`
	Logs                     []*DescribeOfficeSitesResponseBodyOfficeSitesLogs         `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	MfaEnabled               *bool                                                     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	Name                     *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedVerifyLoginRisk      *bool                                                     `json:"NeedVerifyLoginRisk,omitempty" xml:"NeedVerifyLoginRisk,omitempty"`
	NeedVerifyZeroDevice     *bool                                                     `json:"NeedVerifyZeroDevice,omitempty" xml:"NeedVerifyZeroDevice,omitempty"`
	NetworkPackageId         *string                                                   `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	OfficeSiteId             *string                                                   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteType           *string                                                   `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	OuName                   *string                                                   `json:"OuName,omitempty" xml:"OuName,omitempty"`
	ProtocolType             *string                                                   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	SsoEnabled               *bool                                                     `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	SsoType                  *string                                                   `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
	Status                   *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	SubDnsAddress            []*string                                                 `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" type:"Repeated"`
	SubDomainName            *string                                                   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	TrustPassword            *string                                                   `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
	VSwitchIds               []*string                                                 `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId                    *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcType                  *string                                                   `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetADConnectors(v []*DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ADConnectors = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetAdHostname(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.AdHostname = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetBandwidth(v int32) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Bandwidth = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCenId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CenId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCidrBlock(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CidrBlock = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCloudBoxOfficeSite(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CloudBoxOfficeSite = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCreationTime(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CreationTime = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCustomSecurityGroupId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopAccessType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopCount(v int64) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopCount = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopVpcEndpoint(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDnsAddress(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DnsAddress = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDnsUserName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DnsUserName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainPassword(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainPassword = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainUserName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainUserName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableAdminAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableCrossDesktopAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableInternetAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetFileSystemIds(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.FileSystemIds = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetLogs(v []*DescribeOfficeSitesResponseBodyOfficeSitesLogs) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Logs = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetMfaEnabled(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Name = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetNeedVerifyLoginRisk(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.NeedVerifyLoginRisk = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetNeedVerifyZeroDevice(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.NeedVerifyZeroDevice = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetNetworkPackageId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOuName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OuName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetProtocolType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ProtocolType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSsoEnabled(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSsoType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SsoType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetStatus(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Status = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSubDnsAddress(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSubDomainName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SubDomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetTrustPassword(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.TrustPassword = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetVSwitchIds(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.VSwitchIds = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetVpcId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.VpcId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetVpcType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.VpcType = &v
	return s
}

type DescribeOfficeSitesResponseBodyOfficeSitesADConnectors struct {
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty"`
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	Specification      *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	TrustKey           *string `json:"TrustKey,omitempty" xml:"TrustKey,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetADConnectorAddress(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetConnectorStatus(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetNetworkInterfaceId(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetSpecification(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.Specification = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetTrustKey(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.TrustKey = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetVSwitchId(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.VSwitchId = &v
	return s
}

type DescribeOfficeSitesResponseBodyOfficeSitesLogs struct {
	Level     *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Step      *string `json:"Step,omitempty" xml:"Step,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesLogs) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetLevel(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Level = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetMessage(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Message = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetStep(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Step = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetTimeStamp(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.TimeStamp = &v
	return s
}

type DescribeOfficeSitesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOfficeSitesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOfficeSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponse) SetHeaders(v map[string]*string) *DescribeOfficeSitesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOfficeSitesResponse) SetBody(v *DescribeOfficeSitesResponseBody) *DescribeOfficeSitesResponse {
	s.Body = v
	return s
}

type DescribePolicyGroupsRequest struct {
	MaxResults    *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsRequest) SetMaxResults(v int32) *DescribePolicyGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetNextToken(v string) *DescribePolicyGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DescribePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

func (s *DescribePolicyGroupsRequest) SetRegionId(v string) *DescribePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

type DescribePolicyGroupsResponseBody struct {
	DescribePolicyGroups []*DescribePolicyGroupsResponseBodyDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" type:"Repeated"`
	NextToken            *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBody) SetDescribePolicyGroups(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroups) *DescribePolicyGroupsResponseBody {
	s.DescribePolicyGroups = v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetNextToken(v string) *DescribePolicyGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetRequestId(v string) *DescribePolicyGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroups struct {
	AuthorizeAccessPolicyRules   []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules   `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	AuthorizeSecurityPolicyRules []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules `json:"AuthorizeSecurityPolicyRules,omitempty" xml:"AuthorizeSecurityPolicyRules,omitempty" type:"Repeated"`
	ClientTypes                  []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes                  `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	Clipboard                    *string                                                                             `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	DomainList                   *string                                                                             `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	EdsCount                     *int32                                                                              `json:"EdsCount,omitempty" xml:"EdsCount,omitempty"`
	GpuAcceleration              *string                                                                             `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	Html5Access                  *string                                                                             `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	Html5FileTransfer            *string                                                                             `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	LocalDrive                   *string                                                                             `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	Name                         *string                                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyGroupId                *string                                                                             `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PolicyGroupType              *string                                                                             `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty"`
	PolicyStatus                 *string                                                                             `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	PreemptLogin                 *string                                                                             `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	PreemptLoginUsers            []*string                                                                           `json:"PreemptLoginUsers,omitempty" xml:"PreemptLoginUsers,omitempty" type:"Repeated"`
	PrinterRedirection           *string                                                                             `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	UsbRedirect                  *string                                                                             `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	UsbSupplyRedirectRule        []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule        `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	VisualQuality                *string                                                                             `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	Watermark                    *string                                                                             `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	WatermarkCustomText          *string                                                                             `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency        *string                                                                             `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	WatermarkType                *string                                                                             `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroups) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAuthorizeAccessPolicyRules(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAuthorizeSecurityPolicyRules(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AuthorizeSecurityPolicyRules = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClientTypes(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ClientTypes = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClipboard(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Clipboard = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDomainList(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DomainList = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetEdsCount(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.EdsCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetGpuAcceleration(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.GpuAcceleration = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHtml5Access(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Html5Access = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHtml5FileTransfer(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Html5FileTransfer = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetLocalDrive(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.LocalDrive = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetName(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Name = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyGroupId(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyGroupType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyGroupType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyStatus(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyStatus = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPreemptLogin(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PreemptLogin = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPreemptLoginUsers(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PreemptLoginUsers = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPrinterRedirection(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PrinterRedirection = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetUsbRedirect(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.UsbRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetUsbSupplyRedirectRule(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVisualQuality(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VisualQuality = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermark(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Watermark = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkCustomText(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkCustomText = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkTransparency(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkTransparency = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkType = &v
	return s
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetCidrIp(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetCidrIp(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetIpProtocol(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.IpProtocol = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPolicy(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Policy = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPortRange(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.PortRange = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPriority(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Priority = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Type = &v
	return s
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes struct {
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) SetClientType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes {
	s.ClientType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) SetStatus(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes {
	s.Status = &v
	return s
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceClass     *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	DeviceSubclass  *string `json:"DeviceSubclass,omitempty" xml:"DeviceSubclass,omitempty"`
	ProductId       *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	UsbRedirectType *int64  `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	UsbRuleType     *int64  `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	VendorId        *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetDeviceClass(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.DeviceClass = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetDeviceSubclass(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.DeviceSubclass = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetProductId(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetUsbRedirectType(v int64) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetUsbRuleType(v int64) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetVendorId(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

type DescribePolicyGroupsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePolicyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponse) SetHeaders(v map[string]*string) *DescribePolicyGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyGroupsResponse) SetBody(v *DescribePolicyGroupsResponseBody) *DescribePolicyGroupsResponse {
	s.Body = v
	return s
}

type DescribePriceRequest struct {
	Amount                   *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Bandwidth                *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	HardwareVersion          *string `json:"HardwareVersion,omitempty" xml:"HardwareVersion,omitempty"`
	InstanceType             *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType       *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OsType                   *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PackageSize              *int32  `json:"PackageSize,omitempty" xml:"PackageSize,omitempty"`
	Period                   *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit               *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId              *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType             *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	RootDiskSizeGib          *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	UserDiskSizeGib          *int32  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetAmount(v int32) *DescribePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequest) SetBandwidth(v int32) *DescribePriceRequest {
	s.Bandwidth = &v
	return s
}

func (s *DescribePriceRequest) SetHardwareVersion(v string) *DescribePriceRequest {
	s.HardwareVersion = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetOsType(v string) *DescribePriceRequest {
	s.OsType = &v
	return s
}

func (s *DescribePriceRequest) SetPackageSize(v int32) *DescribePriceRequest {
	s.PackageSize = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int32) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetPeriodUnit(v string) *DescribePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePriceRequest) SetPromotionId(v string) *DescribePriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceType(v string) *DescribePriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePriceRequest) SetRootDiskPerformanceLevel(v string) *DescribePriceRequest {
	s.RootDiskPerformanceLevel = &v
	return s
}

func (s *DescribePriceRequest) SetRootDiskSizeGib(v int32) *DescribePriceRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribePriceRequest) SetUserDiskPerformanceLevel(v string) *DescribePriceRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *DescribePriceRequest) SetUserDiskSizeGib(v int32) *DescribePriceRequest {
	s.UserDiskSizeGib = &v
	return s
}

type DescribePriceResponseBody struct {
	PriceInfo *DescribePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

type DescribePriceResponseBodyPriceInfo struct {
	FreeCdsQuota *bool                                      `json:"FreeCdsQuota,omitempty" xml:"FreeCdsQuota,omitempty"`
	FreeCdsSize  *int64                                     `json:"FreeCdsSize,omitempty" xml:"FreeCdsSize,omitempty"`
	Price        *DescribePriceResponseBodyPriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules        []*DescribePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfo) SetFreeCdsQuota(v bool) *DescribePriceResponseBodyPriceInfo {
	s.FreeCdsQuota = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetFreeCdsSize(v int64) *DescribePriceResponseBodyPriceInfo {
	s.FreeCdsSize = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetPrice(v *DescribePriceResponseBodyPriceInfoPrice) *DescribePriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetRules(v []*DescribePriceResponseBodyPriceInfoRules) *DescribePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

type DescribePriceResponseBodyPriceInfoPrice struct {
	Currency      *string                                              `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32                                             `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32                                             `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*DescribePriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	TradePrice    *float32                                             `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribePriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribePriceResponseBodyPriceInfoPricePromotions) *DescribePriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

type DescribePriceResponseBodyPriceInfoPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

type DescribePriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type DescribePriceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRenewalPriceRequest struct {
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds  []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	Period       *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit   *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId  *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRenewalPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewalPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceRequest) SetInstanceId(v string) *DescribeRenewalPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetInstanceIds(v []*string) *DescribeRenewalPriceRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPeriod(v int32) *DescribeRenewalPriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPeriodUnit(v string) *DescribeRenewalPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPromotionId(v string) *DescribeRenewalPriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetRegionId(v string) *DescribeRenewalPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceType(v string) *DescribeRenewalPriceRequest {
	s.ResourceType = &v
	return s
}

type DescribeRenewalPriceResponseBody struct {
	PriceInfo *DescribeRenewalPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRenewalPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewalPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBody) SetPriceInfo(v *DescribeRenewalPriceResponseBodyPriceInfo) *DescribeRenewalPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeRenewalPriceResponseBody) SetRequestId(v string) *DescribeRenewalPriceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRenewalPriceResponseBodyPriceInfo struct {
	Price *DescribeRenewalPriceResponseBodyPriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules []*DescribeRenewalPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetPrice(v *DescribeRenewalPriceResponseBodyPriceInfoPrice) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetRules(v []*DescribeRenewalPriceResponseBodyPriceInfoRules) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

type DescribeRenewalPriceResponseBodyPriceInfoPrice struct {
	Currency      *string                                                     `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32                                                    `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32                                                    `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*DescribeRenewalPriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	TradePrice    *float32                                                    `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

type DescribeRenewalPriceResponseBodyPriceInfoPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

type DescribeRenewalPriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribeRenewalPriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribeRenewalPriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type DescribeRenewalPriceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRenewalPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRenewalPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewalPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponse) SetHeaders(v map[string]*string) *DescribeRenewalPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRenewalPriceResponse) SetBody(v *DescribeRenewalPriceResponseBody) *DescribeRenewalPriceResponse {
	s.Body = v
	return s
}

type DescribeScanTaskProgressRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId   *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScanTaskProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressRequest) SetRegionId(v string) *DescribeScanTaskProgressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScanTaskProgressRequest) SetTaskId(v int64) *DescribeScanTaskProgressRequest {
	s.TaskId = &v
	return s
}

type DescribeScanTaskProgressResponseBody struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeScanTaskProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanTaskProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressResponseBody) SetCreateTime(v string) *DescribeScanTaskProgressResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeScanTaskProgressResponseBody) SetRequestId(v string) *DescribeScanTaskProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScanTaskProgressResponseBody) SetTaskStatus(v string) *DescribeScanTaskProgressResponseBody {
	s.TaskStatus = &v
	return s
}

type DescribeScanTaskProgressResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScanTaskProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScanTaskProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressResponse) SetHeaders(v map[string]*string) *DescribeScanTaskProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeScanTaskProgressResponse) SetBody(v *DescribeScanTaskProgressResponseBody) *DescribeScanTaskProgressResponse {
	s.Body = v
	return s
}

type DescribeSecurityEventOperationStatusRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityEventId []*string `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty" type:"Repeated"`
	TaskId          *int64    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSecurityEventOperationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusRequest) SetRegionId(v string) *DescribeSecurityEventOperationStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetSecurityEventId(v []*string) *DescribeSecurityEventOperationStatusRequest {
	s.SecurityEventId = v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetTaskId(v int64) *DescribeSecurityEventOperationStatusRequest {
	s.TaskId = &v
	return s
}

type DescribeSecurityEventOperationStatusResponseBody struct {
	RequestId                      *string                                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityEventOperationStatuses []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses `json:"SecurityEventOperationStatuses,omitempty" xml:"SecurityEventOperationStatuses,omitempty" type:"Repeated"`
	TaskStatus                     *string                                                                           `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeSecurityEventOperationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetRequestId(v string) *DescribeSecurityEventOperationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetSecurityEventOperationStatuses(v []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) *DescribeSecurityEventOperationStatusResponseBody {
	s.SecurityEventOperationStatuses = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetTaskStatus(v string) *DescribeSecurityEventOperationStatusResponseBody {
	s.TaskStatus = &v
	return s
}

type DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses struct {
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) SetErrorCode(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) SetSecurityEventId(v int64) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) SetStatus(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses {
	s.Status = &v
	return s
}

type DescribeSecurityEventOperationStatusResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityEventOperationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventOperationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponse) SetBody(v *DescribeSecurityEventOperationStatusResponseBody) *DescribeSecurityEventOperationStatusResponse {
	s.Body = v
	return s
}

type DescribeSecurityEventOperationsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
}

func (s DescribeSecurityEventOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsRequest) SetRegionId(v string) *DescribeSecurityEventOperationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) SetSecurityEventId(v int64) *DescribeSecurityEventOperationsRequest {
	s.SecurityEventId = &v
	return s
}

type DescribeSecurityEventOperationsResponseBody struct {
	RequestId               *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityEventOperations []*DescribeSecurityEventOperationsResponseBodySecurityEventOperations `json:"SecurityEventOperations,omitempty" xml:"SecurityEventOperations,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBody) SetRequestId(v string) *DescribeSecurityEventOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBody) SetSecurityEventOperations(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperations) *DescribeSecurityEventOperationsResponseBody {
	s.SecurityEventOperations = v
	return s
}

type DescribeSecurityEventOperationsResponseBodySecurityEventOperations struct {
	OperationCode   *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	UserCanOperate  *bool   `json:"UserCanOperate,omitempty" xml:"UserCanOperate,omitempty"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperations) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperations) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperations) SetOperationCode(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperations {
	s.OperationCode = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperations) SetOperationParams(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperations {
	s.OperationParams = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperations) SetUserCanOperate(v bool) *DescribeSecurityEventOperationsResponseBodySecurityEventOperations {
	s.UserCanOperate = &v
	return s
}

type DescribeSecurityEventOperationsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityEventOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventOperationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventOperationsResponse) SetBody(v *DescribeSecurityEventOperationsResponseBody) *DescribeSecurityEventOperationsResponse {
	s.Body = v
	return s
}

type DescribeSnapshotsRequest struct {
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetDesktopId(v string) *DescribeSnapshotsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int32) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotId(v string) *DescribeSnapshotsRequest {
	s.SnapshotId = &v
	return s
}

type DescribeSnapshotsResponseBody struct {
	NextToken *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots []*DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) SetNextToken(v string) *DescribeSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

type DescribeSnapshotsResponseBodySnapshots struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RemainTime     *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotType   *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	SourceDiskSize *string `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetCreationTime(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDesktopId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshots {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskSize(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

type DescribeSnapshotsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotsResponse) SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeSuspEventOverviewRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSuspEventOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventOverviewRequest) SetRegionId(v string) *DescribeSuspEventOverviewRequest {
	s.RegionId = &v
	return s
}

type DescribeSuspEventOverviewResponseBody struct {
	RemindCount     *int32  `json:"RemindCount,omitempty" xml:"RemindCount,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SeriousCount    *int32  `json:"SeriousCount,omitempty" xml:"SeriousCount,omitempty"`
	SuspiciousCount *int32  `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty"`
}

func (s DescribeSuspEventOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventOverviewResponseBody) SetRemindCount(v int32) *DescribeSuspEventOverviewResponseBody {
	s.RemindCount = &v
	return s
}

func (s *DescribeSuspEventOverviewResponseBody) SetRequestId(v string) *DescribeSuspEventOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventOverviewResponseBody) SetSeriousCount(v int32) *DescribeSuspEventOverviewResponseBody {
	s.SeriousCount = &v
	return s
}

func (s *DescribeSuspEventOverviewResponseBody) SetSuspiciousCount(v int32) *DescribeSuspEventOverviewResponseBody {
	s.SuspiciousCount = &v
	return s
}

type DescribeSuspEventOverviewResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventOverviewResponse) SetHeaders(v map[string]*string) *DescribeSuspEventOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventOverviewResponse) SetBody(v *DescribeSuspEventOverviewResponseBody) *DescribeSuspEventOverviewResponse {
	s.Body = v
	return s
}

type DescribeSuspEventQuaraFilesRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSuspEventQuaraFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesRequest) SetCurrentPage(v int32) *DescribeSuspEventQuaraFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetOfficeSiteId(v string) *DescribeSuspEventQuaraFilesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetPageSize(v int32) *DescribeSuspEventQuaraFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetRegionId(v string) *DescribeSuspEventQuaraFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetStatus(v string) *DescribeSuspEventQuaraFilesRequest {
	s.Status = &v
	return s
}

type DescribeSuspEventQuaraFilesResponseBody struct {
	CurrentPage *int32                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QuaraFiles  []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles `json:"QuaraFiles,omitempty" xml:"QuaraFiles,omitempty" type:"Repeated"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSuspEventQuaraFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetCurrentPage(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetPageSize(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetQuaraFiles(v []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) *DescribeSuspEventQuaraFilesResponseBody {
	s.QuaraFiles = v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetRequestId(v string) *DescribeSuspEventQuaraFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetTotalCount(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSuspEventQuaraFilesResponseBodyQuaraFiles struct {
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	EventName   *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType   *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Id          *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Md5         *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag         *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetDesktopId(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.DesktopId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetDesktopName(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.DesktopName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetEventName(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.EventName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetEventType(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.EventType = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetId(v int32) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetMd5(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Md5 = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetModifyTime(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.ModifyTime = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetPath(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Path = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetStatus(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetTag(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Tag = &v
	return s
}

type DescribeSuspEventQuaraFilesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventQuaraFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventQuaraFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponse) SetHeaders(v map[string]*string) *DescribeSuspEventQuaraFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetBody(v *DescribeSuspEventQuaraFilesResponseBody) *DescribeSuspEventQuaraFilesResponse {
	s.Body = v
	return s
}

type DescribeSuspEventsRequest struct {
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Dealed          *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Levels          *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	OfficeSiteId    *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentEventType *string `json:"ParentEventType,omitempty" xml:"ParentEventType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSuspEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsRequest) SetAlarmUniqueInfo(v string) *DescribeSuspEventsRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetCurrentPage(v int32) *DescribeSuspEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetDealed(v string) *DescribeSuspEventsRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLang(v string) *DescribeSuspEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLevels(v string) *DescribeSuspEventsRequest {
	s.Levels = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetOfficeSiteId(v string) *DescribeSuspEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetPageSize(v int32) *DescribeSuspEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetParentEventType(v string) *DescribeSuspEventsRequest {
	s.ParentEventType = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetRegionId(v string) *DescribeSuspEventsRequest {
	s.RegionId = &v
	return s
}

type DescribeSuspEventsResponseBody struct {
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *string                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuspEvents  []*DescribeSuspEventsResponseBodySuspEvents `json:"SuspEvents,omitempty" xml:"SuspEvents,omitempty" type:"Repeated"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSuspEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBody) SetCurrentPage(v int32) *DescribeSuspEventsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetPageSize(v string) *DescribeSuspEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetRequestId(v string) *DescribeSuspEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetSuspEvents(v []*DescribeSuspEventsResponseBodySuspEvents) *DescribeSuspEventsResponseBody {
	s.SuspEvents = v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetTotalCount(v int32) *DescribeSuspEventsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSuspEventsResponseBodySuspEvents struct {
	AlarmEventName        *string                                            `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	AlarmEventNameDisplay *string                                            `json:"AlarmEventNameDisplay,omitempty" xml:"AlarmEventNameDisplay,omitempty"`
	AlarmEventType        *string                                            `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	AlarmEventTypeDisplay *string                                            `json:"AlarmEventTypeDisplay,omitempty" xml:"AlarmEventTypeDisplay,omitempty"`
	AlarmUniqueInfo       *string                                            `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	CanBeDealOnLine       *string                                            `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	CanCancelFault        *bool                                              `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	DataSource            *string                                            `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	Desc                  *string                                            `json:"Desc,omitempty" xml:"Desc,omitempty"`
	DesktopId             *string                                            `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName           *string                                            `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	Details               []*DescribeSuspEventsResponseBodySuspEventsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	EventStatus           *int32                                             `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	EventSubType          *string                                            `json:"EventSubType,omitempty" xml:"EventSubType,omitempty"`
	Id                    *int64                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	LastTime              *string                                            `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	Level                 *string                                            `json:"Level,omitempty" xml:"Level,omitempty"`
	Name                  *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	OccurrenceTime        *string                                            `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	OperateErrorCode      *string                                            `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	OperateMsg            *string                                            `json:"OperateMsg,omitempty" xml:"OperateMsg,omitempty"`
	UniqueInfo            *string                                            `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEvents) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventNameDisplay(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventNameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventType(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventTypeDisplay(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventTypeDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmUniqueInfo(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetCanBeDealOnLine(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetCanCancelFault(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDataSource(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DataSource = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDesc(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Desc = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDesktopId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDesktopName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DesktopName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDetails(v []*DescribeSuspEventsResponseBodySuspEventsDetails) *DescribeSuspEventsResponseBodySuspEvents {
	s.Details = v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventStatus(v int32) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventStatus = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventSubType(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventSubType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetId(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLastTime(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.LastTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLevel(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Level = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOccurrenceTime(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateErrorCode(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateMsg(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateMsg = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetUniqueInfo(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.UniqueInfo = &v
	return s
}

type DescribeSuspEventsResponseBodySuspEventsDetails struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameDisplay  *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEventsDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEventsDetails) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetName(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetNameDisplay(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.NameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetType(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetValue(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Value = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetValueDisplay(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.ValueDisplay = &v
	return s
}

type DescribeSuspEventsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponse) SetHeaders(v map[string]*string) *DescribeSuspEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventsResponse) SetBody(v *DescribeSuspEventsResponseBody) *DescribeSuspEventsResponse {
	s.Body = v
	return s
}

type DescribeUserConnectionRecordsRequest struct {
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserType    *string `json:"EndUserType,omitempty" xml:"EndUserType,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserConnectionRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsRequest) SetDesktopGroupId(v string) *DescribeUserConnectionRecordsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetEndUserId(v string) *DescribeUserConnectionRecordsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetEndUserType(v string) *DescribeUserConnectionRecordsRequest {
	s.EndUserType = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetMaxResults(v int32) *DescribeUserConnectionRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetNextToken(v string) *DescribeUserConnectionRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetRegionId(v string) *DescribeUserConnectionRecordsRequest {
	s.RegionId = &v
	return s
}

type DescribeUserConnectionRecordsResponseBody struct {
	ConnectionRecords []*DescribeUserConnectionRecordsResponseBodyConnectionRecords `json:"ConnectionRecords,omitempty" xml:"ConnectionRecords,omitempty" type:"Repeated"`
	NextToken         *string                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserConnectionRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponseBody) SetConnectionRecords(v []*DescribeUserConnectionRecordsResponseBodyConnectionRecords) *DescribeUserConnectionRecordsResponseBody {
	s.ConnectionRecords = v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBody) SetNextToken(v string) *DescribeUserConnectionRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBody) SetRequestId(v string) *DescribeUserConnectionRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserConnectionRecordsResponseBodyConnectionRecords struct {
	ConnectDuration    *string `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	ConnectEndTime     *string `json:"ConnectEndTime,omitempty" xml:"ConnectEndTime,omitempty"`
	ConnectStartTime   *string `json:"ConnectStartTime,omitempty" xml:"ConnectStartTime,omitempty"`
	ConnectionRecordId *string `json:"ConnectionRecordId,omitempty" xml:"ConnectionRecordId,omitempty"`
	DesktopId          *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName        *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
}

func (s DescribeUserConnectionRecordsResponseBodyConnectionRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponseBodyConnectionRecords) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectDuration(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectDuration = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectEndTime(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectEndTime = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectStartTime(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectStartTime = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectionRecordId(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectionRecordId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetDesktopId(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.DesktopId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetDesktopName(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.DesktopName = &v
	return s
}

type DescribeUserConnectionRecordsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserConnectionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserConnectionRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponse) SetHeaders(v map[string]*string) *DescribeUserConnectionRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserConnectionRecordsResponse) SetBody(v *DescribeUserConnectionRecordsResponseBody) *DescribeUserConnectionRecordsResponse {
	s.Body = v
	return s
}

type DescribeUsersInGroupRequest struct {
	ConnectState   *int32  `json:"ConnectState,omitempty" xml:"ConnectState,omitempty"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUsersInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupRequest) SetConnectState(v int32) *DescribeUsersInGroupRequest {
	s.ConnectState = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetDesktopGroupId(v string) *DescribeUsersInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetEndUserId(v string) *DescribeUsersInGroupRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetMaxResults(v int32) *DescribeUsersInGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetNextToken(v string) *DescribeUsersInGroupRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetRegionId(v string) *DescribeUsersInGroupRequest {
	s.RegionId = &v
	return s
}

type DescribeUsersInGroupResponseBody struct {
	EndUsers         []*DescribeUsersInGroupResponseBodyEndUsers `json:"EndUsers,omitempty" xml:"EndUsers,omitempty" type:"Repeated"`
	NextToken        *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OnlineUsersCount *int32                                      `json:"OnlineUsersCount,omitempty" xml:"OnlineUsersCount,omitempty"`
	RequestId        *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsersCount       *int32                                      `json:"UsersCount,omitempty" xml:"UsersCount,omitempty"`
}

func (s DescribeUsersInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBody) SetEndUsers(v []*DescribeUsersInGroupResponseBodyEndUsers) *DescribeUsersInGroupResponseBody {
	s.EndUsers = v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetNextToken(v string) *DescribeUsersInGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetOnlineUsersCount(v int32) *DescribeUsersInGroupResponseBody {
	s.OnlineUsersCount = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetRequestId(v string) *DescribeUsersInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetUsersCount(v int32) *DescribeUsersInGroupResponseBody {
	s.UsersCount = &v
	return s
}

type DescribeUsersInGroupResponseBodyEndUsers struct {
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName      *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	EndUserEmail     *string `json:"EndUserEmail,omitempty" xml:"EndUserEmail,omitempty"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserName      *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
	EndUserPhone     *string `json:"EndUserPhone,omitempty" xml:"EndUserPhone,omitempty"`
	EndUserType      *string `json:"EndUserType,omitempty" xml:"EndUserType,omitempty"`
}

func (s DescribeUsersInGroupResponseBodyEndUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupResponseBodyEndUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetConnectionStatus(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetDesktopId(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.DesktopId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetDesktopName(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.DesktopName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserEmail(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserEmail = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserId(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserName(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserPhone(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserPhone = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserType(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserType = &v
	return s
}

type DescribeUsersInGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUsersInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUsersInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponse) SetHeaders(v map[string]*string) *DescribeUsersInGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersInGroupResponse) SetBody(v *DescribeUsersInGroupResponseBody) *DescribeUsersInGroupResponse {
	s.Body = v
	return s
}

type DescribeVirtualMFADevicesRequest struct {
	DirectoryId  *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId    []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	MaxResults   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeVirtualMFADevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesRequest) SetDirectoryId(v string) *DescribeVirtualMFADevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetEndUserId(v []*string) *DescribeVirtualMFADevicesRequest {
	s.EndUserId = v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetMaxResults(v int32) *DescribeVirtualMFADevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetNextToken(v string) *DescribeVirtualMFADevicesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetOfficeSiteId(v string) *DescribeVirtualMFADevicesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetRegionId(v string) *DescribeVirtualMFADevicesRequest {
	s.RegionId = &v
	return s
}

type DescribeVirtualMFADevicesResponseBody struct {
	NextToken         *string                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VirtualMFADevices []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" type:"Repeated"`
}

func (s DescribeVirtualMFADevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseBody) SetNextToken(v string) *DescribeVirtualMFADevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) SetRequestId(v string) *DescribeVirtualMFADevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) SetVirtualMFADevices(v []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) *DescribeVirtualMFADevicesResponseBody {
	s.VirtualMFADevices = v
	return s
}

type DescribeVirtualMFADevicesResponseBodyVirtualMFADevices struct {
	ConsecutiveFails *int32  `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty"`
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	GmtEnabled       *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty"`
	GmtUnlock        *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty"`
	OfficeSiteId     *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetConsecutiveFails(v int32) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.ConsecutiveFails = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetDirectoryId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetEndUserId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.EndUserId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetGmtEnabled(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.GmtEnabled = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetGmtUnlock(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetOfficeSiteId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetSerialNumber(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.SerialNumber = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetStatus(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.Status = &v
	return s
}

type DescribeVirtualMFADevicesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVirtualMFADevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVirtualMFADevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponse) SetHeaders(v map[string]*string) *DescribeVirtualMFADevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) SetBody(v *DescribeVirtualMFADevicesResponseBody) *DescribeVirtualMFADevicesResponse {
	s.Body = v
	return s
}

type DescribeVulDetailsRequest struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsRequest) SetAliasName(v string) *DescribeVulDetailsRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetLang(v string) *DescribeVulDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetName(v string) *DescribeVulDetailsRequest {
	s.Name = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetRegionId(v string) *DescribeVulDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetType(v string) *DescribeVulDetailsRequest {
	s.Type = &v
	return s
}

type DescribeVulDetailsResponseBody struct {
	Cves      []*DescribeVulDetailsResponseBodyCves `json:"Cves,omitempty" xml:"Cves,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVulDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBody) SetCves(v []*DescribeVulDetailsResponseBodyCves) *DescribeVulDetailsResponseBody {
	s.Cves = v
	return s
}

func (s *DescribeVulDetailsResponseBody) SetRequestId(v string) *DescribeVulDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVulDetailsResponseBodyCves struct {
	CveId     *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	CvssScore *string `json:"CvssScore,omitempty" xml:"CvssScore,omitempty"`
	Summary   *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeVulDetailsResponseBodyCves) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponseBodyCves) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBodyCves) SetCveId(v string) *DescribeVulDetailsResponseBodyCves {
	s.CveId = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCvssScore(v string) *DescribeVulDetailsResponseBodyCves {
	s.CvssScore = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetSummary(v string) *DescribeVulDetailsResponseBodyCves {
	s.Summary = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetTitle(v string) *DescribeVulDetailsResponseBodyCves {
	s.Title = &v
	return s
}

type DescribeVulDetailsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponse) SetHeaders(v map[string]*string) *DescribeVulDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulDetailsResponse) SetBody(v *DescribeVulDetailsResponseBody) *DescribeVulDetailsResponse {
	s.Body = v
	return s
}

type DescribeVulListRequest struct {
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Dealed       *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Necessity    *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulListRequest) SetAliasName(v string) *DescribeVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListRequest) SetCurrentPage(v int32) *DescribeVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListRequest) SetDealed(v string) *DescribeVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeVulListRequest) SetLang(v string) *DescribeVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulListRequest) SetNecessity(v string) *DescribeVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListRequest) SetOfficeSiteId(v string) *DescribeVulListRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVulListRequest) SetPageSize(v int32) *DescribeVulListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListRequest) SetRegionId(v string) *DescribeVulListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVulListRequest) SetType(v string) *DescribeVulListRequest {
	s.Type = &v
	return s
}

type DescribeVulListResponseBody struct {
	CurrentPage *int32                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VulRecords  []*DescribeVulListResponseBodyVulRecords `json:"VulRecords,omitempty" xml:"VulRecords,omitempty" type:"Repeated"`
}

func (s DescribeVulListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBody) SetCurrentPage(v int32) *DescribeVulListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListResponseBody) SetPageSize(v int32) *DescribeVulListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListResponseBody) SetRequestId(v string) *DescribeVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulListResponseBody) SetTotalCount(v int32) *DescribeVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulListResponseBody) SetVulRecords(v []*DescribeVulListResponseBodyVulRecords) *DescribeVulListResponseBody {
	s.VulRecords = v
	return s
}

type DescribeVulListResponseBodyVulRecords struct {
	AliasName         *string                                                 `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	DesktopId         *string                                                 `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName       *string                                                 `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	ExtendContentJson *DescribeVulListResponseBodyVulRecordsExtendContentJson `json:"ExtendContentJson,omitempty" xml:"ExtendContentJson,omitempty" type:"Struct"`
	FirstTs           *int64                                                  `json:"FirstTs,omitempty" xml:"FirstTs,omitempty"`
	LastTs            *int64                                                  `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	ModifyTs          *int64                                                  `json:"ModifyTs,omitempty" xml:"ModifyTs,omitempty"`
	Name              *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Necessity         *string                                                 `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Online            *bool                                                   `json:"Online,omitempty" xml:"Online,omitempty"`
	OsVersion         *string                                                 `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	Related           *string                                                 `json:"Related,omitempty" xml:"Related,omitempty"`
	RepairTs          *int64                                                  `json:"RepairTs,omitempty" xml:"RepairTs,omitempty"`
	ResultCode        *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage     *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Status            *int32                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag               *string                                                 `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Type              *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecords) SetAliasName(v string) *DescribeVulListResponseBodyVulRecords {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetDesktopId(v string) *DescribeVulListResponseBodyVulRecords {
	s.DesktopId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetDesktopName(v string) *DescribeVulListResponseBodyVulRecords {
	s.DesktopName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetExtendContentJson(v *DescribeVulListResponseBodyVulRecordsExtendContentJson) *DescribeVulListResponseBodyVulRecords {
	s.ExtendContentJson = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetFirstTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.FirstTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetLastTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.LastTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetModifyTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.ModifyTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetName(v string) *DescribeVulListResponseBodyVulRecords {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetNecessity(v string) *DescribeVulListResponseBodyVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOnline(v bool) *DescribeVulListResponseBodyVulRecords {
	s.Online = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOsVersion(v string) *DescribeVulListResponseBodyVulRecords {
	s.OsVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRelated(v string) *DescribeVulListResponseBodyVulRecords {
	s.Related = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRepairTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.RepairTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetResultCode(v string) *DescribeVulListResponseBodyVulRecords {
	s.ResultCode = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetResultMessage(v string) *DescribeVulListResponseBodyVulRecords {
	s.ResultMessage = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetStatus(v int32) *DescribeVulListResponseBodyVulRecords {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetTag(v string) *DescribeVulListResponseBodyVulRecords {
	s.Tag = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetType(v string) *DescribeVulListResponseBodyVulRecords {
	s.Type = &v
	return s
}

type DescribeVulListResponseBodyVulRecordsExtendContentJson struct {
	RpmEntityList []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList `json:"RpmEntityList,omitempty" xml:"RpmEntityList,omitempty" type:"Repeated"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJson) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJson) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetRpmEntityList(v []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.RpmEntityList = v
	return s
}

type DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList struct {
	FullVersion *string `json:"FullVersion,omitempty" xml:"FullVersion,omitempty"`
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdateCmd   *string `json:"UpdateCmd,omitempty" xml:"UpdateCmd,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetFullVersion(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.FullVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchDetail(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchDetail = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetName(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetPath(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Path = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetUpdateCmd(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.UpdateCmd = &v
	return s
}

type DescribeVulListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponse) SetHeaders(v map[string]*string) *DescribeVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulListResponse) SetBody(v *DescribeVulListResponseBody) *DescribeVulListResponse {
	s.Body = v
	return s
}

type DescribeVulOverviewRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeVulOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulOverviewRequest) SetRegionId(v string) *DescribeVulOverviewRequest {
	s.RegionId = &v
	return s
}

type DescribeVulOverviewResponseBody struct {
	AsapCount  *int32  `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	LaterCount *int32  `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	NntfCount  *int32  `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVulOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulOverviewResponseBody) SetAsapCount(v int32) *DescribeVulOverviewResponseBody {
	s.AsapCount = &v
	return s
}

func (s *DescribeVulOverviewResponseBody) SetLaterCount(v int32) *DescribeVulOverviewResponseBody {
	s.LaterCount = &v
	return s
}

func (s *DescribeVulOverviewResponseBody) SetNntfCount(v int32) *DescribeVulOverviewResponseBody {
	s.NntfCount = &v
	return s
}

func (s *DescribeVulOverviewResponseBody) SetRequestId(v string) *DescribeVulOverviewResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVulOverviewResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulOverviewResponse) SetHeaders(v map[string]*string) *DescribeVulOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulOverviewResponse) SetBody(v *DescribeVulOverviewResponseBody) *DescribeVulOverviewResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetZoneType(v string) *DescribeZonesRequest {
	s.ZoneType = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DetachCenRequest struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachCenRequest) GoString() string {
	return s.String()
}

func (s *DetachCenRequest) SetOfficeSiteId(v string) *DetachCenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DetachCenRequest) SetRegionId(v string) *DetachCenRequest {
	s.RegionId = &v
	return s
}

type DetachCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachCenResponseBody) GoString() string {
	return s.String()
}

func (s *DetachCenResponseBody) SetRequestId(v string) *DetachCenResponseBody {
	s.RequestId = &v
	return s
}

type DetachCenResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachCenResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachCenResponse) GoString() string {
	return s.String()
}

func (s *DetachCenResponse) SetHeaders(v map[string]*string) *DetachCenResponse {
	s.Headers = v
	return s
}

func (s *DetachCenResponse) SetBody(v *DetachCenResponseBody) *DetachCenResponse {
	s.Body = v
	return s
}

type DisableDesktopsInGroupRequest struct {
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopIds     []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableDesktopsInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDesktopsInGroupRequest) GoString() string {
	return s.String()
}

func (s *DisableDesktopsInGroupRequest) SetDesktopGroupId(v string) *DisableDesktopsInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DisableDesktopsInGroupRequest) SetDesktopIds(v []*string) *DisableDesktopsInGroupRequest {
	s.DesktopIds = v
	return s
}

func (s *DisableDesktopsInGroupRequest) SetRegionId(v string) *DisableDesktopsInGroupRequest {
	s.RegionId = &v
	return s
}

type DisableDesktopsInGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDesktopsInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDesktopsInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDesktopsInGroupResponseBody) SetRequestId(v string) *DisableDesktopsInGroupResponseBody {
	s.RequestId = &v
	return s
}

type DisableDesktopsInGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableDesktopsInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableDesktopsInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDesktopsInGroupResponse) GoString() string {
	return s.String()
}

func (s *DisableDesktopsInGroupResponse) SetHeaders(v map[string]*string) *DisableDesktopsInGroupResponse {
	s.Headers = v
	return s
}

func (s *DisableDesktopsInGroupResponse) SetBody(v *DisableDesktopsInGroupResponseBody) *DisableDesktopsInGroupResponse {
	s.Body = v
	return s
}

type ExportClientEventsRequest struct {
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName    *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EventType      *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	OfficeSiteId   *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ExportClientEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportClientEventsRequest) GoString() string {
	return s.String()
}

func (s *ExportClientEventsRequest) SetDesktopId(v string) *ExportClientEventsRequest {
	s.DesktopId = &v
	return s
}

func (s *ExportClientEventsRequest) SetDesktopName(v string) *ExportClientEventsRequest {
	s.DesktopName = &v
	return s
}

func (s *ExportClientEventsRequest) SetEndTime(v string) *ExportClientEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ExportClientEventsRequest) SetEndUserId(v string) *ExportClientEventsRequest {
	s.EndUserId = &v
	return s
}

func (s *ExportClientEventsRequest) SetEventType(v string) *ExportClientEventsRequest {
	s.EventType = &v
	return s
}

func (s *ExportClientEventsRequest) SetMaxResults(v int32) *ExportClientEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *ExportClientEventsRequest) SetOfficeSiteId(v string) *ExportClientEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ExportClientEventsRequest) SetOfficeSiteName(v string) *ExportClientEventsRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *ExportClientEventsRequest) SetRegionId(v string) *ExportClientEventsRequest {
	s.RegionId = &v
	return s
}

func (s *ExportClientEventsRequest) SetStartTime(v string) *ExportClientEventsRequest {
	s.StartTime = &v
	return s
}

type ExportClientEventsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExportClientEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportClientEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportClientEventsResponseBody) SetRequestId(v string) *ExportClientEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportClientEventsResponseBody) SetUrl(v string) *ExportClientEventsResponseBody {
	s.Url = &v
	return s
}

type ExportClientEventsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportClientEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportClientEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportClientEventsResponse) GoString() string {
	return s.String()
}

func (s *ExportClientEventsResponse) SetHeaders(v map[string]*string) *ExportClientEventsResponse {
	s.Headers = v
	return s
}

func (s *ExportClientEventsResponse) SetBody(v *ExportClientEventsResponseBody) *ExportClientEventsResponse {
	s.Body = v
	return s
}

type ExportDesktopGroupInfoRequest struct {
	ChargeType       *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DesktopGroupId   []*string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" type:"Repeated"`
	DesktopGroupName *string   `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DirectoryId      *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId        []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	ExpiredTime      *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	MaxResults       *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId     *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PolicyGroupId    *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExportDesktopGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportDesktopGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *ExportDesktopGroupInfoRequest) SetChargeType(v string) *ExportDesktopGroupInfoRequest {
	s.ChargeType = &v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetDesktopGroupId(v []*string) *ExportDesktopGroupInfoRequest {
	s.DesktopGroupId = v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetDesktopGroupName(v string) *ExportDesktopGroupInfoRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetDirectoryId(v string) *ExportDesktopGroupInfoRequest {
	s.DirectoryId = &v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetEndUserId(v []*string) *ExportDesktopGroupInfoRequest {
	s.EndUserId = v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetExpiredTime(v string) *ExportDesktopGroupInfoRequest {
	s.ExpiredTime = &v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetMaxResults(v int32) *ExportDesktopGroupInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetNextToken(v string) *ExportDesktopGroupInfoRequest {
	s.NextToken = &v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetOfficeSiteId(v string) *ExportDesktopGroupInfoRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetPolicyGroupId(v string) *ExportDesktopGroupInfoRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ExportDesktopGroupInfoRequest) SetRegionId(v string) *ExportDesktopGroupInfoRequest {
	s.RegionId = &v
	return s
}

type ExportDesktopGroupInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExportDesktopGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportDesktopGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ExportDesktopGroupInfoResponseBody) SetRequestId(v string) *ExportDesktopGroupInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportDesktopGroupInfoResponseBody) SetUrl(v string) *ExportDesktopGroupInfoResponseBody {
	s.Url = &v
	return s
}

type ExportDesktopGroupInfoResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportDesktopGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportDesktopGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportDesktopGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *ExportDesktopGroupInfoResponse) SetHeaders(v map[string]*string) *ExportDesktopGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *ExportDesktopGroupInfoResponse) SetBody(v *ExportDesktopGroupInfoResponseBody) *ExportDesktopGroupInfoResponse {
	s.Body = v
	return s
}

type ExportDesktopListInfoRequest struct {
	ChargeType    *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	DesktopName   *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DirectoryId   *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId     []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	ExpiredTime   *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GroupId       *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MaxResults    *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId  *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserName      *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ExportDesktopListInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportDesktopListInfoRequest) GoString() string {
	return s.String()
}

func (s *ExportDesktopListInfoRequest) SetChargeType(v string) *ExportDesktopListInfoRequest {
	s.ChargeType = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetDesktopId(v []*string) *ExportDesktopListInfoRequest {
	s.DesktopId = v
	return s
}

func (s *ExportDesktopListInfoRequest) SetDesktopName(v string) *ExportDesktopListInfoRequest {
	s.DesktopName = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetDesktopStatus(v string) *ExportDesktopListInfoRequest {
	s.DesktopStatus = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetDirectoryId(v string) *ExportDesktopListInfoRequest {
	s.DirectoryId = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetEndUserId(v []*string) *ExportDesktopListInfoRequest {
	s.EndUserId = v
	return s
}

func (s *ExportDesktopListInfoRequest) SetExpiredTime(v string) *ExportDesktopListInfoRequest {
	s.ExpiredTime = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetGroupId(v string) *ExportDesktopListInfoRequest {
	s.GroupId = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetMaxResults(v int32) *ExportDesktopListInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetNextToken(v string) *ExportDesktopListInfoRequest {
	s.NextToken = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetOfficeSiteId(v string) *ExportDesktopListInfoRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetPolicyGroupId(v string) *ExportDesktopListInfoRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetRegionId(v string) *ExportDesktopListInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ExportDesktopListInfoRequest) SetUserName(v string) *ExportDesktopListInfoRequest {
	s.UserName = &v
	return s
}

type ExportDesktopListInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExportDesktopListInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportDesktopListInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ExportDesktopListInfoResponseBody) SetRequestId(v string) *ExportDesktopListInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportDesktopListInfoResponseBody) SetUrl(v string) *ExportDesktopListInfoResponseBody {
	s.Url = &v
	return s
}

type ExportDesktopListInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportDesktopListInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportDesktopListInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportDesktopListInfoResponse) GoString() string {
	return s.String()
}

func (s *ExportDesktopListInfoResponse) SetHeaders(v map[string]*string) *ExportDesktopListInfoResponse {
	s.Headers = v
	return s
}

func (s *ExportDesktopListInfoResponse) SetBody(v *ExportDesktopListInfoResponseBody) *ExportDesktopListInfoResponse {
	s.Body = v
	return s
}

type GetConnectionTicketRequest struct {
	DesktopId            *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EndUserId            *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetDesktopId(v string) *GetConnectionTicketRequest {
	s.DesktopId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEndUserId(v string) *GetConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetOwnerId(v int64) *GetConnectionTicketRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetPassword(v string) *GetConnectionTicketRequest {
	s.Password = &v
	return s
}

func (s *GetConnectionTicketRequest) SetRegionId(v string) *GetConnectionTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerAccount(v string) *GetConnectionTicketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerId(v int64) *GetConnectionTicketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Ticket     *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetHeaders(v map[string]*string) *GetConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionTicketResponse) SetBody(v *GetConnectionTicketResponseBody) *GetConnectionTicketResponse {
	s.Body = v
	return s
}

type GetDesktopGroupDetailRequest struct {
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDesktopGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailRequest) SetDesktopGroupId(v string) *GetDesktopGroupDetailRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailRequest) SetRegionId(v string) *GetDesktopGroupDetailRequest {
	s.RegionId = &v
	return s
}

type GetDesktopGroupDetailResponseBody struct {
	Desktops  *GetDesktopGroupDetailResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDesktopGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBody) SetDesktops(v *GetDesktopGroupDetailResponseBodyDesktops) *GetDesktopGroupDetailResponseBody {
	s.Desktops = v
	return s
}

func (s *GetDesktopGroupDetailResponseBody) SetRequestId(v string) *GetDesktopGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetDesktopGroupDetailResponseBodyDesktops struct {
	AllowAutoSetup     *int32                                                 `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	AllowBufferCount   *int32                                                 `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	BindAmount         *int32                                                 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	Comments           *string                                                `json:"Comments,omitempty" xml:"Comments,omitempty"`
	Cpu                *int32                                                 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime       *string                                                `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Creator            *string                                                `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DataDiskCategory   *string                                                `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskSize       *string                                                `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopGroupId     *string                                                `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopGroupName   *string                                                `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DirectoryId        *string                                                `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryType      *string                                                `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	ExpiredTime        *string                                                `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GpuCount           *float32                                               `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuSpec            *string                                                `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	KeepDuration       *int64                                                 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	LoadPolicy         *int32                                                 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	MaxDesktopsCount   *int32                                                 `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	Memory             *int64                                                 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	MinDesktopsCount   *int32                                                 `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	OfficeSiteId       *string                                                `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName     *string                                                `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	OfficeSiteType     *string                                                `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	OwnBundleId        *string                                                `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	OwnBundleName      *string                                                `json:"OwnBundleName,omitempty" xml:"OwnBundleName,omitempty"`
	OwnType            *int32                                                 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	PayType            *string                                                `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PolicyGroupId      *string                                                `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PolicyGroupName    *string                                                `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	ResType            *int32                                                 `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResetType          *int32                                                 `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	Status             *int32                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemDiskCategory *string                                                `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskSize     *int32                                                 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	TimerInfos         []*GetDesktopGroupDetailResponseBodyDesktopsTimerInfos `json:"TimerInfos,omitempty" xml:"TimerInfos,omitempty" type:"Repeated"`
}

func (s GetDesktopGroupDetailResponseBodyDesktops) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetAllowAutoSetup(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.AllowAutoSetup = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetAllowBufferCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.AllowBufferCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetBindAmount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.BindAmount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetComments(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Comments = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCpu(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCreationTime(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCreator(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Creator = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDataDiskCategory(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDataDiskSize(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDesktopGroupId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDesktopGroupName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DesktopGroupName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDirectoryId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDirectoryType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DirectoryType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetExpiredTime(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetGpuCount(v float32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.GpuCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetGpuSpec(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.GpuSpec = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetKeepDuration(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.KeepDuration = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetLoadPolicy(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.LoadPolicy = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMaxDesktopsCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.MaxDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMemory(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMinDesktopsCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.MinDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOwnBundleId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OwnBundleId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOwnBundleName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OwnBundleName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOwnType(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OwnType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPayType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PayType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPolicyGroupId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPolicyGroupName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PolicyGroupName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetResType(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ResType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetResetType(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ResetType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetStatus(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Status = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetSystemDiskCategory(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetSystemDiskSize(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetTimerInfos(v []*GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) *GetDesktopGroupDetailResponseBodyDesktops {
	s.TimerInfos = v
	return s
}

type GetDesktopGroupDetailResponseBodyDesktopsTimerInfos struct {
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TimerType      *int32  `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) SetCronExpression(v string) *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos {
	s.CronExpression = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) SetStatus(v int32) *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos {
	s.Status = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) SetTimerType(v int32) *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos {
	s.TimerType = &v
	return s
}

type GetDesktopGroupDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDesktopGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDesktopGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponse) SetHeaders(v map[string]*string) *GetDesktopGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDesktopGroupDetailResponse) SetBody(v *GetDesktopGroupDetailResponseBody) *GetDesktopGroupDetailResponse {
	s.Body = v
	return s
}

type GetDirectorySsoStatusRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDirectorySsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusRequest) SetDirectoryId(v string) *GetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectorySsoStatusRequest) SetRegionId(v string) *GetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

type GetDirectorySsoStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SsoStatus *bool   `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty"`
}

func (s GetDirectorySsoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusResponseBody) SetRequestId(v string) *GetDirectorySsoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDirectorySsoStatusResponseBody) SetSsoStatus(v bool) *GetDirectorySsoStatusResponseBody {
	s.SsoStatus = &v
	return s
}

type GetDirectorySsoStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDirectorySsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDirectorySsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusResponse) SetHeaders(v map[string]*string) *GetDirectorySsoStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDirectorySsoStatusResponse) SetBody(v *GetDirectorySsoStatusResponseBody) *GetDirectorySsoStatusResponse {
	s.Body = v
	return s
}

type GetOfficeSiteSsoStatusRequest struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOfficeSiteSsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeSiteSsoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusRequest) SetOfficeSiteId(v string) *GetOfficeSiteSsoStatusRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetOfficeSiteSsoStatusRequest) SetRegionId(v string) *GetOfficeSiteSsoStatusRequest {
	s.RegionId = &v
	return s
}

type GetOfficeSiteSsoStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SsoStatus *bool   `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty"`
}

func (s GetOfficeSiteSsoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeSiteSsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusResponseBody) SetRequestId(v string) *GetOfficeSiteSsoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfficeSiteSsoStatusResponseBody) SetSsoStatus(v bool) *GetOfficeSiteSsoStatusResponseBody {
	s.SsoStatus = &v
	return s
}

type GetOfficeSiteSsoStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficeSiteSsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficeSiteSsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeSiteSsoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusResponse) SetHeaders(v map[string]*string) *GetOfficeSiteSsoStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOfficeSiteSsoStatusResponse) SetBody(v *GetOfficeSiteSsoStatusResponseBody) *GetOfficeSiteSsoStatusResponse {
	s.Body = v
	return s
}

type GetSpMetadataRequest struct {
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSpMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataRequest) GoString() string {
	return s.String()
}

func (s *GetSpMetadataRequest) SetDirectoryId(v string) *GetSpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetSpMetadataRequest) SetOfficeSiteId(v string) *GetSpMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetSpMetadataRequest) SetRegionId(v string) *GetSpMetadataRequest {
	s.RegionId = &v
	return s
}

type GetSpMetadataResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpMetadata *string `json:"SpMetadata,omitempty" xml:"SpMetadata,omitempty"`
}

func (s GetSpMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpMetadataResponseBody) SetRequestId(v string) *GetSpMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpMetadataResponseBody) SetSpMetadata(v string) *GetSpMetadataResponseBody {
	s.SpMetadata = &v
	return s
}

type GetSpMetadataResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataResponse) GoString() string {
	return s.String()
}

func (s *GetSpMetadataResponse) SetHeaders(v map[string]*string) *GetSpMetadataResponse {
	s.Headers = v
	return s
}

func (s *GetSpMetadataResponse) SetBody(v *GetSpMetadataResponseBody) *GetSpMetadataResponse {
	s.Body = v
	return s
}

type HandleSecurityEventsRequest struct {
	OperationCode   *string                                     `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	OperationParams *string                                     `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	RegionId        *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityEvent   []*HandleSecurityEventsRequestSecurityEvent `json:"SecurityEvent,omitempty" xml:"SecurityEvent,omitempty" type:"Repeated"`
}

func (s HandleSecurityEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequest) SetOperationCode(v string) *HandleSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationParams(v string) *HandleSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetRegionId(v string) *HandleSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetSecurityEvent(v []*HandleSecurityEventsRequestSecurityEvent) *HandleSecurityEventsRequest {
	s.SecurityEvent = v
	return s
}

type HandleSecurityEventsRequestSecurityEvent struct {
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	SecurityEventId *string `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
}

func (s HandleSecurityEventsRequestSecurityEvent) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsRequestSecurityEvent) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequestSecurityEvent) SetDesktopId(v string) *HandleSecurityEventsRequestSecurityEvent {
	s.DesktopId = &v
	return s
}

func (s *HandleSecurityEventsRequestSecurityEvent) SetSecurityEventId(v string) *HandleSecurityEventsRequestSecurityEvent {
	s.SecurityEventId = &v
	return s
}

type HandleSecurityEventsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s HandleSecurityEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsResponseBody) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponseBody) SetRequestId(v string) *HandleSecurityEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleSecurityEventsResponseBody) SetTaskId(v int64) *HandleSecurityEventsResponseBody {
	s.TaskId = &v
	return s
}

type HandleSecurityEventsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HandleSecurityEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandleSecurityEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponse) SetHeaders(v map[string]*string) *HandleSecurityEventsResponse {
	s.Headers = v
	return s
}

func (s *HandleSecurityEventsResponse) SetBody(v *HandleSecurityEventsResponseBody) *HandleSecurityEventsResponse {
	s.Body = v
	return s
}

type ListDirectoryUsersRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Filter      *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OUPath      *string `json:"OUPath,omitempty" xml:"OUPath,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDirectoryUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersRequest) SetDirectoryId(v string) *ListDirectoryUsersRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetFilter(v string) *ListDirectoryUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetMaxResults(v int32) *ListDirectoryUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetNextToken(v string) *ListDirectoryUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetOUPath(v string) *ListDirectoryUsersRequest {
	s.OUPath = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetRegionId(v string) *ListDirectoryUsersRequest {
	s.RegionId = &v
	return s
}

type ListDirectoryUsersResponseBody struct {
	NextToken *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*ListDirectoryUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListDirectoryUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponseBody) SetNextToken(v string) *ListDirectoryUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersResponseBody) SetRequestId(v string) *ListDirectoryUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectoryUsersResponseBody) SetUsers(v []*ListDirectoryUsersResponseBodyUsers) *ListDirectoryUsersResponseBody {
	s.Users = v
	return s
}

type ListDirectoryUsersResponseBodyUsers struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EndUser     *string `json:"EndUser,omitempty" xml:"EndUser,omitempty"`
}

func (s ListDirectoryUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponseBodyUsers) SetDisplayName(v string) *ListDirectoryUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListDirectoryUsersResponseBodyUsers) SetEndUser(v string) *ListDirectoryUsersResponseBodyUsers {
	s.EndUser = &v
	return s
}

type ListDirectoryUsersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDirectoryUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDirectoryUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponse) SetHeaders(v map[string]*string) *ListDirectoryUsersResponse {
	s.Headers = v
	return s
}

func (s *ListDirectoryUsersResponse) SetBody(v *ListDirectoryUsersResponseBody) *ListDirectoryUsersResponse {
	s.Body = v
	return s
}

type ListOfficeSiteOverviewRequest struct {
	ForceRefresh *bool     `json:"ForceRefresh,omitempty" xml:"ForceRefresh,omitempty"`
	MaxResults   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	QueryRange   *int32    `json:"QueryRange,omitempty" xml:"QueryRange,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOfficeSiteOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewRequest) SetForceRefresh(v bool) *ListOfficeSiteOverviewRequest {
	s.ForceRefresh = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetMaxResults(v int32) *ListOfficeSiteOverviewRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetNextToken(v string) *ListOfficeSiteOverviewRequest {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetOfficeSiteId(v []*string) *ListOfficeSiteOverviewRequest {
	s.OfficeSiteId = v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetQueryRange(v int32) *ListOfficeSiteOverviewRequest {
	s.QueryRange = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetRegionId(v string) *ListOfficeSiteOverviewRequest {
	s.RegionId = &v
	return s
}

type ListOfficeSiteOverviewResponseBody struct {
	NextToken                 *string                                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteOverviewResults []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults `json:"OfficeSiteOverviewResults,omitempty" xml:"OfficeSiteOverviewResults,omitempty" type:"Repeated"`
	RequestId                 *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOfficeSiteOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponseBody) SetNextToken(v string) *ListOfficeSiteOverviewResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBody) SetOfficeSiteOverviewResults(v []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) *ListOfficeSiteOverviewResponseBody {
	s.OfficeSiteOverviewResults = v
	return s
}

func (s *ListOfficeSiteOverviewResponseBody) SetRequestId(v string) *ListOfficeSiteOverviewResponseBody {
	s.RequestId = &v
	return s
}

type ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults struct {
	HasExpiredEdsCount          *int32  `json:"HasExpiredEdsCount,omitempty" xml:"HasExpiredEdsCount,omitempty"`
	HasExpiredEdsCountForGroup  *int32  `json:"HasExpiredEdsCountForGroup,omitempty" xml:"HasExpiredEdsCountForGroup,omitempty"`
	OfficeSiteId                *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName              *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	OfficeSiteStatus            *string `json:"OfficeSiteStatus,omitempty" xml:"OfficeSiteStatus,omitempty"`
	RegionId                    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RunningEdsCount             *int32  `json:"RunningEdsCount,omitempty" xml:"RunningEdsCount,omitempty"`
	RunningEdsCountForGroup     *int32  `json:"RunningEdsCountForGroup,omitempty" xml:"RunningEdsCountForGroup,omitempty"`
	TotalEdsCount               *int32  `json:"TotalEdsCount,omitempty" xml:"TotalEdsCount,omitempty"`
	TotalEdsCountForGroup       *int32  `json:"TotalEdsCountForGroup,omitempty" xml:"TotalEdsCountForGroup,omitempty"`
	VpcType                     *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	WillExpiredEdsCount         *int32  `json:"WillExpiredEdsCount,omitempty" xml:"WillExpiredEdsCount,omitempty"`
	WillExpiredEdsCountForGroup *int32  `json:"WillExpiredEdsCountForGroup,omitempty" xml:"WillExpiredEdsCountForGroup,omitempty"`
}

func (s ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetHasExpiredEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.HasExpiredEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetHasExpiredEdsCountForGroup(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.HasExpiredEdsCountForGroup = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteId(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteName(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteName = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteStatus(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteStatus = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetRegionId(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetRunningEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.RunningEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetRunningEdsCountForGroup(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.RunningEdsCountForGroup = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetTotalEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.TotalEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetTotalEdsCountForGroup(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.TotalEdsCountForGroup = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetVpcType(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.VpcType = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetWillExpiredEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.WillExpiredEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetWillExpiredEdsCountForGroup(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.WillExpiredEdsCountForGroup = &v
	return s
}

type ListOfficeSiteOverviewResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOfficeSiteOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOfficeSiteOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponse) SetHeaders(v map[string]*string) *ListOfficeSiteOverviewResponse {
	s.Headers = v
	return s
}

func (s *ListOfficeSiteOverviewResponse) SetBody(v *ListOfficeSiteOverviewResponseBody) *ListOfficeSiteOverviewResponse {
	s.Body = v
	return s
}

type ListOfficeSiteUsersRequest struct {
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OUPath       *string `json:"OUPath,omitempty" xml:"OUPath,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOfficeSiteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersRequest) SetFilter(v string) *ListOfficeSiteUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetMaxResults(v int32) *ListOfficeSiteUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetNextToken(v string) *ListOfficeSiteUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetOUPath(v string) *ListOfficeSiteUsersRequest {
	s.OUPath = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetOfficeSiteId(v string) *ListOfficeSiteUsersRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetRegionId(v string) *ListOfficeSiteUsersRequest {
	s.RegionId = &v
	return s
}

type ListOfficeSiteUsersResponseBody struct {
	NextToken *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*ListOfficeSiteUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListOfficeSiteUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponseBody) SetNextToken(v string) *ListOfficeSiteUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBody) SetRequestId(v string) *ListOfficeSiteUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBody) SetUsers(v []*ListOfficeSiteUsersResponseBodyUsers) *ListOfficeSiteUsersResponseBody {
	s.Users = v
	return s
}

type ListOfficeSiteUsersResponseBodyUsers struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EndUser     *string `json:"EndUser,omitempty" xml:"EndUser,omitempty"`
}

func (s ListOfficeSiteUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetDisplayName(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetEndUser(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.EndUser = &v
	return s
}

type ListOfficeSiteUsersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOfficeSiteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOfficeSiteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponse) SetHeaders(v map[string]*string) *ListOfficeSiteUsersResponse {
	s.Headers = v
	return s
}

func (s *ListOfficeSiteUsersResponse) SetBody(v *ListOfficeSiteUsersResponseBody) *ListOfficeSiteUsersResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	MaxResults   *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListUserAdOrganizationUnitsRequest struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUserAdOrganizationUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserAdOrganizationUnitsRequest) GoString() string {
	return s.String()
}

func (s *ListUserAdOrganizationUnitsRequest) SetOfficeSiteId(v string) *ListUserAdOrganizationUnitsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListUserAdOrganizationUnitsRequest) SetRegionId(v string) *ListUserAdOrganizationUnitsRequest {
	s.RegionId = &v
	return s
}

type ListUserAdOrganizationUnitsResponseBody struct {
	OUNames   []*ListUserAdOrganizationUnitsResponseBodyOUNames `json:"OUNames,omitempty" xml:"OUNames,omitempty" type:"Repeated"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserAdOrganizationUnitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserAdOrganizationUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAdOrganizationUnitsResponseBody) SetOUNames(v []*ListUserAdOrganizationUnitsResponseBodyOUNames) *ListUserAdOrganizationUnitsResponseBody {
	s.OUNames = v
	return s
}

func (s *ListUserAdOrganizationUnitsResponseBody) SetRequestId(v string) *ListUserAdOrganizationUnitsResponseBody {
	s.RequestId = &v
	return s
}

type ListUserAdOrganizationUnitsResponseBodyOUNames struct {
	OUName       *string `json:"OUName,omitempty" xml:"OUName,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s ListUserAdOrganizationUnitsResponseBodyOUNames) String() string {
	return tea.Prettify(s)
}

func (s ListUserAdOrganizationUnitsResponseBodyOUNames) GoString() string {
	return s.String()
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) SetOUName(v string) *ListUserAdOrganizationUnitsResponseBodyOUNames {
	s.OUName = &v
	return s
}

func (s *ListUserAdOrganizationUnitsResponseBodyOUNames) SetOfficeSiteId(v string) *ListUserAdOrganizationUnitsResponseBodyOUNames {
	s.OfficeSiteId = &v
	return s
}

type ListUserAdOrganizationUnitsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserAdOrganizationUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserAdOrganizationUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserAdOrganizationUnitsResponse) GoString() string {
	return s.String()
}

func (s *ListUserAdOrganizationUnitsResponse) SetHeaders(v map[string]*string) *ListUserAdOrganizationUnitsResponse {
	s.Headers = v
	return s
}

func (s *ListUserAdOrganizationUnitsResponse) SetBody(v *ListUserAdOrganizationUnitsResponseBody) *ListUserAdOrganizationUnitsResponse {
	s.Body = v
	return s
}

type LockVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s LockVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceRequest) SetRegionId(v string) *LockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *LockVirtualMFADeviceRequest) SetSerialNumber(v string) *LockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type LockVirtualMFADeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceResponseBody) SetRequestId(v string) *LockVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type LockVirtualMFADeviceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LockVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *LockVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *LockVirtualMFADeviceResponse) SetBody(v *LockVirtualMFADeviceResponseBody) *LockVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type ModifyADConnectorDirectoryRequest struct {
	AdHostname          *string   `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	DirectoryId         *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	OUName              *string   `json:"OUName,omitempty" xml:"OUName,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
}

func (s ModifyADConnectorDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryRequest) SetAdHostname(v string) *ModifyADConnectorDirectoryRequest {
	s.AdHostname = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDirectoryId(v string) *ModifyADConnectorDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDirectoryName(v string) *ModifyADConnectorDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDnsAddress(v []*string) *ModifyADConnectorDirectoryRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainName(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainPassword(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainPassword = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainUserName(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainUserName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetMfaEnabled(v bool) *ModifyADConnectorDirectoryRequest {
	s.MfaEnabled = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetOUName(v string) *ModifyADConnectorDirectoryRequest {
	s.OUName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetRegionId(v string) *ModifyADConnectorDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetSubDomainDnsAddress(v []*string) *ModifyADConnectorDirectoryRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetSubDomainName(v string) *ModifyADConnectorDirectoryRequest {
	s.SubDomainName = &v
	return s
}

type ModifyADConnectorDirectoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyADConnectorDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryResponseBody) SetRequestId(v string) *ModifyADConnectorDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type ModifyADConnectorDirectoryResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyADConnectorDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyADConnectorDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryResponse) SetHeaders(v map[string]*string) *ModifyADConnectorDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyADConnectorDirectoryResponse) SetBody(v *ModifyADConnectorDirectoryResponseBody) *ModifyADConnectorDirectoryResponse {
	s.Body = v
	return s
}

type ModifyADConnectorOfficeSiteRequest struct {
	AdHostname          *string   `json:"AdHostname,omitempty" xml:"AdHostname,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	OUName              *string   `json:"OUName,omitempty" xml:"OUName,omitempty"`
	OfficeSiteId        *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName      *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
}

func (s ModifyADConnectorOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteRequest) SetAdHostname(v string) *ModifyADConnectorOfficeSiteRequest {
	s.AdHostname = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainPassword(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainPassword = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainUserName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainUserName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetMfaEnabled(v bool) *ModifyADConnectorOfficeSiteRequest {
	s.MfaEnabled = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOUName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OUName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOfficeSiteId(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOfficeSiteName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetRegionId(v string) *ModifyADConnectorOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetSubDomainDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetSubDomainName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.SubDomainName = &v
	return s
}

type ModifyADConnectorOfficeSiteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyADConnectorOfficeSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteResponseBody) SetRequestId(v string) *ModifyADConnectorOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

type ModifyADConnectorOfficeSiteResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyADConnectorOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyADConnectorOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteResponse) SetHeaders(v map[string]*string) *ModifyADConnectorOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *ModifyADConnectorOfficeSiteResponse) SetBody(v *ModifyADConnectorOfficeSiteResponseBody) *ModifyADConnectorOfficeSiteResponse {
	s.Body = v
	return s
}

type ModifyBundleRequest struct {
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	BundleName  *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleRequest) GoString() string {
	return s.String()
}

func (s *ModifyBundleRequest) SetBundleId(v string) *ModifyBundleRequest {
	s.BundleId = &v
	return s
}

func (s *ModifyBundleRequest) SetBundleName(v string) *ModifyBundleRequest {
	s.BundleName = &v
	return s
}

func (s *ModifyBundleRequest) SetDescription(v string) *ModifyBundleRequest {
	s.Description = &v
	return s
}

func (s *ModifyBundleRequest) SetImageId(v string) *ModifyBundleRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyBundleRequest) SetLanguage(v string) *ModifyBundleRequest {
	s.Language = &v
	return s
}

func (s *ModifyBundleRequest) SetRegionId(v string) *ModifyBundleRequest {
	s.RegionId = &v
	return s
}

type ModifyBundleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBundleResponseBody) SetRequestId(v string) *ModifyBundleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBundleResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleResponse) GoString() string {
	return s.String()
}

func (s *ModifyBundleResponse) SetHeaders(v map[string]*string) *ModifyBundleResponse {
	s.Headers = v
	return s
}

func (s *ModifyBundleResponse) SetBody(v *ModifyBundleResponseBody) *ModifyBundleResponse {
	s.Body = v
	return s
}

type ModifyDesktopChargeTypeRequest struct {
	AutoPay     *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ChargeType  *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DesktopId   []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	Period      *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit  *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDesktopChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopChargeTypeRequest) SetAutoPay(v bool) *ModifyDesktopChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetChargeType(v string) *ModifyDesktopChargeTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetDesktopId(v []*string) *ModifyDesktopChargeTypeRequest {
	s.DesktopId = v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetPeriod(v int32) *ModifyDesktopChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetPeriodUnit(v string) *ModifyDesktopChargeTypeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetPromotionId(v string) *ModifyDesktopChargeTypeRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetRegionId(v string) *ModifyDesktopChargeTypeRequest {
	s.RegionId = &v
	return s
}

type ModifyDesktopChargeTypeResponseBody struct {
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	OrderId   *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopChargeTypeResponseBody) SetDesktopId(v []*string) *ModifyDesktopChargeTypeResponseBody {
	s.DesktopId = v
	return s
}

func (s *ModifyDesktopChargeTypeResponseBody) SetOrderId(v string) *ModifyDesktopChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDesktopChargeTypeResponseBody) SetRequestId(v string) *ModifyDesktopChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopChargeTypeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyDesktopChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopChargeTypeResponse) SetBody(v *ModifyDesktopChargeTypeResponseBody) *ModifyDesktopChargeTypeResponse {
	s.Body = v
	return s
}

type ModifyDesktopGroupRequest struct {
	AllowAutoSetup   *int32  `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	AllowBufferCount *int32  `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	BindAmount       *int64  `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	Comments         *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	DesktopGroupId   *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	ImageId          *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	KeepDuration     *int64  `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	LoadPolicy       *int64  `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	MaxDesktopsCount *int32  `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	MinDesktopsCount *int32  `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	OwnBundleId      *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	PolicyGroupId    *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResetType        *int64  `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	ScaleStrategyId  *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
}

func (s ModifyDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupRequest) SetAllowAutoSetup(v int32) *ModifyDesktopGroupRequest {
	s.AllowAutoSetup = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetAllowBufferCount(v int32) *ModifyDesktopGroupRequest {
	s.AllowBufferCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetBindAmount(v int64) *ModifyDesktopGroupRequest {
	s.BindAmount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetComments(v string) *ModifyDesktopGroupRequest {
	s.Comments = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDesktopGroupId(v string) *ModifyDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDesktopGroupName(v string) *ModifyDesktopGroupRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetImageId(v string) *ModifyDesktopGroupRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetKeepDuration(v int64) *ModifyDesktopGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetLoadPolicy(v int64) *ModifyDesktopGroupRequest {
	s.LoadPolicy = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetMaxDesktopsCount(v int32) *ModifyDesktopGroupRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetMinDesktopsCount(v int32) *ModifyDesktopGroupRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetOwnBundleId(v string) *ModifyDesktopGroupRequest {
	s.OwnBundleId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetRegionId(v string) *ModifyDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetResetType(v int64) *ModifyDesktopGroupRequest {
	s.ResetType = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetScaleStrategyId(v string) *ModifyDesktopGroupRequest {
	s.ScaleStrategyId = &v
	return s
}

type ModifyDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupResponseBody) SetRequestId(v string) *ModifyDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupResponse) SetHeaders(v map[string]*string) *ModifyDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopGroupResponse) SetBody(v *ModifyDesktopGroupResponseBody) *ModifyDesktopGroupResponse {
	s.Body = v
	return s
}

type ModifyDesktopHostNameRequest struct {
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	NewHostName *string `json:"NewHostName,omitempty" xml:"NewHostName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDesktopHostNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopHostNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopHostNameRequest) SetDesktopId(v string) *ModifyDesktopHostNameRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopHostNameRequest) SetNewHostName(v string) *ModifyDesktopHostNameRequest {
	s.NewHostName = &v
	return s
}

func (s *ModifyDesktopHostNameRequest) SetRegionId(v string) *ModifyDesktopHostNameRequest {
	s.RegionId = &v
	return s
}

type ModifyDesktopHostNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopHostNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopHostNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopHostNameResponseBody) SetRequestId(v string) *ModifyDesktopHostNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopHostNameResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopHostNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopHostNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopHostNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopHostNameResponse) SetHeaders(v map[string]*string) *ModifyDesktopHostNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopHostNameResponse) SetBody(v *ModifyDesktopHostNameResponseBody) *ModifyDesktopHostNameResponse {
	s.Body = v
	return s
}

type ModifyDesktopNameRequest struct {
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	NewDesktopName *string `json:"NewDesktopName,omitempty" xml:"NewDesktopName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDesktopNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameRequest) SetDesktopId(v string) *ModifyDesktopNameRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetNewDesktopName(v string) *ModifyDesktopNameRequest {
	s.NewDesktopName = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetRegionId(v string) *ModifyDesktopNameRequest {
	s.RegionId = &v
	return s
}

type ModifyDesktopNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameResponseBody) SetRequestId(v string) *ModifyDesktopNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopNameResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameResponse) SetHeaders(v map[string]*string) *ModifyDesktopNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopNameResponse) SetBody(v *ModifyDesktopNameResponseBody) *ModifyDesktopNameResponse {
	s.Body = v
	return s
}

type ModifyDesktopSpecRequest struct {
	AutoPay                  *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	DesktopId                *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopType              *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	PromotionId              *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RootDiskSizeGib          *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	UserDiskSizeGib          *int32  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s ModifyDesktopSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecRequest) SetAutoPay(v bool) *ModifyDesktopSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetDesktopId(v string) *ModifyDesktopSpecRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetDesktopType(v string) *ModifyDesktopSpecRequest {
	s.DesktopType = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetPromotionId(v string) *ModifyDesktopSpecRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetRegionId(v string) *ModifyDesktopSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetRootDiskSizeGib(v int32) *ModifyDesktopSpecRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetUserDiskPerformanceLevel(v string) *ModifyDesktopSpecRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetUserDiskSizeGib(v int32) *ModifyDesktopSpecRequest {
	s.UserDiskSizeGib = &v
	return s
}

type ModifyDesktopSpecResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecResponseBody) SetOrderId(v string) *ModifyDesktopSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDesktopSpecResponseBody) SetRequestId(v string) *ModifyDesktopSpecResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopSpecResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecResponse) SetHeaders(v map[string]*string) *ModifyDesktopSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopSpecResponse) SetBody(v *ModifyDesktopSpecResponseBody) *ModifyDesktopSpecResponse {
	s.Body = v
	return s
}

type ModifyDesktopsPolicyGroupRequest struct {
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDesktopsPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupRequest) SetDesktopId(v []*string) *ModifyDesktopsPolicyGroupRequest {
	s.DesktopId = v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetRegionId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.RegionId = &v
	return s
}

type ModifyDesktopsPolicyGroupResponseBody struct {
	ModifyResults []*ModifyDesktopsPolicyGroupResponseBodyModifyResults `json:"ModifyResults,omitempty" xml:"ModifyResults,omitempty" type:"Repeated"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopsPolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponseBody) SetModifyResults(v []*ModifyDesktopsPolicyGroupResponseBodyModifyResults) *ModifyDesktopsPolicyGroupResponseBody {
	s.ModifyResults = v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBody) SetRequestId(v string) *ModifyDesktopsPolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopsPolicyGroupResponseBodyModifyResults struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ModifyDesktopsPolicyGroupResponseBodyModifyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponseBodyModifyResults) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetCode(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.Code = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetDesktopId(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetMessage(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.Message = &v
	return s
}

type ModifyDesktopsPolicyGroupResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopsPolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopsPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponse) SetHeaders(v map[string]*string) *ModifyDesktopsPolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponse) SetBody(v *ModifyDesktopsPolicyGroupResponseBody) *ModifyDesktopsPolicyGroupResponse {
	s.Body = v
	return s
}

type ModifyEntitlementRequest struct {
	DesktopId *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyEntitlementRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementRequest) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementRequest) SetDesktopId(v string) *ModifyEntitlementRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyEntitlementRequest) SetEndUserId(v []*string) *ModifyEntitlementRequest {
	s.EndUserId = v
	return s
}

func (s *ModifyEntitlementRequest) SetRegionId(v string) *ModifyEntitlementRequest {
	s.RegionId = &v
	return s
}

type ModifyEntitlementResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEntitlementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementResponseBody) SetRequestId(v string) *ModifyEntitlementResponseBody {
	s.RequestId = &v
	return s
}

type ModifyEntitlementResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyEntitlementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyEntitlementResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementResponse) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementResponse) SetHeaders(v map[string]*string) *ModifyEntitlementResponse {
	s.Headers = v
	return s
}

func (s *ModifyEntitlementResponse) SetBody(v *ModifyEntitlementResponseBody) *ModifyEntitlementResponse {
	s.Body = v
	return s
}

type ModifyImageAttributeRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyImageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) SetDescription(v string) *ModifyImageAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetName(v string) *ModifyImageAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetRegionId(v string) *ModifyImageAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyImageAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponseBody) SetRequestId(v string) *ModifyImageAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageAttributeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyImageAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyImageAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponse) SetHeaders(v map[string]*string) *ModifyImageAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageAttributeResponse) SetBody(v *ModifyImageAttributeResponseBody) *ModifyImageAttributeResponse {
	s.Body = v
	return s
}

type ModifyNASDefaultMountTargetRequest struct {
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyNASDefaultMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNASDefaultMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetRequest) SetFileSystemId(v string) *ModifyNASDefaultMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) SetMountTargetDomain(v string) *ModifyNASDefaultMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) SetRegionId(v string) *ModifyNASDefaultMountTargetRequest {
	s.RegionId = &v
	return s
}

type ModifyNASDefaultMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNASDefaultMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNASDefaultMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetResponseBody) SetRequestId(v string) *ModifyNASDefaultMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNASDefaultMountTargetResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNASDefaultMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNASDefaultMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNASDefaultMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetResponse) SetHeaders(v map[string]*string) *ModifyNASDefaultMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyNASDefaultMountTargetResponse) SetBody(v *ModifyNASDefaultMountTargetResponseBody) *ModifyNASDefaultMountTargetResponse {
	s.Body = v
	return s
}

type ModifyNetworkPackageRequest struct {
	Bandwidth        *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyNetworkPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageRequest) SetBandwidth(v int32) *ModifyNetworkPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyNetworkPackageRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageRequest) SetRegionId(v string) *ModifyNetworkPackageRequest {
	s.RegionId = &v
	return s
}

type ModifyNetworkPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageResponseBody) SetRequestId(v string) *ModifyNetworkPackageResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNetworkPackageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNetworkPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNetworkPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageResponse) SetHeaders(v map[string]*string) *ModifyNetworkPackageResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkPackageResponse) SetBody(v *ModifyNetworkPackageResponseBody) *ModifyNetworkPackageResponse {
	s.Body = v
	return s
}

type ModifyNetworkPackageBandwidthRequest struct {
	AutoPay          *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Bandwidth        *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	PromotionId      *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyNetworkPackageBandwidthRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageBandwidthRequest) SetAutoPay(v bool) *ModifyNetworkPackageBandwidthRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetBandwidth(v int32) *ModifyNetworkPackageBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageBandwidthRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetPromotionId(v string) *ModifyNetworkPackageBandwidthRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthRequest) SetRegionId(v string) *ModifyNetworkPackageBandwidthRequest {
	s.RegionId = &v
	return s
}

type ModifyNetworkPackageBandwidthResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkPackageBandwidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageBandwidthResponseBody) SetOrderId(v string) *ModifyNetworkPackageBandwidthResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyNetworkPackageBandwidthResponseBody) SetRequestId(v string) *ModifyNetworkPackageBandwidthResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNetworkPackageBandwidthResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNetworkPackageBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNetworkPackageBandwidthResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageBandwidthResponse) SetHeaders(v map[string]*string) *ModifyNetworkPackageBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkPackageBandwidthResponse) SetBody(v *ModifyNetworkPackageBandwidthResponseBody) *ModifyNetworkPackageBandwidthResponse {
	s.Body = v
	return s
}

type ModifyNetworkPackageEnabledRequest struct {
	Enabled          *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyNetworkPackageEnabledRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageEnabledRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledRequest) SetEnabled(v bool) *ModifyNetworkPackageEnabledRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageEnabledRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) SetRegionId(v string) *ModifyNetworkPackageEnabledRequest {
	s.RegionId = &v
	return s
}

type ModifyNetworkPackageEnabledResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkPackageEnabledResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledResponseBody) SetRequestId(v string) *ModifyNetworkPackageEnabledResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNetworkPackageEnabledResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNetworkPackageEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNetworkPackageEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageEnabledResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledResponse) SetHeaders(v map[string]*string) *ModifyNetworkPackageEnabledResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkPackageEnabledResponse) SetBody(v *ModifyNetworkPackageEnabledResponseBody) *ModifyNetworkPackageEnabledResponse {
	s.Body = v
	return s
}

type ModifyOfficeSiteAttributeRequest struct {
	DesktopAccessType    *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	NeedVerifyLoginRisk  *bool   `json:"NeedVerifyLoginRisk,omitempty" xml:"NeedVerifyLoginRisk,omitempty"`
	NeedVerifyZeroDevice *bool   `json:"NeedVerifyZeroDevice,omitempty" xml:"NeedVerifyZeroDevice,omitempty"`
	OfficeSiteId         *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteName       *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeRequest) SetDesktopAccessType(v string) *ModifyOfficeSiteAttributeRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetNeedVerifyLoginRisk(v bool) *ModifyOfficeSiteAttributeRequest {
	s.NeedVerifyLoginRisk = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetNeedVerifyZeroDevice(v bool) *ModifyOfficeSiteAttributeRequest {
	s.NeedVerifyZeroDevice = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteAttributeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetOfficeSiteName(v string) *ModifyOfficeSiteAttributeRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetRegionId(v string) *ModifyOfficeSiteAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyOfficeSiteAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeResponseBody) SetRequestId(v string) *ModifyOfficeSiteAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOfficeSiteAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOfficeSiteAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOfficeSiteAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteAttributeResponse) SetBody(v *ModifyOfficeSiteAttributeResponseBody) *ModifyOfficeSiteAttributeResponse {
	s.Body = v
	return s
}

type ModifyOfficeSiteCrossDesktopAccessRequest struct {
	EnableCrossDesktopAccess *bool   `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
	OfficeSiteId             *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteCrossDesktopAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetEnableCrossDesktopAccess(v bool) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetRegionId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.RegionId = &v
	return s
}

type ModifyOfficeSiteCrossDesktopAccessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteCrossDesktopAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponseBody) SetRequestId(v string) *ModifyOfficeSiteCrossDesktopAccessResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOfficeSiteCrossDesktopAccessResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOfficeSiteCrossDesktopAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOfficeSiteCrossDesktopAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteCrossDesktopAccessResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) SetBody(v *ModifyOfficeSiteCrossDesktopAccessResponseBody) *ModifyOfficeSiteCrossDesktopAccessResponse {
	s.Body = v
	return s
}

type ModifyOfficeSiteMfaEnabledRequest struct {
	MfaEnabled   *bool   `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteMfaEnabledRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetMfaEnabled(v bool) *ModifyOfficeSiteMfaEnabledRequest {
	s.MfaEnabled = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteMfaEnabledRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetRegionId(v string) *ModifyOfficeSiteMfaEnabledRequest {
	s.RegionId = &v
	return s
}

type ModifyOfficeSiteMfaEnabledResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteMfaEnabledResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledResponseBody) SetRequestId(v string) *ModifyOfficeSiteMfaEnabledResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOfficeSiteMfaEnabledResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOfficeSiteMfaEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOfficeSiteMfaEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteMfaEnabledResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledResponse) SetBody(v *ModifyOfficeSiteMfaEnabledResponseBody) *ModifyOfficeSiteMfaEnabledResponse {
	s.Body = v
	return s
}

type ModifyOperateVulRequest struct {
	OperateType *string                           `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Reason      *string                           `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RegionId    *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type        *string                           `json:"Type,omitempty" xml:"Type,omitempty"`
	VulInfo     []*ModifyOperateVulRequestVulInfo `json:"VulInfo,omitempty" xml:"VulInfo,omitempty" type:"Repeated"`
}

func (s ModifyOperateVulRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulRequest) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulRequest) SetOperateType(v string) *ModifyOperateVulRequest {
	s.OperateType = &v
	return s
}

func (s *ModifyOperateVulRequest) SetReason(v string) *ModifyOperateVulRequest {
	s.Reason = &v
	return s
}

func (s *ModifyOperateVulRequest) SetRegionId(v string) *ModifyOperateVulRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOperateVulRequest) SetType(v string) *ModifyOperateVulRequest {
	s.Type = &v
	return s
}

func (s *ModifyOperateVulRequest) SetVulInfo(v []*ModifyOperateVulRequestVulInfo) *ModifyOperateVulRequest {
	s.VulInfo = v
	return s
}

type ModifyOperateVulRequestVulInfo struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ModifyOperateVulRequestVulInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulRequestVulInfo) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulRequestVulInfo) SetDesktopId(v string) *ModifyOperateVulRequestVulInfo {
	s.DesktopId = &v
	return s
}

func (s *ModifyOperateVulRequestVulInfo) SetName(v string) *ModifyOperateVulRequestVulInfo {
	s.Name = &v
	return s
}

func (s *ModifyOperateVulRequestVulInfo) SetTag(v string) *ModifyOperateVulRequestVulInfo {
	s.Tag = &v
	return s
}

type ModifyOperateVulResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOperateVulResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponseBody) SetRequestId(v string) *ModifyOperateVulResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOperateVulResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOperateVulResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOperateVulResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulResponse) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponse) SetHeaders(v map[string]*string) *ModifyOperateVulResponse {
	s.Headers = v
	return s
}

func (s *ModifyOperateVulResponse) SetBody(v *ModifyOperateVulResponseBody) *ModifyOperateVulResponse {
	s.Body = v
	return s
}

type ModifyPolicyGroupRequest struct {
	AuthorizeAccessPolicyRule   []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule   `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	AuthorizeSecurityPolicyRule []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	ClientType                  []*ModifyPolicyGroupRequestClientType                  `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	Clipboard                   *string                                                `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	DomainList                  *string                                                `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	GpuAcceleration             *string                                                `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	Html5Access                 *string                                                `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	Html5FileTransfer           *string                                                `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	LocalDrive                  *string                                                `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	Name                        *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyGroupId               *string                                                `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PreemptLogin                *string                                                `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	PreemptLoginUser            []*string                                              `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	PrinterRedirection          *string                                                `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	RegionId                    *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RevokeAccessPolicyRule      []*ModifyPolicyGroupRequestRevokeAccessPolicyRule      `json:"RevokeAccessPolicyRule,omitempty" xml:"RevokeAccessPolicyRule,omitempty" type:"Repeated"`
	RevokeSecurityPolicyRule    []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule    `json:"RevokeSecurityPolicyRule,omitempty" xml:"RevokeSecurityPolicyRule,omitempty" type:"Repeated"`
	UsbRedirect                 *string                                                `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	UsbSupplyRedirectRule       []*ModifyPolicyGroupRequestUsbSupplyRedirectRule       `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	VisualQuality               *string                                                `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	Watermark                   *string                                                `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	WatermarkCustomText         *string                                                `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency       *string                                                `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	WatermarkType               *string                                                `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s ModifyPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequest) SetAuthorizeAccessPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) *ModifyPolicyGroupRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetAuthorizeSecurityPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) *ModifyPolicyGroupRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClientType(v []*ModifyPolicyGroupRequestClientType) *ModifyPolicyGroupRequest {
	s.ClientType = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClipboard(v string) *ModifyPolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetDomainList(v string) *ModifyPolicyGroupRequest {
	s.DomainList = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetGpuAcceleration(v string) *ModifyPolicyGroupRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5Access(v string) *ModifyPolicyGroupRequest {
	s.Html5Access = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5FileTransfer(v string) *ModifyPolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLocalDrive(v string) *ModifyPolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetName(v string) *ModifyPolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPreemptLogin(v string) *ModifyPolicyGroupRequest {
	s.PreemptLogin = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPreemptLoginUser(v []*string) *ModifyPolicyGroupRequest {
	s.PreemptLoginUser = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPrinterRedirection(v string) *ModifyPolicyGroupRequest {
	s.PrinterRedirection = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRegionId(v string) *ModifyPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRevokeAccessPolicyRule(v []*ModifyPolicyGroupRequestRevokeAccessPolicyRule) *ModifyPolicyGroupRequest {
	s.RevokeAccessPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRevokeSecurityPolicyRule(v []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule) *ModifyPolicyGroupRequest {
	s.RevokeSecurityPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetUsbRedirect(v string) *ModifyPolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetUsbSupplyRedirectRule(v []*ModifyPolicyGroupRequestUsbSupplyRedirectRule) *ModifyPolicyGroupRequest {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetVisualQuality(v string) *ModifyPolicyGroupRequest {
	s.VisualQuality = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermark(v string) *ModifyPolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkCustomText(v string) *ModifyPolicyGroupRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkTransparency(v string) *ModifyPolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkType(v string) *ModifyPolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

type ModifyPolicyGroupRequestAuthorizeAccessPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

type ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetType(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

type ModifyPolicyGroupRequestClientType struct {
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyPolicyGroupRequestClientType) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestClientType) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestClientType) SetClientType(v string) *ModifyPolicyGroupRequestClientType {
	s.ClientType = &v
	return s
}

func (s *ModifyPolicyGroupRequestClientType) SetStatus(v string) *ModifyPolicyGroupRequestClientType {
	s.Status = &v
	return s
}

type ModifyPolicyGroupRequestRevokeAccessPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPolicyGroupRequestRevokeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestRevokeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	s.Description = &v
	return s
}

type ModifyPolicyGroupRequestRevokeSecurityPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyPolicyGroupRequestRevokeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetIpProtocol(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPolicy(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPortRange(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPriority(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetType(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Type = &v
	return s
}

type ModifyPolicyGroupRequestUsbSupplyRedirectRule struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceClass     *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	DeviceSubclass  *string `json:"DeviceSubclass,omitempty" xml:"DeviceSubclass,omitempty"`
	ProductId       *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	UsbRedirectType *int64  `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	UsbRuleType     *int64  `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	VendorId        *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s ModifyPolicyGroupRequestUsbSupplyRedirectRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetDescription(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetDeviceClass(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.DeviceClass = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetDeviceSubclass(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.DeviceSubclass = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetProductId(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetUsbRedirectType(v int64) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetUsbRuleType(v int64) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetVendorId(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

type ModifyPolicyGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponseBody) SetRequestId(v string) *ModifyPolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPolicyGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponse) SetHeaders(v map[string]*string) *ModifyPolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyGroupResponse) SetBody(v *ModifyPolicyGroupResponseBody) *ModifyPolicyGroupResponse {
	s.Body = v
	return s
}

type ModifyUserEntitlementRequest struct {
	AuthorizeDesktopId []*string `json:"AuthorizeDesktopId,omitempty" xml:"AuthorizeDesktopId,omitempty" type:"Repeated"`
	EndUserId          []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RevokeDesktopId    []*string `json:"RevokeDesktopId,omitempty" xml:"RevokeDesktopId,omitempty" type:"Repeated"`
}

func (s ModifyUserEntitlementRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserEntitlementRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserEntitlementRequest) SetAuthorizeDesktopId(v []*string) *ModifyUserEntitlementRequest {
	s.AuthorizeDesktopId = v
	return s
}

func (s *ModifyUserEntitlementRequest) SetEndUserId(v []*string) *ModifyUserEntitlementRequest {
	s.EndUserId = v
	return s
}

func (s *ModifyUserEntitlementRequest) SetRegionId(v string) *ModifyUserEntitlementRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserEntitlementRequest) SetRevokeDesktopId(v []*string) *ModifyUserEntitlementRequest {
	s.RevokeDesktopId = v
	return s
}

type ModifyUserEntitlementResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserEntitlementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserEntitlementResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserEntitlementResponseBody) SetRequestId(v string) *ModifyUserEntitlementResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserEntitlementResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserEntitlementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserEntitlementResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserEntitlementResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserEntitlementResponse) SetHeaders(v map[string]*string) *ModifyUserEntitlementResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserEntitlementResponse) SetBody(v *ModifyUserEntitlementResponseBody) *ModifyUserEntitlementResponse {
	s.Body = v
	return s
}

type ModifyUserToDesktopGroupRequest struct {
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	NewEndUserIds  []*string `json:"NewEndUserIds,omitempty" xml:"NewEndUserIds,omitempty" type:"Repeated"`
	OldEndUserIds  []*string `json:"OldEndUserIds,omitempty" xml:"OldEndUserIds,omitempty" type:"Repeated"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyUserToDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserToDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupRequest) SetDesktopGroupId(v string) *ModifyUserToDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetNewEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest {
	s.NewEndUserIds = v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetOldEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest {
	s.OldEndUserIds = v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetRegionId(v string) *ModifyUserToDesktopGroupRequest {
	s.RegionId = &v
	return s
}

type ModifyUserToDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserToDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserToDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupResponseBody) SetRequestId(v string) *ModifyUserToDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserToDesktopGroupResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserToDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserToDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserToDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupResponse) SetHeaders(v map[string]*string) *ModifyUserToDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserToDesktopGroupResponse) SetBody(v *ModifyUserToDesktopGroupResponseBody) *ModifyUserToDesktopGroupResponse {
	s.Body = v
	return s
}

type OperateVulsRequest struct {
	DesktopId    []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	OperateType  *string   `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Precondition *int32    `json:"Precondition,omitempty" xml:"Precondition,omitempty"`
	Reason       *string   `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type         *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	VulName      []*string `json:"VulName,omitempty" xml:"VulName,omitempty" type:"Repeated"`
}

func (s OperateVulsRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateVulsRequest) GoString() string {
	return s.String()
}

func (s *OperateVulsRequest) SetDesktopId(v []*string) *OperateVulsRequest {
	s.DesktopId = v
	return s
}

func (s *OperateVulsRequest) SetOperateType(v string) *OperateVulsRequest {
	s.OperateType = &v
	return s
}

func (s *OperateVulsRequest) SetPrecondition(v int32) *OperateVulsRequest {
	s.Precondition = &v
	return s
}

func (s *OperateVulsRequest) SetReason(v string) *OperateVulsRequest {
	s.Reason = &v
	return s
}

func (s *OperateVulsRequest) SetRegionId(v string) *OperateVulsRequest {
	s.RegionId = &v
	return s
}

func (s *OperateVulsRequest) SetType(v string) *OperateVulsRequest {
	s.Type = &v
	return s
}

func (s *OperateVulsRequest) SetVulName(v []*string) *OperateVulsRequest {
	s.VulName = v
	return s
}

type OperateVulsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateVulsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateVulsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateVulsResponseBody) SetRequestId(v string) *OperateVulsResponseBody {
	s.RequestId = &v
	return s
}

type OperateVulsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateVulsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateVulsResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateVulsResponse) GoString() string {
	return s.String()
}

func (s *OperateVulsResponse) SetHeaders(v map[string]*string) *OperateVulsResponse {
	s.Headers = v
	return s
}

func (s *OperateVulsResponse) SetBody(v *OperateVulsResponseBody) *OperateVulsResponse {
	s.Body = v
	return s
}

type RebootDesktopsRequest struct {
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

type RebootDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponseBody) SetRequestId(v string) *RebootDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type RebootDesktopsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebootDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebootDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponse) SetHeaders(v map[string]*string) *RebootDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebootDesktopsResponse) SetBody(v *RebootDesktopsResponseBody) *RebootDesktopsResponse {
	s.Body = v
	return s
}

type RebuildDesktopsRequest struct {
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	ImageId   *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebuildDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsRequest) SetDesktopId(v []*string) *RebuildDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RebuildDesktopsRequest) SetImageId(v string) *RebuildDesktopsRequest {
	s.ImageId = &v
	return s
}

func (s *RebuildDesktopsRequest) SetRegionId(v string) *RebuildDesktopsRequest {
	s.RegionId = &v
	return s
}

type RebuildDesktopsResponseBody struct {
	RebuildResults []*RebuildDesktopsResponseBodyRebuildResults `json:"RebuildResults,omitempty" xml:"RebuildResults,omitempty" type:"Repeated"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebuildDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponseBody) SetRebuildResults(v []*RebuildDesktopsResponseBodyRebuildResults) *RebuildDesktopsResponseBody {
	s.RebuildResults = v
	return s
}

func (s *RebuildDesktopsResponseBody) SetRequestId(v string) *RebuildDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type RebuildDesktopsResponseBodyRebuildResults struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RebuildDesktopsResponseBodyRebuildResults) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponseBodyRebuildResults) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetCode(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.Code = &v
	return s
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetDesktopId(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.DesktopId = &v
	return s
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetMessage(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.Message = &v
	return s
}

type RebuildDesktopsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebuildDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebuildDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponse) SetHeaders(v map[string]*string) *RebuildDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebuildDesktopsResponse) SetBody(v *RebuildDesktopsResponseBody) *RebuildDesktopsResponse {
	s.Body = v
	return s
}

type RemoveUserFromDesktopGroupRequest struct {
	DesktopGroupId  *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	EndUserIds      []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveUserFromDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupRequest) SetDesktopGroupId(v string) *RemoveUserFromDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetDesktopGroupIds(v []*string) *RemoveUserFromDesktopGroupRequest {
	s.DesktopGroupIds = v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetEndUserIds(v []*string) *RemoveUserFromDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetRegionId(v string) *RemoveUserFromDesktopGroupRequest {
	s.RegionId = &v
	return s
}

type RemoveUserFromDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupResponseBody) SetRequestId(v string) *RemoveUserFromDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUserFromDesktopGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUserFromDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUserFromDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupResponse) SetHeaders(v map[string]*string) *RemoveUserFromDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromDesktopGroupResponse) SetBody(v *RemoveUserFromDesktopGroupResponseBody) *RemoveUserFromDesktopGroupResponse {
	s.Body = v
	return s
}

type RenewDesktopsRequest struct {
	AutoPay     *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	DesktopId   []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	Period      *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit  *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenewDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RenewDesktopsRequest) SetAutoPay(v bool) *RenewDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewDesktopsRequest) SetDesktopId(v []*string) *RenewDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RenewDesktopsRequest) SetPeriod(v int32) *RenewDesktopsRequest {
	s.Period = &v
	return s
}

func (s *RenewDesktopsRequest) SetPeriodUnit(v string) *RenewDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDesktopsRequest) SetPromotionId(v string) *RenewDesktopsRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewDesktopsRequest) SetRegionId(v string) *RenewDesktopsRequest {
	s.RegionId = &v
	return s
}

type RenewDesktopsResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponseBody) SetOrderId(v string) *RenewDesktopsResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewDesktopsResponseBody) SetRequestId(v string) *RenewDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type RenewDesktopsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponse) SetHeaders(v map[string]*string) *RenewDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RenewDesktopsResponse) SetBody(v *RenewDesktopsResponseBody) *RenewDesktopsResponse {
	s.Body = v
	return s
}

type RenewNetworkPackagesRequest struct {
	AutoPay          *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	Period           *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit       *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId      *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenewNetworkPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *RenewNetworkPackagesRequest) SetAutoPay(v bool) *RenewNetworkPackagesRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetNetworkPackageId(v []*string) *RenewNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPeriod(v int32) *RenewNetworkPackagesRequest {
	s.Period = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPeriodUnit(v string) *RenewNetworkPackagesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPromotionId(v string) *RenewNetworkPackagesRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetRegionId(v string) *RenewNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

type RenewNetworkPackagesResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewNetworkPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewNetworkPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewNetworkPackagesResponseBody) SetOrderId(v string) *RenewNetworkPackagesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewNetworkPackagesResponseBody) SetRequestId(v string) *RenewNetworkPackagesResponseBody {
	s.RequestId = &v
	return s
}

type RenewNetworkPackagesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewNetworkPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewNetworkPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *RenewNetworkPackagesResponse) SetHeaders(v map[string]*string) *RenewNetworkPackagesResponse {
	s.Headers = v
	return s
}

func (s *RenewNetworkPackagesResponse) SetBody(v *RenewNetworkPackagesResponseBody) *RenewNetworkPackagesResponse {
	s.Body = v
	return s
}

type ResetNASDefaultMountTargetRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetNASDefaultMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetNASDefaultMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetRequest) SetFileSystemId(v string) *ResetNASDefaultMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ResetNASDefaultMountTargetRequest) SetRegionId(v string) *ResetNASDefaultMountTargetRequest {
	s.RegionId = &v
	return s
}

type ResetNASDefaultMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetNASDefaultMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetNASDefaultMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetResponseBody) SetRequestId(v string) *ResetNASDefaultMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type ResetNASDefaultMountTargetResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetNASDefaultMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetNASDefaultMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetNASDefaultMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetResponse) SetHeaders(v map[string]*string) *ResetNASDefaultMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ResetNASDefaultMountTargetResponse) SetBody(v *ResetNASDefaultMountTargetResponseBody) *ResetNASDefaultMountTargetResponse {
	s.Body = v
	return s
}

type ResetSnapshotRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ResetSnapshotRequest) SetRegionId(v string) *ResetSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSnapshotId(v string) *ResetSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type ResetSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponseBody) SetRequestId(v string) *ResetSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type ResetSnapshotResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponse) SetHeaders(v map[string]*string) *ResetSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ResetSnapshotResponse) SetBody(v *ResetSnapshotResponseBody) *ResetSnapshotResponse {
	s.Body = v
	return s
}

type RollbackSuspEventQuaraFileRequest struct {
	DesktopId    *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	QuaraFieldId *int32  `json:"QuaraFieldId,omitempty" xml:"QuaraFieldId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RollbackSuspEventQuaraFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileRequest) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileRequest) SetDesktopId(v string) *RollbackSuspEventQuaraFileRequest {
	s.DesktopId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetQuaraFieldId(v int32) *RollbackSuspEventQuaraFileRequest {
	s.QuaraFieldId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetRegionId(v string) *RollbackSuspEventQuaraFileRequest {
	s.RegionId = &v
	return s
}

type RollbackSuspEventQuaraFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackSuspEventQuaraFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponseBody) SetRequestId(v string) *RollbackSuspEventQuaraFileResponseBody {
	s.RequestId = &v
	return s
}

type RollbackSuspEventQuaraFileResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackSuspEventQuaraFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackSuspEventQuaraFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponse) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponse) SetHeaders(v map[string]*string) *RollbackSuspEventQuaraFileResponse {
	s.Headers = v
	return s
}

func (s *RollbackSuspEventQuaraFileResponse) SetBody(v *RollbackSuspEventQuaraFileResponseBody) *RollbackSuspEventQuaraFileResponse {
	s.Body = v
	return s
}

type RunCommandRequest struct {
	CommandContent  *string   `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	ContentEncoding *string   `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	DesktopId       []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Timeout         *int64    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type            *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandRequest) SetDesktopId(v []*string) *RunCommandRequest {
	s.DesktopId = v
	return s
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int64) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

type RunCommandResponseBody struct {
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

type RunCommandResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

type SendVerifyCodeRequest struct {
	ExtraInfo        *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VerifyCodeAction *string `json:"VerifyCodeAction,omitempty" xml:"VerifyCodeAction,omitempty"`
}

func (s SendVerifyCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeRequest) SetExtraInfo(v string) *SendVerifyCodeRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SendVerifyCodeRequest) SetRegionId(v string) *SendVerifyCodeRequest {
	s.RegionId = &v
	return s
}

func (s *SendVerifyCodeRequest) SetVerifyCodeAction(v string) *SendVerifyCodeRequest {
	s.VerifyCodeAction = &v
	return s
}

type SendVerifyCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerifyCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeResponseBody) SetRequestId(v string) *SendVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

type SendVerifyCodeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendVerifyCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeResponse) SetHeaders(v map[string]*string) *SendVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *SendVerifyCodeResponse) SetBody(v *SendVerifyCodeResponseBody) *SendVerifyCodeResponse {
	s.Body = v
	return s
}

type SetDirectorySsoStatusRequest struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EnableSso   *bool   `json:"EnableSso,omitempty" xml:"EnableSso,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDirectorySsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusRequest) SetDirectoryId(v string) *SetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetEnableSso(v bool) *SetDirectorySsoStatusRequest {
	s.EnableSso = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetRegionId(v string) *SetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

type SetDirectorySsoStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDirectorySsoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusResponseBody) SetRequestId(v string) *SetDirectorySsoStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetDirectorySsoStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDirectorySsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDirectorySsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusResponse) SetHeaders(v map[string]*string) *SetDirectorySsoStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDirectorySsoStatusResponse) SetBody(v *SetDirectorySsoStatusResponseBody) *SetDirectorySsoStatusResponse {
	s.Body = v
	return s
}

type SetIdpMetadataRequest struct {
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	IdpMetadata  *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetIdpMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataRequest) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataRequest) SetDirectoryId(v string) *SetIdpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetIdpMetadata(v string) *SetIdpMetadataRequest {
	s.IdpMetadata = &v
	return s
}

func (s *SetIdpMetadataRequest) SetOfficeSiteId(v string) *SetIdpMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetRegionId(v string) *SetIdpMetadataRequest {
	s.RegionId = &v
	return s
}

type SetIdpMetadataResponseBody struct {
	IdpEntityId *string `json:"IdpEntityId,omitempty" xml:"IdpEntityId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIdpMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataResponseBody) SetIdpEntityId(v string) *SetIdpMetadataResponseBody {
	s.IdpEntityId = &v
	return s
}

func (s *SetIdpMetadataResponseBody) SetRequestId(v string) *SetIdpMetadataResponseBody {
	s.RequestId = &v
	return s
}

type SetIdpMetadataResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetIdpMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetIdpMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataResponse) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataResponse) SetHeaders(v map[string]*string) *SetIdpMetadataResponse {
	s.Headers = v
	return s
}

func (s *SetIdpMetadataResponse) SetBody(v *SetIdpMetadataResponseBody) *SetIdpMetadataResponse {
	s.Body = v
	return s
}

type SetOfficeSiteSsoStatusRequest struct {
	EnableSso    *bool   `json:"EnableSso,omitempty" xml:"EnableSso,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SsoType      *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
}

func (s SetOfficeSiteSsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetOfficeSiteSsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusRequest) SetEnableSso(v bool) *SetOfficeSiteSsoStatusRequest {
	s.EnableSso = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetOfficeSiteId(v string) *SetOfficeSiteSsoStatusRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetRegionId(v string) *SetOfficeSiteSsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetSsoType(v string) *SetOfficeSiteSsoStatusRequest {
	s.SsoType = &v
	return s
}

type SetOfficeSiteSsoStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetOfficeSiteSsoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetOfficeSiteSsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusResponseBody) SetRequestId(v string) *SetOfficeSiteSsoStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetOfficeSiteSsoStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetOfficeSiteSsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetOfficeSiteSsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetOfficeSiteSsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusResponse) SetHeaders(v map[string]*string) *SetOfficeSiteSsoStatusResponse {
	s.Headers = v
	return s
}

func (s *SetOfficeSiteSsoStatusResponse) SetBody(v *SetOfficeSiteSsoStatusResponseBody) *SetOfficeSiteSsoStatusResponse {
	s.Body = v
	return s
}

type StartDesktopsRequest struct {
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

type StartDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponseBody) SetRequestId(v string) *StartDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StartDesktopsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponse) SetHeaders(v map[string]*string) *StartDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StartDesktopsResponse) SetBody(v *StartDesktopsResponseBody) *StartDesktopsResponse {
	s.Body = v
	return s
}

type StartVirusScanTaskRequest struct {
	DesktopId    []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartVirusScanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartVirusScanTaskRequest) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskRequest) SetDesktopId(v []*string) *StartVirusScanTaskRequest {
	s.DesktopId = v
	return s
}

func (s *StartVirusScanTaskRequest) SetOfficeSiteId(v []*string) *StartVirusScanTaskRequest {
	s.OfficeSiteId = v
	return s
}

func (s *StartVirusScanTaskRequest) SetRegionId(v string) *StartVirusScanTaskRequest {
	s.RegionId = &v
	return s
}

type StartVirusScanTaskResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScanTaskId *int64  `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
}

func (s StartVirusScanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartVirusScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskResponseBody) SetRequestId(v string) *StartVirusScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartVirusScanTaskResponseBody) SetScanTaskId(v int64) *StartVirusScanTaskResponseBody {
	s.ScanTaskId = &v
	return s
}

type StartVirusScanTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartVirusScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartVirusScanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartVirusScanTaskResponse) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskResponse) SetHeaders(v map[string]*string) *StartVirusScanTaskResponse {
	s.Headers = v
	return s
}

func (s *StartVirusScanTaskResponse) SetBody(v *StartVirusScanTaskResponseBody) *StartVirusScanTaskResponse {
	s.Body = v
	return s
}

type StopDesktopsRequest struct {
	DesktopId   []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StoppedMode *string   `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDesktopsRequest) SetStoppedMode(v string) *StopDesktopsRequest {
	s.StoppedMode = &v
	return s
}

type StopDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponseBody) SetRequestId(v string) *StopDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StopDesktopsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponse) SetHeaders(v map[string]*string) *StopDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StopDesktopsResponse) SetBody(v *StopDesktopsResponseBody) *StopDesktopsResponse {
	s.Body = v
	return s
}

type StopInvocationRequest struct {
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	InvokeId  *string   `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) SetDesktopId(v []*string) *StopInvocationRequest {
	s.DesktopId = v
	return s
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetRegionId(v string) *StopInvocationRequest {
	s.RegionId = &v
	return s
}

type StopInvocationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInvocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *StopInvocationResponseBody) SetRequestId(v string) *StopInvocationResponseBody {
	s.RequestId = &v
	return s
}

type StopInvocationResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopInvocationResponse) SetHeaders(v map[string]*string) *StopInvocationResponse {
	s.Headers = v
	return s
}

func (s *StopInvocationResponse) SetBody(v *StopInvocationResponseBody) *StopInvocationResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnlockVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UnlockVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceRequest) SetRegionId(v string) *UnlockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockVirtualMFADeviceRequest) SetSerialNumber(v string) *UnlockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type UnlockVirtualMFADeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceResponseBody) SetRequestId(v string) *UnlockVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnlockVirtualMFADeviceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnlockVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *UnlockVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *UnlockVirtualMFADeviceResponse) SetBody(v *UnlockVirtualMFADeviceResponseBody) *UnlockVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateFotaTaskRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskUid    *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s UpdateFotaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFotaTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateFotaTaskRequest) SetRegionId(v string) *UpdateFotaTaskRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateFotaTaskRequest) SetTaskUid(v string) *UpdateFotaTaskRequest {
	s.TaskUid = &v
	return s
}

func (s *UpdateFotaTaskRequest) SetUserStatus(v string) *UpdateFotaTaskRequest {
	s.UserStatus = &v
	return s
}

type UpdateFotaTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFotaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFotaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFotaTaskResponseBody) SetRequestId(v string) *UpdateFotaTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFotaTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFotaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFotaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFotaTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateFotaTaskResponse) SetHeaders(v map[string]*string) *UpdateFotaTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateFotaTaskResponse) SetBody(v *UpdateFotaTaskResponseBody) *UpdateFotaTaskResponse {
	s.Body = v
	return s
}

type UploadImageRequest struct {
	DataDiskSize  *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GpuCategory   *bool   `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	GpuDriverType *string `json:"GpuDriverType,omitempty" xml:"GpuDriverType,omitempty"`
	ImageName     *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	LicenseType   *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	OsType        *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OssObjectPath *string `json:"OssObjectPath,omitempty" xml:"OssObjectPath,omitempty"`
	ProtocolType  *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UploadImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadImageRequest) GoString() string {
	return s.String()
}

func (s *UploadImageRequest) SetDataDiskSize(v int32) *UploadImageRequest {
	s.DataDiskSize = &v
	return s
}

func (s *UploadImageRequest) SetDescription(v string) *UploadImageRequest {
	s.Description = &v
	return s
}

func (s *UploadImageRequest) SetGpuCategory(v bool) *UploadImageRequest {
	s.GpuCategory = &v
	return s
}

func (s *UploadImageRequest) SetGpuDriverType(v string) *UploadImageRequest {
	s.GpuDriverType = &v
	return s
}

func (s *UploadImageRequest) SetImageName(v string) *UploadImageRequest {
	s.ImageName = &v
	return s
}

func (s *UploadImageRequest) SetLicenseType(v string) *UploadImageRequest {
	s.LicenseType = &v
	return s
}

func (s *UploadImageRequest) SetOsType(v string) *UploadImageRequest {
	s.OsType = &v
	return s
}

func (s *UploadImageRequest) SetOssObjectPath(v string) *UploadImageRequest {
	s.OssObjectPath = &v
	return s
}

func (s *UploadImageRequest) SetProtocolType(v string) *UploadImageRequest {
	s.ProtocolType = &v
	return s
}

func (s *UploadImageRequest) SetRegionId(v string) *UploadImageRequest {
	s.RegionId = &v
	return s
}

type UploadImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponseBody) GoString() string {
	return s.String()
}

func (s *UploadImageResponseBody) SetImageId(v string) *UploadImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *UploadImageResponseBody) SetRequestId(v string) *UploadImageResponseBody {
	s.RequestId = &v
	return s
}

type UploadImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponse) GoString() string {
	return s.String()
}

func (s *UploadImageResponse) SetHeaders(v map[string]*string) *UploadImageResponse {
	s.Headers = v
	return s
}

func (s *UploadImageResponse) SetBody(v *UploadImageResponseBody) *UploadImageResponse {
	s.Body = v
	return s
}

type VerifyCenRequest struct {
	CenId      *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId *int64  `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	CidrBlock  *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s VerifyCenRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyCenRequest) GoString() string {
	return s.String()
}

func (s *VerifyCenRequest) SetCenId(v string) *VerifyCenRequest {
	s.CenId = &v
	return s
}

func (s *VerifyCenRequest) SetCenOwnerId(v int64) *VerifyCenRequest {
	s.CenOwnerId = &v
	return s
}

func (s *VerifyCenRequest) SetCidrBlock(v string) *VerifyCenRequest {
	s.CidrBlock = &v
	return s
}

func (s *VerifyCenRequest) SetRegionId(v string) *VerifyCenRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyCenRequest) SetVerifyCode(v string) *VerifyCenRequest {
	s.VerifyCode = &v
	return s
}

type VerifyCenResponseBody struct {
	CidrBlocks   []*string                            `json:"CidrBlocks,omitempty" xml:"CidrBlocks,omitempty" type:"Repeated"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteEntries []*VerifyCenResponseBodyRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
	Status       *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VerifyCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyCenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCenResponseBody) SetCidrBlocks(v []*string) *VerifyCenResponseBody {
	s.CidrBlocks = v
	return s
}

func (s *VerifyCenResponseBody) SetRequestId(v string) *VerifyCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyCenResponseBody) SetRouteEntries(v []*VerifyCenResponseBodyRouteEntries) *VerifyCenResponseBody {
	s.RouteEntries = v
	return s
}

func (s *VerifyCenResponseBody) SetStatus(v string) *VerifyCenResponseBody {
	s.Status = &v
	return s
}

type VerifyCenResponseBodyRouteEntries struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NextHopInstanceId    *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VerifyCenResponseBodyRouteEntries) String() string {
	return tea.Prettify(s)
}

func (s VerifyCenResponseBodyRouteEntries) GoString() string {
	return s.String()
}

func (s *VerifyCenResponseBodyRouteEntries) SetDestinationCidrBlock(v string) *VerifyCenResponseBodyRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *VerifyCenResponseBodyRouteEntries) SetNextHopInstanceId(v string) *VerifyCenResponseBodyRouteEntries {
	s.NextHopInstanceId = &v
	return s
}

func (s *VerifyCenResponseBodyRouteEntries) SetRegionId(v string) *VerifyCenResponseBodyRouteEntries {
	s.RegionId = &v
	return s
}

func (s *VerifyCenResponseBodyRouteEntries) SetStatus(v string) *VerifyCenResponseBodyRouteEntries {
	s.Status = &v
	return s
}

type VerifyCenResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyCenResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyCenResponse) GoString() string {
	return s.String()
}

func (s *VerifyCenResponse) SetHeaders(v map[string]*string) *VerifyCenResponse {
	s.Headers = v
	return s
}

func (s *VerifyCenResponse) SetBody(v *VerifyCenResponseBody) *VerifyCenResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActivateOfficeSiteWithOptions(request *ActivateOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *ActivateOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateOfficeSite"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateOfficeSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivateOfficeSite(request *ActivateOfficeSiteRequest) (_result *ActivateOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateOfficeSiteResponse{}
	_body, _err := client.ActivateOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserToDesktopGroupWithOptions(request *AddUserToDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *AddUserToDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupIds)) {
		query["DesktopGroupIds"] = request.DesktopGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToDesktopGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserToDesktopGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserToDesktopGroup(request *AddUserToDesktopGroupRequest) (_result *AddUserToDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToDesktopGroupResponse{}
	_body, _err := client.AddUserToDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveFotaUpdateWithOptions(request *ApproveFotaUpdateRequest, runtime *util.RuntimeOptions) (_result *ApproveFotaUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveFotaUpdate"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveFotaUpdate(request *ApproveFotaUpdateRequest) (_result *ApproveFotaUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.ApproveFotaUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachCenWithOptions(request *AttachCenRequest, runtime *util.RuntimeOptions) (_result *AttachCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CenOwnerId)) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachCen"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachCen(request *AttachCenRequest) (_result *AttachCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachCenResponse{}
	_body, _err := client.AttachCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUserTagsWithOptions(request *CheckUserTagsRequest, runtime *util.RuntimeOptions) (_result *CheckUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUserTags"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUserTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUserTags(request *CheckUserTagsRequest) (_result *CheckUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUserTagsResponse{}
	_body, _err := client.CheckUserTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClonePolicyGroupWithOptions(request *ClonePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ClonePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClonePolicyGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClonePolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClonePolicyGroup(request *ClonePolicyGroupRequest) (_result *ClonePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClonePolicyGroupResponse{}
	_body, _err := client.ClonePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigADConnectorTrustWithOptions(request *ConfigADConnectorTrustRequest, runtime *util.RuntimeOptions) (_result *ConfigADConnectorTrustResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TrustKey)) {
		query["TrustKey"] = request.TrustKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigADConnectorTrust"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigADConnectorTrustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigADConnectorTrust(request *ConfigADConnectorTrustRequest) (_result *ConfigADConnectorTrustResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigADConnectorTrustResponse{}
	_body, _err := client.ConfigADConnectorTrustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigADConnectorUserWithOptions(request *ConfigADConnectorUserRequest, runtime *util.RuntimeOptions) (_result *ConfigADConnectorUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainPassword)) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DomainUserName)) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !tea.BoolValue(util.IsUnset(request.OUName)) {
		query["OUName"] = request.OUName
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigADConnectorUser"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigADConnectorUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigADConnectorUser(request *ConfigADConnectorUserRequest) (_result *ConfigADConnectorUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigADConnectorUserResponse{}
	_body, _err := client.ConfigADConnectorUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateADConnectorDirectoryWithOptions(request *CreateADConnectorDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateADConnectorDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopAccessType)) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryName)) {
		query["DirectoryName"] = request.DirectoryName
	}

	if !tea.BoolValue(util.IsUnset(request.DnsAddress)) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainPassword)) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DomainUserName)) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAdminAccess)) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !tea.BoolValue(util.IsUnset(request.MfaEnabled)) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		query["Specification"] = request.Specification
	}

	if !tea.BoolValue(util.IsUnset(request.SubDomainDnsAddress)) {
		query["SubDomainDnsAddress"] = request.SubDomainDnsAddress
	}

	if !tea.BoolValue(util.IsUnset(request.SubDomainName)) {
		query["SubDomainName"] = request.SubDomainName
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateADConnectorDirectory"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateADConnectorDirectory(request *CreateADConnectorDirectoryRequest) (_result *CreateADConnectorDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.CreateADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateADConnectorOfficeSiteWithOptions(request *CreateADConnectorOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdHostname)) {
		query["AdHostname"] = request.AdHostname
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CenOwnerId)) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.CidrBlock)) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopAccessType)) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.DnsAddress)) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainPassword)) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DomainUserName)) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAdminAccess)) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !tea.BoolValue(util.IsUnset(request.EnableInternetAccess)) {
		query["EnableInternetAccess"] = request.EnableInternetAccess
	}

	if !tea.BoolValue(util.IsUnset(request.MfaEnabled)) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteName)) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		query["Specification"] = request.Specification
	}

	if !tea.BoolValue(util.IsUnset(request.SubDomainDnsAddress)) {
		query["SubDomainDnsAddress"] = request.SubDomainDnsAddress
	}

	if !tea.BoolValue(util.IsUnset(request.SubDomainName)) {
		query["SubDomainName"] = request.SubDomainName
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateADConnectorOfficeSite"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateADConnectorOfficeSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateADConnectorOfficeSite(request *CreateADConnectorOfficeSiteRequest) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateADConnectorOfficeSiteResponse{}
	_body, _err := client.CreateADConnectorOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBandwidthResourcePackagesWithOptions(request *CreateBandwidthResourcePackagesRequest, runtime *util.RuntimeOptions) (_result *CreateBandwidthResourcePackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.PackageSize)) {
		query["PackageSize"] = request.PackageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBandwidthResourcePackages"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBandwidthResourcePackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBandwidthResourcePackages(request *CreateBandwidthResourcePackagesRequest) (_result *CreateBandwidthResourcePackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBandwidthResourcePackagesResponse{}
	_body, _err := client.CreateBandwidthResourcePackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBundleWithOptions(request *CreateBundleRequest, runtime *util.RuntimeOptions) (_result *CreateBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleName)) {
		query["BundleName"] = request.BundleName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopType)) {
		query["DesktopType"] = request.DesktopType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RootDiskPerformanceLevel)) {
		query["RootDiskPerformanceLevel"] = request.RootDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RootDiskSizeGib)) {
		query["RootDiskSizeGib"] = request.RootDiskSizeGib
	}

	if !tea.BoolValue(util.IsUnset(request.UserDiskPerformanceLevel)) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.UserDiskSizeGib)) {
		query["UserDiskSizeGib"] = request.UserDiskSizeGib
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBundle"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBundleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBundle(request *CreateBundleRequest) (_result *CreateBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBundleResponse{}
	_body, _err := client.CreateBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopGroupWithOptions(request *CreateDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowAutoSetup)) {
		query["AllowAutoSetup"] = request.AllowAutoSetup
	}

	if !tea.BoolValue(util.IsUnset(request.AllowBufferCount)) {
		query["AllowBufferCount"] = request.AllowBufferCount
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.BindAmount)) {
		query["BindAmount"] = request.BindAmount
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultInitDesktopCount)) {
		query["DefaultInitDesktopCount"] = request.DefaultInitDesktopCount
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupName)) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.KeepDuration)) {
		query["KeepDuration"] = request.KeepDuration
	}

	if !tea.BoolValue(util.IsUnset(request.LoadPolicy)) {
		query["LoadPolicy"] = request.LoadPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.MaxDesktopsCount)) {
		query["MaxDesktopsCount"] = request.MaxDesktopsCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinDesktopsCount)) {
		query["MinDesktopsCount"] = request.MinDesktopsCount
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnType)) {
		query["OwnType"] = request.OwnType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResetType)) {
		query["ResetType"] = request.ResetType
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleStrategyId)) {
		query["ScaleStrategyId"] = request.ScaleStrategyId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDesktopGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDesktopGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktopGroup(request *CreateDesktopGroupRequest) (_result *CreateDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopGroupResponse{}
	_body, _err := client.CreateDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopsWithOptions(request *CreateDesktopsRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopName)) {
		query["DesktopName"] = request.DesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopNameSuffix)) {
		query["DesktopNameSuffix"] = request.DesktopNameSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Hostname)) {
		query["Hostname"] = request.Hostname
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UserAssignMode)) {
		query["UserAssignMode"] = request.UserAssignMode
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktops(request *CreateDesktopsRequest) (_result *CreateDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.CreateDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageResourceType)) {
		query["ImageResourceType"] = request.ImageResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImage"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNASFileSystemWithOptions(request *CreateNASFileSystemRequest, runtime *util.RuntimeOptions) (_result *CreateNASFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNASFileSystem"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNASFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNASFileSystem(request *CreateNASFileSystemRequest) (_result *CreateNASFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNASFileSystemResponse{}
	_body, _err := client.CreateNASFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNetworkPackageWithOptions(request *CreateNetworkPackageRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNetworkPackage"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNetworkPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNetworkPackage(request *CreateNetworkPackageRequest) (_result *CreateNetworkPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkPackageResponse{}
	_body, _err := client.CreateNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderForHardwareWithOptions(request *CreateOrderForHardwareRequest, runtime *util.RuntimeOptions) (_result *CreateOrderForHardwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.CityCode)) {
		query["CityCode"] = request.CityCode
	}

	if !tea.BoolValue(util.IsUnset(request.CityName)) {
		query["CityName"] = request.CityName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		query["CountryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.CountryName)) {
		query["CountryName"] = request.CountryName
	}

	if !tea.BoolValue(util.IsUnset(request.DetailAddress)) {
		query["DetailAddress"] = request.DetailAddress
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictCode)) {
		query["DistrictCode"] = request.DistrictCode
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictName)) {
		query["DistrictName"] = request.DistrictName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.HardwareType)) {
		query["HardwareType"] = request.HardwareType
	}

	if !tea.BoolValue(util.IsUnset(request.HardwareVersion)) {
		query["HardwareVersion"] = request.HardwareVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MobilePhone)) {
		query["MobilePhone"] = request.MobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProvCode)) {
		query["ProvCode"] = request.ProvCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProvName)) {
		query["ProvName"] = request.ProvName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StreetCode)) {
		query["StreetCode"] = request.StreetCode
	}

	if !tea.BoolValue(util.IsUnset(request.StreetName)) {
		query["StreetName"] = request.StreetName
	}

	if !tea.BoolValue(util.IsUnset(request.ZipCode)) {
		query["ZipCode"] = request.ZipCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrderForHardware"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderForHardwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrderForHardware(request *CreateOrderForHardwareRequest) (_result *CreateOrderForHardwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderForHardwareResponse{}
	_body, _err := client.CreateOrderForHardwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePolicyGroupWithOptions(request *CreatePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizeAccessPolicyRule)) {
		query["AuthorizeAccessPolicyRule"] = request.AuthorizeAccessPolicyRule
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizeSecurityPolicyRule)) {
		query["AuthorizeSecurityPolicyRule"] = request.AuthorizeSecurityPolicyRule
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Clipboard)) {
		query["Clipboard"] = request.Clipboard
	}

	if !tea.BoolValue(util.IsUnset(request.DomainList)) {
		query["DomainList"] = request.DomainList
	}

	if !tea.BoolValue(util.IsUnset(request.GpuAcceleration)) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !tea.BoolValue(util.IsUnset(request.Html5Access)) {
		query["Html5Access"] = request.Html5Access
	}

	if !tea.BoolValue(util.IsUnset(request.Html5FileTransfer)) {
		query["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.LocalDrive)) {
		query["LocalDrive"] = request.LocalDrive
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PreemptLogin)) {
		query["PreemptLogin"] = request.PreemptLogin
	}

	if !tea.BoolValue(util.IsUnset(request.PreemptLoginUser)) {
		query["PreemptLoginUser"] = request.PreemptLoginUser
	}

	if !tea.BoolValue(util.IsUnset(request.PrinterRedirection)) {
		query["PrinterRedirection"] = request.PrinterRedirection
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UsbRedirect)) {
		query["UsbRedirect"] = request.UsbRedirect
	}

	if !tea.BoolValue(util.IsUnset(request.UsbSupplyRedirectRule)) {
		query["UsbSupplyRedirectRule"] = request.UsbSupplyRedirectRule
	}

	if !tea.BoolValue(util.IsUnset(request.VisualQuality)) {
		query["VisualQuality"] = request.VisualQuality
	}

	if !tea.BoolValue(util.IsUnset(request.Watermark)) {
		query["Watermark"] = request.Watermark
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkCustomText)) {
		query["WatermarkCustomText"] = request.WatermarkCustomText
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkTransparency)) {
		query["WatermarkTransparency"] = request.WatermarkTransparency
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkType)) {
		query["WatermarkType"] = request.WatermarkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicyGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (_result *CreatePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CreatePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRAMDirectoryWithOptions(request *CreateRAMDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateRAMDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopAccessType)) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryName)) {
		query["DirectoryName"] = request.DirectoryName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAdminAccess)) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !tea.BoolValue(util.IsUnset(request.EnableInternetAccess)) {
		query["EnableInternetAccess"] = request.EnableInternetAccess
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRAMDirectory"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRAMDirectory(request *CreateRAMDirectoryRequest) (_result *CreateRAMDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.CreateRAMDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSimpleOfficeSiteWithOptions(request *CreateSimpleOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CenOwnerId)) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.CidrBlock)) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.CloudBoxOfficeSite)) {
		query["CloudBoxOfficeSite"] = request.CloudBoxOfficeSite
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopAccessType)) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAdminAccess)) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !tea.BoolValue(util.IsUnset(request.EnableInternetAccess)) {
		query["EnableInternetAccess"] = request.EnableInternetAccess
	}

	if !tea.BoolValue(util.IsUnset(request.NeedVerifyZeroDevice)) {
		query["NeedVerifyZeroDevice"] = request.NeedVerifyZeroDevice
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteName)) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSimpleOfficeSite"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSimpleOfficeSite(request *CreateSimpleOfficeSiteRequest) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.CreateSimpleOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDiskType)) {
		query["SourceDiskType"] = request.SourceDiskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSnapshot"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserTagsWithOptions(request *CreateUserTagsRequest, runtime *util.RuntimeOptions) (_result *CreateUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserTags"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserTags(request *CreateUserTagsRequest) (_result *CreateUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserTagsResponse{}
	_body, _err := client.CreateUserTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBundlesWithOptions(request *DeleteBundlesRequest, runtime *util.RuntimeOptions) (_result *DeleteBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBundles"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBundles(request *DeleteBundlesRequest) (_result *DeleteBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.DeleteBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDesktopGroupWithOptions(request *DeleteDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDesktopGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDesktopGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDesktopGroup(request *DeleteDesktopGroupRequest) (_result *DeleteDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDesktopGroupResponse{}
	_body, _err := client.DeleteDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDesktopsWithOptions(request *DeleteDesktopsRequest, runtime *util.RuntimeOptions) (_result *DeleteDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDesktops(request *DeleteDesktopsRequest) (_result *DeleteDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.DeleteDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDirectoriesWithOptions(request *DeleteDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDirectories"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDirectories(request *DeleteDirectoriesRequest) (_result *DeleteDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.DeleteDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImagesWithOptions(request *DeleteImagesRequest, runtime *util.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImages"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImages(request *DeleteImagesRequest) (_result *DeleteImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DeleteImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNASFileSystemsWithOptions(request *DeleteNASFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DeleteNASFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNASFileSystems"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNASFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNASFileSystems(request *DeleteNASFileSystemsRequest) (_result *DeleteNASFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNASFileSystemsResponse{}
	_body, _err := client.DeleteNASFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNetworkPackagesWithOptions(request *DeleteNetworkPackagesRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkPackageId)) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNetworkPackages"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNetworkPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNetworkPackages(request *DeleteNetworkPackagesRequest) (_result *DeleteNetworkPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkPackagesResponse{}
	_body, _err := client.DeleteNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOfficeSitesWithOptions(request *DeleteOfficeSitesRequest, runtime *util.RuntimeOptions) (_result *DeleteOfficeSitesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOfficeSites"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOfficeSitesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOfficeSites(request *DeleteOfficeSitesRequest) (_result *DeleteOfficeSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOfficeSitesResponse{}
	_body, _err := client.DeleteOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePolicyGroupsWithOptions(request *DeletePolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicyGroups"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePolicyGroups(request *DeletePolicyGroupsRequest) (_result *DeletePolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.DeletePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshot"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserTagsWithOptions(request *DeleteUserTagsRequest, runtime *util.RuntimeOptions) (_result *DeleteUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserTags"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserTags(request *DeleteUserTagsRequest) (_result *DeleteUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserTagsResponse{}
	_body, _err := client.DeleteUserTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVirtualMFADevice"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlarmEventStackInfoWithOptions(request *DescribeAlarmEventStackInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmEventStackInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		query["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueInfo)) {
		query["UniqueInfo"] = request.UniqueInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlarmEventStackInfo"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlarmEventStackInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlarmEventStackInfo(request *DescribeAlarmEventStackInfoRequest) (_result *DescribeAlarmEventStackInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmEventStackInfoResponse{}
	_body, _err := client.DescribeAlarmEventStackInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBundlesWithOptions(request *DescribeBundlesRequest, runtime *util.RuntimeOptions) (_result *DescribeBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.BundleType)) {
		query["BundleType"] = request.BundleType
	}

	if !tea.BoolValue(util.IsUnset(request.CheckStock)) {
		query["CheckStock"] = request.CheckStock
	}

	if !tea.BoolValue(util.IsUnset(request.CpuCount)) {
		query["CpuCount"] = request.CpuCount
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopTypeFamily)) {
		query["DesktopTypeFamily"] = request.DesktopTypeFamily
	}

	if !tea.BoolValue(util.IsUnset(request.FromDesktopGroup)) {
		query["FromDesktopGroup"] = request.FromDesktopGroup
	}

	if !tea.BoolValue(util.IsUnset(request.GpuCount)) {
		query["GpuCount"] = request.GpuCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MemorySize)) {
		query["MemorySize"] = request.MemorySize
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBundles"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBundles(request *DescribeBundlesRequest) (_result *DescribeBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.DescribeBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCensWithOptions(request *DescribeCensRequest, runtime *util.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCens"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCens(request *DescribeCensRequest) (_result *DescribeCensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCensResponse{}
	_body, _err := client.DescribeCensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClientEventsWithOptions(request *DescribeClientEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeClientEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopIp)) {
		query["DesktopIp"] = request.DesktopIp
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopName)) {
		query["DesktopName"] = request.DesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteName)) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClientEvents"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClientEvents(request *DescribeClientEventsRequest) (_result *DescribeClientEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.DescribeClientEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopGroupsWithOptions(request *DescribeDesktopGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupName)) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludedEndUserIds)) {
		query["ExcludedEndUserIds"] = request.ExcludedEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnType)) {
		query["OwnType"] = request.OwnType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDesktopGroups"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDesktopGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopGroups(request *DescribeDesktopGroupsRequest) (_result *DescribeDesktopGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopGroupsResponse{}
	_body, _err := client.DescribeDesktopGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopIdsByVulNamesWithOptions(request *DescribeDesktopIdsByVulNamesRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopIdsByVulNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Necessity)) {
		query["Necessity"] = request.Necessity
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VulName)) {
		query["VulName"] = request.VulName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDesktopIdsByVulNames"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDesktopIdsByVulNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopIdsByVulNames(request *DescribeDesktopIdsByVulNamesRequest) (_result *DescribeDesktopIdsByVulNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopIdsByVulNamesResponse{}
	_body, _err := client.DescribeDesktopIdsByVulNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopTypesWithOptions(request *DescribeDesktopTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CpuCount)) {
		query["CpuCount"] = request.CpuCount
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopIdForModify)) {
		query["DesktopIdForModify"] = request.DesktopIdForModify
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopTypeId)) {
		query["DesktopTypeId"] = request.DesktopTypeId
	}

	if !tea.BoolValue(util.IsUnset(request.GpuCount)) {
		query["GpuCount"] = request.GpuCount
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypeFamily)) {
		query["InstanceTypeFamily"] = request.InstanceTypeFamily
	}

	if !tea.BoolValue(util.IsUnset(request.MemorySize)) {
		query["MemorySize"] = request.MemorySize
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDesktopTypes"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopTypes(request *DescribeDesktopTypesRequest) (_result *DescribeDesktopTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.DescribeDesktopTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopsWithOptions(request *DescribeDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopName)) {
		query["DesktopName"] = request.DesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopStatus)) {
		query["DesktopStatus"] = request.DesktopStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludedEndUserId)) {
		query["ExcludedEndUserId"] = request.ExcludedEndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterDesktopGroup)) {
		query["FilterDesktopGroup"] = request.FilterDesktopGroup
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagementFlag)) {
		query["ManagementFlag"] = request.ManagementFlag
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteName)) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryFotaUpdate)) {
		query["QueryFotaUpdate"] = request.QueryFotaUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktops(request *DescribeDesktopsRequest) (_result *DescribeDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DescribeDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopsInGroupWithOptions(request *DescribeDesktopsInGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDesktopsInGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDesktopsInGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopsInGroup(request *DescribeDesktopsInGroupRequest) (_result *DescribeDesktopsInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopsInGroupResponse{}
	_body, _err := client.DescribeDesktopsInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryStatus)) {
		query["DirectoryStatus"] = request.DirectoryStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryType)) {
		query["DirectoryType"] = request.DirectoryType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDirectories"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowMetricWithOptions(request *DescribeFlowMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowMetric"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowMetric(request *DescribeFlowMetricRequest) (_result *DescribeFlowMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowMetricResponse{}
	_body, _err := client.DescribeFlowMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFotaPendingDesktopsWithOptions(request *DescribeFotaPendingDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeFotaPendingDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFotaPendingDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFotaPendingDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFotaPendingDesktops(request *DescribeFotaPendingDesktopsRequest) (_result *DescribeFotaPendingDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFotaPendingDesktopsResponse{}
	_body, _err := client.DescribeFotaPendingDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFotaTasksWithOptions(request *DescribeFotaTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeFotaTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FotaStatus)) {
		query["FotaStatus"] = request.FotaStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUid)) {
		query["TaskUid"] = request.TaskUid
	}

	if !tea.BoolValue(util.IsUnset(request.UserStatus)) {
		query["UserStatus"] = request.UserStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFotaTasks"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFotaTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFotaTasks(request *DescribeFotaTasksRequest) (_result *DescribeFotaTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFotaTasksResponse{}
	_body, _err := client.DescribeFotaTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFrontVulPatchListWithOptions(request *DescribeFrontVulPatchListRequest, runtime *util.RuntimeOptions) (_result *DescribeFrontVulPatchListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VulInfo)) {
		query["VulInfo"] = request.VulInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFrontVulPatchList"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFrontVulPatchListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFrontVulPatchList(request *DescribeFrontVulPatchListRequest) (_result *DescribeFrontVulPatchListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFrontVulPatchListResponse{}
	_body, _err := client.DescribeFrontVulPatchListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupedVulWithOptions(request *DescribeGroupedVulRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupedVulResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Dealed)) {
		query["Dealed"] = request.Dealed
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Necessity)) {
		query["Necessity"] = request.Necessity
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGroupedVul"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGroupedVulResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupedVul(request *DescribeGroupedVulRequest) (_result *DescribeGroupedVulResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupedVulResponse{}
	_body, _err := client.DescribeGroupedVulWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHardwareTerminalsWithOptions(request *DescribeHardwareTerminalsRequest, runtime *util.RuntimeOptions) (_result *DescribeHardwareTerminalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HardwareType)) {
		query["HardwareType"] = request.HardwareType
	}

	if !tea.BoolValue(util.IsUnset(request.HardwareVersion)) {
		query["HardwareVersion"] = request.HardwareVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHardwareTerminals"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHardwareTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHardwareTerminals(request *DescribeHardwareTerminalsRequest) (_result *DescribeHardwareTerminalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHardwareTerminalsResponse{}
	_body, _err := client.DescribeHardwareTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopInstanceType)) {
		query["DesktopInstanceType"] = request.DesktopInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.GpuCategory)) {
		query["GpuCategory"] = request.GpuCategory
	}

	if !tea.BoolValue(util.IsUnset(request.GpuDriverVersion)) {
		query["GpuDriverVersion"] = request.GpuDriverVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageStatus)) {
		query["ImageStatus"] = request.ImageStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageType)) {
		query["LanguageType"] = request.LanguageType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImages"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandType)) {
		query["CommandType"] = request.CommandType
	}

	if !tea.BoolValue(util.IsUnset(request.ContentEncoding)) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeOutput)) {
		query["IncludeOutput"] = request.IncludeOutput
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		query["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeStatus)) {
		query["InvokeStatus"] = request.InvokeStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvocations"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModificationPriceWithOptions(request *DescribeModificationPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeModificationPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RootDiskPerformanceLevel)) {
		query["RootDiskPerformanceLevel"] = request.RootDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RootDiskSizeGib)) {
		query["RootDiskSizeGib"] = request.RootDiskSizeGib
	}

	if !tea.BoolValue(util.IsUnset(request.UserDiskPerformanceLevel)) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.UserDiskSizeGib)) {
		query["UserDiskSizeGib"] = request.UserDiskSizeGib
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeModificationPrice"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeModificationPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModificationPrice(request *DescribeModificationPriceRequest) (_result *DescribeModificationPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModificationPriceResponse{}
	_body, _err := client.DescribeModificationPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNASFileSystemsWithOptions(request *DescribeNASFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DescribeNASFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNASFileSystems"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNASFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNASFileSystems(request *DescribeNASFileSystemsRequest) (_result *DescribeNASFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNASFileSystemsResponse{}
	_body, _err := client.DescribeNASFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkPackagesWithOptions(request *DescribeNetworkPackagesRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkPackageId)) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNetworkPackages"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNetworkPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkPackages(request *DescribeNetworkPackagesRequest) (_result *DescribeNetworkPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkPackagesResponse{}
	_body, _err := client.DescribeNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOfficeSitesWithOptions(request *DescribeOfficeSitesRequest, runtime *util.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteType)) {
		query["OfficeSiteType"] = request.OfficeSiteType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOfficeSites"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (_result *DescribeOfficeSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DescribeOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePolicyGroupsWithOptions(request *DescribePolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyGroups"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePolicyGroups(request *DescribePolicyGroupsRequest) (_result *DescribePolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.DescribePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.HardwareVersion)) {
		query["HardwareVersion"] = request.HardwareVersion
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageSize)) {
		query["PackageSize"] = request.PackageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RootDiskPerformanceLevel)) {
		query["RootDiskPerformanceLevel"] = request.RootDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RootDiskSizeGib)) {
		query["RootDiskSizeGib"] = request.RootDiskSizeGib
	}

	if !tea.BoolValue(util.IsUnset(request.UserDiskPerformanceLevel)) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.UserDiskSizeGib)) {
		query["UserDiskSizeGib"] = request.UserDiskSizeGib
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePrice"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRenewalPriceWithOptions(request *DescribeRenewalPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeRenewalPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRenewalPrice"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRenewalPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRenewalPrice(request *DescribeRenewalPriceRequest) (_result *DescribeRenewalPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRenewalPriceResponse{}
	_body, _err := client.DescribeRenewalPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScanTaskProgressWithOptions(request *DescribeScanTaskProgressRequest, runtime *util.RuntimeOptions) (_result *DescribeScanTaskProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScanTaskProgress"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScanTaskProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScanTaskProgress(request *DescribeScanTaskProgressRequest) (_result *DescribeScanTaskProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScanTaskProgressResponse{}
	_body, _err := client.DescribeScanTaskProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationStatusWithOptions(request *DescribeSecurityEventOperationStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityEventId)) {
		query["SecurityEventId"] = request.SecurityEventId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecurityEventOperationStatus"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationStatus(request *DescribeSecurityEventOperationStatusRequest) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.DescribeSecurityEventOperationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationsWithOptions(request *DescribeSecurityEventOperationsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityEventId)) {
		query["SecurityEventId"] = request.SecurityEventId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecurityEventOperations"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperations(request *DescribeSecurityEventOperationsRequest) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.DescribeSecurityEventOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSnapshots"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventOverviewWithOptions(request *DescribeSuspEventOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSuspEventOverview"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSuspEventOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEventOverview(request *DescribeSuspEventOverviewRequest) (_result *DescribeSuspEventOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventOverviewResponse{}
	_body, _err := client.DescribeSuspEventOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventQuaraFilesWithOptions(request *DescribeSuspEventQuaraFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventQuaraFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSuspEventQuaraFiles"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSuspEventQuaraFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEventQuaraFiles(request *DescribeSuspEventQuaraFilesRequest) (_result *DescribeSuspEventQuaraFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventQuaraFilesResponse{}
	_body, _err := client.DescribeSuspEventQuaraFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventsWithOptions(request *DescribeSuspEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmUniqueInfo)) {
		query["AlarmUniqueInfo"] = request.AlarmUniqueInfo
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Dealed)) {
		query["Dealed"] = request.Dealed
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Levels)) {
		query["Levels"] = request.Levels
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentEventType)) {
		query["ParentEventType"] = request.ParentEventType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSuspEvents"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEvents(request *DescribeSuspEventsRequest) (_result *DescribeSuspEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DescribeSuspEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserConnectionRecordsWithOptions(request *DescribeUserConnectionRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserConnectionRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserType)) {
		query["EndUserType"] = request.EndUserType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserConnectionRecords"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserConnectionRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserConnectionRecords(request *DescribeUserConnectionRecordsRequest) (_result *DescribeUserConnectionRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserConnectionRecordsResponse{}
	_body, _err := client.DescribeUserConnectionRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsersInGroupWithOptions(request *DescribeUsersInGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeUsersInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectState)) {
		query["ConnectState"] = request.ConnectState
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsersInGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsersInGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsersInGroup(request *DescribeUsersInGroupRequest) (_result *DescribeUsersInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsersInGroupResponse{}
	_body, _err := client.DescribeUsersInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVirtualMFADevicesWithOptions(request *DescribeVirtualMFADevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVirtualMFADevices"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVirtualMFADevices(request *DescribeVirtualMFADevicesRequest) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.DescribeVirtualMFADevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulDetailsWithOptions(request *DescribeVulDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeVulDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVulDetails"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVulDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulDetails(request *DescribeVulDetailsRequest) (_result *DescribeVulDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulDetailsResponse{}
	_body, _err := client.DescribeVulDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulListWithOptions(request *DescribeVulListRequest, runtime *util.RuntimeOptions) (_result *DescribeVulListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Dealed)) {
		query["Dealed"] = request.Dealed
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Necessity)) {
		query["Necessity"] = request.Necessity
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVulList"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVulListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulList(request *DescribeVulListRequest) (_result *DescribeVulListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulListResponse{}
	_body, _err := client.DescribeVulListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulOverviewWithOptions(request *DescribeVulOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeVulOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVulOverview"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVulOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulOverview(request *DescribeVulOverviewRequest) (_result *DescribeVulOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulOverviewResponse{}
	_body, _err := client.DescribeVulOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneType)) {
		query["ZoneType"] = request.ZoneType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachCenWithOptions(request *DetachCenRequest, runtime *util.RuntimeOptions) (_result *DetachCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachCen"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachCen(request *DetachCenRequest) (_result *DetachCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachCenResponse{}
	_body, _err := client.DetachCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableDesktopsInGroupWithOptions(request *DisableDesktopsInGroupRequest, runtime *util.RuntimeOptions) (_result *DisableDesktopsInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopIds)) {
		query["DesktopIds"] = request.DesktopIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableDesktopsInGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableDesktopsInGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableDesktopsInGroup(request *DisableDesktopsInGroupRequest) (_result *DisableDesktopsInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDesktopsInGroupResponse{}
	_body, _err := client.DisableDesktopsInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportClientEventsWithOptions(request *ExportClientEventsRequest, runtime *util.RuntimeOptions) (_result *ExportClientEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopName)) {
		query["DesktopName"] = request.DesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteName)) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportClientEvents"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportClientEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportClientEvents(request *ExportClientEventsRequest) (_result *ExportClientEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportClientEventsResponse{}
	_body, _err := client.ExportClientEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportDesktopGroupInfoWithOptions(request *ExportDesktopGroupInfoRequest, runtime *util.RuntimeOptions) (_result *ExportDesktopGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupName)) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportDesktopGroupInfo"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportDesktopGroupInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportDesktopGroupInfo(request *ExportDesktopGroupInfoRequest) (_result *ExportDesktopGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportDesktopGroupInfoResponse{}
	_body, _err := client.ExportDesktopGroupInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportDesktopListInfoWithOptions(request *ExportDesktopListInfoRequest, runtime *util.RuntimeOptions) (_result *ExportDesktopListInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopName)) {
		query["DesktopName"] = request.DesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopStatus)) {
		query["DesktopStatus"] = request.DesktopStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportDesktopListInfo"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportDesktopListInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportDesktopListInfo(request *ExportDesktopListInfoRequest) (_result *ExportDesktopListInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportDesktopListInfoResponse{}
	_body, _err := client.ExportDesktopListInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionTicket"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDesktopGroupDetailWithOptions(request *GetDesktopGroupDetailRequest, runtime *util.RuntimeOptions) (_result *GetDesktopGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDesktopGroupDetail"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDesktopGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDesktopGroupDetail(request *GetDesktopGroupDetailRequest) (_result *GetDesktopGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDesktopGroupDetailResponse{}
	_body, _err := client.GetDesktopGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDirectorySsoStatusWithOptions(request *GetDirectorySsoStatusRequest, runtime *util.RuntimeOptions) (_result *GetDirectorySsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectorySsoStatus"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDirectorySsoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDirectorySsoStatus(request *GetDirectorySsoStatusRequest) (_result *GetDirectorySsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectorySsoStatusResponse{}
	_body, _err := client.GetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficeSiteSsoStatusWithOptions(request *GetOfficeSiteSsoStatusRequest, runtime *util.RuntimeOptions) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOfficeSiteSsoStatus"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOfficeSiteSsoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficeSiteSsoStatus(request *GetOfficeSiteSsoStatusRequest) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOfficeSiteSsoStatusResponse{}
	_body, _err := client.GetOfficeSiteSsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpMetadataWithOptions(request *GetSpMetadataRequest, runtime *util.RuntimeOptions) (_result *GetSpMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSpMetadata"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpMetadata(request *GetSpMetadataRequest) (_result *GetSpMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.GetSpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandleSecurityEventsWithOptions(request *HandleSecurityEventsRequest, runtime *util.RuntimeOptions) (_result *HandleSecurityEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationCode)) {
		query["OperationCode"] = request.OperationCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperationParams)) {
		query["OperationParams"] = request.OperationParams
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityEvent)) {
		query["SecurityEvent"] = request.SecurityEvent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HandleSecurityEvents"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandleSecurityEvents(request *HandleSecurityEventsRequest) (_result *HandleSecurityEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.HandleSecurityEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDirectoryUsersWithOptions(request *ListDirectoryUsersRequest, runtime *util.RuntimeOptions) (_result *ListDirectoryUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OUPath)) {
		query["OUPath"] = request.OUPath
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDirectoryUsers"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDirectoryUsers(request *ListDirectoryUsersRequest) (_result *ListDirectoryUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.ListDirectoryUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfficeSiteOverviewWithOptions(request *ListOfficeSiteOverviewRequest, runtime *util.RuntimeOptions) (_result *ListOfficeSiteOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceRefresh)) {
		query["ForceRefresh"] = request.ForceRefresh
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryRange)) {
		query["QueryRange"] = request.QueryRange
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOfficeSiteOverview"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOfficeSiteOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfficeSiteOverview(request *ListOfficeSiteOverviewRequest) (_result *ListOfficeSiteOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfficeSiteOverviewResponse{}
	_body, _err := client.ListOfficeSiteOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfficeSiteUsersWithOptions(request *ListOfficeSiteUsersRequest, runtime *util.RuntimeOptions) (_result *ListOfficeSiteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OUPath)) {
		query["OUPath"] = request.OUPath
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOfficeSiteUsers"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOfficeSiteUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfficeSiteUsers(request *ListOfficeSiteUsersRequest) (_result *ListOfficeSiteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfficeSiteUsersResponse{}
	_body, _err := client.ListOfficeSiteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserAdOrganizationUnitsWithOptions(request *ListUserAdOrganizationUnitsRequest, runtime *util.RuntimeOptions) (_result *ListUserAdOrganizationUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserAdOrganizationUnits"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserAdOrganizationUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserAdOrganizationUnits(request *ListUserAdOrganizationUnitsRequest) (_result *ListUserAdOrganizationUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserAdOrganizationUnitsResponse{}
	_body, _err := client.ListUserAdOrganizationUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockVirtualMFADeviceWithOptions(request *LockVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *LockVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LockVirtualMFADevice"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockVirtualMFADevice(request *LockVirtualMFADeviceRequest) (_result *LockVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.LockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyADConnectorDirectoryWithOptions(request *ModifyADConnectorDirectoryRequest, runtime *util.RuntimeOptions) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdHostname)) {
		query["AdHostname"] = request.AdHostname
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryName)) {
		query["DirectoryName"] = request.DirectoryName
	}

	if !tea.BoolValue(util.IsUnset(request.DnsAddress)) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainPassword)) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DomainUserName)) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !tea.BoolValue(util.IsUnset(request.MfaEnabled)) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OUName)) {
		query["OUName"] = request.OUName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubDomainDnsAddress)) {
		query["SubDomainDnsAddress"] = request.SubDomainDnsAddress
	}

	if !tea.BoolValue(util.IsUnset(request.SubDomainName)) {
		query["SubDomainName"] = request.SubDomainName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyADConnectorDirectory"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyADConnectorDirectory(request *ModifyADConnectorDirectoryRequest) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.ModifyADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyADConnectorOfficeSiteWithOptions(request *ModifyADConnectorOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdHostname)) {
		query["AdHostname"] = request.AdHostname
	}

	if !tea.BoolValue(util.IsUnset(request.DnsAddress)) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainPassword)) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DomainUserName)) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !tea.BoolValue(util.IsUnset(request.MfaEnabled)) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OUName)) {
		query["OUName"] = request.OUName
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteName)) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubDomainDnsAddress)) {
		query["SubDomainDnsAddress"] = request.SubDomainDnsAddress
	}

	if !tea.BoolValue(util.IsUnset(request.SubDomainName)) {
		query["SubDomainName"] = request.SubDomainName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyADConnectorOfficeSite"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyADConnectorOfficeSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyADConnectorOfficeSite(request *ModifyADConnectorOfficeSiteRequest) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyADConnectorOfficeSiteResponse{}
	_body, _err := client.ModifyADConnectorOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBundleWithOptions(request *ModifyBundleRequest, runtime *util.RuntimeOptions) (_result *ModifyBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.BundleName)) {
		query["BundleName"] = request.BundleName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBundle"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBundleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBundle(request *ModifyBundleRequest) (_result *ModifyBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBundleResponse{}
	_body, _err := client.ModifyBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopChargeTypeWithOptions(request *ModifyDesktopChargeTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDesktopChargeType"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDesktopChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopChargeType(request *ModifyDesktopChargeTypeRequest) (_result *ModifyDesktopChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopChargeTypeResponse{}
	_body, _err := client.ModifyDesktopChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopGroupWithOptions(request *ModifyDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowAutoSetup)) {
		query["AllowAutoSetup"] = request.AllowAutoSetup
	}

	if !tea.BoolValue(util.IsUnset(request.AllowBufferCount)) {
		query["AllowBufferCount"] = request.AllowBufferCount
	}

	if !tea.BoolValue(util.IsUnset(request.BindAmount)) {
		query["BindAmount"] = request.BindAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupName)) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.KeepDuration)) {
		query["KeepDuration"] = request.KeepDuration
	}

	if !tea.BoolValue(util.IsUnset(request.LoadPolicy)) {
		query["LoadPolicy"] = request.LoadPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.MaxDesktopsCount)) {
		query["MaxDesktopsCount"] = request.MaxDesktopsCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinDesktopsCount)) {
		query["MinDesktopsCount"] = request.MinDesktopsCount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnBundleId)) {
		query["OwnBundleId"] = request.OwnBundleId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResetType)) {
		query["ResetType"] = request.ResetType
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleStrategyId)) {
		query["ScaleStrategyId"] = request.ScaleStrategyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDesktopGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDesktopGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopGroup(request *ModifyDesktopGroupRequest) (_result *ModifyDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopGroupResponse{}
	_body, _err := client.ModifyDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopHostNameWithOptions(request *ModifyDesktopHostNameRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopHostNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.NewHostName)) {
		query["NewHostName"] = request.NewHostName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDesktopHostName"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDesktopHostNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopHostName(request *ModifyDesktopHostNameRequest) (_result *ModifyDesktopHostNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopHostNameResponse{}
	_body, _err := client.ModifyDesktopHostNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopNameWithOptions(request *ModifyDesktopNameRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.NewDesktopName)) {
		query["NewDesktopName"] = request.NewDesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDesktopName"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopName(request *ModifyDesktopNameRequest) (_result *ModifyDesktopNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.ModifyDesktopNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopSpecWithOptions(request *ModifyDesktopSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopType)) {
		query["DesktopType"] = request.DesktopType
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RootDiskSizeGib)) {
		query["RootDiskSizeGib"] = request.RootDiskSizeGib
	}

	if !tea.BoolValue(util.IsUnset(request.UserDiskPerformanceLevel)) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.UserDiskSizeGib)) {
		query["UserDiskSizeGib"] = request.UserDiskSizeGib
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDesktopSpec"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDesktopSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopSpec(request *ModifyDesktopSpecRequest) (_result *ModifyDesktopSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopSpecResponse{}
	_body, _err := client.ModifyDesktopSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopsPolicyGroupWithOptions(request *ModifyDesktopsPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDesktopsPolicyGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopsPolicyGroup(request *ModifyDesktopsPolicyGroupRequest) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.ModifyDesktopsPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEntitlementWithOptions(request *ModifyEntitlementRequest, runtime *util.RuntimeOptions) (_result *ModifyEntitlementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyEntitlement"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEntitlement(request *ModifyEntitlementRequest) (_result *ModifyEntitlementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.ModifyEntitlementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageAttributeWithOptions(request *ModifyImageAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyImageAttribute"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (_result *ModifyImageAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.ModifyImageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNASDefaultMountTargetWithOptions(request *ModifyNASDefaultMountTargetRequest, runtime *util.RuntimeOptions) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNASDefaultMountTarget"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNASDefaultMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNASDefaultMountTarget(request *ModifyNASDefaultMountTargetRequest) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNASDefaultMountTargetResponse{}
	_body, _err := client.ModifyNASDefaultMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkPackageWithOptions(request *ModifyNetworkPackageRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkPackageId)) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNetworkPackage"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNetworkPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkPackage(request *ModifyNetworkPackageRequest) (_result *ModifyNetworkPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkPackageResponse{}
	_body, _err := client.ModifyNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkPackageBandwidthWithOptions(request *ModifyNetworkPackageBandwidthRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkPackageBandwidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkPackageId)) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNetworkPackageBandwidth"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNetworkPackageBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkPackageBandwidth(request *ModifyNetworkPackageBandwidthRequest) (_result *ModifyNetworkPackageBandwidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkPackageBandwidthResponse{}
	_body, _err := client.ModifyNetworkPackageBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkPackageEnabledWithOptions(request *ModifyNetworkPackageEnabledRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkPackageId)) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNetworkPackageEnabled"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNetworkPackageEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkPackageEnabled(request *ModifyNetworkPackageEnabledRequest) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkPackageEnabledResponse{}
	_body, _err := client.ModifyNetworkPackageEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteAttributeWithOptions(request *ModifyOfficeSiteAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopAccessType)) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.NeedVerifyLoginRisk)) {
		query["NeedVerifyLoginRisk"] = request.NeedVerifyLoginRisk
	}

	if !tea.BoolValue(util.IsUnset(request.NeedVerifyZeroDevice)) {
		query["NeedVerifyZeroDevice"] = request.NeedVerifyZeroDevice
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteName)) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyOfficeSiteAttribute"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyOfficeSiteAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteAttribute(request *ModifyOfficeSiteAttributeRequest) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteAttributeResponse{}
	_body, _err := client.ModifyOfficeSiteAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteCrossDesktopAccessWithOptions(request *ModifyOfficeSiteCrossDesktopAccessRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableCrossDesktopAccess)) {
		query["EnableCrossDesktopAccess"] = request.EnableCrossDesktopAccess
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyOfficeSiteCrossDesktopAccess"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyOfficeSiteCrossDesktopAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteCrossDesktopAccess(request *ModifyOfficeSiteCrossDesktopAccessRequest) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteCrossDesktopAccessResponse{}
	_body, _err := client.ModifyOfficeSiteCrossDesktopAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteMfaEnabledWithOptions(request *ModifyOfficeSiteMfaEnabledRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MfaEnabled)) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyOfficeSiteMfaEnabled"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyOfficeSiteMfaEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteMfaEnabled(request *ModifyOfficeSiteMfaEnabledRequest) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteMfaEnabledResponse{}
	_body, _err := client.ModifyOfficeSiteMfaEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOperateVulWithOptions(request *ModifyOperateVulRequest, runtime *util.RuntimeOptions) (_result *ModifyOperateVulResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VulInfo)) {
		query["VulInfo"] = request.VulInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyOperateVul"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyOperateVulResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOperateVul(request *ModifyOperateVulRequest) (_result *ModifyOperateVulResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOperateVulResponse{}
	_body, _err := client.ModifyOperateVulWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPolicyGroupWithOptions(request *ModifyPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizeAccessPolicyRule)) {
		query["AuthorizeAccessPolicyRule"] = request.AuthorizeAccessPolicyRule
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizeSecurityPolicyRule)) {
		query["AuthorizeSecurityPolicyRule"] = request.AuthorizeSecurityPolicyRule
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Clipboard)) {
		query["Clipboard"] = request.Clipboard
	}

	if !tea.BoolValue(util.IsUnset(request.DomainList)) {
		query["DomainList"] = request.DomainList
	}

	if !tea.BoolValue(util.IsUnset(request.GpuAcceleration)) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !tea.BoolValue(util.IsUnset(request.Html5Access)) {
		query["Html5Access"] = request.Html5Access
	}

	if !tea.BoolValue(util.IsUnset(request.Html5FileTransfer)) {
		query["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.LocalDrive)) {
		query["LocalDrive"] = request.LocalDrive
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PreemptLogin)) {
		query["PreemptLogin"] = request.PreemptLogin
	}

	if !tea.BoolValue(util.IsUnset(request.PreemptLoginUser)) {
		query["PreemptLoginUser"] = request.PreemptLoginUser
	}

	if !tea.BoolValue(util.IsUnset(request.PrinterRedirection)) {
		query["PrinterRedirection"] = request.PrinterRedirection
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RevokeAccessPolicyRule)) {
		query["RevokeAccessPolicyRule"] = request.RevokeAccessPolicyRule
	}

	if !tea.BoolValue(util.IsUnset(request.RevokeSecurityPolicyRule)) {
		query["RevokeSecurityPolicyRule"] = request.RevokeSecurityPolicyRule
	}

	if !tea.BoolValue(util.IsUnset(request.UsbRedirect)) {
		query["UsbRedirect"] = request.UsbRedirect
	}

	if !tea.BoolValue(util.IsUnset(request.UsbSupplyRedirectRule)) {
		query["UsbSupplyRedirectRule"] = request.UsbSupplyRedirectRule
	}

	if !tea.BoolValue(util.IsUnset(request.VisualQuality)) {
		query["VisualQuality"] = request.VisualQuality
	}

	if !tea.BoolValue(util.IsUnset(request.Watermark)) {
		query["Watermark"] = request.Watermark
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkCustomText)) {
		query["WatermarkCustomText"] = request.WatermarkCustomText
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkTransparency)) {
		query["WatermarkTransparency"] = request.WatermarkTransparency
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkType)) {
		query["WatermarkType"] = request.WatermarkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolicyGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (_result *ModifyPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.ModifyPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserEntitlementWithOptions(request *ModifyUserEntitlementRequest, runtime *util.RuntimeOptions) (_result *ModifyUserEntitlementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizeDesktopId)) {
		query["AuthorizeDesktopId"] = request.AuthorizeDesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RevokeDesktopId)) {
		query["RevokeDesktopId"] = request.RevokeDesktopId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserEntitlement"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserEntitlementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserEntitlement(request *ModifyUserEntitlementRequest) (_result *ModifyUserEntitlementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserEntitlementResponse{}
	_body, _err := client.ModifyUserEntitlementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserToDesktopGroupWithOptions(request *ModifyUserToDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyUserToDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NewEndUserIds)) {
		query["NewEndUserIds"] = request.NewEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OldEndUserIds)) {
		query["OldEndUserIds"] = request.OldEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserToDesktopGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserToDesktopGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserToDesktopGroup(request *ModifyUserToDesktopGroupRequest) (_result *ModifyUserToDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserToDesktopGroupResponse{}
	_body, _err := client.ModifyUserToDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateVulsWithOptions(request *OperateVulsRequest, runtime *util.RuntimeOptions) (_result *OperateVulsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.Precondition)) {
		query["Precondition"] = request.Precondition
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VulName)) {
		query["VulName"] = request.VulName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateVuls"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateVulsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateVuls(request *OperateVulsRequest) (_result *OperateVulsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateVulsResponse{}
	_body, _err := client.OperateVulsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootDesktops(request *RebootDesktopsRequest) (_result *RebootDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.RebootDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebuildDesktopsWithOptions(request *RebuildDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebuildDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebuildDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebuildDesktops(request *RebuildDesktopsRequest) (_result *RebuildDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.RebuildDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUserFromDesktopGroupWithOptions(request *RemoveUserFromDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopGroupId)) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopGroupIds)) {
		query["DesktopGroupIds"] = request.DesktopGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserFromDesktopGroup"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserFromDesktopGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserFromDesktopGroup(request *RemoveUserFromDesktopGroupRequest) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserFromDesktopGroupResponse{}
	_body, _err := client.RemoveUserFromDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewDesktopsWithOptions(request *RenewDesktopsRequest, runtime *util.RuntimeOptions) (_result *RenewDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewDesktops(request *RenewDesktopsRequest) (_result *RenewDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.RenewDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewNetworkPackagesWithOptions(request *RenewNetworkPackagesRequest, runtime *util.RuntimeOptions) (_result *RenewNetworkPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkPackageId)) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewNetworkPackages"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewNetworkPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewNetworkPackages(request *RenewNetworkPackagesRequest) (_result *RenewNetworkPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewNetworkPackagesResponse{}
	_body, _err := client.RenewNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetNASDefaultMountTargetWithOptions(request *ResetNASDefaultMountTargetRequest, runtime *util.RuntimeOptions) (_result *ResetNASDefaultMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetNASDefaultMountTarget"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetNASDefaultMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetNASDefaultMountTarget(request *ResetNASDefaultMountTargetRequest) (_result *ResetNASDefaultMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetNASDefaultMountTargetResponse{}
	_body, _err := client.ResetNASDefaultMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSnapshotWithOptions(request *ResetSnapshotRequest, runtime *util.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSnapshot"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSnapshot(request *ResetSnapshotRequest) (_result *ResetSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.ResetSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackSuspEventQuaraFileWithOptions(request *RollbackSuspEventQuaraFileRequest, runtime *util.RuntimeOptions) (_result *RollbackSuspEventQuaraFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.QuaraFieldId)) {
		query["QuaraFieldId"] = request.QuaraFieldId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackSuspEventQuaraFile"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackSuspEventQuaraFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackSuspEventQuaraFile(request *RollbackSuspEventQuaraFileRequest) (_result *RollbackSuspEventQuaraFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackSuspEventQuaraFileResponse{}
	_body, _err := client.RollbackSuspEventQuaraFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCommandWithOptions(request *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		query["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.ContentEncoding)) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCommand"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendVerifyCodeWithOptions(request *SendVerifyCodeRequest, runtime *util.RuntimeOptions) (_result *SendVerifyCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCodeAction)) {
		query["VerifyCodeAction"] = request.VerifyCodeAction
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerifyCode"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerifyCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendVerifyCode(request *SendVerifyCodeRequest) (_result *SendVerifyCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerifyCodeResponse{}
	_body, _err := client.SendVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDirectorySsoStatusWithOptions(request *SetDirectorySsoStatusRequest, runtime *util.RuntimeOptions) (_result *SetDirectorySsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSso)) {
		query["EnableSso"] = request.EnableSso
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDirectorySsoStatus"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDirectorySsoStatus(request *SetDirectorySsoStatusRequest) (_result *SetDirectorySsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.SetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetIdpMetadataWithOptions(request *SetIdpMetadataRequest, runtime *util.RuntimeOptions) (_result *SetIdpMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.IdpMetadata)) {
		query["IdpMetadata"] = request.IdpMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetIdpMetadata"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetIdpMetadata(request *SetIdpMetadataRequest) (_result *SetIdpMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.SetIdpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetOfficeSiteSsoStatusWithOptions(request *SetOfficeSiteSsoStatusRequest, runtime *util.RuntimeOptions) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableSso)) {
		query["EnableSso"] = request.EnableSso
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SsoType)) {
		query["SsoType"] = request.SsoType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetOfficeSiteSsoStatus"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetOfficeSiteSsoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetOfficeSiteSsoStatus(request *SetOfficeSiteSsoStatusRequest) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetOfficeSiteSsoStatusResponse{}
	_body, _err := client.SetOfficeSiteSsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *util.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDesktops(request *StartDesktopsRequest) (_result *StartDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDesktopsResponse{}
	_body, _err := client.StartDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartVirusScanTaskWithOptions(request *StartVirusScanTaskRequest, runtime *util.RuntimeOptions) (_result *StartVirusScanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartVirusScanTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartVirusScanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartVirusScanTask(request *StartVirusScanTaskRequest) (_result *StartVirusScanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartVirusScanTaskResponse{}
	_body, _err := client.StartVirusScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *util.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StoppedMode)) {
		query["StoppedMode"] = request.StoppedMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDesktops"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDesktops(request *StopDesktopsRequest) (_result *StopDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDesktopsResponse{}
	_body, _err := client.StopDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInvocationWithOptions(request *StopInvocationRequest, runtime *util.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		query["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInvocation"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInvocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInvocation(request *StopInvocationRequest) (_result *StopInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInvocationResponse{}
	_body, _err := client.StopInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockVirtualMFADeviceWithOptions(request *UnlockVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockVirtualMFADevice"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockVirtualMFADevice(request *UnlockVirtualMFADeviceRequest) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.UnlockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFotaTaskWithOptions(request *UpdateFotaTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateFotaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUid)) {
		query["TaskUid"] = request.TaskUid
	}

	if !tea.BoolValue(util.IsUnset(request.UserStatus)) {
		query["UserStatus"] = request.UserStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFotaTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFotaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFotaTask(request *UpdateFotaTaskRequest) (_result *UpdateFotaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFotaTaskResponse{}
	_body, _err := client.UpdateFotaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadImageWithOptions(request *UploadImageRequest, runtime *util.RuntimeOptions) (_result *UploadImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataDiskSize)) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GpuCategory)) {
		query["GpuCategory"] = request.GpuCategory
	}

	if !tea.BoolValue(util.IsUnset(request.GpuDriverType)) {
		query["GpuDriverType"] = request.GpuDriverType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseType)) {
		query["LicenseType"] = request.LicenseType
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectPath)) {
		query["OssObjectPath"] = request.OssObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadImage"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadImage(request *UploadImageRequest) (_result *UploadImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadImageResponse{}
	_body, _err := client.UploadImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyCenWithOptions(request *VerifyCenRequest, runtime *util.RuntimeOptions) (_result *VerifyCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CenOwnerId)) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.CidrBlock)) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyCen"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyCen(request *VerifyCenRequest) (_result *VerifyCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyCenResponse{}
	_body, _err := client.VerifyCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
