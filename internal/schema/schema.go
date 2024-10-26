package schema

import "fmt"

type ConfigEndPoint struct {
	Url      string `json:"url"`
	Property string `json:"property"`
}

type ConfigDomain struct {
	Type      string `json:"type"`
	Email     string `json:"email"`
	Domain    string `json:"domain"`
	Subdomain string `json:"subdomain"`
	ApiKey    string `json:"apiKey"`
}

func (domain ConfigDomain) GetFullDomain() string {
	if domain.Subdomain == "" {
		return domain.Domain
	}
	return fmt.Sprintf("%s.%s", domain.Subdomain, domain.Domain)
}

type ConfigSchema struct {
	EndPoints []ConfigEndPoint `json:"endPoints"`
	Domains   []ConfigDomain   `json:"domains"`
}
