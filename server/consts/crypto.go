package consts

const (
	KEX_EXPORT_DH    = "KEX_EXPORT_DH"
	KEX_ENROL_DEVICE = "KEX_ENROL_DEVICE"
	KEX_HTTP_SR      = "KEX_HTTP_SR"
)

// defines type of credential provider
type CREDPROV string

const (
	CREDPROV_TSXVAULT CREDPROV = "CREDPROV_TSXVAULT"
	CREDPROV_HCVAULT  CREDPROV = "CREDPROV_TSXVAULT"
	CREDPROV_HCVAULT_TOKEN  CREDPROV = "CREDPROV_HCVAULT_TOKEN"
)