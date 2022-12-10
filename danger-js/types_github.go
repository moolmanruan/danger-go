package danger_js

import "time"

type GitHub struct {
	Issue              GitHubIssue              `json:"issue"`
	PR                 GitHubPR                 `json:"pr"`
	Commits            []GitHubCommit           `json:"commits"`
	Reviews            []GitHubReview           `json:"reviews"`
	RequestedReviewers GitHubRequestedReviewers `json:"requested_reviewers"`
	ThisPR             GitHubThisPR             `json:"thisPR"` // Note: Not include in danger-kotlin
}

type GitHubIssue struct {
	ID           int64              `json:"id"`
	Number       int64              `json:"number"`
	Title        string             `json:"title"`
	User         GitHubUser         `json:"user"`
	State        string             `json:"state"` // "closed" | "open" | "locked"
	Locked       bool               `json:"locked"`
	Body         string             `json:"body,omitempty"`
	CommentCount int                `json:"comments"`
	Assignee     GitHubUser         `json:"assignee,omitempty"`
	Assignees    []GitHubUser       `json:"assignees"`
	Milestone    GitHubMilestone    `json:"milestone"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	ClosedAt     time.Time          `json:"closed_at"`
	Labels       []GitHubIssueLabel `json:"labels"`
}

type GitHubIssueLabel struct {
	ID    int64  `json:"id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type GitHubUser struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Type      string `json:"type"` // "User" | "Organization" | "Bot"
	AvatarURL string `json:"avatar_url"`
}

type GitHubPR struct {
	ID                 int64           `json:"id"`
	Number             int             `json:"number"`
	Title              string          `json:"title"`
	Body               string          `json:"body"`
	User               GitHubUser      `json:"user"`
	Assignee           GitHubUser      `json:"assignee"`
	Assignees          []GitHubUser    `json:"assignees"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	ClosedAt           time.Time       `json:"closed_at,omitempty"`
	MergedAt           time.Time       `json:"merged_at,omitempty"`
	Head               GitHubMergeRef  `json:"head"`
	Base               GitHubMergeRef  `json:"base"`
	State              string          `json:"state"` // "closed" | "open" | "merged" | "locked"
	Locked             bool            `json:"locked"`
	Merged             bool            `json:"merged"`
	CommitCount        int             `json:"commits"`
	CommentCount       int             `json:"comments"`
	ReviewCommentCount int             `json:"review_comments"`
	Additions          int             `json:"additions"`
	Deletions          int             `json:"deletions"`
	ChangedFiles       int             `json:"changed_files"`
	Milestone          GitHubMilestone `json:"milestone"`
	HTMLURL            string          `json:"html_url"`
}
type GitHubMergeRef struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	SHA   string     `json:"sha"`
	User  GitHubUser `json:"user"`
	Repo  GitHubRepo `json:"repo"`
}

type GitHubRepo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	IsPrivate   bool   `json:"private"`
	Description string `json:"description,omitempty"`
	IsFork      bool   `json:"fork"`
	HTMLURL     string `json:"html_url"`
}

type GitHubCommit struct {
	SHA       string     `json:"sha"`
	URL       string     `json:"url"`
	Author    GitHubUser `json:"author"`
	Commit    Commit     `json:"commit"` // Should we create a custom Commit struct?
	Committer GitHubUser `json:"committer"`
}

type GitHubReview struct {
	User     GitHubUser `json:"user"`
	ID       int64      `json:"id"`
	Body     string     `json:"body"`
	CommitID string     `json:"commit_id"`
	State    string     `json:"state"` // APPROVED | CHANGES_REQUESTED | COMMENTED | PENDING | DISMISSED
}

type GitHubRequestedReviewers struct {
	Users []string `json:"users"`
	Teams []string `json:"teams"`
}

type GitHubMilestone struct {
	ID           int64      `json:"id"`
	Number       int64      `json:"number"`
	State        string     `json:"state"` // "closed" | "open" | "all"
	Title        string     `json:"title"`
	Description  string     `json:"description,omitempty"`
	Creator      GitHubUser `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	ClosedAt     time.Time  `json:"closed_at"`
	DueOn        time.Time  `json:"due_on"`
}

type GitHubThisPR struct {
	Number int    `json:"number"`
	Repo   string `json:"repo"`
	Owner  string `json:"owner"`
}
