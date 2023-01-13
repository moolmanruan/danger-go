package dangerJs

type GitCommit struct {
	SHA       string          `json:"sha"`
	Author    GitCommitAuthor `json:"author"`
	Committer GitCommitAuthor `json:"committer"`
	Message   string          `json:"message"`
	Tree      any             `json:"tree"`
	Parents   []string        `json:"parents,omitempty"`
	URL       string          `json:"url"`
}

type GitCommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}
