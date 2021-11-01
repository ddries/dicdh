package main

type CloudflareDnsEntry struct {
	Id string `json:"id"`
	Ttl int `json:"ttl"`
}

type CloudflareDnsResult struct {
	Result []CloudflareDnsEntry `json:"result"`
}

type Dns struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Ttl int `json:"ttl"`
	Proxy bool `json:"proxy"`
}

type ConfigFile struct {
	DnsList []Dns `json:"dns"`
	Key string `json:"key"`
	Zone string `json:"zone"`
}