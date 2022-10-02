package github

import (
	"context"

	"github.com/google/go-github/v47/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

const (
	packageType = "container"
)

type Client struct {
	github *github.Client
	dryrun bool
}

func NewClient(token string, dryrun bool) *Client {
	ctx := context.TODO()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	return &Client{
		github: github.NewClient(tc),
		dryrun: dryrun,
	}
}

func (c *Client) CleanupPackages(username, packageName string, minRetain int) error {
	var (
		active  = "active"
		page    = 1
		perPage = 1
		allpkgs = []*github.PackageVersion{}
	)

	for {
		pkgs, resp, err := c.github.Users.PackageGetAllVersions(context.TODO(), username, packageType, packageName, &github.PackageListOptions{
			State: &active,
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: perPage,
			},
		})
		if err != nil {
			return err
		}

		allpkgs = append(allpkgs, pkgs...)
		if resp.NextPage == 0 {
			break
		}

		page = resp.NextPage
	}

	if len(allpkgs) <= minRetain {
		logrus.Infof("found %d versions for %s/%s which is less than min retain configuration value of %d", len(allpkgs), username, packageName, minRetain)
		return nil
	}

	dryRunMsg := ""
	if c.dryrun {
		dryRunMsg = " specify --yes to actually delete"
	}

	logrus.Infof("will be deleting %d versions for %s/%s.%s", len(allpkgs)-minRetain, username, packageName, dryRunMsg)
	for i, pkg := range allpkgs {
		if i < minRetain {
			continue
		}

		logrus.Debugf("deleting %#v\n", pkg.GetMetadata().Container.Tags)

		if c.dryrun {
			continue
		}

		_, err := c.github.Users.PackageDeleteVersion(context.TODO(), username, packageType, packageName, pkg.GetID())
		if err != nil {
			return err
		}
	}

	return nil
}
