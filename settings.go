package main

import (
	"errors"
	"flag"
	"io/ioutil"
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

func newSettings() (Settings, error) {
	siteRoot, targetName := readFromArgs()
	settings = Settings{
		SiteRoot:      siteRoot,
		TargetRoot:    filepath.Join(siteRoot, targetName),
		TemplatesRoot: filepath.Join(siteRoot, defaultTemplatesDirName),
		ArticlesRoot:  filepath.Join(siteRoot, articleDirName),
	}

	if !PathExists(settings.SiteRoot) {
		return settings, errors.New("Site root expected at '" + settings.SiteRoot + "' but path does not exist.")
	}
	if !PathExists(settings.TemplatesRoot) {
		return settings, errors.New("Templates expected at '" + settings.TemplatesRoot + "' but path does not exist.")
	}
	if !PathExists(settings.ArticlesRoot) {
		return settings, errors.New("Articles expected at '" + settings.ArticlesRoot + "' but path does not exist.")
	}

	createOrPurgeTargetFolder(settings.TargetRoot)
	log.Println("Environment: reading from", settings.SiteRoot, "and writing to", settings.TargetRoot)

	return settings, nil
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

func ListFilesAndDirs(dir string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)
	return allFiles
}

func readFromArgs() (string, string) {

	var siteDir = flag.String("s", defaultSiteDir, "site directory")
	var targetDir = flag.String("t", defaultTargetDirName, "output directory name; will be purged if already existing")

	flag.Parse()

	return *siteDir, *targetDir
}
