// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import context "context"
import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"

// Resolver is an autogenerated failing mock type for the Resolver type
type Resolver struct {
	err error
}

// NewResolver creates a new Resolver type instance
func NewResolver(err error) *Resolver {
	return &Resolver{err: err}
}

// ContentQuery provides a failing mock function with given fields: ctx, contentType, id
func (_m *Resolver) ContentQuery(ctx context.Context, contentType string, id string) (*gqlschema.JSON, error) {
	var r0 *gqlschema.JSON
	var r1 error
	r1 = _m.err

	return r0, r1
}

// TopicsQuery provides a failing mock function with given fields: ctx, topics, internal
func (_m *Resolver) TopicsQuery(ctx context.Context, topics []gqlschema.InputTopic, internal *bool) ([]gqlschema.TopicEntry, error) {
	var r0 []gqlschema.TopicEntry
	var r1 error
	r1 = _m.err

	return r0, r1
}