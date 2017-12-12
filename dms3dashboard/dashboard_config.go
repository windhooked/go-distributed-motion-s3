// Package dms3dash dashboard configuration structures and variables
//
package dms3dash

import "time"

var dashboardConfig *tomlTables
var dashboardData *deviceData

// tomlTables represents the TOML table(s)
type tomlTables struct {
	Client *clientKeyValues
	Server *serverKeyValues
}

// clientKeyValues represents the k-v pairs in the TOML file
type clientKeyValues struct {
	ImagesFolder string
}

// serverKeyValues represents the k-v pairs in the TOML file
type serverKeyValues struct {
	Enable       bool
	Port         int
	Filename     string
	FileLocation string
	Title        string
}

// deviceData represents all client dashboard elements
type deviceData struct {
	Title   string
	Clients []DeviceMetrics
}

// DeviceMetrics represents device data presented on the dashboard
type DeviceMetrics struct {
	Hostname       string
	Environment    string
	Kernel         string
	LastReport     time.Time
	StartTime      time.Time
	Uptime         string
	CheckInterval  int
	ShowEventCount bool
	EventCount     int
	Type           dashboardType
}

// dashboardType defines the dashboard device type
type dashboardType int

// states of the motion detector application
const (
	Client dashboardType = iota
	Server
)