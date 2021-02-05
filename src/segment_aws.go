package main

import (
	"context"
	"fmt"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type aws struct {
	props           *properties
	env             environmentInfo
	region          string
	profile_name    string
	config_provider func() (awssdk.Config, error)
}

const (
	// DisplayProfileName hides or shows the profile name
	DisplayProfileName Property = "display_profile_name"
	// DisplayRegion hides or show the region id
	DisplayRegion Property = "display_region"
	// Separator is put between the region and profile
	Separator Property = "separator"
)

func (a *aws) init(props *properties, env environmentInfo) {
	a.props = props
	a.env = env

	a.config_provider = func() (cfg awssdk.Config, err error) {
		return config.LoadDefaultConfig(context.TODO())
	}
}

func (a *aws) string() string {
	separator := ""
	if a.regionEnabled() && a.profileEnabled() {
		separator = a.props.getString(Separator, fmt.Sprintf("%s", "<#F8991D>\uf52c</>"))
	}

	return fmt.Sprintf("%s%s%s", a.getProfile(), separator, a.getRegion())
}

func (a *aws) enabled() bool {

	if !a.regionEnabled() && !a.profileEnabled() {
		return false
	}

	cfg, err := a.config_provider()

	if err != nil || cfg.Region == "" {
		return false
	}

	a.region = cfg.Region
	a.profile_name = a.env.getenv("AWS_PROFILE")

	return a.region != "" || a.profile_name != ""
}

func (a *aws) getRegion() string {
	if !a.regionEnabled() {
		return ""
	}

	return a.region
}

func (a *aws) getProfile() string {
	if !a.profileEnabled() {
		return ""
	}

	return a.profile_name
}

func (a *aws) regionEnabled() bool {
	return a.props.getBool(DisplayRegion, true)
}

func (a *aws) profileEnabled() bool {
	return a.props.getBool(DisplayProfileName, true)
}
