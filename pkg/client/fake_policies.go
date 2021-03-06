package client

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/watch"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

type FakePolicies struct {
	Fake *Fake
}

func (c *FakePolicies) List(label labels.Selector, field fields.Selector) (*authorizationapi.PolicyList, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "list-policies"}, &authorizationapi.PolicyList{})
	return obj.(*authorizationapi.PolicyList), err
}

func (c *FakePolicies) Get(name string) (*authorizationapi.Policy, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "get-policy"}, &authorizationapi.Policy{})
	return obj.(*authorizationapi.Policy), err
}

func (c *FakePolicies) Delete(name string) error {
	c.Fake.Actions = append(c.Fake.Actions, FakeAction{Action: "delete-policy", Value: name})
	return nil
}

func (c *FakePolicies) Watch(label labels.Selector, field fields.Selector, resourceVersion string) (watch.Interface, error) {
	c.Fake.Actions = append(c.Fake.Actions, FakeAction{Action: "watch-policy"})
	return nil, nil
}
