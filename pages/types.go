package pages

type Page struct {
	ArticleData    ArticleData
	TemplatingConf TemplatingConf
}

type ArticleData struct {
	isDraft       bool
	BucketName    string
	Title         string
	Description   string
	PublishedDate string
}

type TemplateDataContext struct {
	// a list of all articles
	AllArticles []Page
	// how many articles there are
	ArticleCount int
	// this is the data of the template being built
	LocalPage Page
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
