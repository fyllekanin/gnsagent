package schema

type CloudflareZone struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CloudflareZones struct {
	Result []CloudflareZone `json:"result"`
}

type CloudflareDnsRecord struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type CloudflareDnsRecords struct {
	Result []*CloudflareDnsRecord `json:"result"`
}
