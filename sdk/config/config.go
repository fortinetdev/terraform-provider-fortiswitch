package config

import (
	"net/http"

	"github.com/terraform-providers/terraform-provider-fortiswitch/sdk/auth"
)

// Config provides configuration to a FortiSwitch client instance
// It saves authentication information and a http connection
// for FortiSwitch Client instance to create New connction to FortiSwitch
// and Send data to FortiSwitch,  etc. (needs to be extended later.)
type Config struct {
	Auth     *auth.Auth
	HTTPCon  *http.Client
	FwTarget string
}
