package main

import (
	"fmt"

	"github.com/google/go-github/v29/github"
)

func getAdminStats() (o string) {
	stats, _, err := restClient.Admin.GetAdminStats(ctx)

	if err != nil {
		errorAndExit(err)
	}

	o += getOrgStats(stats.GetOrgs())
	o += getUserStats(stats.GetUsers())
	o += getRepoStats(stats.GetRepos())
	o += getIssueStats(stats.GetIssues())
	o += getMilestoneStats(stats.GetMilestones())
	o += getPullRequestStats(stats.GetPulls())
	o += getCommentStats(stats.GetComments())
	o += getGistStats(stats.GetGists())
	o += getPageStats(stats.GetPages())
	o += getHookStats(stats.GetHooks())

	return
}

func getOrgStats(o *github.OrgStats) string {
	return fmt.Sprintf(`%s
Total..........% 12v
Disabled.......% 12v

%s
Total..........% 12v
Members........% 12v
`,
		bold("Organizations"),
		o.GetTotalOrgs(),
		o.GetDisabledOrgs(),
		bold("Teams"),
		o.GetTotalTeams(),
		o.GetTotalTeamMembers(),
	)
}

func getUserStats(u *github.UserStats) string {
	return fmt.Sprintf(`
%s
Total..........% 12v
Admins.........% 12v
Suspended......% 12v
`,
		bold("Users"),
		u.GetTotalUsers(),
		u.GetAdminUsers(),
		u.GetSuspendedUsers(),
	)
}

func getRepoStats(r *github.RepoStats) string {
	total := r.GetTotalRepos()
	root := r.GetRootRepos()
	forks := r.GetForkRepos()
	orgs := r.GetOrgRepos()
	users := (total - orgs)
	pushes := r.GetTotalPushes()
	wikis := r.GetTotalWikis()

	return fmt.Sprintf(`
%s
Total..........% 12v
Roots..........% 12v
Forks..........% 12v

Organization...% 12v
User...........% 12v

Pushes.........% 12v

Wikis..........% 12v
`,
		bold("Repositories"),
		total,
		root,
		forks,
		orgs,
		users,
		pushes,
		wikis,
	)
}

func getIssueStats(i *github.IssueStats) string {
	return fmt.Sprintf(`
%s
Total..........% 12v
Open...........% 12v
Closed.........% 12v
`,
		bold("Issues"),
		i.GetTotalIssues(),
		i.GetOpenIssues(),
		i.GetClosedIssues(),
	)
}

func getMilestoneStats(m *github.MilestoneStats) string {
	return fmt.Sprintf(`
%s
Total..........% 12v
Open...........% 12v
Closed.........% 12v
`,
		bold("Milestones"),
		m.GetTotalMilestones(),
		m.GetOpenMilestones(),
		m.GetClosedMilestones(),
	)
}

func getPullRequestStats(p *github.PullStats) string {
	return fmt.Sprintf(`
%s
Total..........% 12v
Open...........% 12v
Merged.........% 12v
Unmergable.....% 12v
`,
		bold("Pull Requests"),
		p.GetTotalPulls(),
		p.GetMergablePulls(),
		p.GetMergedPulls(),
		p.GetUnmergablePulls(),
	)
}

func getCommentStats(c *github.CommentStats) string {
	return fmt.Sprintf(`
%s
Issues.........% 12v
Pull Requests..% 12v
Commit.........% 12v
Gists..........% 12v
`,
		bold("Comments"),
		c.GetTotalIssueComments(),
		c.GetTotalPullRequestComments(),
		c.GetTotalCommitComments(),
		c.GetTotalGistComments(),
	)
}

func getGistStats(g *github.GistStats) string {
	return fmt.Sprintf(`
%s
Total..........% 12v
Public.........% 12v
Private........% 12v
`,
		bold("Gists"),
		g.GetTotalGists(),
		g.GetPublicGists(),
		g.GetPrivateGists(),
	)
}

func getPageStats(p *github.PageStats) string {
	return fmt.Sprintf(`
%s
Total..........% 12v
`,
		bold("Pages"),
		p.GetTotalPages(),
	)
}

func getHookStats(h *github.HookStats) string {
	return fmt.Sprintf(`
%s
Total..........% 12v
Active.........% 12v
Inactive.......% 12v
`,
		bold("Hooks"),
		h.GetTotalHooks(),
		h.GetActiveHooks(),
		h.GetInactiveHooks(),
	)
}
