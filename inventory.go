package main

import (
	"log"
	"path/filepath"
)

func collectAllBundles(settings Settings) []bundle {
	pages := collectArticleBundles(settings)
	pages = append(pages, collectStaticBundles(settings)...)
	return pages
}

func collectArticleBundles(settings Settings) []bundle {
	articlesRootPath := settings.ArticlesRoot
	var bundles []bundle
	for _, bundle := range ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bundle.Name(), `article.md`)
		if Exists(articlePath) {
			page := newBundle(articlesRootPath, bundle.Name(), ReadFileContent(articlePath))
			if !page.ArticleData.IsDraft {
				log.Printf("Proceeding with non-draft article '%s' at '%s'", page.ArticleData.Title, articlePath)
				bundles = append(bundles, page)
			} else {
				log.Printf("Skipping draft article '%s.'", bundle.Name())
			}
		} else {
			log.Printf("Skipping bundle '%s', no 'article.md' found.", bundle.Name())
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
