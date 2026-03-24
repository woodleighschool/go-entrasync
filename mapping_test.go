//nolint:testpackage // tests require access to unexported functions (userFromGraph)
package entrasync

import (
	"testing"

	"github.com/google/uuid"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
)

func strPtr(s string) *string { return &s }

func TestUserFromGraph_CoreFields(t *testing.T) {
	id := uuid.New()
	u := msgraphmodels.NewUser()
	u.SetId(strPtr(id.String()))
	u.SetDisplayName(strPtr("Alice Smith"))
	u.SetUserPrincipalName(strPtr("alice@example.com"))

	got, ok := userFromGraph(u)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if got.ID != id {
		t.Errorf("ID = %v, want %v", got.ID, id)
	}
	if got.DisplayName != "Alice Smith" {
		t.Errorf("DisplayName = %q, want %q", got.DisplayName, "Alice Smith")
	}
	if got.UPN != "alice@example.com" {
		t.Errorf("UPN = %q", got.UPN)
	}
}

func TestUserFromGraph_ExtendedFields(t *testing.T) {
	id := uuid.New()
	u := msgraphmodels.NewUser()
	u.SetId(strPtr(id.String()))
	u.SetDisplayName(strPtr("Bob Jones"))
	u.SetUserPrincipalName(strPtr("bob@example.com"))
	u.SetDepartment(strPtr("Engineering"))
	u.SetMailNickname(strPtr("bjones"))
	u.SetOfficeLocation(strPtr("Building A"))
	u.SetEmployeeId(strPtr("EMP001"))
	u.SetCompanyName(strPtr("Woodleigh"))

	got, ok := userFromGraph(u)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if got.Department != "Engineering" {
		t.Errorf("Department = %q", got.Department)
	}
	if got.MailNickname != "bjones" {
		t.Errorf("MailNickname = %q", got.MailNickname)
	}
	if got.OfficeLocation != "Building A" {
		t.Errorf("OfficeLocation = %q", got.OfficeLocation)
	}
	if got.EmployeeID != "EMP001" {
		t.Errorf("EmployeeID = %q", got.EmployeeID)
	}
	if got.CompanyName != "Woodleigh" {
		t.Errorf("CompanyName = %q", got.CompanyName)
	}
}

func TestUserFromGraph_UnrequestedFieldsAreZero(t *testing.T) {
	id := uuid.New()
	u := msgraphmodels.NewUser()
	u.SetId(strPtr(id.String()))
	u.SetDisplayName(strPtr("Carol"))
	u.SetUserPrincipalName(strPtr("carol@example.com"))
	// Extended fields deliberately not set — simulates Graph omitting fields not in $select

	got, ok := userFromGraph(u)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if got.Department != "" {
		t.Errorf("Department should be empty, got %q", got.Department)
	}
	if got.MailNickname != "" {
		t.Errorf("MailNickname should be empty, got %q", got.MailNickname)
	}
}

func TestUserFromGraph_NilInput(t *testing.T) {
	_, ok := userFromGraph(nil)
	if ok {
		t.Error("expected ok=false for nil input")
	}
}

func TestUserFromGraph_MissingID(t *testing.T) {
	u := msgraphmodels.NewUser()
	u.SetDisplayName(strPtr("No ID"))
	_, ok := userFromGraph(u)
	if ok {
		t.Error("expected ok=false when ID is missing")
	}
}
