package model

// TenantID is the value object used to provide unique tenant identifier.
type TenantID string

// IsZero will check if tenant identifier is zero value.
func (t TenantID) IsZero() bool {
	return t == ""
}
