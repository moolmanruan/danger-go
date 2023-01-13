package dangerJs

type DSL struct {
	Git    Git    `json:"git"`
	GitHub GitHub `json:"github,omitempty"`
	GitLab GitLab `json:"gitlab,omitempty"`
	// TODO: bitbucket_server
	// TODO: bitbucket_cloud
	Settings Settings `json:"settings"`
}

type FilePath = string

type Git struct {
	ModifiedFiles []FilePath  `json:"modified_files"`
	CreateFiles   []FilePath  `json:"created_files"`
	DeletedFiles  []FilePath  `json:"deleted_files"`
	Commits       []GitCommit `json:"commits"`
}

type Settings struct {
	GitHub struct {
		AccessToken       string `json:"accessToken"`
		BaseURL           string `json:"baseURL"`
		AdditionalHeaders any    `json:"additionalHeaders"`
	} `json:"github"`
	CLIArgs CLIArgs `json:"cliArgs"`
}

type CLIArgs struct {
	Base               string `json:"base"`
	Verbose            string `json:"verbose"`
	ExternalCIProvider string `json:"externalCiProvider"`
	TextOnly           bool   `json:"textOnly"` // JS has this as string
	Dangerfile         string `json:"dangerfile"`
	ID                 string `json:"id"`
	Staging            bool   `json:"staging"`
}
