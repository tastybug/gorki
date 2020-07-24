package pages

import (
	"log"
	"os"
	"path/filepath"
)

const defaultSiteDir string = `site` // relative to CWD
const defaultTargetDirName string = `target`
const defaultTemplatesDirName string = `templates`
const articleDirName string = `posts`

var siteRootPath = ``

func GetSiteRootDirectory() string {
	if siteRootPath == `` {
		siteRootPath = defaultSiteDir
		if len(os.Args) == 2 {
			siteRootPath = os.Args[1]
			log.Printf("Using site directory '%s'.", siteRootPath)
		} else if len(os.Args) > 2 {
			log.Printf("Usage: bloggo [path-to-site-directory]\n\nIf omitted, site is expected at $CWD/%s", defaultSiteDir)
			os.Exit(0)
		} else {
			log.Printf("Using default site directory '%s'.", defaultSiteDir)
		}
	}
	return siteRootPath
}

func GetTargetRootDirectory() string {
	return filepath.Join(GetSiteRootDirectory(), defaultTargetDirName)
}

func GetTemplatesRootDirectory() string {
	return filepath.Join(GetSiteRootDirectory(), defaultTemplatesDirName)
}

func GetArticlesRootDirectory() string {
	return filepath.Join(GetSiteRootDirectory(), articleDirName)
}
