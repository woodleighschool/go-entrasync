//nolint:testpackage // tests require access to unexported functions (applyDefaults, fetchAllGroupMembers) and types (Client)
package entrasync

import (
	"context"
	"testing"
)

func TestNewClientDefaults(t *testing.T) {
	c := &Client{}
	applyDefaults(c)

	if c.concurrency != defaultGroupMemberConcurrency {
		t.Errorf("concurrency = %d, want %d", c.concurrency, defaultGroupMemberConcurrency)
	}
	want := []string{"id", "displayName", "userPrincipalName"}
	if len(c.userSelectFields) != len(want) {
		t.Fatalf("userSelectFields = %v, want %v", c.userSelectFields, want)
	}
	for i, f := range want {
		if c.userSelectFields[i] != f {
			t.Errorf("userSelectFields[%d] = %q, want %q", i, c.userSelectFields[i], f)
		}
	}
	if c.fetchMemberships {
		t.Error("fetchMemberships should default to false")
	}
}

func TestWithUserFields(t *testing.T) {
	c := &Client{}
	applyDefaults(c)
	WithUserFields(FieldDepartment, FieldMailNickname)(c)

	want := []string{"id", "displayName", "userPrincipalName", "department", "mailNickname"}
	if len(c.userSelectFields) != len(want) {
		t.Fatalf("userSelectFields = %v, want %v", c.userSelectFields, want)
	}
	for i, f := range want {
		if c.userSelectFields[i] != f {
			t.Errorf("[%d] = %q, want %q", i, c.userSelectFields[i], f)
		}
	}
}

func TestWithTransitiveMemberships(t *testing.T) {
	c := &Client{}
	applyDefaults(c)
	WithTransitiveMemberships()(c)

	if !c.fetchMemberships {
		t.Error("fetchMemberships should be true after WithTransitiveMemberships")
	}
}

func TestWithGroupMemberConcurrency(t *testing.T) {
	c := &Client{}
	applyDefaults(c)
	WithGroupMemberConcurrency(16)(c)

	if c.concurrency != 16 {
		t.Errorf("concurrency = %d, want 16", c.concurrency)
	}
}

func TestWithGroupMemberConcurrency_ZeroIgnored(t *testing.T) {
	c := &Client{}
	applyDefaults(c)
	WithGroupMemberConcurrency(0)(c)

	if c.concurrency != defaultGroupMemberConcurrency {
		t.Errorf("concurrency = %d, want default %d", c.concurrency, defaultGroupMemberConcurrency)
	}
}

func TestFetchAllGroupMembers_EmptyGroupsReturnsEmpty(t *testing.T) {
	c := &Client{}
	applyDefaults(c)
	result, err := c.fetchAllGroupMembers(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected empty map, got %v", result)
	}
}
