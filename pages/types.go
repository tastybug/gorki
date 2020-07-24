package pages

type ContentPage struct {
	isArticle      bool
	BucketName     string
	Title          string
	Description    string
	PublishedDate  string
	templatingConf TemplatingConf
}

type Articles struct {
	Articles     []ContentPage
	ArticleCount int
}

type TemplatingConf struct {
	extraContent     string
	assetFolderPath  string
	templateFolder   string
	templateFileName string
	resultFolderName string
	resultFileName   string
}
