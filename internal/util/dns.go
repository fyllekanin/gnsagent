package util

import (
	"errors"
	"fmt"

	"github.com/fyllekanin/gnsagent/internal/logger"
	"github.com/fyllekanin/gnsagent/internal/schema"
	"github.com/fyllekanin/gnsagent/internal/util/dns"
)

func UpdateDnsService(ip string, domain schema.ConfigDomain) error {
	switch domain.Type {
	case "CLOUDFLARE":
		err := dns.UpdateCloudflareRecord(ip, domain)
		if err != nil {
			logger.Error(fmt.Sprintf("failed updating dns record: %s", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("updated %s to new IP %s", domain.GetFullDomain(), ip))
		}
	default:
		logger.Error(fmt.Sprintf("%s domain type is not supported", domain.Type))
		return errors.New("unsupported domain type")
	}
	return nil
}
