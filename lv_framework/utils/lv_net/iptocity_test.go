package lv_net

import "testing"

func TestGetCityByIp(t *testing.T) {
	t.Log(GetCityByIp("175.0.118.197"))
}
