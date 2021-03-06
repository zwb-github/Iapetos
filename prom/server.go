package prom

import (
	"net"
	"net/http"

	"io"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespaceCore = "iapetos"
	//HostName is the string "host_name"
	HostName = "host_name"
	//ServiceDescription is the string "service_description"
	ServiceDescription = "service_description"
	//CommandName is the string "command_name"
	CommandName = "command_name"
	//Type is the string "type"
	Type = "type"
)

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
<!DOCTYPE html>
<HTML>
	<HEAD><TITLE>Iapetos</TITLE></HEAD>
<BODY>
	<h2>Hi, this is the landingpage of Iapetos.</h2>
	<p>The data you are looking for: <a href="/metrics">metrics</a></p>
	<p><a href="https://github.com/Griesbacher/Iapetos">Github</a></p>
</BODY>
</HTML>`)
}

//InitPrometheus starts the prometheus server
func InitPrometheus(address string) (net.Listener, error) {
	var err error
	prometheusListener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", prometheus.Handler())
		mux.HandleFunc("/", handleMainPage)
		http.Serve(prometheusListener, mux)
	}()
	initCheckData()
	initIapetos()
	initNotificationCheckData()
	initContactNotificationCheckData()
	initStatisticsHost()
	initStatisticsService()
	return prometheusListener, nil
}
