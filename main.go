package goosinfo

import "runtime"

type OSInfo struct {
	Architecture string
	Name         string
	Version      string
	Distribution string
}

// Retrieve OS information.
func Retrieve() (*OSInfo, error) {
	osinfo := OSInfo{
		Architecture: runtime.GOARCH,
		Name:         runtime.GOOS,
	}
	err := osinfo.retrieve()
	if err != nil {
		return nil, err
	}
	return &osinfo, nil
}
