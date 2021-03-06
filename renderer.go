package main

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

func renderPages(settings Settings, bundles []bundle) []renderedPage {
	articleBundles := filterArticlesAndSortByDate(bundles, true)

	var pages []renderedPage
	for _, bundle := range bundles {
		pages = append(pages,
			renderAndPackage(
				bundle,
				articleBundles,
				settings.TemplatesRoot),
		)
	}
	return pages
}

func getPartialsPaths(templatesDir string) []string {
	return []string{
		filepath.Join(templatesDir, "footer.html"),
		filepath.Join(templatesDir, "navigation.html"),
		filepath.Join(templatesDir, "head.html"),
	}
}

func filterArticlesAndSortByDate(bundles []bundle, sortedDesc bool) []bundle {
	articleMap := make(map[string]bundle)
	var dateStrings []string
	var sortedArticles []bundle
	for _, page := range bundles {
		if page.ArticleData.PublishedDate != `` {
			articleMap[page.ArticleData.PublishedDate] = page
			dateStrings = append(dateStrings, page.ArticleData.PublishedDate)
		}
	}
	dateStrings = SortStrings(dateStrings, sortedDesc)
	for _, dateString := range dateStrings {
		sortedArticles = append(sortedArticles, articleMap[dateString])
	}

	return sortedArticles
}

func renderAndPackage(page bundle, pagesThatAreArticles []bundle, templatesRoot string) renderedPage {
	conf := page.TemplatingConf
	paths := []string{filepath.Join(templatesRoot, conf.templateFolder, conf.templateFileName)}
	if conf.extraContent != `` {
		extraContentTemplate := createContentTemplate(conf.extraContent)
		paths = append(paths, extraContentTemplate.Name())
		defer RemoveFile(*extraContentTemplate)
	}
	paths = append(paths, getPartialsPaths(templatesRoot)...)

	var htmlString bytes.Buffer
	t, _ := template.New(conf.templateFileName).Funcs(template.FuncMap{
		"ToRssDate": func(isoDate string) string {
			return ISODateToRSSDateTime(isoDate)
		},
		"GetNowAsRSSDateTime": func() string {
			return GetNowAsRSSDateTime()
		},
	}).ParseFiles(paths...)

	err := t.Execute(
		&htmlString,
		templateDataContext{
			AllArticles:  pagesThatAreArticles,
			ArticleCount: len(pagesThatAreArticles),
			LocalPage:    page,
		})
	PanicOnError(err)

	return renderedPage{
		FolderName:  conf.resultFolderName,
		HtmlContent: htmlString.String(),
		FileName:    conf.ResultFileName,
		assets: collectAssets(
			conf.assetFolderPath,
			conf.resultFolderName),
	}
}

func createContentTemplate(content string) *os.File {
	return WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func collectAssets(assetFolderPath, resultFolderName string) []asset {
	var assets []asset
	for _, assetFile := range ListFilesMatching(assetFolderPath, `.*\.[^mdhtml]+`) {
		assets = append(assets,
			asset{FileName: assetFile.Name(),
				FolderName:   resultFolderName,
				CopyFromPath: filepath.Join(assetFolderPath, assetFile.Name())})
	}
	if len(assets) > 0 {
		return assets
	} else {
		return nil
	}
}

type templateDataContext struct {
	// a list of all articles
	AllArticles []bundle
	// how many articles there are
	ArticleCount int
	// this is the data of the template being built
	LocalPage bundle
}

type renderedPage struct {
	FolderName  string
	FileName    string
	HtmlContent string
	assets      []asset
}

type asset struct {
	FolderName   string
	FileName     string
	CopyFromPath string
}
