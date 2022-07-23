package unifiapi

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type UDMProCollector struct {
	Broker *APIBroker
}

func NewUDMProCollector(broker *APIBroker) *UDMProCollector {
	collector := UDMProCollector{Broker: broker}

	return &collector
}

func (collector *UDMProCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- temperatureGauge
}

func (collector *UDMProCollector) Collect(ch chan<- prometheus.Metric) {
	if !collector.Broker.IsLoggedIn() {
		loginStatus, err := collector.Broker.Login()
		if err != nil || loginStatus != 200 {
			log.Fatal(err)
		}
	}

	collector.CollectDeviceMetrics(ch)

	logoutStatus, err := collector.Broker.Logout()
	if err != nil || logoutStatus != 200 {
		log.Fatal(err)
	}
}

func (collector *UDMProCollector) CollectDeviceMetrics(ch chan<- prometheus.Metric) {
	deviceResponse, err := collector.Broker.Device()

	if err != nil {
		log.Fatal(err)
	}

	if len(deviceResponse.Data) > 0 {
		device := deviceResponse.Data[0]

		for _, temp := range device.Temperatures {
			ch <- prometheus.MustNewConstMetric(
				temperatureGauge, prometheus.GaugeValue,
				float64(temp.Value), temp.Name, temp.Type,
			)
		}

		for _, storage := range device.Storage {
			ch <- prometheus.MustNewConstMetric(
				storageAvailableGauge, prometheus.GaugeValue,
				float64(storage.Size), storage.MountPoint, storage.Name, storage.Type,
			)

			ch <- prometheus.MustNewConstMetric(
				storageUsedGauge, prometheus.GaugeValue,
				float64(storage.Used), storage.MountPoint, storage.Name, storage.Type,
			)
		}

		ch <- prometheus.MustNewConstMetric(
			uptimeCounter, prometheus.CounterValue, float64(device.Uptime),
		)

		ch <- prometheus.MustNewConstMetric(
			loadAvg1Guage, prometheus.GaugeValue, float64(device.SysStats.LoadAvg1),
		)

		ch <- prometheus.MustNewConstMetric(
			loadAvg5Guage, prometheus.GaugeValue, float64(device.SysStats.LoadAvg5),
		)

		ch <- prometheus.MustNewConstMetric(
			loadAvg15Guage, prometheus.GaugeValue, float64(device.SysStats.LoadAvg15),
		)

		ch <- prometheus.MustNewConstMetric(
			memoryTotalGauge, prometheus.GaugeValue, float64(device.SysStats.MemTotal),
		)

		ch <- prometheus.MustNewConstMetric(
			memoryBufferlGauge, prometheus.GaugeValue, float64(device.SysStats.MemBuffer),
		)

		ch <- prometheus.MustNewConstMetric(
			memoryUsedGauge, prometheus.GaugeValue, float64(device.SysStats.MemUsed),
		)

		ch <- prometheus.MustNewConstMetric(
			memoryPercentGauge, prometheus.GaugeValue, float64(device.SystemStats.Mem),
		)

		ch <- prometheus.MustNewConstMetric(
			cpuPercentGauge, prometheus.GaugeValue, float64(device.SystemStats.CPU),
		)

		ch <- prometheus.MustNewConstMetric(
			deviceTxBytesGauge, prometheus.GaugeValue, float64(device.TransmitBytes),
		)

		ch <- prometheus.MustNewConstMetric(
			deviceRxBytesGauge, prometheus.GaugeValue, float64(device.ReceiveBytes),
		)

		ch <- prometheus.MustNewConstMetric(
			deviceBytesGauge, prometheus.GaugeValue, float64(device.Bytes),
		)
	}
}
