package schema

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

type ConfigSchema struct {
	EndPoints []ConfigEndPoint `json:"endPoints"`
	Domains   []ConfigDomain   `json:"domains"`
}
