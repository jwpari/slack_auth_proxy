package main

import (
	"github.com/jwpari/slack_auth_proxy/slack"
)

func NewValidator() func(*slack.Auth, *UpstreamConfiguration) bool {
	validator := func(auth *slack.Auth, upstream *UpstreamConfiguration) bool {
		return len(upstream.Users) == 0 || upstream.FindUsername(auth.User.Id) != ""
	}
	return validator
}
