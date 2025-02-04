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
package pubsub

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Topic struct {
	Name                 *string                    `json:"name"`
	KmsKeyName           *string                    `json:"kmsKeyName"`
	Labels               map[string]string          `json:"labels"`
	MessageStoragePolicy *TopicMessageStoragePolicy `json:"messageStoragePolicy"`
	Project              *string                    `json:"project"`
}

func (r *Topic) String() string {
	return dcl.SprintResource(r)
}

type TopicMessageStoragePolicy struct {
	empty                     bool     `json:"-"`
	AllowedPersistenceRegions []string `json:"allowedPersistenceRegions"`
}

type jsonTopicMessageStoragePolicy TopicMessageStoragePolicy

func (r *TopicMessageStoragePolicy) UnmarshalJSON(data []byte) error {
	var res jsonTopicMessageStoragePolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTopicMessageStoragePolicy
	} else {

		r.AllowedPersistenceRegions = res.AllowedPersistenceRegions

	}
	return nil
}

// This object is used to assert a desired state where this TopicMessageStoragePolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTopicMessageStoragePolicy *TopicMessageStoragePolicy = &TopicMessageStoragePolicy{empty: true}

func (r *TopicMessageStoragePolicy) Empty() bool {
	return r.empty
}

func (r *TopicMessageStoragePolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *TopicMessageStoragePolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Topic) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "pubsub",
		Type:    "Topic",
		Version: "pubsub",
	}
}

func (r *Topic) ID() (string, error) {
	if err := extractTopicFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                 dcl.ValueOrEmptyString(nr.Name),
		"kmsKeyName":           dcl.ValueOrEmptyString(nr.KmsKeyName),
		"labels":               dcl.ValueOrEmptyString(nr.Labels),
		"messageStoragePolicy": dcl.ValueOrEmptyString(nr.MessageStoragePolicy),
		"project":              dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/topics/{{name}}", params), nil
}

const TopicMaxPage = -1

type TopicList struct {
	Items []*Topic

	nextToken string

	pageSize int32

	resource *Topic
}

func (l *TopicList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TopicList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTopic(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTopic(ctx context.Context, r *Topic) (*TopicList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTopicWithMaxResults(ctx, r, TopicMaxPage)

}

func (c *Client) ListTopicWithMaxResults(ctx context.Context, r *Topic, pageSize int32) (*TopicList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listTopic(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TopicList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetTopic(ctx context.Context, r *Topic) (*Topic, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getTopicRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTopic(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTopicNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTopic(ctx context.Context, r *Topic) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Topic resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Topic...")
	deleteOp := deleteTopicOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTopic deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTopic(ctx context.Context, project string, filter func(*Topic) bool) error {
	r := &Topic{
		Project: &project,
	}
	listObj, err := c.ListTopic(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllTopic(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTopic(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTopic(ctx context.Context, rawDesired *Topic, opts ...dcl.ApplyOption) (*Topic, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Topic
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTopicHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyTopicHelper(c *Client, ctx context.Context, rawDesired *Topic, opts ...dcl.ApplyOption) (*Topic, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyTopic...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractTopicFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.topicDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTopicDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
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
	var ops []topicApiOperation
	if create {
		ops = append(ops, &createTopicOperation{})
	} else if recreate {
		ops = append(ops, &deleteTopicOperation{})
		ops = append(ops, &createTopicOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeTopicDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetTopic(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTopicOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTopic(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTopicNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTopicNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTopicDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTopic(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
