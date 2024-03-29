package markdown

type Config struct {
	ReferenceURL        string
	ApiName             string
	TitleFormat         string
	SortFilePath        bool
	InlineProperties    bool
	OutputDir           string
	AllowExternalRefs   bool
	IncludeRestrictions bool
	IncludeExamples     bool
}
