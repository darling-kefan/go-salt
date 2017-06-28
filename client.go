// Refer: https://github.com/r3labs/go-salt/blob/master/client.go
package salt

import (
	"fmt"
	"encoding/json"
)

type Client struct {
	Connector *Connector
}

func NewClient(config Config) (*Client, error) {
	c := Client{}
	c.Connector = NewConnector(config)
	err := c.Connector.Authenticate()

	return &c, err
}

func (c *Client) Minions() (map[string]Minion, error) {
	m := MinionsResponse{}

	resp, err := c.Connector.Get("/minions")
	if err != nil {
		return m.Minions[0], err
	}

	data, err := parseResponse(resp)
	// fmt.Println(string(*data))
	if err != nil {
		return m.Minions[0], err
	}

	err = json.Unmarshal(*data, &m)

	return m.Minions[0], err
}

func (c *Client) Minion(id string) (Minion, error) {
	var mr MinionsResponse
	var m  Minion

	uri := fmt.Sprintf("/minions/%s", id)
	resp, err := c.Connector.Get(uri)
	if err != nil {
		return m, err
	}

	data, err := parseResponse(resp)
	// fmt.Println(string(*data))
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(*data, &mr)
	m = mr.Minions[0][id]

	return m, err
}

func (c *Client) Jobs() ([]map[string]Job, error) {
	var jsr JobsResponse

	resp, err := c.Connector.Get("/jobs")
	if err != nil {
		return jsr.Jobs, err
	}

	data, err := parseResponse(resp)
	// fmt.Printf("%v\n", string(*data))
	if err != nil {
		return jsr.Jobs, err
	}

	err = json.Unmarshal(*data, &jsr)

	return jsr.Jobs, err
}

// @todo: Job结果的缓存时间是多久？
func (c *Client) Job(id string) (Job, error) {
	var jr JobResponse

	uri := fmt.Sprintf("/jobs/%s", id)
	resp, err := c.Connector.Get(uri)
	if err != nil {
		return Job{}, err
	}

	data, err := parseResponse(resp)
	if err != nil {
		return Job{}, err
	}
	fmt.Printf("%v\n", string(*data))

	err = json.Unmarshal(*data, &jr)
	// fmt.Printf("%v\n", jr)

	job := jr.Info[0]
	job.Result = jr.Return[0]

	return job, err
}

// 异步执行salt cmd模块
//
// Start an execution command and immediately return the job id
//
// 参考：https://docs.saltstack.com/en/latest/ref/clients/#salt.client.LocalClient
//
//
// salt.script使用示例：
// - salt '1.2.3.4' cmd.script   "salt://scripts/test.sh" "ok test" runas="user"
// - curl -k https://127.0.0.1:8000/ -H "Accept: application/x-yaml" -H "X-Auth-Token: 786fb622f6ef8b8f719de57dd6145df2d1981899" -d client='local' -d tgt='1,2,3,4' -d fun='cmd.script' -d arg='salt://scripts//test.sh' -d arg="ok hello" -d arg='runas=user'
//
//
// tgt - salt target '192.168.110.132' 'os:Linux'
// fun - 执行模块 test.ping
// arg=() - 模块参数
// exprForm - target的类型
// glob - Bash glob completion - Default
// pcre - Perl style regular expression
// list - Python list of hosts
// grain - Match based on a grain comparison
// grain_pcre - Grain comparison with a regex
// pillar - Pillar data comparison
// pillar_pcre - Pillar data comparison with a regex
// nodegroup - Match on nodegroup
// range - Use a Range server for matching
// compound - Pass a compound match string
// ipcidr - Match based on Subnet (CIDR notation) or IPv4 address.
func (c *Client) CmdAsync(tgt string, fun string, arg []string, exprForm string) (jobId string, err error) {
	args := struct{
		Client   string   `json:"client"`
		Tgt      string   `json:"tgt"`
		Fun      string   `json:"fun"`
		Arg      []string `json:"arg"`
		ExprForm string   `json:"exprForm"`
	}{
		Client:   "local_async",
		Tgt:      tgt,
		Fun:      fun,
		Arg:      arg,
		ExprForm: exprForm,
	}
	param, err := json.Marshal(args)

	// 两种请求方式任选其一
	// resp, err := c.Connector.Post("/", param)
	resp, err := c.Connector.Post("/minions", param)
	if err != nil {
		return
	}

	data, err := parseResponse(resp)
	fmt.Println(string(*data))
	if err != nil {
		return
	}

	er := ExecutionResponse{}
	err = json.Unmarshal(*data, &er)
	
	fmt.Printf("%#v\n", er)
	jobId = er.Job[0].ID
	
	return
}

// 异步执行salt cmd模块
func (c *Client) Cmd(tgt string, fun string, arg []string, exprForm string) (result Result, err error) {
	args := struct{
		Client   string   `json:"client"`
		Tgt      string   `json:"tgt"`
		Fun      string   `json:"fun"`
		Arg      []string `json:"arg"`
		ExprForm string   `json:"exprForm"`
	}{
		Client:   "local",
		Tgt:      tgt,
		Fun:      fun,
		Arg:      arg,
		ExprForm: exprForm,
	}
	param, err := json.Marshal(args)
	// 两种请求方式任选其一
	// resp, err := c.Connector.Post("/", param)
	resp, err := c.Connector.Post("/", param)
	if err != nil {
		return
	}

	data, err := parseResponse(resp)
	// fmt.Printf("%v\n", string(*data))
	if err != nil {
		return
	}

	var rr ResultResponse
	err = json.Unmarshal(*data, &rr)
	result = rr.Return[0]
	
	return
}
