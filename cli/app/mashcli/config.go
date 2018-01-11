package mashcli

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/tcnksm/go-input"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func New(userid string, password string, apikey string, apikeysecret string, name string, area string, tm string, ccurl string) *Config {
	return &Config{
		UserId:       userid,
		Password:     password,
		ApiKey:       apikey,
		ApiKeySecret: apikeysecret,
		Name:         name,
		Area:         area,
		Tm:           tm,
		CcUrl:        ccurl,
	}
}

func (c *Config) Save() error {

	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	filename := c.Name + ".config"

	configPath := filepath.Join(UserHomeDir(), "/.mashcli")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.MkdirAll(configPath,os.ModePerm)
	}

	filePathname := filepath.Join(configPath,filename)
	return ioutil.WriteFile(filePathname, bytes, 0644)
}

func Load(name string) (*Config, error) {

	var filename, filePathname="",""

	configPath := filepath.Join(UserHomeDir(), "/.mashcli")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("warning: unable to local config directory. Please run 'mashcli config add' to add first configuration.")
		return nil, nil
	}

	if len(name) > 0 {
		filename = name + ".config"
	} else {
		// Default configuration
		filename = "mashcli.config"
	}

	filePathname = filepath.Join(configPath,filename)

	bytes, err := ioutil.ReadFile(filePathname)
	c := new(Config)

	if err != nil {
		return nil, fmt.Errorf("config: unable to load configuration :  %s ", filePathname)
	}

	err = json.Unmarshal(bytes, c)
	if err != nil {
		return nil, fmt.Errorf("config: unable to load configuration :  %s ", filePathname)
	}

	return c,nil
}

func (c *Config) PrettyPrint() {

	data := []string{c.Name,c.Area,c.Tm,c.CcUrl,c.ApiKey,c.ApiKeySecret,c.UserId,c.Password}
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Name", "Area", "Traffic Manager", "CC URL", "API Key", "API Secret", "User ID", "User Password"})
	table.Append(data)
	table.Render()

	return

}


func Add() *Config {

	ui := &input.UI{}

	name, err := ui.Ask("Configuration Name?", &input.Options{
		Default:  "mashcli",
		Required: true,
	})
	if err != nil {
		return nil
	}

	area, err := ui.Ask("Area ID?", &input.Options{
		Default:  "",
		Required: true,
	})
	if err != nil {
		return nil
	}

	tm, err := ui.Ask("Traffic Manager?", &input.Options{
		Default:  "",
		Required: true,
	})
	if err != nil {
		return nil
	}

	ccurl, err := ui.Ask("Control Centre URL?", &input.Options{
		Default:  "https://<<area>>.admin.mashery.com/control-center",
		Required: true,
	})
	if err != nil {
		return nil
	}

	apikey, err := ui.Ask("API Key?", &input.Options{
		Default:  "",
		Required: true,
	})
	if err != nil {
		return nil
	}

	apikeysecret, err := ui.Ask("API Key Secret?", &input.Options{
		Default:  "",
		Required: true,
	})
	if err != nil {
		return nil
	}

	userid, err := ui.Ask("User ID?", &input.Options{
		Default:  "",
		Required: true,
	})
	if err != nil {
		return nil
	}

	password, err := ui.Ask("User Password?", &input.Options{
		Default:  "",
		Required: true,
	})
	if err != nil {
		return nil
	}

	c := New(userid, password, apikey, apikeysecret, name, area, tm, ccurl)
	c.Save()

	return c

}
