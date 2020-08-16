package pages

import (
	"bytes"
	"gorki/util"
	"html/template"
	"os"
	"path/filepath"
)

func CreatePacks(settings util.Settings, pages []Page) []ContentPack {
	pagesThatAreArticles := getSortedArticlePages(pages, true)

	var packs []ContentPack
	for _, page := range pages {
		packs = append(packs,
			renderAndPackage(
				page,
				pagesThatAreArticles,
				settings.TemplatesRoot),
		)
	}
	return packs
}

func getPartialsPaths(templatesDir string) []string {
	return []string{
		filepath.Join(templatesDir, "footer.html"),
		filepath.Join(templatesDir, "navigation.html"),
		filepath.Join(templatesDir, "head.html"),
	}
}

func getSortedArticlePages(pages []Page, sortedDesc bool) []Page {
	articleMap := make(map[string]Page)
	var dateStrings []string
	var sortedArticles []Page
	for _, page := range pages {
		if page.ArticleData.PublishedDate != `` {
			articleMap[page.ArticleData.PublishedDate] = page
			dateStrings = append(dateStrings, page.ArticleData.PublishedDate)
		}
	}
	dateStrings = util.SortStrings(dateStrings, sortedDesc)
	for _, dateString := range dateStrings {
		sortedArticles = append(sortedArticles, articleMap[dateString])
	}

	return sortedArticles
}

func renderAndPackage(page Page, pagesThatAreArticles []Page, templatesRoot string) ContentPack {
	conf := page.TemplatingConf
	paths := []string{filepath.Join(templatesRoot, conf.templateFolder, conf.templateFileName)}
	if conf.extraContent != `` {
		extraContentTemplate := createContentTemplate(conf.extraContent)
		paths = append(paths, extraContentTemplate.Name())
		defer os.Remove(extraContentTemplate.Name())
	}
	paths = append(paths, getPartialsPaths(templatesRoot)...)

	var htmlString bytes.Buffer
	t, _ := template.New(conf.templateFileName).Funcs(template.FuncMap{
		"ToRssDate": func(isoDate string) string {
			return util.ISODateToRSSDateTime(isoDate)
		},
		"GetNowAsRSSDateTime": func() string {
			return util.GetNowAsRSSDateTime()
		},
	}).ParseFiles(paths...)

	err := t.Execute(
		&htmlString,
		TemplateDataContext{
			AllArticles:  pagesThatAreArticles,
			ArticleCount: len(pagesThatAreArticles),
			LocalPage:    page,
		})
	util.PanicOnError(err)

	return ContentPack{
		FolderName:  conf.resultFolderName,
		HtmlContent: htmlString.String(),
		FileName:    conf.ResultFileName,
		assets: collectAssets(
			conf.assetFolderPath,
			conf.resultFolderName),
	}
}

func createContentTemplate(content string) *os.File {
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func collectAssets(assetFolderPath, resultFolderName string) []Asset {
	var assets []Asset
	for _, assetFile := range util.ListFilesMatching(assetFolderPath, `.*\.[^mdhtml]+`) {
		assets = append(assets,
			Asset{FileName: assetFile.Name(),
				FolderName:   resultFolderName,
				CopyFromPath: filepath.Join(assetFolderPath, assetFile.Name())})
	}
	if len(assets) > 0 {
		return assets
	} else {
		return nil
	}
}
