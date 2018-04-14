package mock

import "github.com/maurofran/iam"

// GroupRepository is the group repository mock.
type GroupRepository struct {
	AddFn             func(*iam.Group) error
	AddInvoked        bool
	UpdateFn          func(*iam.Group) error
	UpdateInvoked     bool
	RemoveFn          func(*iam.Group) error
	RemoveInvoked     bool
	GroupNamedFn      func(iam.TenantID, string) (*iam.Group, error)
	GroupNamedInvoked bool
	AllGroupsFn       func(iam.TenantID) (iam.Groups, error)
	AllGroupsInvoked  bool
}

// Add is the mock method.
func (g *GroupRepository) Add(group *iam.Group) error {
	g.AddInvoked = true
	return g.AddFn(group)
}

// Update is the mock method.
func (g *GroupRepository) Update(group *iam.Group) error {
	g.UpdateInvoked = true
	return g.UpdateFn(group)
}

// Remove is the mock method.
func (g *GroupRepository) Remove(group *iam.Group) error {
	g.RemoveInvoked = true
	return g.RemoveFn(group)
}

// GroupNamed is the mock method.
func (g *GroupRepository) GroupNamed(tenantID iam.TenantID, name string) (*iam.Group, error) {
	g.GroupNamedInvoked = true
	return g.GroupNamedFn(tenantID, name)
}

// AllGroups is the mock method.
func (g *GroupRepository) AllGroups(tenantID iam.TenantID) (iam.Groups, error) {
	g.AllGroupsInvoked = true
	return g.AllGroupsFn(tenantID)
}
