package gorki

import (
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

func CollectAllBundles(settings Settings) []bundle {
	pages := collectUsableArticleBundles(settings)
	pages = append(pages, collectStaticBundles(settings)...)
	return pages
}

func collectUsableArticleBundles(settings Settings) []bundle {
	var bundles []bundle
	for _, bundle := range listDirs(settings.ArticlesRoot) {
		b, err := newBundle(settings.ArticlesRoot, bundle.Name())
		if err != nil {
			log.Println("Skipping broken bundle:", err)
		} else if b.isToBeRendered() {
			bundles = append(bundles, b)
		} else {
			log.Println("Skipping bundle not to be rendered:", bundle.Name())
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

func listDirs(dir string) []os.FileInfo {
	if allFiles, err := ioutil.ReadDir(dir); err != nil {
		panic(err)
	} else {
		isDir := func(file os.FileInfo) bool { return file.IsDir() }
		return filter(allFiles, isDir)
	}
}
