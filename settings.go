package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

type Settings struct {
	SiteRoot      string
	TargetRoot    string
	TemplatesRoot string
	ArticlesRoot  string
}

const defaultSiteDir string = `site` // relative to CWD
const defaultTargetDirName string = `target`
const defaultTemplatesDirName string = `templates`
const articleDirName string = `posts`

var settings Settings

func PrepareEnvironment() Settings {
	if settings.SiteRoot == `` {
		siteRoot, targetName := readFromArgs()
		settings = Settings{
			SiteRoot:      siteRoot,
			TargetRoot:    filepath.Join(siteRoot, targetName),
			TemplatesRoot: filepath.Join(siteRoot, defaultTemplatesDirName),
			ArticlesRoot:  filepath.Join(siteRoot, articleDirName),
		}
	}
	createOrPurgeTargetFolder(settings.TargetRoot)
	log.Println("Environment: reading from", settings.SiteRoot, "and writing to", settings.TargetRoot)

	return settings
}

func createOrPurgeTargetFolder(dir string) {
	if _, err := os.Stat(dir); err == nil {
		log.Printf("Emptying target folder '%s'.\n", dir)
		for _, toBeRemoved := range ListFilesAndDirs(dir) {
			name := toBeRemoved.Name()
			PanicOnError(os.RemoveAll(filepath.Join(dir, name)))
		}
	} else {
		log.Printf("Creating non-existent target folder '%s'.\n", dir)
		err := os.Mkdir(dir, os.FileMode(0740))
		PanicOnError(err)
	}
}

func readFromArgs() (string, string) {
	var siteDir = flag.String("s", defaultSiteDir, "site directory (relative to CWD)")
	var targetDir = flag.String("t", defaultTargetDirName, "target directory name; will be created in site directory")

	flag.Parse()

	return *siteDir, *targetDir
}
