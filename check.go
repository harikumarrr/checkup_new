package checkup

import (
	"encoding/json"
	"fmt"

	"checkup/check/dns"
	"checkup/check/exec"
	"checkup/check/http"
	"checkup/check/tcp"
	"checkup/check/tls"
)

func checkerDecode(typeName string, config json.RawMessage) (Checker, error) {
	switch typeName {
	case dns.Type:
		return dns.New(config)
	case exec.Type:
		return exec.New(config)
	case http.Type:
		return http.New(config)
	case tcp.Type:
		return tcp.New(config)
	case tls.Type:
		return tls.New(config)
	default:
		return nil, fmt.Errorf(errUnknownCheckerType, typeName)
	}
}
