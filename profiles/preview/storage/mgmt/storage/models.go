// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package storage

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessTier = original.AccessTier

const (
	Cool AccessTier = original.Cool
	Hot  AccessTier = original.Hot
)

type AccountExpand = original.AccountExpand

const (
	AccountExpandGeoReplicationStats AccountExpand = original.AccountExpandGeoReplicationStats
)

type AccountStatus = original.AccountStatus

const (
	Available   AccountStatus = original.Available
	Unavailable AccountStatus = original.Unavailable
)

type Action = original.Action

const (
	Allow Action = original.Allow
)

type Action1 = original.Action1

const (
	Acquire Action1 = original.Acquire
	Break   Action1 = original.Break
	Change  Action1 = original.Change
	Release Action1 = original.Release
	Renew   Action1 = original.Renew
)

type Bypass = original.Bypass

const (
	AzureServices Bypass = original.AzureServices
	Logging       Bypass = original.Logging
	Metrics       Bypass = original.Metrics
	None          Bypass = original.None
)

type DefaultAction = original.DefaultAction

const (
	DefaultActionAllow DefaultAction = original.DefaultActionAllow
	DefaultActionDeny  DefaultAction = original.DefaultActionDeny
)

type DirectoryServiceOptions = original.DirectoryServiceOptions

const (
	DirectoryServiceOptionsAADDS DirectoryServiceOptions = original.DirectoryServiceOptionsAADDS
	DirectoryServiceOptionsAD    DirectoryServiceOptions = original.DirectoryServiceOptionsAD
	DirectoryServiceOptionsNone  DirectoryServiceOptions = original.DirectoryServiceOptionsNone
)

type GeoReplicationStatus = original.GeoReplicationStatus

const (
	GeoReplicationStatusBootstrap   GeoReplicationStatus = original.GeoReplicationStatusBootstrap
	GeoReplicationStatusLive        GeoReplicationStatus = original.GeoReplicationStatusLive
	GeoReplicationStatusUnavailable GeoReplicationStatus = original.GeoReplicationStatusUnavailable
)

type HTTPProtocol = original.HTTPProtocol

const (
	HTTPS     HTTPProtocol = original.HTTPS
	Httpshttp HTTPProtocol = original.Httpshttp
)

type ImmutabilityPolicyState = original.ImmutabilityPolicyState

const (
	Locked   ImmutabilityPolicyState = original.Locked
	Unlocked ImmutabilityPolicyState = original.Unlocked
)

type ImmutabilityPolicyUpdateType = original.ImmutabilityPolicyUpdateType

const (
	Extend ImmutabilityPolicyUpdateType = original.Extend
	Lock   ImmutabilityPolicyUpdateType = original.Lock
	Put    ImmutabilityPolicyUpdateType = original.Put
)

type KeyPermission = original.KeyPermission

const (
	Full KeyPermission = original.Full
	Read KeyPermission = original.Read
)

type KeySource = original.KeySource

const (
	MicrosoftKeyvault KeySource = original.MicrosoftKeyvault
	MicrosoftStorage  KeySource = original.MicrosoftStorage
)

type Kind = original.Kind

const (
	BlobStorage      Kind = original.BlobStorage
	BlockBlobStorage Kind = original.BlockBlobStorage
	FileStorage      Kind = original.FileStorage
	Storage          Kind = original.Storage
	StorageV2        Kind = original.StorageV2
)

type LargeFileSharesState = original.LargeFileSharesState

const (
	Disabled LargeFileSharesState = original.Disabled
	Enabled  LargeFileSharesState = original.Enabled
)

type LeaseDuration = original.LeaseDuration

const (
	Fixed    LeaseDuration = original.Fixed
	Infinite LeaseDuration = original.Infinite
)

type LeaseState = original.LeaseState

const (
	LeaseStateAvailable LeaseState = original.LeaseStateAvailable
	LeaseStateBreaking  LeaseState = original.LeaseStateBreaking
	LeaseStateBroken    LeaseState = original.LeaseStateBroken
	LeaseStateExpired   LeaseState = original.LeaseStateExpired
	LeaseStateLeased    LeaseState = original.LeaseStateLeased
)

type LeaseStatus = original.LeaseStatus

const (
	LeaseStatusLocked   LeaseStatus = original.LeaseStatusLocked
	LeaseStatusUnlocked LeaseStatus = original.LeaseStatusUnlocked
)

type ListKeyExpand = original.ListKeyExpand

const (
	Kerb ListKeyExpand = original.Kerb
)

type Permissions = original.Permissions

const (
	A Permissions = original.A
	C Permissions = original.C
	D Permissions = original.D
	L Permissions = original.L
	P Permissions = original.P
	R Permissions = original.R
	U Permissions = original.U
	W Permissions = original.W
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	Creating  PrivateEndpointConnectionProvisioningState = original.Creating
	Deleting  PrivateEndpointConnectionProvisioningState = original.Deleting
	Failed    PrivateEndpointConnectionProvisioningState = original.Failed
	Succeeded PrivateEndpointConnectionProvisioningState = original.Succeeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	Approved PrivateEndpointServiceConnectionStatus = original.Approved
	Pending  PrivateEndpointServiceConnectionStatus = original.Pending
	Rejected PrivateEndpointServiceConnectionStatus = original.Rejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCreating     ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateResolvingDNS ProvisioningState = original.ProvisioningStateResolvingDNS
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
)

type PublicAccess = original.PublicAccess

const (
	PublicAccessBlob      PublicAccess = original.PublicAccessBlob
	PublicAccessContainer PublicAccess = original.PublicAccessContainer
	PublicAccessNone      PublicAccess = original.PublicAccessNone
)

type Reason = original.Reason

const (
	AccountNameInvalid Reason = original.AccountNameInvalid
	AlreadyExists      Reason = original.AlreadyExists
)

type ReasonCode = original.ReasonCode

const (
	NotAvailableForSubscription ReasonCode = original.NotAvailableForSubscription
	QuotaID                     ReasonCode = original.QuotaID
)

type RoutingChoice = original.RoutingChoice

const (
	InternetRouting  RoutingChoice = original.InternetRouting
	MicrosoftRouting RoutingChoice = original.MicrosoftRouting
)

type Services = original.Services

const (
	B Services = original.B
	F Services = original.F
	Q Services = original.Q
	T Services = original.T
)

type SignedResource = original.SignedResource

const (
	SignedResourceB SignedResource = original.SignedResourceB
	SignedResourceC SignedResource = original.SignedResourceC
	SignedResourceF SignedResource = original.SignedResourceF
	SignedResourceS SignedResource = original.SignedResourceS
)

type SignedResourceTypes = original.SignedResourceTypes

const (
	SignedResourceTypesC SignedResourceTypes = original.SignedResourceTypesC
	SignedResourceTypesO SignedResourceTypes = original.SignedResourceTypesO
	SignedResourceTypesS SignedResourceTypes = original.SignedResourceTypesS
)

type SkuName = original.SkuName

const (
	PremiumLRS     SkuName = original.PremiumLRS
	PremiumZRS     SkuName = original.PremiumZRS
	StandardGRS    SkuName = original.StandardGRS
	StandardGZRS   SkuName = original.StandardGZRS
	StandardLRS    SkuName = original.StandardLRS
	StandardRAGRS  SkuName = original.StandardRAGRS
	StandardRAGZRS SkuName = original.StandardRAGZRS
	StandardZRS    SkuName = original.StandardZRS
)

type SkuTier = original.SkuTier

const (
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type State = original.State

const (
	StateDeprovisioning       State = original.StateDeprovisioning
	StateFailed               State = original.StateFailed
	StateNetworkSourceDeleted State = original.StateNetworkSourceDeleted
	StateProvisioning         State = original.StateProvisioning
	StateSucceeded            State = original.StateSucceeded
)

type UsageUnit = original.UsageUnit

const (
	Bytes           UsageUnit = original.Bytes
	BytesPerSecond  UsageUnit = original.BytesPerSecond
	Count           UsageUnit = original.Count
	CountsPerSecond UsageUnit = original.CountsPerSecond
	Percent         UsageUnit = original.Percent
	Seconds         UsageUnit = original.Seconds
)

type Account = original.Account
type AccountCheckNameAvailabilityParameters = original.AccountCheckNameAvailabilityParameters
type AccountCreateParameters = original.AccountCreateParameters
type AccountInternetEndpoints = original.AccountInternetEndpoints
type AccountKey = original.AccountKey
type AccountListKeysResult = original.AccountListKeysResult
type AccountListResult = original.AccountListResult
type AccountListResultIterator = original.AccountListResultIterator
type AccountListResultPage = original.AccountListResultPage
type AccountMicrosoftEndpoints = original.AccountMicrosoftEndpoints
type AccountProperties = original.AccountProperties
type AccountPropertiesCreateParameters = original.AccountPropertiesCreateParameters
type AccountPropertiesUpdateParameters = original.AccountPropertiesUpdateParameters
type AccountRegenerateKeyParameters = original.AccountRegenerateKeyParameters
type AccountSasParameters = original.AccountSasParameters
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountsClient = original.AccountsClient
type AccountsCreateFuture = original.AccountsCreateFuture
type AccountsFailoverFuture = original.AccountsFailoverFuture
type ActiveDirectoryProperties = original.ActiveDirectoryProperties
type AzureEntityResource = original.AzureEntityResource
type AzureFilesIdentityBasedAuthentication = original.AzureFilesIdentityBasedAuthentication
type BaseClient = original.BaseClient
type BlobContainer = original.BlobContainer
type BlobContainersClient = original.BlobContainersClient
type BlobServiceItems = original.BlobServiceItems
type BlobServiceProperties = original.BlobServiceProperties
type BlobServicePropertiesProperties = original.BlobServicePropertiesProperties
type BlobServicesClient = original.BlobServicesClient
type ChangeFeed = original.ChangeFeed
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ContainerProperties = original.ContainerProperties
type CorsRule = original.CorsRule
type CorsRules = original.CorsRules
type CustomDomain = original.CustomDomain
type DateAfterCreation = original.DateAfterCreation
type DateAfterModification = original.DateAfterModification
type DeleteRetentionPolicy = original.DeleteRetentionPolicy
type Dimension = original.Dimension
type Encryption = original.Encryption
type EncryptionService = original.EncryptionService
type EncryptionServices = original.EncryptionServices
type Endpoints = original.Endpoints
type ErrorResponse = original.ErrorResponse
type FileServiceItems = original.FileServiceItems
type FileServiceProperties = original.FileServiceProperties
type FileServicePropertiesProperties = original.FileServicePropertiesProperties
type FileServicesClient = original.FileServicesClient
type FileShare = original.FileShare
type FileShareItem = original.FileShareItem
type FileShareItems = original.FileShareItems
type FileShareItemsIterator = original.FileShareItemsIterator
type FileShareItemsPage = original.FileShareItemsPage
type FileShareProperties = original.FileShareProperties
type FileSharesClient = original.FileSharesClient
type GeoReplicationStats = original.GeoReplicationStats
type IPRule = original.IPRule
type Identity = original.Identity
type ImmutabilityPolicy = original.ImmutabilityPolicy
type ImmutabilityPolicyProperties = original.ImmutabilityPolicyProperties
type ImmutabilityPolicyProperty = original.ImmutabilityPolicyProperty
type KeyVaultProperties = original.KeyVaultProperties
type LeaseContainerRequest = original.LeaseContainerRequest
type LeaseContainerResponse = original.LeaseContainerResponse
type LegalHold = original.LegalHold
type LegalHoldProperties = original.LegalHoldProperties
type ListAccountSasResponse = original.ListAccountSasResponse
type ListContainerItem = original.ListContainerItem
type ListContainerItems = original.ListContainerItems
type ListContainerItemsIterator = original.ListContainerItemsIterator
type ListContainerItemsPage = original.ListContainerItemsPage
type ListServiceSasResponse = original.ListServiceSasResponse
type ManagementPoliciesClient = original.ManagementPoliciesClient
type ManagementPolicy = original.ManagementPolicy
type ManagementPolicyAction = original.ManagementPolicyAction
type ManagementPolicyBaseBlob = original.ManagementPolicyBaseBlob
type ManagementPolicyDefinition = original.ManagementPolicyDefinition
type ManagementPolicyFilter = original.ManagementPolicyFilter
type ManagementPolicyProperties = original.ManagementPolicyProperties
type ManagementPolicyRule = original.ManagementPolicyRule
type ManagementPolicySchema = original.ManagementPolicySchema
type ManagementPolicySnapShot = original.ManagementPolicySnapShot
type MetricSpecification = original.MetricSpecification
type NetworkRuleSet = original.NetworkRuleSet
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type Restriction = original.Restriction
type RoutingPreference = original.RoutingPreference
type SKUCapability = original.SKUCapability
type ServiceSasParameters = original.ServiceSasParameters
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type SkuInformation = original.SkuInformation
type SkuListResult = original.SkuListResult
type SkusClient = original.SkusClient
type TagProperty = original.TagProperty
type TrackedResource = original.TrackedResource
type UpdateHistoryProperty = original.UpdateHistoryProperty
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageName = original.UsageName
type UsagesClient = original.UsagesClient
type VirtualNetworkRule = original.VirtualNetworkRule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountListResultIterator(page AccountListResultPage) AccountListResultIterator {
	return original.NewAccountListResultIterator(page)
}
func NewAccountListResultPage(getNextPage func(context.Context, AccountListResult) (AccountListResult, error)) AccountListResultPage {
	return original.NewAccountListResultPage(getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBlobContainersClient(subscriptionID string) BlobContainersClient {
	return original.NewBlobContainersClient(subscriptionID)
}
func NewBlobContainersClientWithBaseURI(baseURI string, subscriptionID string) BlobContainersClient {
	return original.NewBlobContainersClientWithBaseURI(baseURI, subscriptionID)
}
func NewBlobServicesClient(subscriptionID string) BlobServicesClient {
	return original.NewBlobServicesClient(subscriptionID)
}
func NewBlobServicesClientWithBaseURI(baseURI string, subscriptionID string) BlobServicesClient {
	return original.NewBlobServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFileServicesClient(subscriptionID string) FileServicesClient {
	return original.NewFileServicesClient(subscriptionID)
}
func NewFileServicesClientWithBaseURI(baseURI string, subscriptionID string) FileServicesClient {
	return original.NewFileServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFileShareItemsIterator(page FileShareItemsPage) FileShareItemsIterator {
	return original.NewFileShareItemsIterator(page)
}
func NewFileShareItemsPage(getNextPage func(context.Context, FileShareItems) (FileShareItems, error)) FileShareItemsPage {
	return original.NewFileShareItemsPage(getNextPage)
}
func NewFileSharesClient(subscriptionID string) FileSharesClient {
	return original.NewFileSharesClient(subscriptionID)
}
func NewFileSharesClientWithBaseURI(baseURI string, subscriptionID string) FileSharesClient {
	return original.NewFileSharesClientWithBaseURI(baseURI, subscriptionID)
}
func NewListContainerItemsIterator(page ListContainerItemsPage) ListContainerItemsIterator {
	return original.NewListContainerItemsIterator(page)
}
func NewListContainerItemsPage(getNextPage func(context.Context, ListContainerItems) (ListContainerItems, error)) ListContainerItemsPage {
	return original.NewListContainerItemsPage(getNextPage)
}
func NewManagementPoliciesClient(subscriptionID string) ManagementPoliciesClient {
	return original.NewManagementPoliciesClient(subscriptionID)
}
func NewManagementPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ManagementPoliciesClient {
	return original.NewManagementPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessTierValues() []AccessTier {
	return original.PossibleAccessTierValues()
}
func PossibleAccountExpandValues() []AccountExpand {
	return original.PossibleAccountExpandValues()
}
func PossibleAccountStatusValues() []AccountStatus {
	return original.PossibleAccountStatusValues()
}
func PossibleAction1Values() []Action1 {
	return original.PossibleAction1Values()
}
func PossibleActionValues() []Action {
	return original.PossibleActionValues()
}
func PossibleBypassValues() []Bypass {
	return original.PossibleBypassValues()
}
func PossibleDefaultActionValues() []DefaultAction {
	return original.PossibleDefaultActionValues()
}
func PossibleDirectoryServiceOptionsValues() []DirectoryServiceOptions {
	return original.PossibleDirectoryServiceOptionsValues()
}
func PossibleGeoReplicationStatusValues() []GeoReplicationStatus {
	return original.PossibleGeoReplicationStatusValues()
}
func PossibleHTTPProtocolValues() []HTTPProtocol {
	return original.PossibleHTTPProtocolValues()
}
func PossibleImmutabilityPolicyStateValues() []ImmutabilityPolicyState {
	return original.PossibleImmutabilityPolicyStateValues()
}
func PossibleImmutabilityPolicyUpdateTypeValues() []ImmutabilityPolicyUpdateType {
	return original.PossibleImmutabilityPolicyUpdateTypeValues()
}
func PossibleKeyPermissionValues() []KeyPermission {
	return original.PossibleKeyPermissionValues()
}
func PossibleKeySourceValues() []KeySource {
	return original.PossibleKeySourceValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLargeFileSharesStateValues() []LargeFileSharesState {
	return original.PossibleLargeFileSharesStateValues()
}
func PossibleLeaseDurationValues() []LeaseDuration {
	return original.PossibleLeaseDurationValues()
}
func PossibleLeaseStateValues() []LeaseState {
	return original.PossibleLeaseStateValues()
}
func PossibleLeaseStatusValues() []LeaseStatus {
	return original.PossibleLeaseStatusValues()
}
func PossibleListKeyExpandValues() []ListKeyExpand {
	return original.PossibleListKeyExpandValues()
}
func PossiblePermissionsValues() []Permissions {
	return original.PossiblePermissionsValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicAccessValues() []PublicAccess {
	return original.PossiblePublicAccessValues()
}
func PossibleReasonCodeValues() []ReasonCode {
	return original.PossibleReasonCodeValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleRoutingChoiceValues() []RoutingChoice {
	return original.PossibleRoutingChoiceValues()
}
func PossibleServicesValues() []Services {
	return original.PossibleServicesValues()
}
func PossibleSignedResourceTypesValues() []SignedResourceTypes {
	return original.PossibleSignedResourceTypesValues()
}
func PossibleSignedResourceValues() []SignedResource {
	return original.PossibleSignedResourceValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleUsageUnitValues() []UsageUnit {
	return original.PossibleUsageUnitValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
