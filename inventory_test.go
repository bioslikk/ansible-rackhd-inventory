package main

import (
	"os"
	"testing"
)

const (
	rackhdUrl  = "http://127.0.0.1:5000"
	configPath = "./config.test.yaml"
)

func setEnvironmentVars(t *testing.T) error {
	err := os.Setenv(RackHdApiUrlEnvVarName, rackhdUrl)
	if err != nil {

		return err
	}
	err = os.Setenv(AnsibleRackHdConfigPath, configPath)
	if err != nil {

		return err
	}

	return nil
}

func unsetEnvironmentVars(t *testing.T) error {
	err := os.Unsetenv(RackHdApiUrlEnvVarName)
	if err != nil {

		return err
	}
	err = os.Unsetenv(AnsibleRackHdConfigPath)
	if err != nil {

		return err
	}

	return nil
}

func TestConfigReads(t *testing.T) {
	err := setEnvironmentVars(t)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	props := getPropsFromConfig()
	if props.rackhdUrl != rackhdUrl {
		t.Errorf("\n%s  \n%s", props.rackhdUrl, rackhdUrl)
	}
	if len(props.groups) != 3 {
		t.Errorf("\n%d  \n%d", len(props.groups), 3)
	}
	if props.groups[2] != "test_group_2" {
		t.Errorf("\n%s  \n%s", props.groups[2], "test_group_2")
	}
	if props.filterGroup != "new" {
		t.Errorf("\n%s  \n%s", props.filterGroup, "new")
	}

	err = unsetEnvironmentVars(t)
	if err != nil {
		t.Errorf("%s\n", err)
	}
}

func TestHandleList(t *testing.T) {
	err := setEnvironmentVars(t)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	props := getPropsFromConfig()
	output, err := handleList(props)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	if value, ok := output["all"]; !ok {
		t.Errorf("Expected key 'all' got %s\n", value)
	}
	if value, ok := output["ungrouped"]; !ok {
		t.Errorf("Expected key 'ungrouped' got %s\n", value)
	}
	err = unsetEnvironmentVars(t)
	if err != nil {
		t.Errorf("%s\n", err)
	}
}
