//go:build linux

package goosinfo

import (
	"os"
	"strings"
)

func (i *OSInfo) retrieve() error {
	file, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return err
	}
	fileString := strings.Split(string(file), "\n")
	var kv = make(map[string]string)
	for _, line := range fileString {
		splits := strings.Split(line, "=")
		if len(splits) > 1 {
			key := splits[0]
			val := strings.Trim(splits[1], "\"")
			kv[key] = val
		}
	}
	i.Distribution = kv["ID"]
	i.Version = kv["VERSION_ID"]

	return nil
}
