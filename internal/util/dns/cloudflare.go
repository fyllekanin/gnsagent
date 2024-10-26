package dns

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/fyllekanin/gnsagent/internal/schema"
)

func UpdateCloudflareRecord(ip string, domain schema.ConfigDomain) error {
	zoneId, err := getZoneId(domain)
	if err != nil {
		return errors.New("failed getting zone id for domain: " + domain.Domain)
	}

	dnsRecord, err := getDnsRecord(domain, zoneId)
	if err != nil {
		return errors.New("failed getting dns record id for domain: " + domain.Domain)
	}
	if dnsRecord.Content == ip {
		fmt.Printf("Ip for domain %s is already updated correctly", domain.Domain)
	}

	err = updateDnsRecord(domain, zoneId, dnsRecord.Id, ip)
	if err != nil {
		return errors.New("failed to update the dns record")
	}
	return nil
}

func updateDnsRecord(domain schema.ConfigDomain, zoneId string, recordId string, ip string) error {
	body := make(map[string]string)
	body["content"] = ip
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return errors.New("failed creating body to update record")
	}
	reader := bytes.NewReader(bodyBytes)

	_, err = getResponseBytesForRequest(domain, "PATCH", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", zoneId, recordId), reader)
	if err != nil {
		return errors.New("failed to update DNS record")
	}
	return nil
}

func getDnsRecord(domain schema.ConfigDomain, zoneId string) (*schema.CloudflareDnsRecord, error) {
	bytes, err := getResponseBytesForRequest(domain, "GET", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", zoneId), nil)
	if err != nil {
		return nil, errors.New("failed reading the dns records body")
	}

	var zones schema.CloudflareDnsRecords
	err = json.Unmarshal(bytes, &zones)
	if err != nil {
		return nil, errors.New("failed parsing the dns records body")
	}

	var subdomainName = domain.Subdomain + "." + domain.Domain
	for _, item := range zones.Result {
		if domain.Subdomain == "" && item.Name == domain.Domain {
			return item, nil
		}
		if subdomainName == item.Name {
			return item, nil
		}
	}
	return nil, errors.New("failed getting dns records for domain: " + domain.Domain)
}

func getZoneId(domain schema.ConfigDomain) (string, error) {
	bytes, err := getResponseBytesForRequest(domain, "GET", "https://api.cloudflare.com/client/v4/zones", nil)
	if err != nil {
		return "", errors.New("failed reading the zones body")
	}

	var zones schema.CloudflareZones
	err = json.Unmarshal(bytes, &zones)
	if err != nil {
		return "", errors.New("failed parsing the zones body")
	}
	for _, item := range zones.Result {
		if item.Name == domain.Domain {
			return item.Id, nil
		}
	}
	return "", errors.New("failed getting zone id for domain: " + domain.Domain)
}

func getResponseBytesForRequest(domain schema.ConfigDomain, requestType string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(requestType, url, body)
	if err != nil {
		return nil, errors.New("failed creating new request for zones")
	}
	req.Header.Set("X-Auth-Email", domain.Email)
	req.Header.Set("X-Auth-Key", domain.ApiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("failed getting zones from cloudflare")
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("failed reading the zones body")
	}
	return bytes, nil
}
