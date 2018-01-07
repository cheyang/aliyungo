package common

import (
	"os"
)

//Modify with your Access Key Id and Access Key Secret

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestRegionID        = "cn-zhangjiakou"

	TestInstanceId      = "MY_TEST_INSTANCEID"
	TestSecurityGroupId = "MY_TEST_SECURITY_GROUP_ID"
	TestImageId         = "MY_IMAGE_ID"
	TestAccountId       = "MY_TEST_ACCOUNT_ID" //Get from https://account.console.aliyun.com
	TestInstanceType    = "ecs.n4.large"
	TestVSwitchID       = "MY_TEST_VSWITCHID"

	TestIAmRich = false
	TestQuick   = false
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}

var testLocationClient *Client

func NetTestLocationClientForDebug() *Client {
	if testLocationClient == nil {
		testLocationClient = NewECSClient(TestAccessKeyId, TestAccessKeySecret, TestRegionID)
		testLocationClient.SetDebug(true)
	}

	return testLocationClient
}
