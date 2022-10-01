package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

type Client struct {
	github *github.Client
	dryrun bool
}

func NewClient(dryrun bool) *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)

	tc := oauth2.NewClient(ctx, ts)
	return &Client{
		github: github.NewClient(tc),
		dryrun: dryrun,
	}
}

func (c *Client) CleanupPackages(packageName string, minRetain int) error {
	active := "active"
	pkgs, _, err := c.github.Users.PackageGetAllVersions(context.TODO(), "rajatjindal", "container", packageName, &github.PackageListOptions{
		State: &active,
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	})
	if err != nil {
		return err
	}

	for i, pkg := range pkgs {
		if i < minRetain {
			continue
		}

		fmt.Printf("deleting %#v\n", pkg.GetMetadata().Container.Tags)

		if c.dryrun {
			continue
		}

		_, err = c.github.Users.PackageDeleteVersion(context.TODO(), "rajatjindal", "container", packageName, pkg.GetID())
		if err != nil {
			return err
		}
	}

	return nil
}
