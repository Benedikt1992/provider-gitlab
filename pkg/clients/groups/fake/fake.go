/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	"github.com/xanzy/go-gitlab"

	"github.com/crossplane-contrib/provider-gitlab/pkg/clients/groups"
)

var _ groups.Client = &MockClient{}

// MockClient is a fake implementation of groups.Client.
type MockClient struct {
	groups.Client

	MockGetGroup    func(pid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error)
	MockCreateGroup func(opt *gitlab.CreateGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error)
	MockUpdateGroup func(pid interface{}, opt *gitlab.UpdateGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error)
	MockDeleteGroup func(pid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error)

	MockGetMember    func(gid interface{}, user int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)
	MockAddMember    func(gid interface{}, opt *gitlab.AddGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)
	MockEditMember   func(gid interface{}, user int, opt *gitlab.EditGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)
	MockRemoveMember func(gid interface{}, user int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error)

	MockListDeployTokens  func(gid interface{}, opt *gitlab.ListGroupDeployTokensOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.DeployToken, *gitlab.Response, error)
	MockCreateDeployToken func(gid interface{}, opt *gitlab.CreateGroupDeployTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.DeployToken, *gitlab.Response, error)
	MockDeleteDeployToken func(gid interface{}, deployToken int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error)
}

// GetGroup calls the underlying MockGetGroup method.
func (c *MockClient) GetGroup(pid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error) {
	return c.MockGetGroup(pid)
}

// CreateGroup calls the underlying MockCreateGroup method
func (c *MockClient) CreateGroup(opt *gitlab.CreateGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error) {
	return c.MockCreateGroup(opt)
}

// UpdateGroup calls the underlying MockUpdateGroup method
func (c *MockClient) UpdateGroup(pid interface{}, opt *gitlab.UpdateGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error) {
	return c.MockUpdateGroup(pid, opt)
}

// DeleteGroup calls the underlying MockDeleteGroup method
func (c *MockClient) DeleteGroup(pid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	return c.MockDeleteGroup(pid)
}

// GetGroupMember calls the underlying MockGetMember method.
func (c *MockClient) GetGroupMember(gid interface{}, user int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	return c.MockGetMember(gid, user)
}

// AddGroupMember calls the underlying MockAddMember method.
func (c *MockClient) AddGroupMember(gid interface{}, opt *gitlab.AddGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	return c.MockAddMember(gid, opt)
}

// EditGroupMember calls the underlying MockEditMember method.
func (c *MockClient) EditGroupMember(gid interface{}, user int, opt *gitlab.EditGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	return c.MockEditMember(gid, user, opt)
}

// RemoveGroupMember calls the underlying MockRemoveMember method.
func (c *MockClient) RemoveGroupMember(gid interface{}, user int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	return c.MockRemoveMember(gid, user)
}

// ListGroupDeployTokens calls the underlying MockListGroupDeployTokens method.
func (c *MockClient) ListGroupDeployTokens(gid interface{}, opt *gitlab.ListGroupDeployTokensOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.DeployToken, *gitlab.Response, error) {
	return c.MockListDeployTokens(gid, opt)
}

// CreateGroupDeployToken calls the underlying MockCreateGroupDeployToken method.
func (c *MockClient) CreateGroupDeployToken(gid interface{}, opt *gitlab.CreateGroupDeployTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.DeployToken, *gitlab.Response, error) {
	return c.MockCreateDeployToken(gid, opt)
}

// DeleteGroupDeployToken calls the underlying MockDeleteGroupDeployToken method.
func (c *MockClient) DeleteGroupDeployToken(gid interface{}, deployToken int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	return c.MockDeleteDeployToken(gid, deployToken)
}
