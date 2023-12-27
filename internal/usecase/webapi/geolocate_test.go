package webapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIPLocation(t *testing.T) {
	geoApi := NewIPGeoWebAPI()
	successResp := []string{
		"24.48.0.1", "21.107.203.106",
		"",
	}

	expectEmptyResp := []string{
		"192.168.52.24", "10.120.130.140", "127.0.0.1",
	}

	for _, ip := range successResp {
		geoOutput, err := geoApi.GetIPLocation(ip)
		assert.NoError(t, err, "successResp error not expected")
		assert.NotEmptyf(t, geoOutput, "empty response 'country: '%s'' for ip: '%s' wasn't as expected\n",
			geoOutput.Country, ip)
	}
	for _, ip := range expectEmptyResp {
		geoOutput, err := geoApi.GetIPLocation(ip)
		assert.NoError(t, err, "expectEmptyResp error not expected")
		assert.Emptyf(t, geoOutput, "success response 'country: '%s'' for ip: '%s' wasn't as expected\n",
			geoOutput.Country, ip)
	}
}
