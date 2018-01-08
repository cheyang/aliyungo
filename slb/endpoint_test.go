package slb

import (
	"testing"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/metadata"
)

// Run this case in the ECS with RamRole
func TestSLBEndpointWithToken(t *testing.T) {
	m := metadata.NewMetaData(nil)
	roleName, err := m.RoleName()
	if err != nil {
		t.Logf("It's not a ramRole ECS: %v", err)
		return
	}

	auth, err := m.RamRoleToken(roleName)
	if err != nil {
		t.Errorf("Faile to get ramRole Token due to %v", err)
		t.FailNow()
	}

	region, err := m.Region()
	if err != nil {
		t.Errorf("Faile to get regionId due to %v", err)
		t.FailNow()
	}

	regionID := common.Region(region)

	slbclient := NewSLBClientWithSecurityToken(auth.AccessKeyId,
		auth.AccessKeySecret,
		auth.SecurityToken,
		region)

	endpoint := slbclient.GetEndpoint()

	t.Logf("endpoint: %s", endpoint)

}

func TestSLBEndpointWithOutToken(t *testing.T) {
	slbclient := NewSLBClientWithSecurityToken(TestAccessKeyId,
		TestAccessKeySecret,
		"",
		TestRegionID)

	endpoint := slbclient.GetEndpoint()

	t.Logf("endpoint: %s", endpoint)
}
