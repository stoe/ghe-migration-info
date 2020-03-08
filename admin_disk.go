package main

import (
	"context"
	"fmt"
	"strings"

	graphql "github.com/shurcooL/githubv4"
)

var (
	tQuery struct {
		Organizations struct {
			Nodes []struct {
				Repositories struct {
					TotalDiskUsage graphql.Int
				} `graphql:"repositories(first: 1)"`
			}
			PageInfo struct {
				EndCursor   graphql.String
				HasNextPage bool
			}
		} `graphql:"organizations(first: 100, after: $organizationsPage)"`
	}
)

const (
	b  = iota             // 0
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
	gb = 1 << (iota * 10) // 1 << (3 * 10)
	tb = 1 << (iota * 10) // 1 << (4 * 10)
)

func bytesConvert(b uint64) string {
	u := ""
	v := float32(b)

	switch {
	case b >= tb:
		u = "TB"
		v = v / tb
	case b >= gb:
		u = "GB"
		v = v / gb
	case b >= mb:
		u = "MB"
		v = v / mb
	case b >= kb:
		u = "kb"
		v = v / kb
	case b >= mb:
		u = "B"
	case b == gb:
		return "0"
	}

	s := strings.TrimSuffix(
		fmt.Sprintf("%.2v", v), ".00",
	)

	return fmt.Sprintf("%s %s", s, u)
}

func getTotalDiskUsage() string {
	variables := map[string]interface{}{
		"organizationsPage": (*graphql.String)(nil),
	}

	var t graphql.Int

	for {
		if err := graphqlClient.Query(context.Background(), &tQuery, variables); err != nil {
			panic(err)
		}

		for _, n := range tQuery.Organizations.Nodes {
			t += n.Repositories.TotalDiskUsage
		}

		// break on last page
		if !tQuery.Organizations.PageInfo.HasNextPage {
			break
		}

		variables["organizationsPage"] = graphql.NewString(tQuery.Organizations.PageInfo.EndCursor)
	}

	return fmt.Sprintf(`
%s...% 12s
`,
		bold("Size on disk"),
		// result is in kilobyte, but bytesConvert expects b
		bytesConvert(uint64(t*kb)),
	)
}
