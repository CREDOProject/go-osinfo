//go:build windows

package goosinfo

import "strings"

func (i *OSInfo) retrieve() error {
	i.Distribution = "Windows"
	path, err := commander.LookPath("cmd")
	if err != nil {
		return err
	}
	cmd := commander.Command(path, "/C", "ver")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	i.Version = strings.TrimSpace(string(out))
	return nil
}
