package entrasync

import (
	"github.com/google/uuid"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
)

func extractMemberIDs(users []msgraphmodels.Userable) []uuid.UUID {
	var ids []uuid.UUID

	for _, u := range users {
		if u == nil {
			continue
		}

		if id := uuidFromStringPtr(u.GetId()); id != uuid.Nil {
			ids = append(ids, id)
		}
	}

	return ids
}

func userFromGraph(u msgraphmodels.Userable) (User, bool) {
	if u == nil {
		return User{}, false
	}

	id := uuidFromStringPtr(u.GetId())
	if id == uuid.Nil {
		return User{}, false
	}

	return User{
		ID:          id,
		DisplayName: stringOrEmpty(u.GetDisplayName()),
		UPN:         stringOrEmpty(u.GetUserPrincipalName()),
		Department:  stringOrEmpty(u.GetDepartment()),
	}, true
}

func uuidFromStringPtr(v *string) uuid.UUID {
	if v == nil || *v == "" {
		return uuid.Nil
	}

	parsed, err := uuid.Parse(*v)
	if err != nil {
		return uuid.Nil
	}

	return parsed
}

func stringOrEmpty(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}
