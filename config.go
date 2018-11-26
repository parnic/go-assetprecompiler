package precompiler

// Config holds information on how to run the compiler
type Config struct {
	Files     []string
	Minify    bool
	OutputDir string
}
