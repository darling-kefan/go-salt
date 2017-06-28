package salt_test

import (
	"testing"
	"maintain/go-salt"
)

const (
	SALT_HOST = "192.168.56.66"
	SALT_PORT = "8888"
	SALT_USERNAME = "saltapi"
	SALT_PASSWORD = "123456"
)

func _TestMinions(t *testing.T) {
	conf := salt.Config{
		Host:          SALT_HOST,
		Port:          SALT_PORT,
		Username:      SALT_USERNAME,
		Password:      SALT_PASSWORD,
		Debug:         false,
		SSLSkipVerify: true,
	}
	c, err := salt.NewClient(conf)
	if err != nil {
		t.Errorf("%v", err)
	}

	minion, err := c.Minions()
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%#v\n", minion)
}

func _TestMinion(t *testing.T) {
	conf := salt.Config{
		Host:          SALT_HOST,
		Port:          SALT_PORT,
		Username:      SALT_USERNAME,
		Password:      SALT_PASSWORD,
		Debug:         false,
		SSLSkipVerify: true,
	}
	c, err := salt.NewClient(conf)
	if err != nil {
		t.Errorf("%v", err)
	}

	minion, err := c.Minion("vm-centos7.1-2")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%#v\n", minion)
}

func _TestJobs(t *testing.T) {
	conf := salt.Config{
		Host:          SALT_HOST,
		Port:          SALT_PORT,
		Username:      SALT_USERNAME,
		Password:      SALT_PASSWORD,
		Debug:         false,
		SSLSkipVerify: true,
	}
	c, err := salt.NewClient(conf)
	if err != nil {
		t.Errorf("%v", err)
	}

	ret, err := c.Jobs()
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%#v\n", ret)
}

func TestJob(t *testing.T) {
	conf := salt.Config{
		Host:          SALT_HOST,
		Port:          SALT_PORT,
		Username:      SALT_USERNAME,
		Password:      SALT_PASSWORD,
		Debug:         false,
		SSLSkipVerify: true,
	}
	c, err := salt.NewClient(conf)
	if err != nil {
		t.Errorf("%v", err)
	}

	ret, err := c.Job("20170616100128794607")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%#v\n", ret)
	t.Logf("%v, %v\n", ret.Minions, ret.Result)
}

func _TestCmdAsync(t *testing.T) {
	conf := salt.Config{
		Host:          SALT_HOST,
		Port:          SALT_PORT,
		Username:      SALT_USERNAME,
		Password:      SALT_PASSWORD,
		Debug:         false,
		SSLSkipVerify: true,
	}
	c, err := salt.NewClient(conf)
	if err != nil {
		t.Errorf("%v", err)
	}

	ret, err := c.CmdAsync("*", "cmd.run", "ifconfig", "glob")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%#v\n", ret)
}

func _TestCmd(t *testing.T) {
	conf := salt.Config{
		Host:          SALT_HOST,
		Port:          SALT_PORT,
		Username:      SALT_USERNAME,
		Password:      SALT_PASSWORD,
		Debug:         false,
		SSLSkipVerify: true,
	}
	c, err := salt.NewClient(conf)
	if err != nil {
		t.Errorf("%v", err)
	}

	ret, err := c.Cmd("*", "cmd.run", "uptime", "glob")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%#v\n", ret)
}
