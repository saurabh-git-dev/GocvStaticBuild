package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/yusufpapurcu/wmi"
)

// Win32_PnPEntity maps relevant fields from the WMI class
type Win32_PnPEntity struct {
	Caption                     string
	Description                 string
	InstallDate                 *time.Time // nullable datetime
	Name                        string
	Status                      string
	Availability                *uint16
	ConfigManagerErrorCode      *uint32
	ConfigManagerUserConfig     *bool
	CreationClassName           string
	DeviceID                    string
	ErrorCleared                *bool
	ErrorDescription            *string
	LastErrorCode               *uint32
	PNPDeviceID                 string
	PowerManagementCapabilities []uint16
	PowerManagementSupported    *bool
	StatusInfo                  *uint16
	SystemCreationClassName     string
	SystemName                  string
	ClassGuid                   string
	CompatibleID                []string
	HardwareID                  []string
	Manufacturer                string
	PNPClass                    string
	Present                     *bool
	Service                     string
}

func list_cameras() {
	var devices []Win32_PnPEntity

	// Run the WMI query
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devices)
	if err != nil {
		log.Fatalf("WMI query failed: %v", err)
	}

	// Filter for webcams (commonly use "usbvideo" service)
	for _, d := range devices {
		if d.Service == "usbvideo" {
			fmt.Printf("-------------------------------\n")
			// marshal to the json format
			jsonData, err := json.MarshalIndent(d, "", "  ")
			if err != nil {
				log.Printf("Error marshaling device %s: %v", d.Name, err)
				continue
			}
			fmt.Printf("%+v\n", string(jsonData))
		}
	}
}
