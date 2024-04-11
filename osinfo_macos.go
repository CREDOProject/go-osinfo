//go:build darwin

package goosinfo

func (i *OSInfo) retrieve() error {
	i.Distribution = "macOS"
	// i.Version =
	return nil
}
