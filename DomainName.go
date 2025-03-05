package CrazyLabelling

const (
	accessDomainDev       = "localhost:55001"
	accessDomainProd      = "access.elderlab.int"
	ssoBackDev            = "localhost:55002"
	ssoBackProd           = "sso-backend.elderlab.int"
	ssoFrontDev           = "localhost:55003"
	ssoFrontProd          = "sso-frontend.elderlab.int"
	apacheDev             = "localhost:55004"
	apacheProd            = "apache.elderlab.int"
	gSheetKillerBackDev   = "localhost:55005"
	gSheetKillerBackProd  = "gsheetkiller-backend.elderlab.int"
	gSheetKillerFrontDev  = "localhost:55006"
	gSheetKillerFrontProd = "gsheetkiller-frontend.elderlab.int"
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
)

func calculateDomainName() {
	if !IsProd {
		AccessDomain = accessDomainDev
		SSOBack = ssoBackDev
		SSOFront = ssoFrontDev
		Apache = apacheDev
		GSheetKillerBack = gSheetKillerBackDev
		GSheetKillerFront = gSheetKillerFrontDev
	}
}
