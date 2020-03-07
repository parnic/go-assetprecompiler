package precompiler

// Config holds information on how to run the compiler
type Config struct {
	// Files is a list of files to precompile together
	Files []string
	// Minify specifies whether minification should happen along with concatenation
	Minify bool
	// OutputDir if specified will cause the result to be written to this directory
	OutputDir string
}
