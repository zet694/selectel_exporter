package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	primaryPrice *prometheus.Desc
	storagePrice *prometheus.Desc
	vpcPrice     *prometheus.Desc
	vmwarePrice  *prometheus.Desc
}

func newCollector() *Collector {
	return &Collector{
		primaryPrice: prometheus.NewDesc(
			"PrimaryPrice", "Full price of Selectel", nil, nil,
		),
		storagePrice: prometheus.NewDesc(
			"StoragePrice", "Full price of Storage", nil, nil,
		),
		vpcPrice: prometheus.NewDesc(
			"VpcPrice", "Full price of VPC", nil, nil,
		),
		vmwarePrice: prometheus.NewDesc(
			"VmwarePrice", "Full price of Vmware", nil, nil,
		),
	}
}

func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.primaryPrice
	ch <- collector.storagePrice
	ch <- collector.vpcPrice
	ch <- collector.vmwarePrice
}

func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	my_json := new(Price)
	getJson(urls[0], my_json)

	primary := float64(my_json.Data.Primary.Main)
	storage := float64(my_json.Data.Storage.Main)
	vpc := float64(my_json.Data.Vpc.Main)
	vmware := float64(my_json.Data.Vmware.Main)
	ch <- prometheus.MustNewConstMetric(collector.primaryPrice, prometheus.GaugeValue, primary)
	ch <- prometheus.MustNewConstMetric(collector.storagePrice, prometheus.GaugeValue, storage)
	ch <- prometheus.MustNewConstMetric(collector.vpcPrice, prometheus.GaugeValue, vpc)
	ch <- prometheus.MustNewConstMetric(collector.vmwarePrice, prometheus.GaugeValue, vmware)
}

var urls = [...]string{"https://api.selectel.ru/v3/billing/balance"}

// Paste API KEY HERE
var headers = map[string]string{"X-Token": "API_KEY", "Content-Type": "application/json"}

func getJson(url string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	r, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
