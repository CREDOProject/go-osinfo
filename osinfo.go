package goosinfo

import (
	"fmt"
	"runtime"

	shell "github.com/CREDOProject/sharedutils/shell"
)

var commander = shell.New()

type OSInfo struct {
	Architecture string
	Name         string
	Version      string
	Distribution string
	Like         []string
}

func (i *OSInfo) String() string {
	return fmt.Sprintf("%s %s %s %s %v",
		i.Name,
		i.Distribution,
		i.Version,
		i.Architecture,
		i.Like,
	)
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
