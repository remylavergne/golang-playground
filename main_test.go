package main

import (
	"testing"
)

func TestFilterUniqueSlice(t *testing.T) {
	test := unique([]string{"1", "1", "2"})

	if len(test) != 2 {
		t.Fatalf("Error")
	}
}

func TestExtractUrlsFromHtmlEmpty(t *testing.T) {
	output := extractUrls("", "google.com")

	if len(output) != 0 {
		t.Fatalf("Output must be empty!")
	}
}

func TestExtractUrlsFromHtmlNoUrlAvailable(t *testing.T) {
	output := extractUrls("<html><title>Test without url</title></html>", "google.com")

	if len(output) != 0 {
		t.Fatalf("Output must be empty!")
	}
}

func TestExtractUrlsFromHtml(t *testing.T) {
	output := extractUrls(getHtmlSourceWithUrl(), "google.com")

	if len(output) != 152 {
		t.Fatalf("Output must not be empty!")
	}
}

func TestUrlsFilteredByExtension(t *testing.T) {
	output := extractUrls(getHtmlSourceWithUrl(), "google.com")
	urlsFiltered := filterUrlByExtension(output, Jpg)

	if len(urlsFiltered) != 145 {
		t.Fatalf("Urls not filtered properly")
	}
}

func TestUrlFormating_MissingHttpProtocol(t *testing.T) {
	urlExpected := "https://www.google.com/file.jpg"
	domain := "www.google.com"
	url := "//www.google.com/file.jpg"

	urlFormated := formatUrl(url, domain)

	if urlFormated != urlExpected {
		t.Fatalf("Urls have to be the same")
	}
}

func TestUrlFormating_RelativePath(t *testing.T) {
	urlExpected := "https://www.google.com/files/picture/file.jpg"
	domain := "www.google.com"
	url := "/files/picture/file.jpg"

	urlFormated := formatUrl(url, domain)

	if urlFormated != urlExpected {
		t.Fatalf("Urls have to be the same")
	}
}

func TestExtractUrlDomain(t *testing.T) {
	domainExpected := "www.google.com"
	url := "https://www.google.com/file.jpg"

	domain := extractDomainName(url)

	if domain != domainExpected {
		t.Fatalf("Domain extracted doesn't match")
	}
}

func TestExtractUrlDomain_second(t *testing.T) {
	domainExpected := "google.com"
	url := "http://google.com/file.jpg"

	domain := extractDomainName(url)

	if domain != domainExpected {
		t.Fatalf("Domain extracted doesn't match")
	}
}

func TestExtractUrlDomain_third(t *testing.T) {
	domainExpected := "www2.google.co.jp"
	url := "http://www2.google.co.jp/file.jpg"

	domain := extractDomainName(url)

	if domain != domainExpected {
		t.Fatalf("Domain extracted doesn't match")
	}
}

func TestExtractUrlDomain_fourth(t *testing.T) {
	domainExpected := "google.com"
	url := "google.com/file.jpg"

	domain := extractDomainName(url)

	if domain != domainExpected {
		t.Fatalf("Domain extracted doesn't match")
	}
}

func TestExtractUrlDomain_fifth(t *testing.T) {
	domainExpected := "www.google.com"
	url := "www.google.com/file.jpg"

	domain := extractDomainName(url)

	if domain != domainExpected {
		t.Fatalf("Domain extracted doesn't match")
	}
}

func getHtmlSourceWithUrl() string {
	return `https://google.com/file/1542938707009.jpg
	https://www.google.com/file/1542938707009s.jpg
	https://www.google.com/file/1543007707295s.webm
	https://www.google.com/file/1543031449785s.jpg
	https://www.google.com/file/1543731791729.jpg
	https://www.google.com/file/1543731791729s.jpg
	https://www.google.com/file/1544081575818s.jpg
	https://google.com/file/1546106903399.jpg
	https://www.google.com/file/1546106903399s.jpg
	https://www.google.com/file/1546107028663.jpg
	https://www.google.com/file/1546107028663s.jpg
	https://www.google.com/file/1546107190697.jpg
	https://www.google.com/file/1546107190697s.gif
	https://www.google.com/file/1546107435147.jpg
	https://www.google.com/file/1546107435147s.jpg
	https://google.com/file/1546147981198.jpg
	https://www.google.com/file/1546147981198s.jpg
	https://www.google.com/file/1546151598301.jpg
	https://www.google.com/file/1546151598301s.jpg
	https://www.google.com/file/1546152892125s.gif
	https://google.com/file/1546201841312.jpg
	https://www.google.com/file/1546201841312s.jpg
	https://www.google.com/file/1546207404643s.jpg
	https://www.google.com/file/1546227530037.webm
	https://www.google.com/file/1546227530037s.gif
	https://www.google.com/file/1546228358181.jpg
	https://www.google.com/file/1546228358181s.jpg
	https://www.google.com/file/1546229029211.jpg
	https://www.google.com/file/1546229029211s.jpg
	https://google.com/file/1547510854536.jpg
	https://www.google.com/file/1547510854536s.jpg
	https://www.google.com/file/1548024179370.jpg
	https://www.google.com/file/1548024179370s.jpg
	https://www.google.com/file/1548024259438s.jpg
	https://google.com/file/1548619574469.jpg
	https://www.google.com/file/1548619574469s.jpg
	https://www.google.com/file/1550203755543.jpg
	https://www.google.com/file/1550203755543s.jpg
	https://www.google.com/file/1550778204435.jpg
	https://www.google.com/file/1550778204435s.jpg
	https://google.com/file/1551899922318.jpg
	https://www.google.com/file/1551899922318s.jpg
	https://www.google.com/file/1553198472361s.jpg
	https://www.google.com/file/1553198597053s.jpg
	https://google.com/file/1553843133073.jpg
	https://www.google.com/file/1553843133073s.jpg
	https://www.google.com/file/1553843201588s.jpg
	https://google.com/file/1553843282321.jpg
	https://www.google.com/file/1553843282321s.jpg
	https://www.google.com/file/1554182508406.jpg
	https://www.google.com/file/1554182508406s.jpg
	https://google.com/file/1556976960860.jpg
	https://www.google.com/file/1556976960860s.jpg
	https://www.google.com/file/1558238707043.jpg
	https://www.google.com/file/1558238707043s.jpg
	https://www.google.com/file/1558458727039.webm
	https://www.google.com/file/1558458727039s.jpg
	https://www.google.com/file/1558580482092.webm
	https://www.google.com/file/1558580482092s.jpg
	https://google.com/file/1558668778099.jpg
	https://www.google.com/file/1558668778099s.jpg
	https://www.google.com/file/1558728274327s.jpg
	https://www.google.com/file/1558728432920s.jpg
	https://www.google.com/file/1558731910999.jpg
	https://www.google.com/file/1558731910999s.jpg
	https://www.google.com/file/1561012683100s.jpg
	https://www.google.com/file/1562094426163s.jpg
	https://www.google.com/file/1562094850922s.jpg
	https://www.google.com/file/1562094965072s.jpg
	https://www.google.com/file/1562095041432s.jpg
	https://www.google.com/file/1565691264515.jpg
	https://www.google.com/file/1565691264515s.jpg
	https://www.google.com/file/1567109298321s.jpg
	https://www.google.com/file/1567550297223s.jpg
	https://google.com/file/1568981406669.jpg
	https://www.google.com/file/1568981406669s.jpg
	https://google.com/file/1568981778612.jpg
	https://www.google.com/file/1568981778612s.jpg
	https://google.com/file/1569789640228.jpg
	https://www.google.com/file/1569789640228s.jpg
	https://www.google.com/file/1571520500720s.jpg
	https://google.com/file/1576942093962.jpg
	https://www.google.com/file/1576942093962s.jpg
	https://www.google.com/file/1578817241994s.jpg
	https://google.com/file/1580451607372.jpg
	https://www.google.com/file/1580451607372s.jpg
	https://www.google.com/file/1580509273384s.jpg
	https://www.google.com/file/1581877268507s.jpg
	https://google.com/file/1582359372393.jpg
	https://www.google.com/file/1582359372393s.jpg
	http://www.google.com/file/1582359818106s.jpg
	http://google.com/file/1583916748627.jpg
	http://www.google.com/file/1583916748627s.jpg
	http://www.google.com/file/1583916812344.jpg
	http://www.google.com/file/1583916812344s.jpg
	http://www.google.com/file/1583916873928.jpg
	http://www.google.com/file/1583916873928s.jpg
	http://www.google.com/file/1583916935635s.jpg
	https://www.google.com/file/1583916999114s.jpg
	https://google.com/file/1583917188265.jpg
	https://www.google.com/file/1583917188265s.jpg
	https://google.com/file/1588914097276.jpg
	https://www.google.com/file/1588914097276s.jpg
	https://google.com/file/1588914175909.jpg
	https://www.google.com/file/1588914175909s.jpg
	https://google.com/file/1588914256730.jpg
	https://www.google.com/file/1588914256730s.jpg
	https://www.google.com/file/1588914532699.jpg
	https://www.google.com/file/1588914532699s.jpg
	https://www.google.com/file/1588914597837.jpg
	https://www.google.com/file/1588914597837s.jpg
	https://www.google.com/file/1588914663532.jpg
	https://www.google.com/file/1588914663532s.jpg
	https://www.google.com/file/1588914729256.jpg
	https://www.google.com/file/1588914729256s.jpg
	https://google.com/file/1592363950347.jpg
	https://www.google.com/file/1592363950347s.jpg
	https://www.google.com/file/1593392126491s.jpg
	https://www.google.com/file/1593726632179s.jpg
	https://www.google.com/file/1594571673271.jpg
	https://www.google.com/file/1594571673271s.jpg
	https://www.google.com/file/1595526703035s.jpg
	www.google.com/file/1596154802371s.jpg
	www.google.com/file/1600548231760.jpg
	www.google.com/file/1600548231760s.jpg
	www.google.com/file/1602283533381s.jpg
	google.com/file/1605476373991.jpg
	www.google.com/file/1605476373991s.jpg
	www.google.com/file/1606930963800.jpg
	www.google.com/file/1606930963800s.jpg
	www.google.com/file/1606937904160s.jpg
	google.com/file/1607244419249.jpg
	https://www.google.com/file/1607244419249s.jpg
	https://www.google.com/file/1608929572470s.jpg
	https://www.google.com/file/1611550935207s.jpg
	https://google.com/file/1612276410247.jpg
	https://www.google.com/file/1612276410247s.jpg
	https://www.google.com/file/1612463395986.jpg
	https://www.google.com/file/1612463395986s.jpg
	https://www.google.com/file/1612746694847s.jpg
	https://www.google.com/file/1613786798096s.jpg
	https://google.com/file/1615501448508.jpg
	https://www.google.com/file/1607244419249s.jpg
	https://www.google.com/file/1608929572470s.jpg
	https://www.google.com/file/1611550935207s.jpg
	https://google.com/file/1612276410247.jpg
	https://www.google.com/file/1612276410247s.jpg
	https://www.google.com/file/1612463395986.jpg
	https://www.google.com/file/1612463395986s.jpg
	https://www.google.com/file/1612746694847s.jpg
	https://www.google.com/file/1613786798096s.jpg
	https://google.com/file/1615501448508.jpg
	`
}
