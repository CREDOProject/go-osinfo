//go:build windows

package goosinfo

func (i *OSInfo) retrieve() error {
	i.Distribution = "Windows"

	// i.Version =
	return nil
}
