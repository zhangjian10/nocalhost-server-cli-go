package api

import (
	"nocalhost/server/utils/assert"

	"github.com/go-resty/resty/v2"
)

type vcluster struct {
	Id int `json:"id"`
}

func CreateVcluster(cluster_id int) int {

	assert.NotEmpty(UserId, "CreateVcluster UserId")

	request := &Request{
		request(V1).SetBody(map[string]interface{}{
			"cluster_id":           cluster_id,
			"cluster_admin":        0,
			"user_id":              UserId,
			"space_name":           "test-",
			"space_resource_limit": nil,
			"dev_space_type":       3,
			"virtual_cluster": map[string]interface{}{
				"service_type": "ClusterIP",
				"version":      "0.5.2",
				"values":       nil,
			},
		}),
	}

	var r vcluster

	res := request.Execute(resty.MethodPost, "/dev_space", &r)

	if res.isSuccess() {
		assert.NotEmpty(r, "create vcluster")

		return r.Id
	}

	return 0
}
