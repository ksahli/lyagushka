package initialize

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Run(path string) error {
	if err := initializeSrc(path); err != nil {
		return err
	}
	if err := initializeTpl(path); err != nil {
		return err
	}
	if err := initialize(path, "dst"); err != nil {
		return err
	}
	return nil
}

func initializeSrc(path string) error {
	if err := initialize(path, "src"); err != nil {
		return err
	}
	homePath := filepath.Join(path, "src", "home.html")
	homeSrc := `
		{{ define "main" }}
			This is the home page
		{{ end }}
	`
	homeContent := []byte(homeSrc)
	if err := ioutil.WriteFile(homePath, homeContent, 0664); err != nil {
		return err
	}
	return nil
}

func initializeTpl(path string) error {
	if err := initialize(path, "tpl"); err != nil {
		return err
	}
	basePath := filepath.Join(path, "tpl", "base.html")
	baseSrc := `
		<html>
			<head>
				<title>site title</title>
			</head>
			<body>
				<header>site header</header>
				<section>
					<div> Welcome !<div>
					{{ block "main" . }}{{ end }}
				</section>
				<footer>site footer</footer>
			</body>
		</html>`
	baseContent := []byte(baseSrc)
	if err := ioutil.WriteFile(basePath, baseContent, 0664); err != nil {
		return err
	}
	return nil
}

func initialize(path, name string) error {
	fullpath := filepath.Join(path, name)
	if err := os.Mkdir(fullpath, os.ModePerm); err != nil {
		return err
	}
	return nil
}
