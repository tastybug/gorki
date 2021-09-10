package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
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
	dateStrings = sort(dateStrings, sortedDesc)
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
		defer removeFile(*extraContentTemplate)
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

	if err := t.Execute(
		&htmlString,
		templateDataContext{
			AllArticles:  pagesThatAreArticles,
			ArticleCount: len(pagesThatAreArticles),
			LocalPage:    page,
		}); err != nil {
		panic(err)
	}

	return renderedPage{
		FolderName:  conf.resultFolderName,
		HtmlContent: htmlString.String(),
		FileName:    conf.ResultFileName,
		assets: collectAssets(
			conf.assetFolderPath,
			conf.resultFolderName),
	}
}

func removeFile(f os.File) {
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
}

func createContentTemplate(content string) *os.File {
	if tmpFile, err := ioutil.TempFile(os.TempDir(), "gorki-"); err != nil {
		panic(err)
	} else {
		fileWriter := bufio.NewWriter(tmpFile)
		if _, err = fileWriter.Write([]byte("{{define \"content\"}}" + content + "{{end}}")); err != nil {
			panic(err)
		}
		if err = fileWriter.Flush(); err != nil {
			panic(err)
		}
		return tmpFile
	}
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

func ListFilesMatching(dir, pattern string) []os.FileInfo {
	if allFiles, err := ioutil.ReadDir(dir); err != nil {
		panic(err)
	} else {
		return filter(allFiles, func(file os.FileInfo) bool {
			return matches(file.Name(), pattern) && !file.IsDir()
		})
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

func ExtractGroupOrFailOnMismatch(data string, pattern string, groupName string) string {

	r := regexp.MustCompile(pattern)
	result := r.FindStringSubmatch(data)

	for index, value := range r.SubexpNames() {
		if value == groupName && len(result) >= index {
			return result[index]
		}
	}
	panic(fmt.Errorf("no match for group '%s' in pattern '%s': %s", groupName, pattern, data))
}

func matches(data, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(data)
}

func ISODateToRSSDateTime(isoDate string) string {
	if dateTime, err := time.Parse(`2006-01-02`, isoDate); err != nil {
		panic(err)
	} else {
		// RSS asks for RFC822 date formats, see https://www.w3.org/Protocols/rfc822/#z28
		// nonetheless the RSS validator at https://validator.w3.org/feed/check.cgi asks for day of the week
		// which you get with RC1123 only, so using this instead of the 822 formatter
		return dateTime.Format(time.RFC1123Z)
	}
}

func GetNowAsRSSDateTime() string {
	return time.Now().Format(time.RFC1123Z)
}
