package pages

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

const templatesFolderName = `templates`

func CollectMainPagesContentPacks(articles Articles, mainPages []MainPage, siteDir string) []ContentPack {

	templatesDir := filepath.Join(siteDir, templatesFolderName)
	var packs []ContentPack
	for _, mainPage := range mainPages {
		packs = append(packs,
			assemblePack(
				append([]string{filepath.Join(templatesDir, mainPage.bucketName, mainPage.bucketName+`.html`)},
					getPartialsPaths(templatesDir)...),
				mainPage,
				siteDir,
				articles),
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

func assemblePack(paths []string, mainPage MainPage, siteDir string, articles Articles) ContentPack {
	tmpl := template.Must(template.ParseFiles(paths...))
	var buffer bytes.Buffer
	tmpl.Execute(&buffer, articles) // articles are provided in case the template wants to show something here

	var folderToBePutIt string
	if !mainPage.goesToRoot {
		folderToBePutIt = mainPage.bucketName
	}

	return ContentPack{
		HtmlContent: buffer.String(),
		FileName:    mainPage.bucketName + ".html",
		Folders:     folderToBePutIt,
		assets: collectAssets(
			filepath.Join(siteDir, `templates`, mainPage.bucketName),
			mainPage.bucketName,
			`.html`,
			!mainPage.goesToRoot)}
}

func TurnArticlesIntoContentPack(articles Articles, siteDir string) []ContentPack {
	templatesFolder := filepath.Join(siteDir, `templates`)
	postsDir := filepath.Join(siteDir, "posts")

	var contentPacks []ContentPack
	for _, postable := range articles.Articles {
		contentPacks = append(contentPacks, toContentPack(postable, postsDir, templatesFolder))
	}
	return contentPacks
}

func toContentPack(article ArticlePage, bucketFolderPath string, templatesFolder string) ContentPack {
	var htmlString bytes.Buffer

	contentTemplate := createContentTemplate(article.ContentAsHtml)
	t, _ := template.ParseFiles(
		append([]string{
			filepath.Join(templatesFolder, `blogpost`, `blogpost.html`),
			filepath.Join(contentTemplate.Name()),
		},
			getPartialsPaths(templatesFolder)...)...)
	err := t.Execute(&htmlString, article)
	util.PanicOnError(err)

	defer os.Remove(contentTemplate.Name())

	return ContentPack{
		Folders:     article.BucketName,
		HtmlContent: htmlString.String(),
		FileName:    "article.html",
		assets: collectAssets(
			filepath.Join(bucketFolderPath, article.BucketName),
			article.BucketName,
			`.md`,
			true),
	}
}

func createContentTemplate(content string) *os.File {
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func collectAssets(assetFolderPath, bucketName, fileExtToAvoid string, putIntoBucket bool) []Asset {
	var folderToBePutIt string
	if putIntoBucket {
		folderToBePutIt = bucketName
	}
	var assets []Asset
	for _, assetFile := range util.ListFilesWithoutSuffix(assetFolderPath, fileExtToAvoid) {
		assets = append(assets,
			Asset{Filename: assetFile.Name(),
				Context:      folderToBePutIt,
				CopyFromPath: filepath.Join(assetFolderPath, assetFile.Name())})
	}
	if len(assets) > 0 {
		return assets
	} else {
		return nil
	}
}
