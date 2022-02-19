package api

import (
	"nocalhost/server/utils/assert"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
)

type DevSpace struct {
	ID             int64          `json:"id"`
	DevSpaceType   int64          `json:"dev_space_type"`
	VirtualCluster VirtualCluster `json:"virtual_cluster"`
}

type StatusType string

const (
	InstallStatus StatusType = "install"
	UpgradeStatus StatusType = "upgrade"
	ErrorStatus   StatusType = "error"
	ReadyStatus   StatusType = "Ready"
)

type VirtualCluster struct {
	Status StatusType `json:"status"`
}

func (r *DevSpace) waitReady() (*DevSpace, error) {

	if r.DevSpaceType == 3 {

		for {
			r = GetDevSpace(strconv.FormatInt(r.ID, 10))

			v := r.VirtualCluster

			if v.Status == ReadyStatus || v.Status == ErrorStatus {
				break
			}

			time.Sleep(time.Second * 5)
		}
	}

	return r, nil
}

func GetDevSpace(id string) *DevSpace {
	request := &Request{
		request(V2).SetQueryParam("cluster_user_id", string(id)),
	}

	var r []DevSpace
	res := request.Execute(resty.MethodGet, "/dev_space/detail", &r)

	if res.isSuccess() {
		assert.NotEmpty(r, "get dev space")

		return &r[0]
	}
	return nil
}

func CreateVcluster(cluster_id int) *DevSpace {

	assert.NotEmpty(UserId, "CreateVcluster UserId")

	request := &Request{
		request(V1).SetBody(map[string]interface{}{
			"cluster_id":     cluster_id,
			"user_id":        UserId,
			"space_name":     "test-" + strings.ReplaceAll(uuid.NewV1().String(), "-", "")[:6],
			"dev_space_type": 3,
			"virtual_cluster": map[string]interface{}{
				"service_type": "ClusterIP",
				"version":      "0.5.2",
			},
		}),
	}

	var r DevSpace

	res := request.Execute(resty.MethodPost, "/dev_space", &r)

	if res.isSuccess() {
		assert.NotEmpty(r, "create vcluster")

		rr, err := r.waitReady()

		assert.Empty(err)

		return rr
	}

	return nil
}
