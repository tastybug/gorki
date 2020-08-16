package pages

import (
	"bytes"
	"gorki/util"
	"html/template"
	"os"
	"path/filepath"
)

func RenderPages(settings util.Settings, pages []page) []renderedPage {
	pagesThatAreArticles := getSortedArticlePages(pages, true)

	var packs []renderedPage
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

func getSortedArticlePages(pages []page, sortedDesc bool) []page {
	articleMap := make(map[string]page)
	var dateStrings []string
	var sortedArticles []page
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

func renderAndPackage(page page, pagesThatAreArticles []page, templatesRoot string) renderedPage {
	conf := page.TemplatingConf
	paths := []string{filepath.Join(templatesRoot, conf.templateFolder, conf.templateFileName)}
	if conf.extraContent != `` {
		extraContentTemplate := createContentTemplate(conf.extraContent)
		paths = append(paths, extraContentTemplate.Name())
		defer util.RemoveFile(*extraContentTemplate)
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
		templateDataContext{
			AllArticles:  pagesThatAreArticles,
			ArticleCount: len(pagesThatAreArticles),
			LocalPage:    page,
		})
	util.PanicOnError(err)

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
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func collectAssets(assetFolderPath, resultFolderName string) []asset {
	var assets []asset
	for _, assetFile := range util.ListFilesMatching(assetFolderPath, `.*\.[^mdhtml]+`) {
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
	AllArticles []page
	// how many articles there are
	ArticleCount int
	// this is the data of the template being built
	LocalPage page
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
