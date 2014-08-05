package meh

import (
        "github.com/PuerkitoBio/goquery"
        "encoding/json"
        "log"
        "strings"
)

type MehItemOfTheDay struct {
        // Basic
        Name     string
        Price    string
        Features []string
        Images   []string

        // Links
        ForumLink string
        VideoLink string

        // Potentially unnecessary
        StoryTitle string
        Story      []string
}

func GetMeh() (response string) {
        var doc *goquery.Document
        var e error

        // var useragent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_5) AppleWebKit/537.77.4 (KHTML, like Gecko) Version/7.0.5 Safari/537.77.4";

        if doc, e = goquery.NewDocument("https://meh.com/"); e != nil {
                log.Fatal(e)
        }

        i_features := make([]string, 0)
        i_story := make([]string, 0)
        i_pictures := make([]string, 0)
        i_videolink, _ := doc.Find("#video section iframe").Attr("src")
        i_forumlink, _ := doc.Find("section.features a").Attr("href")
        i_name := doc.Find("section.features h2").Text()
        i_price := doc.Find("#hero-buttons button.buy-button span").Text()
        i_title := doc.Find("section.story h1").Text()

        doc.Find("#gallery .photos .photo").Each(func(i int, s *goquery.Selection) {
                img, _ := s.Attr("style")
                // figured this was faster than regexing out a url
                img = strings.Replace(img, "background-image: url('", "", 1)
                img = strings.Replace(img, "')", "", 1)

                i_pictures = append(i_pictures, img)
        })

        doc.Find(".story p").Each(func(i int, s *goquery.Selection) {
                text := s.Text()
                i_story = append(i_story, text)
        })

        doc.Find(".features ul li").Each(func(i int, s *goquery.Selection) {
                feature := s.Text()
                i_features = append(i_features, feature)
        })

        data := &MehItemOfTheDay{Name: i_name, Price: i_price, VideoLink: i_videolink, ForumLink: i_forumlink, Features: i_features, StoryTitle: i_title, Images: i_pictures, Story: i_story}
        json_data, _ := json.MarshalIndent(data, "", "  ")
        return string(json_data)
}
