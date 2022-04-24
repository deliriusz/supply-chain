package model

type UserRole int64

const (
	Admin UserRole = iota
	DashboardViewer
	Client
	Unauthorized
)

func (u UserRole) String() string {
	switch u {
	case Admin:
		return "ADMIN"
	case DashboardViewer:
		return "DASHBOARD_VIEWER"
	case Client:
		return "CLIENT"
	}

	return "UNAUTHORIZED"
}
