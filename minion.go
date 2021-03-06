package salt

type MinionsResponse struct {
	Minions []map[string]Minion `json:"return"`
}

type Minion struct {
	Id            string    `json:"id"`
	Name          string    `json:"nodename"`
	Host          string    `json:"host"`
	Domain        string    `json:"domain"`
	OS            string    `json:"os"`
	OSRelease     string    `json:"osrelease"`
	OSName        string    `json:"osfullname"`
	Kernel        string    `json:"kernel"`
	KernelRelease string    `json:"kernelrelease"`
	Shell         string    `json:"shell"`
	ARCH          string    `json:"osarch"`
	CPUS          int       `json:"num_cpus"`
	RAM           int       `json:"mem_total"`
	CPUModel      string    `json:"cpu_model"`
	CPUFlags      []string  `json:"cpu_flags"`
	Virtual       string    `json:"virtual"`
	Ipv4          []string  `json:"ipv4"`
	Ipv6          []string  `json:"ipv6"`
	Path          string    `json:"path"`
	ServerID      int       `json:"server_id"`
}
