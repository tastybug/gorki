package pages

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

const templatesFolderName = `templates`

func getPartialsPaths(templatesDir string) []string {
	return []string{
		filepath.Join(templatesDir, "footer.html"),
		filepath.Join(templatesDir, "navigation.html"),
		filepath.Join(templatesDir, "head.html"),
	}
}

func createArticles(pages []ContentPage) Articles {
	var articlesOnly []ContentPage
	for _, page := range pages {
		if page.isArticle {
			articlesOnly = append(articlesOnly, page)
		}
	}
	return Articles{Articles: articlesOnly, ArticleCount: len(articlesOnly)}
}

func CreatePacks(contentPages []ContentPage, siteDir string) []ContentPack {
	templatesDir := filepath.Join(siteDir, templatesFolderName)
	articles := createArticles(contentPages)

	var packs []ContentPack
	for _, page := range contentPages {
		packs = append(packs,
			renderAndPackage(
				page,
				articles,
				templatesDir),
		)
	}
	return packs
}

func renderAndPackage(page ContentPage, articles Articles, templatesFolder string) ContentPack {
	conf := page.templatingConf
	paths := []string{filepath.Join(templatesFolder, conf.templateFolder, conf.templateFileName)}
	if conf.extraContent != `` {
		extraContentTemplate := createContentTemplate(conf.extraContent)
		paths = append(paths, extraContentTemplate.Name())
		defer os.Remove(extraContentTemplate.Name())
	}
	paths = append(paths, getPartialsPaths(templatesFolder)...)

	var htmlString bytes.Buffer
	t, _ := template.ParseFiles(paths...)
	err := t.Execute(&htmlString, articles)
	util.PanicOnError(err)

	return ContentPack{
		Folders:     conf.resultFolderName,
		HtmlContent: htmlString.String(),
		FileName:    conf.resultFileName,
		assets: collectAssets(
			conf.assetFolderPath,
			`.html`,
			conf.resultFolderName),
	}
}

func createContentTemplate(content string) *os.File {
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func collectAssets(assetFolderPath, fileExtToAvoid, resultFolderName string) []Asset {
	var assets []Asset
	for _, assetFile := range util.ListFilesWithoutSuffix(assetFolderPath, fileExtToAvoid) {
		assets = append(assets,
			Asset{Filename: assetFile.Name(),
				Context:      resultFolderName,
				CopyFromPath: filepath.Join(assetFolderPath, assetFile.Name())})
	}
	if len(assets) > 0 {
		return assets
	} else {
		return nil
	}
}
