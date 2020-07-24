package pages

import (
	"log"
	"os"
	"path/filepath"
)

const defaultSiteDir string = "site" // relative to CWD
const defaultTargetDirName string = "target"
const articleDirName string = `posts`

func GetSiteRootDirectory() string {
	siteDir := defaultSiteDir
	if len(os.Args) == 2 {
		siteDir = os.Args[1]
		log.Printf("Using site directory '%s'.", siteDir)
	} else if len(os.Args) > 2 {
		log.Printf("Usage: bloggo [path-to-site-directory]\n\nIf omitted, site is expected at $CWD/%s", defaultSiteDir)
		os.Exit(0)
	} else {
		log.Printf("Using default site directory '%s'.", defaultSiteDir)
	}
	return siteDir
}

func GetTargetRootDirectory() string {
	return filepath.Join(GetSiteRootDirectory(), defaultTargetDirName)
}

func GetTemplatesRootDirectory() string {
	return filepath.Join(GetSiteRootDirectory(), `templates`)
}

func GetArticlesRootDirectory() string {
	return filepath.Join(GetSiteRootDirectory(), articleDirName)
}
