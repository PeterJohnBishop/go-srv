package integrations

import "github.com/gin-gonic/gin"

func AddGitHubRoutes(r *gin.Engine) {
	gh := r.Group("/gh")
	{
		gh.GET("/repos", func(c *gin.Context) {
			GetAuthenticatedUserRepos(c)
		})
		gh.GET("/:owner/:repo/commits", func(c *gin.Context) {
			owner := c.Param("owner")
			repo := c.Param("repo")
			GetCommits(owner, repo, c)
		})
		gh.GET("/:owner/:repo/commmit/:ref", func(c *gin.Context) {
			owner := c.Param("owner")
			repo := c.Param("repo")
			ref := c.Param("ref")
			GetCommitByRef(owner, repo, ref, c)
		})
	}
}
