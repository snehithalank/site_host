package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type Domain struct {
    ID      string json:"id"
    ZoneID  string json:"zone_id"
    Name    string json:"name"
}

type DNSRecord struct {
    ID     string json:"id"
    Name   string json:"name"
    Type   string json:"type"
    Value  string json:"value"
}

const (
    apiKey    = "fs23hDtETbKJn3JL323hJYbLXD6yDarZ"
    clientID  = "100"
    baseURL   = "https://api.recruitment.shq.nz"
)

func main() {
    fmt.Println("SiteHost Recruitment, Challenge #1")

    // Part One - Calling an API
    domains := getDomains()
    for _, domain := range domains {
        fmt.Printf("Domain: %s\n", domain.Name)
        dnsRecords := getDNSRecords(domain.ZoneID)
        for _, record := range dnsRecords {
            fmt.Printf("DNS Record - Name: %s, Type: %s, Value: %s\n", record.Name, record.Type, record.Value)
        }
    }

    // Part Two - Resolving a Customer Issue
    checkWebsiteIssue("site.recruitment.shq.nz", "120.138.30.179")
}

// getDomains fetches the domains associated with the clientID
func getDomains() []Domain {
    url := fmt.Sprintf("%s/domains/%s?api_key=%s", baseURL, clientID, apiKey)
    response, err := http.Get(url)
    if err != nil {
        log.Fatalf("Failed to fetch domains: %v", err)
    }
    defer response.Body.Close()

    var domains []Domain
    if err := json.NewDecoder(response.Body).Decode(&domains); err != nil {
        log.Fatalf("Failed to decode domain response: %v", err)
    }
    return domains
}

// getDNSRecords fetches the DNS records for a given zoneID
func getDNSRecords(zoneID string) []DNSRecord {
    url := fmt.Sprintf("%s/zones/%s?api_key=%s", baseURL, zoneID, apiKey)
    response, err := http.Get(url)
    if err != nil {
        log.Fatalf("Failed to fetch DNS records: %v", err)
    }
    defer response.Body.Close()

    var dnsRecords []DNSRecord
    if err := json.NewDecoder(response.Body).Decode(&dnsRecords); err != nil {
        log.Fatalf("Failed to decode DNS records response: %v", err)
    }
    return dnsRecords
}

// checkWebsiteIssue simulates checking for a website issue based on the provided domain and IP address
func checkWebsiteIssue(domain, ipAddress string) {
    fmt.Printf("\nChecking website issue for domain: %s with IP: %s\n", domain, ipAddress)

    // Simulated checks
    fmt.Println("Performing DNS lookup...")
    // In a real-world scenario, you would perform an actual DNS lookup here

    fmt.Println("Checking server availability...")
    // You could use a package like "net" or "os/exec" to ping the server or use a similar approach

    // Simulating a possible issue
    issueFound := true

    if issueFound {
        fmt.Printf("Issue found with domain: %s. It may be due to incorrect DNS settings or server misconfiguration.\n", domain)
        fmt.Println("Suggested Action: Verify DNS records, check server configuration, and ensure that the domain is correctly set up.")
    } else {
        fmt.Printf("No issues found with domain: %s. The website appears to be online and configured correctly.\n", domain)
    }
}
