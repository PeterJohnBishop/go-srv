package integrations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Type              string `json:"type"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	SiteAdmin         bool   `json:"site_admin"`
	UserViewType      string `json:"user_view_type"`
	Login             string `json:"login"`
	Id                int    `json:"id"`
	FollowersUrl      string `json:"followers_url"`
	StarredUrl        string `json:"starred_url"`
	EventsUrl         string `json:"events_url"`
	ReposUrl          string `json:"repos_url"`
	NodeId            string `json:"node_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
}

type Permissions struct {
	Push     bool `json:"push"`
	Triage   bool `json:"triage"`
	Pull     bool `json:"pull"`
	Admin    bool `json:"admin"`
	Maintain bool `json:"maintain"`
}

type Repo struct {
	Name                     string      `json:"name"`
	CollaboratorsUrl         string      `json:"collaborators_url"`
	ContentsUrl              string      `json:"contents_url"`
	ArchiveUrl               string      `json:"archive_url"`
	PushedAt                 string      `json:"pushed_at"`
	GitUrl                   string      `json:"git_url"`
	CloneUrl                 string      `json:"clone_url"`
	WatchersCount            int         `json:"watchers_count"`
	Watchers                 int         `json:"watchers"`
	EventsUrl                string      `json:"events_url"`
	StargazersUrl            string      `json:"stargazers_url"`
	LabelsUrl                string      `json:"labels_url"`
	TreesUrl                 string      `json:"trees_url"`
	IssuesUrl                string      `json:"issues_url"`
	SshUrl                   string      `json:"ssh_url"`
	HasProjects              bool        `json:"has_projects"`
	ForksCount               int         `json:"forks_count"`
	Owner                    User        `json:"owner"`
	Description              string      `json:"description"`
	Url                      string      `json:"url"`
	StatusesUrl              string      `json:"statuses_url"`
	IssueCommentUrl          string      `json:"issue_comment_url"`
	MirrorUrl                any         `json:"mirror_url"`
	License                  any         `json:"license"`
	DownloadsUrl             string      `json:"downloads_url"`
	Size                     int         `json:"size"`
	Disabled                 bool        `json:"disabled"`
	AllowForking             bool        `json:"allow_forking"`
	NodeId                   string      `json:"node_id"`
	GitTagsUrl               string      `json:"git_tags_url"`
	Language                 string      `json:"language"`
	Id                       int         `json:"id"`
	AssigneesUrl             string      `json:"assignees_url"`
	ContributorsUrl          string      `json:"contributors_url"`
	CreatedAt                string      `json:"created_at"`
	StargazersCount          int         `json:"stargazers_count"`
	DefaultBranch            string      `json:"default_branch"`
	TeamsUrl                 string      `json:"teams_url"`
	MergesUrl                string      `json:"merges_url"`
	DeploymentsUrl           string      `json:"deployments_url"`
	SvnUrl                   string      `json:"svn_url"`
	Homepage                 any         `json:"homepage"`
	HooksUrl                 string      `json:"hooks_url"`
	PullsUrl                 string      `json:"pulls_url"`
	KeysUrl                  string      `json:"keys_url"`
	IssueEventsUrl           string      `json:"issue_events_url"`
	TagsUrl                  string      `json:"tags_url"`
	CompareUrl               string      `json:"compare_url"`
	HasDownloads             bool        `json:"has_downloads"`
	HasPages                 bool        `json:"has_pages"`
	OpenIssuesCount          int         `json:"open_issues_count"`
	Permissions              Permissions `json:"permissions"`
	SubscribersUrl           string      `json:"subscribers_url"`
	SubscriptionUrl          string      `json:"subscription_url"`
	IsTemplate               bool        `json:"is_template"`
	Forks                    int         `json:"forks"`
	FullName                 string      `json:"full_name"`
	BlobsUrl                 string      `json:"blobs_url"`
	LanguagesUrl             string      `json:"languages_url"`
	MilestonesUrl            string      `json:"milestones_url"`
	HasIssues                bool        `json:"has_issues"`
	Topics                   []any       `json:"topics"`
	Private                  bool        `json:"private"`
	BranchesUrl              string      `json:"branches_url"`
	GitRefsUrl               string      `json:"git_refs_url"`
	CommitsUrl               string      `json:"commits_url"`
	NotificationsUrl         string      `json:"notifications_url"`
	UpdatedAt                string      `json:"updated_at"`
	ForksUrl                 string      `json:"forks_url"`
	GitCommitsUrl            string      `json:"git_commits_url"`
	WebCommitSignoffRequired bool        `json:"web_commit_signoff_required"`
	Fork                     bool        `json:"fork"`
	CommentsUrl              string      `json:"comments_url"`
	ReleasesUrl              string      `json:"releases_url"`
	HasDiscussions           bool        `json:"has_discussions"`
	HtmlUrl                  string      `json:"html_url"`
	HasWiki                  bool        `json:"has_wiki"`
	Archived                 bool        `json:"archived"`
	Visibility               string      `json:"visibility"`
	OpenIssues               int         `json:"open_issues"`
}

type Commit struct {
	Committer    Committer    `json:"committer"`
	Message      string       `json:"message"`
	Tree         Tree         `json:"tree"`
	CommentCount int          `json:"comment_count"`
	Verification Verification `json:"verification"`
	Url          string       `json:"url"`
	Author       User         `json:"author"`
}

type Parent struct {
	Url string `json:"url"`
	Sha string `json:"sha"`
}

type CommitData struct {
	Url         string    `json:"url"`
	CommentsUrl string    `json:"comments_url"`
	Commit      Commit    `json:"commit"`
	Parents     []Parent  `json:"parents"`
	Sha         string    `json:"sha"`
	NodeId      string    `json:"node_id"`
	HtmlUrl     string    `json:"html_url"`
	Author      User      `json:"author"`
	Committer   Committer `json:"committer"`
}

type Committer struct {
	Login             string `json:"login"`
	FollowingUrl      string `json:"following_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	SiteAdmin         bool   `json:"site_admin"`
	GravatarId        string `json:"gravatar_id"`
	StarredUrl        string `json:"starred_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	ReposUrl          string `json:"repos_url"`
	Type              string `json:"type"`
	Id                int    `json:"id"`
	GistsUrl          string `json:"gists_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
}

type Tree struct {
	Url string `json:"url"`
	Sha string `json:"sha"`
}

type Verification struct {
	Signature  any    `json:"signature"`
	Payload    any    `json:"payload"`
	VerifiedAt any    `json:"verified_at"`
	Verified   bool   `json:"verified"`
	Reason     string `json:"reason"`
}

func GetAuthenticatedUserRepos(c *gin.Context) {
	token := os.Getenv("GITHUB")
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Github authentication token is missing"})
		return
	}

	url := "https://api.github.com/user/repos"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to connect to GitHub"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "GitHub API rejected the request"})
		return
	}

	var repos []Repo

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse GitHub response"})
		return
	}

	c.JSON(http.StatusOK, repos)
}

func GetCommits(owner string, repo string, c *gin.Context) {
	token := os.Getenv("GITHUB")
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Github authentication token is missing"})
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to connect to GitHub"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "GitHub API rejected the request"})
		return
	}

	var commits []CommitData
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse GitHub response"})
		return
	}

	c.JSON(http.StatusOK, commits)
}

func GetCommitByRef(owner string, repo string, ref string, c *gin.Context) {
	token := os.Getenv("GITHUB")
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Github authentication token is missing"})
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s", owner, repo, ref)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to connect to GitHub"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "GitHub API rejected the request"})
		return
	}

	var commit CommitData
	if err := json.NewDecoder(resp.Body).Decode(&commit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse GitHub response"})
		return
	}

	c.JSON(http.StatusOK, commit)
}
