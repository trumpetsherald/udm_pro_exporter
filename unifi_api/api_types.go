package unifiapi

type LoginAPIErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SystemInfo struct {
	Version          string   `json:"version"`
	UDMVersion       string   `json:"udm_version"`
	Uptime           int      `json:"uptime"`
	IpAddresses      []string `json:"ip_addrs"`
	UpdateAvailable  bool     `json:"update_available"`
	UpdateDownloaded bool     `json:"update_downloaded"`
}

type SysInfoAPIResponse struct {
	Meta struct {
		RC string
	} `json:"meta"`

	Data []SystemInfo `json:"data"`
}

type Stats struct {
	SiteID              string `json:"site_id"`
	MAC                 string `json:"mac"`
	FirstSeen           int    `json:"first_seen"`
	LastSeen            int    `json:"last_seen"`
	Hostname            string `json:"hostname"`
	DisconnectTimestamp int    `json:"disconnect_timestamp"`
	FixedIP             string `json:"fixed_ip"`
	IP                  string `json:"ip"`
	Network             string `json:"network"`
	Uptime              int    `json:"uptime"`
	WiredTxBytes        int    `json:"wired-tx_bytes"`
	WiredRxBytes        int    `json:"wired-rx_bytes"`
	WiredTxPackets      int    `json:"wired-tx_packets"`
	WiredRxPackets      int    `json:"wired-rx_packets"`
}

type StatsAPIResponse struct {
	Meta struct {
		RC string
	} `json:"meta"`

	Data []Stats `json:"data"`
}

type Temperature struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Value float32 `json:"value"`
}

type Storage struct {
	MountPoint string `json:"mount_point"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Type       string `json:"type"`
	Used       int    `json:"used"`
}

type SysStats struct {
	LoadAvg1  float32 `json:"loadavg_1,string"`
	LoadAvg15 float32 `json:"loadavg_15,string"`
	LoadAvg5  float32 `json:"loadavg_5,string"`
	MemBuffer int     `json:"mem_buffer"`
	MemTotal  int     `json:"mem_total"`
	MemUsed   int     `json:"mem_used"`
}

type SystemStats struct {
	CPU    float32 `json:"cpu,string"`
	Mem    float32 `json:"mem,string"`
	Uptime int     `json:"uptime,string"`
}

type UptimeStats struct {
	Availability   float32 `json:"availability"`
	LatencyAverage int     `json:"latency_average"`
	TimePeriod     int     `json:"time_period"`
	Downtime       int     `json:"downtime"`
}

type SpeedTestStatus struct {
	Latency         int     `json:"latency"`
	RunDate         int     `json:"rundate"`
	RunTime         int     `json:"runtime"`
	SourceInterface string  `json:"source_interface"`
	StatusDownload  int     `json:"status_download"`
	StatusPing      int     `json:"status_ping"`
	StatusSummary   int     `json:"status_summary"`
	StatusUpload    int     `json:"status_upload"`
	XputDownload    float32 `json:"xput_download"`
	XputUpload      float32 `json:"xput_upload"`
}

type Device struct {
	IP               string                 `json:"ip"`
	MAC              string                 `json:"mac"`
	Version          string                 `json:"version"`
	KernelVersion    string                 `json:"kernel_version"`
	Temperatures     []Temperature          `json:"temperatures"`
	Storage          []Storage              `json:"storage"`
	Uptime           int                    `json:"uptime"`
	SysStats         SysStats               `json:"sys_stats"`
	SystemStats      SystemStats            `json:"system-stats"`
	StartupTimestamp int                    `json:"startup_timestamp"`
	UptimeStats      map[string]UptimeStats `json:"uptime_stats"`
	SpeedTestStatus  SpeedTestStatus        `json:"speedtest-status"`
	TransmitBytes    int                    `json:"tx_bytes"`
	ReceiveBytes     int                    `json:"rx_bytes"`
	Bytes            int                    `json:"bytes"`
}

type DeviceAPIResponse struct {
	Meta struct {
		RC string
	} `json:"meta"`

	Data []Device `json:"data"`
}
