package Member

var roles = []string{
	"player",
	"coach",
	"manager",
}

// GetRoles returns the list of roles
func GetRoles() []string {
	return roles
}
