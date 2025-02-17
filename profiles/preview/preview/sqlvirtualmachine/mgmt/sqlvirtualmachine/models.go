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

package sqlvirtualmachine

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/sqlvirtualmachine/mgmt/2017-03-01-preview/sqlvirtualmachine"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BackupScheduleType = original.BackupScheduleType

const (
	Automated BackupScheduleType = original.Automated
	Manual    BackupScheduleType = original.Manual
)

type ClusterConfiguration = original.ClusterConfiguration

const (
	Domainful ClusterConfiguration = original.Domainful
)

type ClusterManagerType = original.ClusterManagerType

const (
	WSFC ClusterManagerType = original.WSFC
)

type ConnectivityType = original.ConnectivityType

const (
	LOCAL   ConnectivityType = original.LOCAL
	PRIVATE ConnectivityType = original.PRIVATE
	PUBLIC  ConnectivityType = original.PUBLIC
)

type DayOfWeek = original.DayOfWeek

const (
	Friday    DayOfWeek = original.Friday
	Monday    DayOfWeek = original.Monday
	Saturday  DayOfWeek = original.Saturday
	Sunday    DayOfWeek = original.Sunday
	Thursday  DayOfWeek = original.Thursday
	Tuesday   DayOfWeek = original.Tuesday
	Wednesday DayOfWeek = original.Wednesday
)

type DiskConfigurationType = original.DiskConfigurationType

const (
	ADD    DiskConfigurationType = original.ADD
	EXTEND DiskConfigurationType = original.EXTEND
	NEW    DiskConfigurationType = original.NEW
)

type FullBackupFrequencyType = original.FullBackupFrequencyType

const (
	Daily  FullBackupFrequencyType = original.Daily
	Weekly FullBackupFrequencyType = original.Weekly
)

type IdentityType = original.IdentityType

const (
	SystemAssigned IdentityType = original.SystemAssigned
)

type OperationOrigin = original.OperationOrigin

const (
	System OperationOrigin = original.System
	User   OperationOrigin = original.User
)

type SQLImageSku = original.SQLImageSku

const (
	Developer  SQLImageSku = original.Developer
	Enterprise SQLImageSku = original.Enterprise
	Express    SQLImageSku = original.Express
	Standard   SQLImageSku = original.Standard
	Web        SQLImageSku = original.Web
)

type SQLManagementMode = original.SQLManagementMode

const (
	Full        SQLManagementMode = original.Full
	LightWeight SQLManagementMode = original.LightWeight
	NoAgent     SQLManagementMode = original.NoAgent
)

type SQLServerLicenseType = original.SQLServerLicenseType

const (
	AHUB SQLServerLicenseType = original.AHUB
	DR   SQLServerLicenseType = original.DR
	PAYG SQLServerLicenseType = original.PAYG
)

type SQLVMGroupImageSku = original.SQLVMGroupImageSku

const (
	SQLVMGroupImageSkuDeveloper  SQLVMGroupImageSku = original.SQLVMGroupImageSkuDeveloper
	SQLVMGroupImageSkuEnterprise SQLVMGroupImageSku = original.SQLVMGroupImageSkuEnterprise
)

type SQLWorkloadType = original.SQLWorkloadType

const (
	DW      SQLWorkloadType = original.DW
	GENERAL SQLWorkloadType = original.GENERAL
	OLTP    SQLWorkloadType = original.OLTP
)

type ScaleType = original.ScaleType

const (
	HA ScaleType = original.HA
)

type StorageWorkloadType = original.StorageWorkloadType

const (
	StorageWorkloadTypeDW      StorageWorkloadType = original.StorageWorkloadTypeDW
	StorageWorkloadTypeGENERAL StorageWorkloadType = original.StorageWorkloadTypeGENERAL
	StorageWorkloadTypeOLTP    StorageWorkloadType = original.StorageWorkloadTypeOLTP
)

type AdditionalFeaturesServerConfigurations = original.AdditionalFeaturesServerConfigurations
type AutoBackupSettings = original.AutoBackupSettings
type AutoPatchingSettings = original.AutoPatchingSettings
type AvailabilityGroupListener = original.AvailabilityGroupListener
type AvailabilityGroupListenerListResult = original.AvailabilityGroupListenerListResult
type AvailabilityGroupListenerListResultIterator = original.AvailabilityGroupListenerListResultIterator
type AvailabilityGroupListenerListResultPage = original.AvailabilityGroupListenerListResultPage
type AvailabilityGroupListenerProperties = original.AvailabilityGroupListenerProperties
type AvailabilityGroupListenersClient = original.AvailabilityGroupListenersClient
type AvailabilityGroupListenersCreateOrUpdateFuture = original.AvailabilityGroupListenersCreateOrUpdateFuture
type AvailabilityGroupListenersDeleteFuture = original.AvailabilityGroupListenersDeleteFuture
type BaseClient = original.BaseClient
type Group = original.Group
type GroupListResult = original.GroupListResult
type GroupListResultIterator = original.GroupListResultIterator
type GroupListResultPage = original.GroupListResultPage
type GroupProperties = original.GroupProperties
type GroupUpdate = original.GroupUpdate
type GroupsClient = original.GroupsClient
type GroupsCreateOrUpdateFuture = original.GroupsCreateOrUpdateFuture
type GroupsDeleteFuture = original.GroupsDeleteFuture
type GroupsUpdateFuture = original.GroupsUpdateFuture
type KeyVaultCredentialSettings = original.KeyVaultCredentialSettings
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type LoadBalancerConfiguration = original.LoadBalancerConfiguration
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PrivateIPAddress = original.PrivateIPAddress
type Properties = original.Properties
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceIdentity = original.ResourceIdentity
type SQLConnectivityUpdateSettings = original.SQLConnectivityUpdateSettings
type SQLStorageSettings = original.SQLStorageSettings
type SQLStorageUpdateSettings = original.SQLStorageUpdateSettings
type SQLVirtualMachine = original.SQLVirtualMachine
type SQLVirtualMachinesClient = original.SQLVirtualMachinesClient
type SQLVirtualMachinesCreateOrUpdateFutureType = original.SQLVirtualMachinesCreateOrUpdateFutureType
type SQLVirtualMachinesDeleteFutureType = original.SQLVirtualMachinesDeleteFutureType
type SQLVirtualMachinesUpdateFutureType = original.SQLVirtualMachinesUpdateFutureType
type SQLWorkloadTypeUpdateSettings = original.SQLWorkloadTypeUpdateSettings
type ServerConfigurationsManagementSettings = original.ServerConfigurationsManagementSettings
type StorageConfigurationSettings = original.StorageConfigurationSettings
type TrackedResource = original.TrackedResource
type Update = original.Update
type WsfcDomainCredentials = original.WsfcDomainCredentials
type WsfcDomainProfile = original.WsfcDomainProfile

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailabilityGroupListenerListResultIterator(page AvailabilityGroupListenerListResultPage) AvailabilityGroupListenerListResultIterator {
	return original.NewAvailabilityGroupListenerListResultIterator(page)
}
func NewAvailabilityGroupListenerListResultPage(getNextPage func(context.Context, AvailabilityGroupListenerListResult) (AvailabilityGroupListenerListResult, error)) AvailabilityGroupListenerListResultPage {
	return original.NewAvailabilityGroupListenerListResultPage(getNextPage)
}
func NewAvailabilityGroupListenersClient(subscriptionID string) AvailabilityGroupListenersClient {
	return original.NewAvailabilityGroupListenersClient(subscriptionID)
}
func NewAvailabilityGroupListenersClientWithBaseURI(baseURI string, subscriptionID string) AvailabilityGroupListenersClient {
	return original.NewAvailabilityGroupListenersClientWithBaseURI(baseURI, subscriptionID)
}
func NewGroupListResultIterator(page GroupListResultPage) GroupListResultIterator {
	return original.NewGroupListResultIterator(page)
}
func NewGroupListResultPage(getNextPage func(context.Context, GroupListResult) (GroupListResult, error)) GroupListResultPage {
	return original.NewGroupListResultPage(getNextPage)
}
func NewGroupsClient(subscriptionID string) GroupsClient {
	return original.NewGroupsClient(subscriptionID)
}
func NewGroupsClientWithBaseURI(baseURI string, subscriptionID string) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSQLVirtualMachinesClient(subscriptionID string) SQLVirtualMachinesClient {
	return original.NewSQLVirtualMachinesClient(subscriptionID)
}
func NewSQLVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string) SQLVirtualMachinesClient {
	return original.NewSQLVirtualMachinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleBackupScheduleTypeValues() []BackupScheduleType {
	return original.PossibleBackupScheduleTypeValues()
}
func PossibleClusterConfigurationValues() []ClusterConfiguration {
	return original.PossibleClusterConfigurationValues()
}
func PossibleClusterManagerTypeValues() []ClusterManagerType {
	return original.PossibleClusterManagerTypeValues()
}
func PossibleConnectivityTypeValues() []ConnectivityType {
	return original.PossibleConnectivityTypeValues()
}
func PossibleDayOfWeekValues() []DayOfWeek {
	return original.PossibleDayOfWeekValues()
}
func PossibleDiskConfigurationTypeValues() []DiskConfigurationType {
	return original.PossibleDiskConfigurationTypeValues()
}
func PossibleFullBackupFrequencyTypeValues() []FullBackupFrequencyType {
	return original.PossibleFullBackupFrequencyTypeValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossibleSQLImageSkuValues() []SQLImageSku {
	return original.PossibleSQLImageSkuValues()
}
func PossibleSQLManagementModeValues() []SQLManagementMode {
	return original.PossibleSQLManagementModeValues()
}
func PossibleSQLServerLicenseTypeValues() []SQLServerLicenseType {
	return original.PossibleSQLServerLicenseTypeValues()
}
func PossibleSQLVMGroupImageSkuValues() []SQLVMGroupImageSku {
	return original.PossibleSQLVMGroupImageSkuValues()
}
func PossibleSQLWorkloadTypeValues() []SQLWorkloadType {
	return original.PossibleSQLWorkloadTypeValues()
}
func PossibleScaleTypeValues() []ScaleType {
	return original.PossibleScaleTypeValues()
}
func PossibleStorageWorkloadTypeValues() []StorageWorkloadType {
	return original.PossibleStorageWorkloadTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
