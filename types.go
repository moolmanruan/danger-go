package danger

type Results struct {
	Fails    []Violation `json:"fails"`
	Warnings []Violation `json:"warnings"`
	Messages []Violation `json:"messages"`
	// Markdowns are messages to attach at the bottom of the comment
	Markdowns []Violation `json:"markdowns"`

	GitHub *GitHubResults `json:"github,omitempty"`
	Meta   *MetaResults   `json:"meta,omitempty"`
}

type Violation struct {
	Message string `json:"message"`
	File    string `json:"file,omitempty"`
	Line    int    `json:"line,omitempty"`
	// Icon is an optional icon for table (Only valid for messages).
	Icon string `json:"icon,omitempty"`
}

type GitHubResults struct {
	// StepSummary is Markdown text which gets added as a summary in the first
	// page which you see when you click through to the PR results.
	StepSummary string `json:"stepSummary,omitempty"`
}

type MetaResults struct {
	// RuntimeRef e.g. "https://danger.systems/js"
	RuntimeRef string `json:"runtimeRef"`
	// RuntimeName E.g. "dangerJS", or "Danger Swift"
	RuntimeName string `json:"runtimeName"`
}
