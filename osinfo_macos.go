//go:build darwin

package goosinfo

import "strings"

func (i *OSInfo) retrieve() error {
	i.Distribution = "macOS"
	path, err := commander.LookPath("sw_vers")
	if err != nil {
		return err
	}
	cmd := commander.Command(path, "-productVersion")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	i.Version = strings.TrimSpace(string(out))
	return nil
}
