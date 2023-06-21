package getver

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func GetVersion() ([]string, error) {
	response, err := http.Get("https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions.html")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve EKS versions. Status code: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	var versions []string

	doc.Find("table").Each(func(index int, tableHtml *goquery.Selection) {
		tableHtml.Find("tr").Each(func(index int, rowHtml *goquery.Selection) {
			version := rowHtml.Find("td").First().Text()

			if version != "" {
				versions = append(versions, version)
			}
		})
	})

	return versions, nil
}
