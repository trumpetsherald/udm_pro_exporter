package unifiapi

import "github.com/prometheus/client_golang/prometheus"

var (
	temperatureGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "temperature_celsius"),
		"Current temperature of various components.",
		[]string{"name", "type"}, nil,
	)

	storageAvailableGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "storage_available_bytes"),
		"The amount of storage available on a mount point.",
		[]string{"mount_point", "name", "type"}, nil,
	)

	storageUsedGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "storage_used_bytes"),
		"The amount of storage used on a mount point.",
		[]string{"mount_point", "name", "type"}, nil,
	)

	uptimeCounter = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "uptime_seconds"),
		"The length of time the system has been running.",
		[]string{}, nil,
	)

	loadAvg1Guage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "load_avg_1"),
		"The amount of load on the system, averaged over 1 minute.",
		[]string{}, nil,
	)

	loadAvg5Guage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "load_avg_5"),
		"The amount of load on the system, averaged over 5 minutes.",
		[]string{}, nil,
	)

	loadAvg15Guage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "load_avg_15"),
		"The amount of load on the system, averaged over 15 minutes.",
		[]string{}, nil,
	)

	memoryTotalGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "memory_total_bytes"),
		"The amount of memory available on the device.",
		[]string{}, nil,
	)

	memoryBufferlGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "memory_buffer_bytes"),
		"The amount of memory allocatted to the buffer on the device.",
		[]string{}, nil,
	)

	memoryUsedGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "memory_used_bytes"),
		"The amount of memory in use on the device.",
		[]string{}, nil,
	)

	memoryPercentGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "memory_used_percent"),
		"The percent of memory in use on the device.",
		[]string{}, nil,
	)

	cpuPercentGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cpu_used_percent"),
		"The percent of CPU utilization on the device.",
		[]string{}, nil,
	)

	deviceTxBytesGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_tx_bytes"),
		"The number of bytes transmitted by the device.",
		[]string{}, nil,
	)

	deviceRxBytesGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_rx_bytes"),
		"The number of bytes received by the device.",
		[]string{}, nil,
	)

	deviceBytesGauge = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "device_bytes"),
		"The total number of bytes sent and received by the device.",
		[]string{}, nil,
	)
)
