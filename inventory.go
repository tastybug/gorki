package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func collectAllBundles(settings Settings) []bundle {
	pages := collectUsableArticleBundles(settings)
	pages = append(pages, collectStaticBundles(settings)...)
	return pages
}

func collectUsableArticleBundles(settings Settings) []bundle {
	articlesRootPath := settings.ArticlesRoot
	var bundles []bundle
	for _, bundle := range ListDirectories(articlesRootPath) {
		page, err := newBundle(articlesRootPath, bundle.Name())
		if err != nil {
			log.Println("Skipping broken bucket:", err)
		} else if page.isToBeRendered() {
			bundles = append(bundles, page)
		} else {
			log.Println("Skipping bucket not to be rendered:", bundle.Name())
		}
	}
	return bundles
}

func collectStaticBundles(settings Settings) []bundle {
	templatesFolderPath := settings.TemplatesRoot
	return []bundle{
		{
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesFolderPath, `index`),
				`index`,
				`index.html`,
				``,
				`index.html`},
		},
		{
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesFolderPath, `about`),
				`about`,
				`about.html`,
				`about`,
				`about.html`},
		},
		{
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesFolderPath, `privacy-imprint`),
				`privacy-imprint`,
				`privacy-imprint.html`,
				`privacy-imprint`,
				`privacy-imprint.html`},
		},
		{
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesFolderPath, `rss`),
				`rss`,
				`feed.tpl`,
				``,
				`rss.xml`},
		},
	}
}

func ListDirectories(dir string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)

	isDir := func(file os.FileInfo) bool { return file.IsDir() }
	return filter(allFiles, isDir)
}
