package gorkify

import (
	"gorki/util"
	"log"
	"path/filepath"
)

func collectAllBundles(settings util.Settings) []bundle {
	pages := collectArticleBundles(settings)
	pages = append(pages, collectStaticBundles(settings)...)
	return pages
}

func collectArticleBundles(settings util.Settings) []bundle {
	articlesRootPath := settings.ArticlesRoot
	var bundles []bundle
	for _, bundle := range util.ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bundle.Name(), `article.md`)
		if util.Exists(articlePath) {
			page := newBundle(articlesRootPath, bundle.Name(), util.ReadFileContent(articlePath))
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

func collectStaticBundles(settings util.Settings) []bundle {
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
