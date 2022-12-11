package danger_js

type PR struct {
	Danger struct {
		Git      Git      `json:"git"`
		Github   GitHub   `json:"github,omitempty"`
		Gitlab   Gitlab   `json:"gitlab,omitempty"`
		Settings Settings `json:"settings"`
	} `json:"danger"`
}

type FilePath = string

type Git struct {
	ModifiedFiles []FilePath `json:"modified_files"`
	CreateFiles   []FilePath `json:"created_files"`
	DeletedFiles  []FilePath `json:"deleted_files"`
	Commits       []Commit   `json:"commits"`
}

type Commit struct {
	SHA       string       `json:"sha,omitempty"`
	Parents   []string     `json:"parents"`
	Author    CommitAuthor `json:"author"`
	Committer CommitAuthor `json:"committer"`
	Message   string       `json:"message"`
	//Tree      struct {
	//	SHA string `json:"sha"`
	//	URL string `json:"url"`
	//} `json:"tree"`
	Tree string `json:"tree"` // local
	URL  string `json:"url"`
}

type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"` //time.Time
}

type Settings struct {
	Github struct {
		AccessToken       string   `json:"accessToken"`
		AdditionalHeaders struct{} `json:"additionalHeaders"`
	} `json:"github"`
	CLIArgs struct{} `json:"cliArgs"`
}
