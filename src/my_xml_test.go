package hello

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"time"
)

func slurpFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func TestParseSimpleXML(t *testing.T) {
	assert := assert.New(t)

	myXML := []byte(`
	<Person>
	  <ID>123</ID>
	  <Name>
	    <First>Hoge</First>
	    <Last>Fuga</Last>
	  </Name>
	</Person>
	`)

	type PersonName struct {
		First string `xml:"First"`
		Last  string `xml:"Last"`
	}
	type Person struct {
		ID   int         `xml:"ID"`
		Name *PersonName `xml:"Name"`
	}

	p := Person{}
	xml.Unmarshal(myXML, &p)

	assert.Equal(123, p.ID)
	assert.Equal("Hoge", p.Name.First)
	assert.Equal("Fuga", p.Name.Last)
}

func TestParseRSS(t *testing.T) {
	assert := assert.New(t)

	myXML, _ := slurpFile("../data/rss.rdf")

	type RSSChannel struct {
		Title string `xml:"title"`
		About string `xml:"about,attr"`
	}
	type RSSItem struct {
		Title string `xml:"title"`
		Link  string `xml:"link"`
		Date  string `xml:"date"`
	}
	type RSS struct {
		Channel *RSSChannel `xml:"channel"`
		Items   []*RSSItem  `xml:"item"`
	}

	rss := RSS{}
	xml.Unmarshal(myXML, &rss)

	assert.Equal("痛いニュース(ﾉ∀`)", rss.Channel.Title)
	assert.Equal("http://blog.livedoor.jp/dqnplus/", rss.Channel.About)

	assert.Equal(15, len(rss.Items))

	assert.Equal("【画像】 『ドラクエ11』の2D版のグラフィックがスーファミ以下と話題に", rss.Items[0].Title)
	assert.Equal("http://blog.livedoor.jp/dqnplus/archives/1926718.html", rss.Items[0].Link)
	assert.Equal("2017-05-28T20:38:56+09:00", rss.Items[0].Date)

	t0, _ := time.Parse("2006-01-02T15:04:05+09:00", rss.Items[0].Date)

	y, m, d := t0.Date()

	assert.Equal(2017, y)
	assert.Equal(time.Month(5), m)
	assert.Equal(28, d)
}

func TestGenerateXML(t *testing.T) {
	assert := assert.New(t)

	type RSSItem struct {
		XMLName xml.Name `xml:"item"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}
	type RSSChannel struct {
		XMLName     xml.Name `xml:"channel"`
		About       string   `xml:"rdf:about,attr"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Items       []*RSSItem
	}
	type RSS struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Channel *RSSChannel
	}

	item1 := &RSSItem{
		Title: "あああ",
		Link:  "http://hogefuga",
	}
	item2 := &RSSItem{
		Title: "いいい",
		Link:  "http://foobar",
	}
	channel := &RSSChannel{
		About:       "http://foo.bar",
		Title:       "ほげ",
		Description: "せつめいぶん",
		Items:       []*RSSItem{item1, item2},
	}
	rss := &RSS{
		Channel: channel,
		Version: "2.0",
	}

	output, _ := xml.MarshalIndent(rss, "", "  ")

	assert.Equal(`<rss version="2.0">
  <channel rdf:about="http://foo.bar">
    <title>ほげ</title>
    <description>せつめいぶん</description>
    <item>
      <title>あああ</title>
      <link>http://hogefuga</link>
    </item>
    <item>
      <title>いいい</title>
      <link>http://foobar</link>
    </item>
  </channel>
</rss>`, string(output))
}
