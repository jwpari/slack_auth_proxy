package main

import (
	"net/url"
)

type UpstreamConfiguration struct {
	Host    string   `yaml:"host,omitempty"`
	HostURL *url.URL `yaml:"-"`

	Users  []string `yaml:"users,omitempty"`
	Groups []string `yaml:"groups,omitempty"`
}

type UpstreamConfigurationMap map[string]*UpstreamConfiguration

func (configMap UpstreamConfigurationMap) Find(pattern string) *UpstreamConfiguration {
	upstreamConfig := configMap[pattern]
	if upstreamConfig == nil {
		upstreamConfig = configMap["/"]
	}


	return upstreamConfig
}

func (c *UpstreamConfiguration) Parse() (err error) {
	c.HostURL, err = url.Parse(c.Host)

	return
}

func (c *UpstreamConfiguration) FindUsername(name string) string {
	var user = ""

	for _, u := range c.Users {
		if u == name {
			user = u
			break

		}
	}

	return user
}
