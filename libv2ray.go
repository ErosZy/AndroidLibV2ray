package libv2ray

//go:generate make all

import (
	"fmt"

	"github.com/xiaokangwang/AndroidLibV2ray/CoreI"

	"github.com/ErosZy/v2ray-core"
	_ "github.com/ErosZy/v2ray-core/main/distro/all"
)

/*CheckVersion int
This func will return libv2ray binding version.
*/
func CheckVersion() int {
	return CoreI.CheckVersion()
}

/*CheckVersionX string
This func will return libv2ray binding version and V2Ray version used.
*/
func CheckVersionX() string {
	return fmt.Sprintf("Libv2ray rev. %d, along with V2Ray %s", CheckVersion(), core.Version())
}
