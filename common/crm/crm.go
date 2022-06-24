package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Crm struct {
	url             string
	accessKey       string
	accessKeySecret string
}
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type Member struct {
	Id             string   `json:"id"`
	State          string   `json:"state"`
	EnablePassword bool     `json:"enablePassword"`
	ActivateTime   string   `json:"activateTime"`
	RegisterTime   string   `json:"registerTime"`
	RegisterScene  string   `json:"registerScene"`
	ExpiredDate    string   `json:"expiredDate"`
	Name           string   `json:"name"`
	NickName       string   `json:"nickName"`
	Gender         string   `json:"gender"`
	Birthday       string   `json:"birthday"`
	SpareMobile    string   `json:"spareMobile"`
	OwnStore       OwnStore `json:"ownStore"`
}
type OwnStore struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type MemberRes struct {
	Result
	Data Member `json:"data"`
}
type IntegralReq struct {
	Channel       Channel `json:"channel"`
	Event         string  `json:"event"`
	ExpiredTime   string  `json:"expiredTime"`
	MemberId      string  `json:"memberId"`
	OccurredOrgId string  `json:"occurredOrgId"`
	OccurredTime  string  `json:"occurredTime"`
	Operator      string  `json:"operator"`
	Points        float32 `json:"points"`
	TransId       TransId `json:"transId"`
}
type PayReq struct {
	Amount       float32 `json:"amount"`
	Channel      Channel `json:"channel"`
	MemberId     string  `json:"memberId"`
	OccurredOrg  string  `json:"occurredOrg"`
	OccurredTime string  `json:"occurredTime"`
	Operator     string  `json:"operator"`
	Password     string  `json:"password"`
	PayMode      string  `json:"payMode"`
	TransId      TransId `json:"transId"`
}
type Channel struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
type TransId struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace" `
}

func NewCrm(url string, accessKey string, accessKeySecret string) *Crm {
	return &Crm{
		url:             url,
		accessKey:       accessKey,
		accessKeySecret: accessKeySecret,
	}
}

/**
 * GetIntegral
 * @Description:查询会员积分余额
 */
func (c *Crm) GetIntegral(memberId string) (bool, error) {
	path := "/points-web/v1/balance/balance/get?memberId=" + memberId
	var res Result
	resp, err := c.sendRequest(path, "GET", "")
	if err != nil {
		return false, err
	}
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		return false, err
	}
	if res.Code == 2000 {
		return true, nil
	} else {
		return false, errors.New(res.Msg)
	}
}

/**
 * GetMember
 * @Description:查询会员信息
 */
func (c *Crm) GetMember(memberId string) (Member, error) {
	path := "/member-web/v1/member/get?memberId=" + memberId
	var res MemberRes
	resp, err := c.sendRequest(path, "GET", "")
	if err != nil {
		return Member{}, err
	}
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		return Member{}, err
	}
	if res.Code == 2000 {
		return res.Data, nil
	} else {
		return Member{}, errors.New(res.Msg)
	}
}

/**
 * UpdateIntegral
 * @Description:调整会员积分
 */
func (c *Crm) UpdateIntegral(memberId string, transId string) (bool, error) {
	path := "/points-web/v1/account/points-account/adjust"
	req := IntegralReq{
		Channel:       Channel{"yike", "third"},
		Event:         "score_adjust",
		ExpiredTime:   "2023-01-01T00:00:00Z",
		MemberId:      memberId,
		OccurredOrgId: "-",
		OccurredTime:  "2022-06-16T00:00:00Z",
		Points:        1.0,
		Operator:      "ludao-test",
		TransId:       TransId{transId, "-"},
	}
	reqByte, err := json.Marshal(req)

	if err != nil {
		return false, err
	}
	resp, err := c.sendRequest(path, "POST", string(reqByte))
	if err != nil {
		return false, err
	}
	var res Result
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		return false, err
	}
	if res.Code == 2000 {
		return true, nil
	} else {
		return false, errors.New(res.Msg)
	}

}

/**
 * PayByMember
 * @Description:会员储值密码支付
 */
func (c *Crm) PayByMember(memberId string, transId string, amount float32, password string, payMod string) (bool, error) {
	path := "/prepay-web/v1/prepay/member/balance/payByMemberId"
	req := PayReq{
		Channel:      Channel{"yike", "third"},
		Amount:       amount,
		MemberId:     memberId,
		OccurredOrg:  "-",
		OccurredTime: "2022-06-16T00:00:00Z",
		Password:     password,
		PayMode:      payMod,
		Operator:     "ludao-test",
		TransId:      TransId{transId, "-"},
	}
	reqByte, err := json.Marshal(req)

	if err != nil {
		return false, err
	}
	resp, err := c.sendRequest(path, "POST", string(reqByte))
	if err != nil {
		return false, err
	}
	var res Result
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		return false, err
	}
	if res.Code == 2000 {
		return true, nil
	} else {
		return false, errors.New(res.Msg)
	}
}

/**
 * sendRequest
 * @Description:发起http请求接口
 */
func (c *Crm) sendRequest(path string, method string, data string) (string, error) {
	client := &http.Client{}
	body := strings.NewReader(data)
	req, err := http.NewRequest(method, c.url+path, body)
	if err != nil {
		return "", err
	}
	contentType := ""
	if method == "POST" {
		contentType = "application/json;charset=UTF-8"
		req.Header.Add("Content-Type", contentType)
	}
	authorization := c.sign(path, method, contentType, data)
	req.Header.Add("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(res))
	return string(res), nil
}

/**
 * sign
 * @Description:通过请求参数生成签名
 */
func (c *Crm) sign(url string, method string, contentType string, data string) (authorization string) {
	bodyMd5Base64 := ""
	if data != "" && method == "POST" {
		md5String := md5.Sum([]byte(data))
		bodyMD5 := strings.ToUpper(fmt.Sprintf("%x", md5String))
		bodyMd5Base64 = base64.StdEncoding.EncodeToString([]byte(bodyMD5))
	}
	ciphertext := method + "\n" + bodyMd5Base64 + "\n" + contentType + "\n\n" + url
	hmacObj := hmac.New(sha1.New, []byte(c.accessKeySecret))
	hmacObj.Write([]byte(ciphertext))
	signatureBase64 := base64.StdEncoding.EncodeToString(hmacObj.Sum([]byte("")))
	authorization = "Sign " + c.accessKey + ":" + signatureBase64
	return
}

func main() {
	crm := NewCrm("http://crmtest.ludaostore.net:10080", "2jx9n326", "CIIDBJMNTX")
	//res,err:= crm.GetMember("305a4169771c48ff8622c9d1ab0f220f")
	//res, err := crm.UpdateIntegral("d98f005a25af457bb6e640a116c227e9", "-")
	res, err := crm.PayByMember("305a4169771c48ff8622c9d1ab0f220f", "-", 1, "111111", "FIX_PAY")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
