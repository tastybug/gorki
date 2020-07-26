<rss version="2.0" xmlns:atom="https://www.w3.org/2005/Atom">
    <channel>
        <title>
            TastyBug
        </title>
        <description>
            Occasionally struggling to articulate half-baked ideas.
        </description>
        <link>https://www.tastybug.com/</link>
        <lastBuildDate>{{GetNowAsRSSDateTime}}</lastBuildDate>
        <pubDate>{{GetNowAsRSSDateTime}}</pubDate>
        <atom:link href="https://www.tastybug.com/rss.xml" rel="self" type="application/rss+xml" />
        <image>
            <url>https://www.tastybug.com/favicon.png</url>
            <title>TastyBug</title>
            <link>https://www.tastybug.com/</link>
        </image>
        {{range .AllArticles}}
        <item>
            <guid isPermaLink="false">{{.BucketName}}</guid>
            <title>{{.Title}}</title>
            <description>
                {{.Description}}
            </description>
            <link>https://www.tastybug.com/{{.BucketName}}/{{.TemplatingConf.ResultFileName}}</link>
            <pubDate>{{ToRssDate .PublishedDate}}</pubDate>
        </item>
        {{end}}
    </channel>
</rss>