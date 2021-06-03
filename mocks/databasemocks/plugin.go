// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package databasemocks

import (
	context "context"

	config "github.com/kaleido-io/firefly/internal/config"

	database "github.com/kaleido-io/firefly/pkg/database"

	fftypes "github.com/kaleido-io/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// Capabilities provides a mock function with given fields:
func (_m *Plugin) Capabilities() *database.Capabilities {
	ret := _m.Called()

	var r0 *database.Capabilities
	if rf, ok := ret.Get(0).(func() *database.Capabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Capabilities)
		}
	}

	return r0
}

// DeleteBlocked provides a mock function with given fields: ctx, id
func (_m *Plugin) DeleteBlocked(ctx context.Context, id *fftypes.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOffset provides a mock function with given fields: ctx, t, ns, name
func (_m *Plugin) DeleteOffset(ctx context.Context, t fftypes.LowerCasedType, ns string, name string) error {
	ret := _m.Called(ctx, t, ns, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, fftypes.LowerCasedType, string, string) error); ok {
		r0 = rf(ctx, t, ns, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSubscriptionByID provides a mock function with given fields: ctx, id
func (_m *Plugin) DeleteSubscriptionByID(ctx context.Context, id *fftypes.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBatchByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetBatchByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Batch, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Batch
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Batch); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatches provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetBatches(ctx context.Context, filter database.Filter) ([]*fftypes.Batch, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Batch
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Batch); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlocked provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetBlocked(ctx context.Context, filter database.Filter) ([]*fftypes.Blocked, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Blocked
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Blocked); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Blocked)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockedByContext provides a mock function with given fields: ctx, ns, _a2, groupID
func (_m *Plugin) GetBlockedByContext(ctx context.Context, ns string, _a2 string, groupID *fftypes.UUID) (*fftypes.Blocked, error) {
	ret := _m.Called(ctx, ns, _a2, groupID)

	var r0 *fftypes.Blocked
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *fftypes.UUID) *fftypes.Blocked); ok {
		r0 = rf(ctx, ns, _a2, groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Blocked)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *fftypes.UUID) error); ok {
		r1 = rf(ctx, ns, _a2, groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetData provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetData(ctx context.Context, filter database.Filter) ([]*fftypes.Data, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Data
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Data); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Data)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataByID provides a mock function with given fields: ctx, id, withValue
func (_m *Plugin) GetDataByID(ctx context.Context, id *fftypes.UUID, withValue bool) (*fftypes.Data, error) {
	ret := _m.Called(ctx, id, withValue)

	var r0 *fftypes.Data
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, bool) *fftypes.Data); ok {
		r0 = rf(ctx, id, withValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Data)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, bool) error); ok {
		r1 = rf(ctx, id, withValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataRefs provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetDataRefs(ctx context.Context, filter database.Filter) (fftypes.DataRefs, error) {
	ret := _m.Called(ctx, filter)

	var r0 fftypes.DataRefs
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) fftypes.DataRefs); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fftypes.DataRefs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatatypeByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetDatatypeByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Datatype, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Datatype
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Datatype); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Datatype)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatatypeByName provides a mock function with given fields: ctx, ns, name, version
func (_m *Plugin) GetDatatypeByName(ctx context.Context, ns string, name string, version string) (*fftypes.Datatype, error) {
	ret := _m.Called(ctx, ns, name, version)

	var r0 *fftypes.Datatype
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *fftypes.Datatype); ok {
		r0 = rf(ctx, ns, name, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Datatype)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, ns, name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatatypes provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetDatatypes(ctx context.Context, filter database.Filter) ([]*fftypes.Datatype, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Datatype
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Datatype); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Datatype)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetEventByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Event, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Event
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Event); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetEvents(ctx context.Context, filter database.Filter) ([]*fftypes.Event, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Event
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Event); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessageByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetMessageByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Message, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Message
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Message); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessageRefs provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetMessageRefs(ctx context.Context, filter database.Filter) ([]*fftypes.MessageRef, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.MessageRef
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.MessageRef); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.MessageRef)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessages provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetMessages(ctx context.Context, filter database.Filter) ([]*fftypes.Message, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Message
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Message); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessagesForData provides a mock function with given fields: ctx, dataID, filter
func (_m *Plugin) GetMessagesForData(ctx context.Context, dataID *fftypes.UUID, filter database.Filter) ([]*fftypes.Message, error) {
	ret := _m.Called(ctx, dataID, filter)

	var r0 []*fftypes.Message
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Filter) []*fftypes.Message); ok {
		r0 = rf(ctx, dataID, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, database.Filter) error); ok {
		r1 = rf(ctx, dataID, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespace provides a mock function with given fields: ctx, name
func (_m *Plugin) GetNamespace(ctx context.Context, name string) (*fftypes.Namespace, error) {
	ret := _m.Called(ctx, name)

	var r0 *fftypes.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string) *fftypes.Namespace); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespaces provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetNamespaces(ctx context.Context, filter database.Filter) ([]*fftypes.Namespace, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Namespace); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNode provides a mock function with given fields: ctx, identity
func (_m *Plugin) GetNode(ctx context.Context, identity string) (*fftypes.Node, error) {
	ret := _m.Called(ctx, identity)

	var r0 *fftypes.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) *fftypes.Node); ok {
		r0 = rf(ctx, identity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, identity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetNodeByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Node, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Node
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Node); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodes provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetNodes(ctx context.Context, filter database.Filter) ([]*fftypes.Node, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Node
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Node); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOffset provides a mock function with given fields: ctx, t, ns, name
func (_m *Plugin) GetOffset(ctx context.Context, t fftypes.LowerCasedType, ns string, name string) (*fftypes.Offset, error) {
	ret := _m.Called(ctx, t, ns, name)

	var r0 *fftypes.Offset
	if rf, ok := ret.Get(0).(func(context.Context, fftypes.LowerCasedType, string, string) *fftypes.Offset); ok {
		r0 = rf(ctx, t, ns, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Offset)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, fftypes.LowerCasedType, string, string) error); ok {
		r1 = rf(ctx, t, ns, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOffsets provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetOffsets(ctx context.Context, filter database.Filter) ([]*fftypes.Offset, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Offset
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Offset); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Offset)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperationByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetOperationByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Operation, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Operation); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperations provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetOperations(ctx context.Context, filter database.Filter) ([]*fftypes.Operation, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Operation
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Operation); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizationByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetOrganizationByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Organization, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Organization
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Organization); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizationByIdentity provides a mock function with given fields: ctx, identity
func (_m *Plugin) GetOrganizationByIdentity(ctx context.Context, identity string) (*fftypes.Organization, error) {
	ret := _m.Called(ctx, identity)

	var r0 *fftypes.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string) *fftypes.Organization); ok {
		r0 = rf(ctx, identity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, identity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizationByName provides a mock function with given fields: ctx, name
func (_m *Plugin) GetOrganizationByName(ctx context.Context, name string) (*fftypes.Organization, error) {
	ret := _m.Called(ctx, name)

	var r0 *fftypes.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string) *fftypes.Organization); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizations provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetOrganizations(ctx context.Context, filter database.Filter) ([]*fftypes.Organization, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Organization
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Organization); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetSubscriptionByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Subscription, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Subscription); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionByName provides a mock function with given fields: ctx, ns, name
func (_m *Plugin) GetSubscriptionByName(ctx context.Context, ns string, name string) (*fftypes.Subscription, error) {
	ret := _m.Called(ctx, ns, name)

	var r0 *fftypes.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *fftypes.Subscription); ok {
		r0 = rf(ctx, ns, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ns, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptions provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetSubscriptions(ctx context.Context, filter database.Filter) ([]*fftypes.Subscription, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Subscription); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByID provides a mock function with given fields: ctx, id
func (_m *Plugin) GetTransactionByID(ctx context.Context, id *fftypes.UUID) (*fftypes.Transaction, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *fftypes.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactions provides a mock function with given fields: ctx, filter
func (_m *Plugin) GetTransactions(ctx context.Context, filter database.Filter) ([]*fftypes.Transaction, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*fftypes.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter) []*fftypes.Transaction); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: ctx, prefix, callbacks
func (_m *Plugin) Init(ctx context.Context, prefix config.Prefix, callbacks database.Callbacks) error {
	ret := _m.Called(ctx, prefix, callbacks)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Prefix, database.Callbacks) error); ok {
		r0 = rf(ctx, prefix, callbacks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitPrefix provides a mock function with given fields: prefix
func (_m *Plugin) InitPrefix(prefix config.Prefix) {
	_m.Called(prefix)
}

// Name provides a mock function with given fields:
func (_m *Plugin) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RunAsGroup provides a mock function with given fields: ctx, fn
func (_m *Plugin) RunAsGroup(ctx context.Context, fn func(context.Context) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBatch provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateBatch(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBlocked provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateBlocked(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateData provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateData(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDatatype provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateDatatype(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEvent provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateEvent(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMessage provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateMessage(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMessages provides a mock function with given fields: ctx, filter, update
func (_m *Plugin) UpdateMessages(ctx context.Context, filter database.Filter, update database.Update) error {
	ret := _m.Called(ctx, filter, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Filter, database.Update) error); ok {
		r0 = rf(ctx, filter, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNamespace provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateNamespace(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNode provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateNode(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOffset provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateOffset(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOperation provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateOperation(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrganization provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateOrganization(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSubscription provides a mock function with given fields: ctx, ns, name, update
func (_m *Plugin) UpdateSubscription(ctx context.Context, ns string, name string, update database.Update) error {
	ret := _m.Called(ctx, ns, name, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, database.Update) error); ok {
		r0 = rf(ctx, ns, name, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTransaction provides a mock function with given fields: ctx, id, update
func (_m *Plugin) UpdateTransaction(ctx context.Context, id *fftypes.UUID, update database.Update) error {
	ret := _m.Called(ctx, id, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, database.Update) error); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertBatch provides a mock function with given fields: ctx, data, allowExisting, allowHashUpdate
func (_m *Plugin) UpsertBatch(ctx context.Context, data *fftypes.Batch, allowExisting bool, allowHashUpdate bool) error {
	ret := _m.Called(ctx, data, allowExisting, allowHashUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Batch, bool, bool) error); ok {
		r0 = rf(ctx, data, allowExisting, allowHashUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertBlocked provides a mock function with given fields: ctx, blocked, allowExisting
func (_m *Plugin) UpsertBlocked(ctx context.Context, blocked *fftypes.Blocked, allowExisting bool) error {
	ret := _m.Called(ctx, blocked, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Blocked, bool) error); ok {
		r0 = rf(ctx, blocked, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertData provides a mock function with given fields: ctx, data, allowExisting, allowHashUpdate
func (_m *Plugin) UpsertData(ctx context.Context, data *fftypes.Data, allowExisting bool, allowHashUpdate bool) error {
	ret := _m.Called(ctx, data, allowExisting, allowHashUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Data, bool, bool) error); ok {
		r0 = rf(ctx, data, allowExisting, allowHashUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertDatatype provides a mock function with given fields: ctx, datadef, allowExisting
func (_m *Plugin) UpsertDatatype(ctx context.Context, datadef *fftypes.Datatype, allowExisting bool) error {
	ret := _m.Called(ctx, datadef, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Datatype, bool) error); ok {
		r0 = rf(ctx, datadef, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertEvent provides a mock function with given fields: ctx, data, allowExisting
func (_m *Plugin) UpsertEvent(ctx context.Context, data *fftypes.Event, allowExisting bool) error {
	ret := _m.Called(ctx, data, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Event, bool) error); ok {
		r0 = rf(ctx, data, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertMessage provides a mock function with given fields: ctx, message, allowExisting, allowHashUpdate
func (_m *Plugin) UpsertMessage(ctx context.Context, message *fftypes.Message, allowExisting bool, allowHashUpdate bool) error {
	ret := _m.Called(ctx, message, allowExisting, allowHashUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Message, bool, bool) error); ok {
		r0 = rf(ctx, message, allowExisting, allowHashUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertNamespace provides a mock function with given fields: ctx, data, allowExisting
func (_m *Plugin) UpsertNamespace(ctx context.Context, data *fftypes.Namespace, allowExisting bool) error {
	ret := _m.Called(ctx, data, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Namespace, bool) error); ok {
		r0 = rf(ctx, data, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertNode provides a mock function with given fields: ctx, data, allowExisting
func (_m *Plugin) UpsertNode(ctx context.Context, data *fftypes.Node, allowExisting bool) error {
	ret := _m.Called(ctx, data, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Node, bool) error); ok {
		r0 = rf(ctx, data, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertOffset provides a mock function with given fields: ctx, data, allowExisting
func (_m *Plugin) UpsertOffset(ctx context.Context, data *fftypes.Offset, allowExisting bool) error {
	ret := _m.Called(ctx, data, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Offset, bool) error); ok {
		r0 = rf(ctx, data, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertOperation provides a mock function with given fields: ctx, operation, allowExisting
func (_m *Plugin) UpsertOperation(ctx context.Context, operation *fftypes.Operation, allowExisting bool) error {
	ret := _m.Called(ctx, operation, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Operation, bool) error); ok {
		r0 = rf(ctx, operation, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertOrganization provides a mock function with given fields: ctx, data, allowExisting
func (_m *Plugin) UpsertOrganization(ctx context.Context, data *fftypes.Organization, allowExisting bool) error {
	ret := _m.Called(ctx, data, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Organization, bool) error); ok {
		r0 = rf(ctx, data, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertSubscription provides a mock function with given fields: ctx, data, allowExisting
func (_m *Plugin) UpsertSubscription(ctx context.Context, data *fftypes.Subscription, allowExisting bool) error {
	ret := _m.Called(ctx, data, allowExisting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Subscription, bool) error); ok {
		r0 = rf(ctx, data, allowExisting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertTransaction provides a mock function with given fields: ctx, data, allowExisting, allowHashUpdate
func (_m *Plugin) UpsertTransaction(ctx context.Context, data *fftypes.Transaction, allowExisting bool, allowHashUpdate bool) error {
	ret := _m.Called(ctx, data, allowExisting, allowHashUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Transaction, bool, bool) error); ok {
		r0 = rf(ctx, data, allowExisting, allowHashUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
