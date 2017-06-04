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

func TestParseSimpleXml(t *testing.T) {
	assert := assert.New(t)

	myXml := []byte(`
	<Person>
	  <Id>123</Id>
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
		Id   int         `xml:"Id"`
		Name *PersonName `xml:"Name"`
	}

	p := Person{}
	xml.Unmarshal(myXml, &p)

	assert.Equal(123, p.Id)
	assert.Equal("Hoge", p.Name.First)
	assert.Equal("Fuga", p.Name.Last)
}

func TestParseRss(t *testing.T) {
	assert := assert.New(t)

	myXml, _ := slurpFile("../data/rss.rdf")

	type RssChannel struct {
		Title string `xml:"title"`
		About string `xml:"about,attr"`
	}
	type RssItem struct {
		Title string `xml:"title"`
		Link  string `xml:"link"`
		Date  string `xml:"date"`
	}
	type Rss struct {
		Channel *RssChannel `xml:"channel"`
		Items   []*RssItem  `xml:"item"`
	}

	rss := Rss{}
	xml.Unmarshal(myXml, &rss)

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

func TestGenerateXml(t *testing.T) {
	assert := assert.New(t)

	type RssItem struct {
		XMLName xml.Name `xml:"item"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}
	type RssChannel struct {
		XMLName     xml.Name `xml:"channel"`
		About       string   `xml:"rdf:about,attr"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Items       []*RssItem
	}
	type Rss struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Channel *RssChannel
	}

	item1 := &RssItem{
		Title: "あああ",
		Link:  "http://hogefuga",
	}
	item2 := &RssItem{
		Title: "いいい",
		Link:  "http://foobar",
	}
	channel := &RssChannel{
		About:       "http://foo.bar",
		Title:       "ほげ",
		Description: "せつめいぶん",
		Items:       []*RssItem{item1, item2},
	}
	rss := &Rss{
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
