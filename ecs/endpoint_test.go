package ecs

import (
	"testing"

	"github.com/denverdino/aliyungo/metadata"
)

func TestLoadEndpointFromFile(t *testing.T) {

}

// Run this case in the ECS with RamRole
func TestECSEndpointWithToekn(t *testing.T) {
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

	ecsclient := NewECSClientWithSecurityToken(auth.AccessKeyId,
		auth.AccessKeySecret,
		auth.SecurityToken,
		region)

	endpoint := ecsclient.GetEndpoint()

	t.Logf("endpoint: %s", endpoint)
}

func TestECSEndpointWithOutToekn(t *testing.T) {
	ecsclient := NewECSClientWithSecurityToken(TestAccessKeyId,
		TestAccessKeySecret,
		"",
		TestRegionID)

	endpoint := ecsclient.GetEndpoint()

	t.Logf("endpoint: %s", endpoint)
}
