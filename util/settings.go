package util

import (
	"flag"
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

func GetSettings() Settings {
	if settings.SiteRoot == `` {
		siteRoot, targetName := readFromArgs()
		settings = Settings{
			SiteRoot:      siteRoot,
			TargetRoot:    filepath.Join(siteRoot, targetName),
			TemplatesRoot: filepath.Join(siteRoot, defaultTemplatesDirName),
			ArticlesRoot:  filepath.Join(siteRoot, articleDirName),
		}
	}
	return settings
}

func readFromArgs() (string, string) {
	var siteDir = flag.String("s", defaultSiteDir, "site directory (relative to CWD)")
	var targetDir = flag.String("t", defaultTargetDirName, "target directory name; will be created in site directory")

	flag.Parse()

	return *siteDir, *targetDir
}
