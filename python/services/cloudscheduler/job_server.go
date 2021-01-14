// Copyright 2021 Google LLC. All Rights Reserved.
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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	cloudschedulerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudscheduler/cloudscheduler_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler"
)

// Server implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobAppEngineHttpTargetHttpMethodEnum converts a JobAppEngineHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerJobAppEngineHttpTargetHttpMethodEnum(e cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum) *cloudscheduler.JobAppEngineHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := cloudscheduler.JobAppEngineHttpTargetHttpMethodEnum(n[len("JobAppEngineHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobHttpTargetHttpMethodEnum converts a JobHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerJobHttpTargetHttpMethodEnum(e cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum) *cloudscheduler.JobHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := cloudscheduler.JobHttpTargetHttpMethodEnum(n[len("JobHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStateEnum converts a JobStateEnum enum from its proto representation.
func ProtoToCloudschedulerJobStateEnum(e cloudschedulerpb.CloudschedulerJobStateEnum) *cloudscheduler.JobStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudschedulerpb.CloudschedulerJobStateEnum_name[int32(e)]; ok {
		e := cloudscheduler.JobStateEnum(n[len("JobStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobViewEnum converts a JobViewEnum enum from its proto representation.
func ProtoToCloudschedulerJobViewEnum(e cloudschedulerpb.CloudschedulerJobViewEnum) *cloudscheduler.JobViewEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudschedulerpb.CloudschedulerJobViewEnum_name[int32(e)]; ok {
		e := cloudscheduler.JobViewEnum(n[len("JobViewEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobLabels converts a JobLabels resource from its proto representation.
func ProtoToCloudschedulerJobLabels(p *cloudschedulerpb.CloudschedulerJobLabels) *cloudscheduler.JobLabels {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobLabels{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToJobPubsubTarget converts a JobPubsubTarget resource from its proto representation.
func ProtoToCloudschedulerJobPubsubTarget(p *cloudschedulerpb.CloudschedulerJobPubsubTarget) *cloudscheduler.JobPubsubTarget {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobPubsubTarget{
		TopicName: dcl.StringOrNil(p.TopicName),
		Data:      dcl.StringOrNil(p.Data),
	}
	for _, r := range p.GetAttributes() {
		obj.Attributes = append(obj.Attributes, *ProtoToCloudschedulerJobPubsubTargetAttributes(r))
	}
	return obj
}

// ProtoToJobPubsubTargetAttributes converts a JobPubsubTargetAttributes resource from its proto representation.
func ProtoToCloudschedulerJobPubsubTargetAttributes(p *cloudschedulerpb.CloudschedulerJobPubsubTargetAttributes) *cloudscheduler.JobPubsubTargetAttributes {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobPubsubTargetAttributes{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToJobAppEngineHttpTarget converts a JobAppEngineHttpTarget resource from its proto representation.
func ProtoToCloudschedulerJobAppEngineHttpTarget(p *cloudschedulerpb.CloudschedulerJobAppEngineHttpTarget) *cloudscheduler.JobAppEngineHttpTarget {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobAppEngineHttpTarget{
		HttpMethod:       ProtoToCloudschedulerJobAppEngineHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		AppEngineRouting: ProtoToCloudschedulerJobAppEngineHttpTargetAppEngineRouting(p.GetAppEngineRouting()),
		RelativeUri:      dcl.StringOrNil(p.RelativeUri),
		Body:             dcl.StringOrNil(p.Body),
	}
	for _, r := range p.GetHeaders() {
		obj.Headers = append(obj.Headers, *ProtoToCloudschedulerJobAppEngineHttpTargetHeaders(r))
	}
	return obj
}

// ProtoToJobAppEngineHttpTargetAppEngineRouting converts a JobAppEngineHttpTargetAppEngineRouting resource from its proto representation.
func ProtoToCloudschedulerJobAppEngineHttpTargetAppEngineRouting(p *cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetAppEngineRouting) *cloudscheduler.JobAppEngineHttpTargetAppEngineRouting {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.StringOrNil(p.Service),
		Version:  dcl.StringOrNil(p.Version),
		Instance: dcl.StringOrNil(p.Instance),
		Host:     dcl.StringOrNil(p.Host),
	}
	return obj
}

// ProtoToJobAppEngineHttpTargetHeaders converts a JobAppEngineHttpTargetHeaders resource from its proto representation.
func ProtoToCloudschedulerJobAppEngineHttpTargetHeaders(p *cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHeaders) *cloudscheduler.JobAppEngineHttpTargetHeaders {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobAppEngineHttpTargetHeaders{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToJobHttpTarget converts a JobHttpTarget resource from its proto representation.
func ProtoToCloudschedulerJobHttpTarget(p *cloudschedulerpb.CloudschedulerJobHttpTarget) *cloudscheduler.JobHttpTarget {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobHttpTarget{
		Uri:        dcl.StringOrNil(p.Uri),
		HttpMethod: ProtoToCloudschedulerJobHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		Body:       dcl.StringOrNil(p.Body),
		OAuthToken: ProtoToCloudschedulerJobHttpTargetOAuthToken(p.GetOauthToken()),
		OidcToken:  ProtoToCloudschedulerJobHttpTargetOidcToken(p.GetOidcToken()),
	}
	for _, r := range p.GetHeaders() {
		obj.Headers = append(obj.Headers, *ProtoToCloudschedulerJobHttpTargetHeaders(r))
	}
	return obj
}

// ProtoToJobHttpTargetHeaders converts a JobHttpTargetHeaders resource from its proto representation.
func ProtoToCloudschedulerJobHttpTargetHeaders(p *cloudschedulerpb.CloudschedulerJobHttpTargetHeaders) *cloudscheduler.JobHttpTargetHeaders {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobHttpTargetHeaders{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToJobHttpTargetOAuthToken converts a JobHttpTargetOAuthToken resource from its proto representation.
func ProtoToCloudschedulerJobHttpTargetOAuthToken(p *cloudschedulerpb.CloudschedulerJobHttpTargetOAuthToken) *cloudscheduler.JobHttpTargetOAuthToken {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
		Scope:               dcl.StringOrNil(p.Scope),
	}
	return obj
}

// ProtoToJobHttpTargetOidcToken converts a JobHttpTargetOidcToken resource from its proto representation.
func ProtoToCloudschedulerJobHttpTargetOidcToken(p *cloudschedulerpb.CloudschedulerJobHttpTargetOidcToken) *cloudscheduler.JobHttpTargetOidcToken {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
		Audience:            dcl.StringOrNil(p.Audience),
	}
	return obj
}

// ProtoToJobStatus converts a JobStatus resource from its proto representation.
func ProtoToCloudschedulerJobStatus(p *cloudschedulerpb.CloudschedulerJobStatus) *cloudscheduler.JobStatus {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToCloudschedulerJobStatusDetails(r))
	}
	return obj
}

// ProtoToJobStatusDetails converts a JobStatusDetails resource from its proto representation.
func ProtoToCloudschedulerJobStatusDetails(p *cloudschedulerpb.CloudschedulerJobStatusDetails) *cloudscheduler.JobStatusDetails {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToJobRetryConfig converts a JobRetryConfig resource from its proto representation.
func ProtoToCloudschedulerJobRetryConfig(p *cloudschedulerpb.CloudschedulerJobRetryConfig) *cloudscheduler.JobRetryConfig {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobRetryConfig{
		RetryCount:         dcl.Int64OrNil(p.RetryCount),
		MaxRetryDuration:   ProtoToCloudschedulerJobRetryConfigMaxRetryDuration(p.GetMaxRetryDuration()),
		MinBackoffDuration: ProtoToCloudschedulerJobRetryConfigMinBackoffDuration(p.GetMinBackoffDuration()),
		MaxBackoffDuration: ProtoToCloudschedulerJobRetryConfigMaxBackoffDuration(p.GetMaxBackoffDuration()),
		MaxDoublings:       dcl.Int64OrNil(p.MaxDoublings),
	}
	return obj
}

// ProtoToJobRetryConfigMaxRetryDuration converts a JobRetryConfigMaxRetryDuration resource from its proto representation.
func ProtoToCloudschedulerJobRetryConfigMaxRetryDuration(p *cloudschedulerpb.CloudschedulerJobRetryConfigMaxRetryDuration) *cloudscheduler.JobRetryConfigMaxRetryDuration {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobRetryConfigMaxRetryDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToJobRetryConfigMinBackoffDuration converts a JobRetryConfigMinBackoffDuration resource from its proto representation.
func ProtoToCloudschedulerJobRetryConfigMinBackoffDuration(p *cloudschedulerpb.CloudschedulerJobRetryConfigMinBackoffDuration) *cloudscheduler.JobRetryConfigMinBackoffDuration {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobRetryConfigMinBackoffDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToJobRetryConfigMaxBackoffDuration converts a JobRetryConfigMaxBackoffDuration resource from its proto representation.
func ProtoToCloudschedulerJobRetryConfigMaxBackoffDuration(p *cloudschedulerpb.CloudschedulerJobRetryConfigMaxBackoffDuration) *cloudscheduler.JobRetryConfigMaxBackoffDuration {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobRetryConfigMaxBackoffDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToJobAttemptDeadline converts a JobAttemptDeadline resource from its proto representation.
func ProtoToCloudschedulerJobAttemptDeadline(p *cloudschedulerpb.CloudschedulerJobAttemptDeadline) *cloudscheduler.JobAttemptDeadline {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobAttemptDeadline{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *cloudschedulerpb.CloudschedulerJob) *cloudscheduler.Job {
	obj := &cloudscheduler.Job{
		Name:                 dcl.StringOrNil(p.Name),
		Description:          dcl.StringOrNil(p.Description),
		PubsubTarget:         ProtoToCloudschedulerJobPubsubTarget(p.GetPubsubTarget()),
		AppEngineHttpTarget:  ProtoToCloudschedulerJobAppEngineHttpTarget(p.GetAppEngineHttpTarget()),
		HttpTarget:           ProtoToCloudschedulerJobHttpTarget(p.GetHttpTarget()),
		Schedule:             dcl.StringOrNil(p.Schedule),
		TimeZone:             dcl.StringOrNil(p.TimeZone),
		UserUpdateTime:       dcl.StringOrNil(p.GetUserUpdateTime()),
		State:                ProtoToCloudschedulerJobStateEnum(p.GetState()),
		Status:               ProtoToCloudschedulerJobStatus(p.GetStatus()),
		TotalAttemptCount:    dcl.Int64OrNil(p.TotalAttemptCount),
		FailedAttemptCount:   dcl.Int64OrNil(p.FailedAttemptCount),
		TotalExecutionCount:  dcl.Int64OrNil(p.TotalExecutionCount),
		FailedExecutionCount: dcl.Int64OrNil(p.FailedExecutionCount),
		View:                 ProtoToCloudschedulerJobViewEnum(p.GetView()),
		ScheduleTime:         dcl.StringOrNil(p.GetScheduleTime()),
		LastAttemptTime:      dcl.StringOrNil(p.GetLastAttemptTime()),
		RetryConfig:          ProtoToCloudschedulerJobRetryConfig(p.GetRetryConfig()),
		AttemptDeadline:      ProtoToCloudschedulerJobAttemptDeadline(p.GetAttemptDeadline()),
		Project:              dcl.StringOrNil(p.Project),
		Location:             dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToCloudschedulerJobLabels(r))
	}
	return obj
}

// JobAppEngineHttpTargetHttpMethodEnumToProto converts a JobAppEngineHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerJobAppEngineHttpTargetHttpMethodEnumToProto(e *cloudscheduler.JobAppEngineHttpTargetHttpMethodEnum) cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum {
	if e == nil {
		return cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum(0)
	}
	if v, ok := cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum_value["JobAppEngineHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum(v)
	}
	return cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum(0)
}

// JobHttpTargetHttpMethodEnumToProto converts a JobHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerJobHttpTargetHttpMethodEnumToProto(e *cloudscheduler.JobHttpTargetHttpMethodEnum) cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum {
	if e == nil {
		return cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum(0)
	}
	if v, ok := cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum_value["JobHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum(v)
	}
	return cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum(0)
}

// JobStateEnumToProto converts a JobStateEnum enum to its proto representation.
func CloudschedulerJobStateEnumToProto(e *cloudscheduler.JobStateEnum) cloudschedulerpb.CloudschedulerJobStateEnum {
	if e == nil {
		return cloudschedulerpb.CloudschedulerJobStateEnum(0)
	}
	if v, ok := cloudschedulerpb.CloudschedulerJobStateEnum_value["JobStateEnum"+string(*e)]; ok {
		return cloudschedulerpb.CloudschedulerJobStateEnum(v)
	}
	return cloudschedulerpb.CloudschedulerJobStateEnum(0)
}

// JobViewEnumToProto converts a JobViewEnum enum to its proto representation.
func CloudschedulerJobViewEnumToProto(e *cloudscheduler.JobViewEnum) cloudschedulerpb.CloudschedulerJobViewEnum {
	if e == nil {
		return cloudschedulerpb.CloudschedulerJobViewEnum(0)
	}
	if v, ok := cloudschedulerpb.CloudschedulerJobViewEnum_value["JobViewEnum"+string(*e)]; ok {
		return cloudschedulerpb.CloudschedulerJobViewEnum(v)
	}
	return cloudschedulerpb.CloudschedulerJobViewEnum(0)
}

// JobLabelsToProto converts a JobLabels resource to its proto representation.
func CloudschedulerJobLabelsToProto(o *cloudscheduler.JobLabels) *cloudschedulerpb.CloudschedulerJobLabels {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobLabels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// JobPubsubTargetToProto converts a JobPubsubTarget resource to its proto representation.
func CloudschedulerJobPubsubTargetToProto(o *cloudscheduler.JobPubsubTarget) *cloudschedulerpb.CloudschedulerJobPubsubTarget {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobPubsubTarget{
		TopicName: dcl.ValueOrEmptyString(o.TopicName),
		Data:      dcl.ValueOrEmptyString(o.Data),
	}
	for _, r := range o.Attributes {
		p.Attributes = append(p.Attributes, CloudschedulerJobPubsubTargetAttributesToProto(&r))
	}
	return p
}

// JobPubsubTargetAttributesToProto converts a JobPubsubTargetAttributes resource to its proto representation.
func CloudschedulerJobPubsubTargetAttributesToProto(o *cloudscheduler.JobPubsubTargetAttributes) *cloudschedulerpb.CloudschedulerJobPubsubTargetAttributes {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobPubsubTargetAttributes{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// JobAppEngineHttpTargetToProto converts a JobAppEngineHttpTarget resource to its proto representation.
func CloudschedulerJobAppEngineHttpTargetToProto(o *cloudscheduler.JobAppEngineHttpTarget) *cloudschedulerpb.CloudschedulerJobAppEngineHttpTarget {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobAppEngineHttpTarget{
		HttpMethod:       CloudschedulerJobAppEngineHttpTargetHttpMethodEnumToProto(o.HttpMethod),
		AppEngineRouting: CloudschedulerJobAppEngineHttpTargetAppEngineRoutingToProto(o.AppEngineRouting),
		RelativeUri:      dcl.ValueOrEmptyString(o.RelativeUri),
		Body:             dcl.ValueOrEmptyString(o.Body),
	}
	for _, r := range o.Headers {
		p.Headers = append(p.Headers, CloudschedulerJobAppEngineHttpTargetHeadersToProto(&r))
	}
	return p
}

// JobAppEngineHttpTargetAppEngineRoutingToProto converts a JobAppEngineHttpTargetAppEngineRouting resource to its proto representation.
func CloudschedulerJobAppEngineHttpTargetAppEngineRoutingToProto(o *cloudscheduler.JobAppEngineHttpTargetAppEngineRouting) *cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetAppEngineRouting {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.ValueOrEmptyString(o.Service),
		Version:  dcl.ValueOrEmptyString(o.Version),
		Instance: dcl.ValueOrEmptyString(o.Instance),
		Host:     dcl.ValueOrEmptyString(o.Host),
	}
	return p
}

// JobAppEngineHttpTargetHeadersToProto converts a JobAppEngineHttpTargetHeaders resource to its proto representation.
func CloudschedulerJobAppEngineHttpTargetHeadersToProto(o *cloudscheduler.JobAppEngineHttpTargetHeaders) *cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHeaders {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHeaders{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// JobHttpTargetToProto converts a JobHttpTarget resource to its proto representation.
func CloudschedulerJobHttpTargetToProto(o *cloudscheduler.JobHttpTarget) *cloudschedulerpb.CloudschedulerJobHttpTarget {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobHttpTarget{
		Uri:        dcl.ValueOrEmptyString(o.Uri),
		HttpMethod: CloudschedulerJobHttpTargetHttpMethodEnumToProto(o.HttpMethod),
		Body:       dcl.ValueOrEmptyString(o.Body),
		OauthToken: CloudschedulerJobHttpTargetOAuthTokenToProto(o.OAuthToken),
		OidcToken:  CloudschedulerJobHttpTargetOidcTokenToProto(o.OidcToken),
	}
	for _, r := range o.Headers {
		p.Headers = append(p.Headers, CloudschedulerJobHttpTargetHeadersToProto(&r))
	}
	return p
}

// JobHttpTargetHeadersToProto converts a JobHttpTargetHeaders resource to its proto representation.
func CloudschedulerJobHttpTargetHeadersToProto(o *cloudscheduler.JobHttpTargetHeaders) *cloudschedulerpb.CloudschedulerJobHttpTargetHeaders {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobHttpTargetHeaders{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// JobHttpTargetOAuthTokenToProto converts a JobHttpTargetOAuthToken resource to its proto representation.
func CloudschedulerJobHttpTargetOAuthTokenToProto(o *cloudscheduler.JobHttpTargetOAuthToken) *cloudschedulerpb.CloudschedulerJobHttpTargetOAuthToken {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.ValueOrEmptyString(o.ServiceAccountEmail),
		Scope:               dcl.ValueOrEmptyString(o.Scope),
	}
	return p
}

// JobHttpTargetOidcTokenToProto converts a JobHttpTargetOidcToken resource to its proto representation.
func CloudschedulerJobHttpTargetOidcTokenToProto(o *cloudscheduler.JobHttpTargetOidcToken) *cloudschedulerpb.CloudschedulerJobHttpTargetOidcToken {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.ValueOrEmptyString(o.ServiceAccountEmail),
		Audience:            dcl.ValueOrEmptyString(o.Audience),
	}
	return p
}

// JobStatusToProto converts a JobStatus resource to its proto representation.
func CloudschedulerJobStatusToProto(o *cloudscheduler.JobStatus) *cloudschedulerpb.CloudschedulerJobStatus {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, CloudschedulerJobStatusDetailsToProto(&r))
	}
	return p
}

// JobStatusDetailsToProto converts a JobStatusDetails resource to its proto representation.
func CloudschedulerJobStatusDetailsToProto(o *cloudscheduler.JobStatusDetails) *cloudschedulerpb.CloudschedulerJobStatusDetails {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// JobRetryConfigToProto converts a JobRetryConfig resource to its proto representation.
func CloudschedulerJobRetryConfigToProto(o *cloudscheduler.JobRetryConfig) *cloudschedulerpb.CloudschedulerJobRetryConfig {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobRetryConfig{
		RetryCount:         dcl.ValueOrEmptyInt64(o.RetryCount),
		MaxRetryDuration:   CloudschedulerJobRetryConfigMaxRetryDurationToProto(o.MaxRetryDuration),
		MinBackoffDuration: CloudschedulerJobRetryConfigMinBackoffDurationToProto(o.MinBackoffDuration),
		MaxBackoffDuration: CloudschedulerJobRetryConfigMaxBackoffDurationToProto(o.MaxBackoffDuration),
		MaxDoublings:       dcl.ValueOrEmptyInt64(o.MaxDoublings),
	}
	return p
}

// JobRetryConfigMaxRetryDurationToProto converts a JobRetryConfigMaxRetryDuration resource to its proto representation.
func CloudschedulerJobRetryConfigMaxRetryDurationToProto(o *cloudscheduler.JobRetryConfigMaxRetryDuration) *cloudschedulerpb.CloudschedulerJobRetryConfigMaxRetryDuration {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobRetryConfigMaxRetryDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// JobRetryConfigMinBackoffDurationToProto converts a JobRetryConfigMinBackoffDuration resource to its proto representation.
func CloudschedulerJobRetryConfigMinBackoffDurationToProto(o *cloudscheduler.JobRetryConfigMinBackoffDuration) *cloudschedulerpb.CloudschedulerJobRetryConfigMinBackoffDuration {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobRetryConfigMinBackoffDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// JobRetryConfigMaxBackoffDurationToProto converts a JobRetryConfigMaxBackoffDuration resource to its proto representation.
func CloudschedulerJobRetryConfigMaxBackoffDurationToProto(o *cloudscheduler.JobRetryConfigMaxBackoffDuration) *cloudschedulerpb.CloudschedulerJobRetryConfigMaxBackoffDuration {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobRetryConfigMaxBackoffDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// JobAttemptDeadlineToProto converts a JobAttemptDeadline resource to its proto representation.
func CloudschedulerJobAttemptDeadlineToProto(o *cloudscheduler.JobAttemptDeadline) *cloudschedulerpb.CloudschedulerJobAttemptDeadline {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobAttemptDeadline{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *cloudscheduler.Job) *cloudschedulerpb.CloudschedulerJob {
	p := &cloudschedulerpb.CloudschedulerJob{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		PubsubTarget:         CloudschedulerJobPubsubTargetToProto(resource.PubsubTarget),
		AppEngineHttpTarget:  CloudschedulerJobAppEngineHttpTargetToProto(resource.AppEngineHttpTarget),
		HttpTarget:           CloudschedulerJobHttpTargetToProto(resource.HttpTarget),
		Schedule:             dcl.ValueOrEmptyString(resource.Schedule),
		TimeZone:             dcl.ValueOrEmptyString(resource.TimeZone),
		UserUpdateTime:       dcl.ValueOrEmptyString(resource.UserUpdateTime),
		State:                CloudschedulerJobStateEnumToProto(resource.State),
		Status:               CloudschedulerJobStatusToProto(resource.Status),
		TotalAttemptCount:    dcl.ValueOrEmptyInt64(resource.TotalAttemptCount),
		FailedAttemptCount:   dcl.ValueOrEmptyInt64(resource.FailedAttemptCount),
		TotalExecutionCount:  dcl.ValueOrEmptyInt64(resource.TotalExecutionCount),
		FailedExecutionCount: dcl.ValueOrEmptyInt64(resource.FailedExecutionCount),
		View:                 CloudschedulerJobViewEnumToProto(resource.View),
		ScheduleTime:         dcl.ValueOrEmptyString(resource.ScheduleTime),
		LastAttemptTime:      dcl.ValueOrEmptyString(resource.LastAttemptTime),
		RetryConfig:          CloudschedulerJobRetryConfigToProto(resource.RetryConfig),
		AttemptDeadline:      CloudschedulerJobAttemptDeadlineToProto(resource.AttemptDeadline),
		Project:              dcl.ValueOrEmptyString(resource.Project),
		Location:             dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Labels {
		p.Labels = append(p.Labels, CloudschedulerJobLabelsToProto(&r))
	}

	return p
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *cloudscheduler.Client, request *cloudschedulerpb.ApplyCloudschedulerJobRequest) (*cloudschedulerpb.CloudschedulerJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyCloudschedulerJob(ctx context.Context, request *cloudschedulerpb.ApplyCloudschedulerJobRequest) (*cloudschedulerpb.CloudschedulerJob, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteCloudschedulerJob(ctx context.Context, request *cloudschedulerpb.DeleteCloudschedulerJobRequest) (*emptypb.Empty, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))
}

// ListJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListCloudschedulerJob(ctx context.Context, request *cloudschedulerpb.ListCloudschedulerJobRequest) (*cloudschedulerpb.ListCloudschedulerJobResponse, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*cloudschedulerpb.CloudschedulerJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	return &cloudschedulerpb.ListCloudschedulerJobResponse{Items: protos}, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*cloudscheduler.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudscheduler.NewClient(conf), nil
}
