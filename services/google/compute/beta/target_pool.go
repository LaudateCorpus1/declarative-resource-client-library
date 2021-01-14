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
package beta

import (
	"context"
	"fmt"
	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type TargetPool struct {
	BackupPool      *string                        `json:"backupPool"`
	Description     *string                        `json:"description"`
	FailoverRatio   *float64                       `json:"failoverRatio"`
	HealthCheck     []string                       `json:"healthCheck"`
	Instance        []string                       `json:"instance"`
	Name            *string                        `json:"name"`
	Region          *string                        `json:"region"`
	SelfLink        *string                        `json:"selfLink"`
	SessionAffinity *TargetPoolSessionAffinityEnum `json:"sessionAffinity"`
	Project         *string                        `json:"project"`
}

func (r *TargetPool) String() string {
	return dcl.SprintResource(r)
}

// The enum TargetPoolSessionAffinityEnum.
type TargetPoolSessionAffinityEnum string

// TargetPoolSessionAffinityEnumRef returns a *TargetPoolSessionAffinityEnum with the value of string s
// If the empty string is provided, nil is returned.
func TargetPoolSessionAffinityEnumRef(s string) *TargetPoolSessionAffinityEnum {
	if s == "" {
		return nil
	}

	v := TargetPoolSessionAffinityEnum(s)
	return &v
}

func (v TargetPoolSessionAffinityEnum) Validate() error {
	for _, s := range []string{"NONE", "CLIENT_IP", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "CLIENT_IP_PORT_PROTO", "HTTP_COOKIE", "HEADER_FIELD"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TargetPoolSessionAffinityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *TargetPool) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "TargetPool",
		Version: "beta",
	}
}

const TargetPoolMaxPage = -1

type TargetPoolList struct {
	Items []*TargetPool

	nextToken string

	pageSize int32

	project string

	region string
}

func (l *TargetPoolList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TargetPoolList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTargetPool(ctx, l.project, l.region, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTargetPool(ctx context.Context, project, region string) (*TargetPoolList, error) {

	return c.ListTargetPoolWithMaxResults(ctx, project, region, TargetPoolMaxPage)

}

func (c *Client) ListTargetPoolWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*TargetPoolList, error) {
	items, token, err := c.listTargetPool(ctx, project, region, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TargetPoolList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		region: region,
	}, nil
}

func (c *Client) GetTargetPool(ctx context.Context, r *TargetPool) (*TargetPool, error) {
	b, err := c.getTargetPoolRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTargetPool(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTargetPoolNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTargetPool(ctx context.Context, r *TargetPool) error {
	if r == nil {
		return fmt.Errorf("TargetPool resource is nil")
	}
	c.Config.Logger.Info("Deleting TargetPool...")
	deleteOp := deleteTargetPoolOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTargetPool deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTargetPool(ctx context.Context, project, region string, filter func(*TargetPool) bool) error {
	listObj, err := c.ListTargetPool(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllTargetPool(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTargetPool(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTargetPool(ctx context.Context, rawDesired *TargetPool, opts ...dcl.ApplyOption) (*TargetPool, error) {
	c.Config.Logger.Info("Beginning ApplyTargetPool...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.targetPoolDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []targetPoolApiOperation
	if create {
		ops = append(ops, &createTargetPoolOperation{})
	} else if recreate {
		ops = append(ops, &deleteTargetPoolOperation{})
		ops = append(ops, &createTargetPoolOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeTargetPoolDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetTargetPool(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTargetPoolNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTargetPoolDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTargetPool(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
