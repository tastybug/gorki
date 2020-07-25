<rss version="2.0">
    <channel>
        <title>
            An RSS Blog - Daily News and Information Related to RSS Feeds, Syndication and Aggregation.
        </title>
        <description>
            Daily RSS Blog and news related to RSS, really simple syndication, rdf, blogs, syndication and news aggregation. Information related to marketing RSS, new RSS software releases, beta test opportunities, new RSS directories and discussions of opportunities related to RSS.
        </description>
        <link>http://www.rss-specifications.com/blog.htm</link>
        <lastBuildDate>Tue, 26 May 2020 11:22:00 -0400</lastBuildDate> <!-- TODO -->
        <pubDate>Tue, 28 Nov 2006 09:00:00 -0500</pubDate>  <!-- TODO -->
        <image>
            <url>https://www.tastybug.com/favicon.png</url>
            <title>TastyBug RSS</title>
            <link>https://www.tastybug.com/</link>
        </image>
        {{range .AllArticles}}
        <item>
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