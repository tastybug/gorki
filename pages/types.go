package pages

type ContentPage struct {
	isDraft        bool
	isArticle      bool
	BucketName     string
	Title          string
	Description    string
	PublishedDate  string
	TemplatingConf TemplatingConf
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
	ResultFileName   string
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
