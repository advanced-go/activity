package testrsc

import (
	"embed"
	"github.com/advanced-go/common/iox"
)

//go:embed files
var f embed.FS

func init() {
	iox.Mount(f)
}
