package pages

type ContentPage struct {
	isArticle      bool
	BucketName     string
	Title          string
	Description    string
	PublishedDate  string
	templatingConf TemplatingConf
}

type TemplateDataContext struct {
	// a list of all articles
	AllArticles []ContentPage
	// how many articles there are
	ArticleCount int
	// this is the data of the template being built
	LocalPage ContentPage
}

type TemplatingConf struct {
	extraContent     string
	assetFolderPath  string
	templateFolder   string
	templateFileName string
	resultFolderName string
	resultFileName   string
}

type ContentPack struct {
	FolderName  string
	FileName    string
	HtmlContent string
	assets      []Asset
}

type Asset struct {
	FolderName   string
	FileName     string
	CopyFromPath string
}
