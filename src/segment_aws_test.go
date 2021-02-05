package main

import (
	"errors"
	"testing"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

type awsArgs struct {
	enabled            bool
	separator          string
	profile            string
	displayRegion      bool
	displayProfileName bool
	config_provider    func() (awssdk.Config, error)
}

func setupAwsTests(args *awsArgs) *aws {

	env := new(MockedEnvironment)
	env.On("getenv", "AWS_PROFILE").Return(args.profile)

	props := &properties{
		values: map[Property]interface{}{
			Separator:          args.separator,
			DisplayRegion:      args.displayRegion,
			DisplayProfileName: args.displayProfileName,
		},
	}

	a := &aws{
		env:             env,
		props:           props,
		config_provider: args.config_provider,
	}

	return a
}

func TestAwsDisabled(t *testing.T) {
	args := &awsArgs{
		enabled: false,
	}

	aws := setupAwsTests(args)

	assert.False(t, aws.enabled())
}

func TestAwsEnabledAndNoAwsData(t *testing.T) {
	args := &awsArgs{
		enabled:            true,
		displayRegion:      true,
		displayProfileName: true,
		config_provider: func() (cfg awssdk.Config, err error) {
			cfg.Region = ""
			err = nil
			return
		},
	}

	aws := setupAwsTests(args)

	assert.False(t, aws.enabled())
}

func TestAwsEnabledAndErrorGettingAwsData(t *testing.T) {
	args := &awsArgs{
		enabled:            true,
		displayRegion:      true,
		displayProfileName: true,
		config_provider: func() (cfg awssdk.Config, err error) {
			err = errors.New("testing")
			return
		},
	}

	aws := setupAwsTests(args)

	assert.False(t, aws.enabled())
}

func TestAwsWriteRegion(t *testing.T) {
	expected := "@testRegion"

	args := &awsArgs{
		enabled:            true,
		separator:          "@",
		displayRegion:      true,
		displayProfileName: true,
		config_provider: func() (awssdk.Config, error) {
			cfg := &awssdk.Config{
				Region: "testRegion",
			}

			return *cfg, nil
		},
	}

	aws := setupAwsTests(args)

	assert.True(t, aws.enabled())
	assert.Equal(t, expected, aws.string())
}

func TestAwsWriteRegionAndProfile(t *testing.T) {
	expected := "testProfile@testRegion"

	args := &awsArgs{
		enabled:            true,
		displayRegion:      true,
		displayProfileName: true,
		separator:          "@",
		profile:            "testProfile",
		config_provider: func() (cfg awssdk.Config, err error) {
			cfg.Region = "testRegion"
			err = nil
			return
		},
	}

	aws := setupAwsTests(args)

	assert.True(t, aws.enabled())
	assert.Equal(t, expected, aws.string())
}

func TestAwsWriteRegionAndProfileWhenRegionDisabled(t *testing.T) {
	expected := "testProfile"

	args := &awsArgs{
		enabled:            true,
		displayRegion:      false,
		displayProfileName: true,
		separator:          "@",
		profile:            "testProfile",
		config_provider: func() (cfg awssdk.Config, err error) {
			cfg.Region = "testRegion"
			err = nil
			return
		},
	}

	aws := setupAwsTests(args)

	assert.True(t, aws.enabled())
	assert.Equal(t, expected, aws.string())
}

func TestAwsWriteRegionAndProfileWhenProfileDisabled(t *testing.T) {
	expected := "testRegion"

	args := &awsArgs{
		enabled:            true,
		displayRegion:      true,
		displayProfileName: false,
		separator:          "@",
		profile:            "testProfile",
		config_provider: func() (cfg awssdk.Config, err error) {
			cfg.Region = "testRegion"
			err = nil
			return
		},
	}

	aws := setupAwsTests(args)

	assert.True(t, aws.enabled())
	assert.Equal(t, expected, aws.string())
}
