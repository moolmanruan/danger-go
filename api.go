package danger

type T struct {
	Results Results
}

func New() *T {
	return &T{
		Results: Results{
			Fails:     []Violation{},
			Messages:  []Violation{},
			Warnings:  []Violation{},
			Markdowns: []Violation{},
		},
	}
}

// Message adds the message to the Danger table. The only difference between
// this and Warn is the emoji which shows in the table.
func (s *T) Message(message string, file string, line int) {
	s.Results.Messages = append(s.Results.Messages,
		Violation{
			Message: message,
			File:    file,
			Line:    line,
		})
}

// Warn adds the message to the Danger table. The message highlights
// low-priority issues, but does not fail the build.
func (s *T) Warn(message string, file string, line int) {
	s.Results.Warnings = append(s.Results.Warnings,
		Violation{
			Message: message,
			File:    file,
			Line:    line,
		})
}

// Fail a build, outputting a specific reason for failing into an HTML table.
func (s *T) Fail(message string, file string, line int) {
	s.Results.Fails = append(s.Results.Fails,
		Violation{
			Message: message,
			File:    file,
			Line:    line,
		})
}

// Markdown adds the message as raw markdown into the Danger comment, under the
// table.
func (s *T) Markdown(message string, file string, line int) {
	s.Results.Markdowns = append(s.Results.Markdowns,
		Violation{
			Message: message,
			File:    file,
			Line:    line,
		})
}
