package util

import (
	"errors"
	"fmt"

	"github.com/fyllekanin/gnsagent/internal/schema"
	"github.com/fyllekanin/gnsagent/internal/util/dns"
)

func UpdateDnsService(ip string, domain schema.ConfigDomain) error {
	switch domain.Type {
	case "CLOUDFLARE":
		dns.UpdateCloudflareRecord(ip, domain)
	default:
		fmt.Printf("Error: %s domain type is not supported", domain.Type)
		return errors.New("unsupported domain type")
	}
	return nil
}
