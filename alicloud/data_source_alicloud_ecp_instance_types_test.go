package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAlicloudEcpInstanceTypesDataSource(t *testing.T) {
	rand := acctest.RandIntRange(100, 999)
	regionIdConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudEcpInstanceTypesDataSourceName(rand, map[string]string{}),
		fakeConfig:  "",
	}

	var existAlicloudEcpZoneDataSourceNameMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"instance_types.#": CHECKSET,
		}
	}
	var fakeEcpZonesMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"instance_types.#": "0",
		}
	}
	var alicloudEcpZonesAccountBusesCheckInfo = dataSourceAttr{
		resourceId:   "data.alicloud_ecp_instance_types.default",
		existMapFunc: existAlicloudEcpZoneDataSourceNameMapFunc,
		fakeMapFunc:  fakeEcpZonesMapFunc,
	}

	preCheck := func() {
		testAccPreCheck(t)
	}

	alicloudEcpZonesAccountBusesCheckInfo.dataSourceTestCheckWithPreCheck(t, rand, preCheck, regionIdConf)
}

func testAccCheckAlicloudEcpInstanceTypesDataSourceName(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}

	config := fmt.Sprintf(`
data "alicloud_ecp_instance_types" "default" {  
   %s
}
`, strings.Join(pairs, " \n "))
	return config
}
