package pages

type MainPage struct {
	bucketName string
	goesToRoot bool
}

func CollectMains() []MainPage {
	return []MainPage{
		{`index`, true},
		{`about`, false},
		{`privacy-imprint`, false},
	}
}
