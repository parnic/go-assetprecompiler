package precompiler

func main() {
	Compile(Config{
		Files: []string{
			"assets/css/bootstrap.min.default.css",
			"assets/css/font-awesome.min.css",
			"assets/css/tempusdominus-bootstrap-4.min.css",
			"assets/css/roboto.css",
			"assets/css/application.css",
			"assets/js/jquery.min.js",
			"assets/js/popper.min.js",
			"assets/js/bootstrap.min.js",
			"assets/js/linkify.min.js",
			"assets/js/linkify-jquery.min.js",
			"assets/js/moment.min.js",
			"assets/js/tempusdominus-bootstrap-4.min.js",
			"assets/js/application.js",
		},
		Minify:    true,
		OutputDir: "assets/",
	})
}
