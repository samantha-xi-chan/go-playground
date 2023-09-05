package play250_http

import (
	"encoding/json"
	"fmt"
	"go-playground/util/util_http"
	"io/ioutil"
	"log"
	"net/http"
)

//# 创建 container
//pref='{"name":"joedval_stress","image":"ubuntu","cmd_str":["/bin/sh","/path/in/docker/cmd.sh", "para01", "para02"],"vol_id":"'
//suff='","vol_path":"/path/in/docker/","exp_id":"random_id_001","exp_path":"/docker/out/"}'
//result=$(printf "%s%s%s" $pref $volId $suff)
//echo "result: $result"
//createContainerResp=$(curl -X POST "http://192.168.31.45:8080/api/v1/container/" -d $result)
//echo "createContainerResp: "$createContainerResp
//containerId=$(echo "$createContainerResp" | jq -r '.data.container_id')
//echo "containerId: $containerId"

func Play() {
	volId, e := PostVol()
	if e != nil {
		log.Println("e: ", e)
	}
	log.Println("volId: ", volId)

	containerId, e := PostContainer(volId, "name")
	if e != nil {
		log.Println("e: ", e)
	}
	log.Println("containerId: ", containerId)

}

type HttpRespBodyPostVol struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data PostVolResp `json:"data"`
}

type PostVolResp struct {
	Id string `json:"id"`
}

type PostContainerReq struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	CmdStr  []string `json:"cmd_str"`
	VolId   string   `json:"vol_id"`
	VolPath string   `json:"vol_path"`

	ExpId   string `json:"exp_id"`
	ExpPath string `json:"exp_path"`
}
type PostContainerResp struct {
	ContainerId string   `json:"container_id"`
	ExportId    []string `json:"export_id"`
}

type HttpRespBodyPostContainer struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data PostContainerResp `json:"data"`
}

func PostVol() (volId string, e error) {
	url := "http://192.168.31.45:8080/api/v1/vol/"
	data := []byte(`{"obj_id":"boot"}`)

	resp, err := util_http.SendHttpReq(url, http.MethodPost, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))

	var httpRespBody HttpRespBodyPostVol
	err = json.Unmarshal(body, &httpRespBody)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	return httpRespBody.Data.Id, nil
}

func PostContainer(volId string, name string) (containerId string, e error) {
	url := "http://192.168.31.45:8080/api/v1/container/"

	req := PostContainerReq{
		Name:    name,
		Image:   "ubuntu",
		CmdStr:  []string{"/bin/sh", "/path/in/docker/cmd.sh", "para01", "para02"},
		VolId:   volId,
		VolPath: "/path/in/docker/",
		ExpId:   "random_id_001",
		ExpPath: "/docker/out/",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := util_http.SendHttpReq(url, http.MethodPost, jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))

	var httpRespBody HttpRespBodyPostContainer
	err = json.Unmarshal(body, &httpRespBody)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	return httpRespBody.Data.ContainerId, nil
}
