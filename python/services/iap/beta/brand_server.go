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
	"errors"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iap/beta/iap_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/beta"
)

// Server implements the gRPC interface for Brand.
type BrandServer struct{}

// ProtoToBrand converts a Brand resource from its proto representation.
func ProtoToBrand(p *betapb.IapBetaBrand) *beta.Brand {
	obj := &beta.Brand{
		ApplicationTitle: dcl.StringOrNil(p.ApplicationTitle),
		Name:             dcl.StringOrNil(p.Name),
		OrgInternalOnly:  dcl.Bool(p.OrgInternalOnly),
		SupportEmail:     dcl.StringOrNil(p.SupportEmail),
		Project:          dcl.StringOrNil(p.Project),
	}
	return obj
}

// BrandToProto converts a Brand resource to its proto representation.
func BrandToProto(resource *beta.Brand) *betapb.IapBetaBrand {
	p := &betapb.IapBetaBrand{
		ApplicationTitle: dcl.ValueOrEmptyString(resource.ApplicationTitle),
		Name:             dcl.ValueOrEmptyString(resource.Name),
		OrgInternalOnly:  dcl.ValueOrEmptyBool(resource.OrgInternalOnly),
		SupportEmail:     dcl.ValueOrEmptyString(resource.SupportEmail),
		Project:          dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyBrand handles the gRPC request by passing it to the underlying Brand Apply() method.
func (s *BrandServer) applyBrand(ctx context.Context, c *beta.Client, request *betapb.ApplyIapBetaBrandRequest) (*betapb.IapBetaBrand, error) {
	p := ProtoToBrand(request.GetResource())
	res, err := c.ApplyBrand(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BrandToProto(res)
	return r, nil
}

// ApplyBrand handles the gRPC request by passing it to the underlying Brand Apply() method.
func (s *BrandServer) ApplyIapBetaBrand(ctx context.Context, request *betapb.ApplyIapBetaBrandRequest) (*betapb.IapBetaBrand, error) {
	cl, err := createConfigBrand(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBrand(ctx, cl, request)
}

// DeleteBrand handles the gRPC request by passing it to the underlying Brand Delete() method.
func (s *BrandServer) DeleteIapBetaBrand(ctx context.Context, request *betapb.DeleteIapBetaBrandRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Brand")

}

// ListIapBetaBrand handles the gRPC request by passing it to the underlying BrandList() method.
func (s *BrandServer) ListIapBetaBrand(ctx context.Context, request *betapb.ListIapBetaBrandRequest) (*betapb.ListIapBetaBrandResponse, error) {
	cl, err := createConfigBrand(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBrand(ctx, ProtoToBrand(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IapBetaBrand
	for _, r := range resources.Items {
		rp := BrandToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListIapBetaBrandResponse{Items: protos}, nil
}

func createConfigBrand(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
