package entrasync

// Option configures optional Client behaviour.
type Option func(*Client)

// UserField is a Graph API user property name that can be requested
// beyond the default set (id, displayName, userPrincipalName).
type UserField string

const (
	FieldDepartment     UserField = "department"
	FieldMailNickname   UserField = "mailNickname"
	FieldOfficeLocation UserField = "officeLocation"
	FieldEmployeeID     UserField = "employeeId"
	FieldCompanyName    UserField = "companyName"
)

// WithUserFields requests additional user properties from Graph beyond the
// default set (id, displayName, userPrincipalName). Fields not requested
// are left as zero values on User.
func WithUserFields(fields ...UserField) Option {
	return func(c *Client) {
		for _, f := range fields {
			c.userSelectFields = append(c.userSelectFields, string(f))
		}
	}
}

// WithTransitiveMemberships fetches transitive group memberships during
// Snapshot. Without this option, Snapshot.Members is nil.
func WithTransitiveMemberships() Option {
	return func(c *Client) {
		c.fetchMemberships = true
	}
}

// WithGroupMemberConcurrency overrides the default parallelism (8) used
// when fetching group members during Snapshot. Values <= 0 are ignored.
func WithGroupMemberConcurrency(n int) Option {
	return func(c *Client) {
		if n > 0 {
			c.concurrency = n
		}
	}
}
