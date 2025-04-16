package CrazyLabelling

const (
	accessDomainDev         = "localhost:55001"
	accessDomainProd        = "access.elderlab.int"
	ssoBackDev              = "localhost:55002"
	ssoBackProd             = "sso-backend.elderlab.int"
	ssoFrontDev             = "localhost:55003"
	ssoFrontProd            = "sso-frontend.elderlab.int"
	apacheDev               = "localhost:55004"
	apacheProd              = "apache.elderlab.int"
	gSheetKillerBackDev     = "localhost:55005"
	gSheetKillerBackProd    = "gsheetkiller-backend.elderlab.int"
	gSheetKillerFrontDev    = "localhost:55006"
	gSheetKillerFrontProd   = "gsheetkiller-frontend.elderlab.int"
	amazingTrackerBackDev   = "localhost:55007"
	amazingTrackerBackProd  = "amazing-tracker-backend.elderlab.int"
	amazingTrackerFrontDev  = "localhost:55008"
	amazingTrackerFrontProd = "amazing-tracker-frontend.elderlab.int"
)

var (
	// AccessDomain : Access Domain
	AccessDomain = accessDomainProd
	// SSOBack : SSO Backend
	SSOBack = ssoBackProd
	// SSOFront : SSO Frontend
	SSOFront = ssoFrontProd
	// Apache : Apache
	Apache = apacheProd
	// GSheetKillerBack : GSheetKiller Backend
	GSheetKillerBack = gSheetKillerBackProd
	// GSheetKillerFront : GSheetKiller Frontend
	GSheetKillerFront = gSheetKillerFrontProd
	// AmazingTrackerBack : Amazing Tracker Backend
	AmazingTrackerBack = amazingTrackerBackProd
	// AmazingTrackerFront : Amazing Tracker Frontend
	AmazingTrackerFront = amazingTrackerFrontProd
)

func calculateDomainName() {
	if !IsProd {
		AccessDomain = accessDomainDev
		SSOBack = ssoBackDev
		SSOFront = ssoFrontDev
		Apache = apacheDev
		GSheetKillerBack = gSheetKillerBackDev
		GSheetKillerFront = gSheetKillerFrontDev
		AmazingTrackerBack = amazingTrackerBackDev
		AmazingTrackerFront = amazingTrackerFrontDev
	}
}
