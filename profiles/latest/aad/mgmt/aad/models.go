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

package aad

import original "github.com/Azure/azure-sdk-for-go/services/aad/mgmt/2018-01-01/aad"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AggregationType = original.AggregationType

const (
	Average AggregationType = original.Average
	Count   AggregationType = original.Count
	Maximum AggregationType = original.Maximum
	Minimum AggregationType = original.Minimum
	None    AggregationType = original.None
	Total   AggregationType = original.Total
)

type Category = original.Category

const (
	AuditLogs  Category = original.AuditLogs
	SignInLogs Category = original.SignInLogs
)

type CategoryType = original.CategoryType

const (
	Logs CategoryType = original.Logs
)

type ResultType = original.ResultType

const (
	Data     ResultType = original.Data
	Metadata ResultType = original.Metadata
)

type Unit = original.Unit

const (
	UnitBitsPerSecond  Unit = original.UnitBitsPerSecond
	UnitBytes          Unit = original.UnitBytes
	UnitByteSeconds    Unit = original.UnitByteSeconds
	UnitBytesPerSecond Unit = original.UnitBytesPerSecond
	UnitCores          Unit = original.UnitCores
	UnitCount          Unit = original.UnitCount
	UnitCountPerSecond Unit = original.UnitCountPerSecond
	UnitMilliCores     Unit = original.UnitMilliCores
	UnitMilliSeconds   Unit = original.UnitMilliSeconds
	UnitNanoCores      Unit = original.UnitNanoCores
	UnitPercent        Unit = original.UnitPercent
	UnitSeconds        Unit = original.UnitSeconds
	UnitUnspecified    Unit = original.UnitUnspecified
)

type BaseClient = original.BaseClient
type DiagnosticSettings = original.DiagnosticSettings
type DiagnosticSettingsCategory = original.DiagnosticSettingsCategory
type DiagnosticSettingsCategoryClient = original.DiagnosticSettingsCategoryClient
type DiagnosticSettingsCategoryResource = original.DiagnosticSettingsCategoryResource
type DiagnosticSettingsCategoryResourceCollection = original.DiagnosticSettingsCategoryResourceCollection
type DiagnosticSettingsClient = original.DiagnosticSettingsClient
type DiagnosticSettingsResource = original.DiagnosticSettingsResource
type DiagnosticSettingsResourceCollection = original.DiagnosticSettingsResourceCollection
type Display = original.Display
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type LocalizableString = original.LocalizableString
type LogSettings = original.LogSettings
type MetadataValue = original.MetadataValue
type Metric = original.Metric
type MetricAvailability = original.MetricAvailability
type MetricDefinition = original.MetricDefinition
type MetricDefinitionCollection = original.MetricDefinitionCollection
type MetricDefinitionsClient = original.MetricDefinitionsClient
type MetricErrorResponse = original.MetricErrorResponse
type MetricValue = original.MetricValue
type MetricsClient = original.MetricsClient
type OperationsClient = original.OperationsClient
type OperationsDiscovery = original.OperationsDiscovery
type OperationsDiscoveryCollection = original.OperationsDiscoveryCollection
type ProxyOnlyResource = original.ProxyOnlyResource
type Response = original.Response
type RetentionPolicy = original.RetentionPolicy
type TimeSeriesElement = original.TimeSeriesElement

func New() BaseClient {
	return original.New()
}
func NewDiagnosticSettingsCategoryClient() DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClient()
}
func NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI string) DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI)
}
func NewDiagnosticSettingsClient() DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClient()
}
func NewDiagnosticSettingsClientWithBaseURI(baseURI string) DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClientWithBaseURI(baseURI)
}
func NewMetricDefinitionsClient() MetricDefinitionsClient {
	return original.NewMetricDefinitionsClient()
}
func NewMetricDefinitionsClientWithBaseURI(baseURI string) MetricDefinitionsClient {
	return original.NewMetricDefinitionsClientWithBaseURI(baseURI)
}
func NewMetricsClient() MetricsClient {
	return original.NewMetricsClient()
}
func NewMetricsClientWithBaseURI(baseURI string) MetricsClient {
	return original.NewMetricsClientWithBaseURI(baseURI)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleAggregationTypeValues() []AggregationType {
	return original.PossibleAggregationTypeValues()
}
func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}
func PossibleCategoryValues() []Category {
	return original.PossibleCategoryValues()
}
func PossibleResultTypeValues() []ResultType {
	return original.PossibleResultTypeValues()
}
func PossibleUnitValues() []Unit {
	return original.PossibleUnitValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
