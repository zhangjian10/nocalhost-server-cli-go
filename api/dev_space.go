package api

import (
	"fmt"
	"log"
	"nocalhost/server/utils/assert"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
)

type DevSpace struct {
	ID             int64          `json:"id"`
	DevSpaceType   int64          `json:"dev_space_type"`
	SpaceName      string         `json:"space_name"`
	Kubeconfig     string         `json:"kubeconfig"`
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
			time.Sleep(time.Second * 5)

			r, err := GetDevSpace(r.ID)

			if err != nil {
				return nil, err
			}

			v := r.VirtualCluster

			if v.Status == ReadyStatus || v.Status == ErrorStatus {
				break
			}

			log.Printf("%s is %s, retry ...", r.SpaceName, v.Status)
		}
	}

	return r, nil
}

func GetDevSpace(id int64) (*DevSpace, error) {

	assert.NotEmpty(id, "get devspace id")

	request := &Request{
		request(V2).SetQueryParam("cluster_user_id", fmt.Sprintf("%d", id)),
	}

	var r []DevSpace
	_, err := request.Execute(resty.MethodGet, "/dev_space/detail", &r)

	if err != nil {
		return nil, err
	}

	assert.NotEmpty(r, "get dev space")

	return &r[0], nil
}

func CreateVcluster(cluster_id int64, vcluster bool) (*DevSpace, error) {

	assert.NotEmpty(cluster_id, "create vcluster cluster_id")

	assert.NotEmpty(UserId, "create vcluster user_id")

	var dev_space_type int = 1
	var virtual_cluster map[string]interface{}

	if vcluster {
		dev_space_type = 3
		virtual_cluster = map[string]interface{}{
			"service_type": "ClusterIP",
			"version":      "0.5.2",
		}
	}
	request := &Request{
		request(V1).SetBody(map[string]interface{}{
			"cluster_id":      cluster_id,
			"user_id":         UserId,
			"space_name":      "test-" + strings.ReplaceAll(uuid.NewV1().String(), "-", "")[:6],
			"dev_space_type":  dev_space_type,
			"virtual_cluster": virtual_cluster,
		}),
	}

	var r DevSpace

	_, err := request.Execute(resty.MethodPost, "/dev_space", &r)

	if err != nil {
		return nil, err
	}

	assert.NotEmpty(r, "create vcluster")

	log.Printf("%s create, wait ready", r.SpaceName)

	rr, err := r.waitReady()

	assert.Empty(err)

	return rr, nil
}

func DeleteDevSpace(id int64) error {
	assert.NotEmpty(id, "delete devSpace id")

	_, err := request(V1).Execute(resty.MethodDelete, fmt.Sprintf("/dev_space/%d", id), nil)

	if err == nil {
		log.Printf("delete devSpace %d", id)
	}

	return err
}

func GetKubeconfig(id int64) (*string, error) {

	assert.NotEmpty(id, "get kubeconfig id")

	request := &Request{
		request(V1).SetQueryParam("user_id", fmt.Sprintf("%d", UserId)),
	}

	var r DevSpace

	_, err := request.Execute(resty.MethodGet, "/dev_space/"+fmt.Sprintf("%d", id)+"/detail", &r)

	if err != nil {
		return nil, err
	}

	assert.NotEmpty(r, "get kubeconfig")

	return &r.Kubeconfig, nil

}
