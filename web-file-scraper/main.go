package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

// TODO: Test main url validity
// TODO: Extract links from a file -> new arg & behavior
// TODO: Fix windows path copy

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkArgsValidity(args []string) bool {
	// Remove first args (command line)
	cleanArgs := args[1:]
	validArgs := map[string]bool{"--url": true, "--dry-run": true, "-e": true}
	areValid := true

	for _, arg := range cleanArgs {
		if strings.Contains(string(arg[0]), "-") {
			_, valid := validArgs[arg]

			if !valid {
				fmt.Println("Arg ->", arg, "is invalid!")
				areValid = false
			}
		}
	}

	return areValid
}

func argsParsing(args []string) {
	// Check if all arguments passed exist
	argsValid := checkArgsValidity(args)

	if !argsValid {
		fmt.Printf("Some args are invalid")
		os.Exit(1)
	}

	for i, arg := range args {
		if arg == "--url" && i+1 < len(args) {
			// TODO: Check url format
			UrlArg = args[i+1]
		}

		if arg == "--extension" || arg == "-e" && i+1 < len(args) {
			rawExt := args[i+1]
			ext, exist := ToFileExtension(rawExt)

			if !exist {
				fmt.Println("Error -> The file extension", rawExt, "isn't available yet!")
				os.Exit(1)
			} else {
				extensionArg = ext
			}
		}

		if arg == "--dry-run" {
			dryRunArg = true
		}
	}
}

func extractDomainName(u string) string {
	// Should start at least with www. or http(s)://
	r, err := regexp.Compile("^https?://")
	check(err)
	if match := r.MatchString(u); !match {
		u = "https://" + u
	}

	fileURL, err := url.Parse(u)
	if err != nil {
		check(err)
	}

	return fileURL.Host
}

func formatUrl(rawUrl string, domain string) string {
	r, err := regexp.Compile("^/[\\w\\d#]+")
	check(err)

	switch {
	// Relative path
	case r.MatchString(rawUrl):
		return "https://" + domain + rawUrl
		// Missing domains
	case strings.HasPrefix(rawUrl, "//"):
		return "https:" + rawUrl
	default:
		return rawUrl
	}
}

func formatUrls(urls []string, domain string) []string {
	us := []string{}

	for _, u := range urls {
		urlFormat := formatUrl(u, domain)
		us = append(us, urlFormat)
	}

	return us
}

func extractUrls(htmlSource string) []string {
	urls := []string{}
	r, _ := regexp.Compile("/?/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()!@:%_\\+.~#?&//=]*)")
	urlsFound := r.FindAllString(string(htmlSource), -1)

	for _, u := range urlsFound {
		urls = append(urls, u)
	}

	return urls
}

func filterUrlByExtension(urls []string, fe FileExtension) []string {
	urlsFiltered := []string{}
	r, _ := regexp.Compile("\\." + fe.String())

	for _, url := range urls {
		if r.MatchString(url) {
			urlsFiltered = append(urlsFiltered, url)
		}
	}

	return urlsFiltered
}

func downloadFile(client *http.Client, fileUrl string, dirPath string, respBodyChannel chan []byte, wg *sync.WaitGroup) {
	defer (*&wg).Done()

	// Extract filename from path
	fileURL, err := url.Parse(fileUrl)
	if err != nil {
		check(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(dirPath + "/" + fileName)
	if err != nil {
		check(err)
	}
	defer file.Close()

	// Put content on file
	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("Unable to download", fileURL)
		respBodyChannel <- make([]byte, 0)
	} else {
		// Read response & send it through channel
		body, error := ioutil.ReadAll(resp.Body)
		check(error)

		respBodyChannel <- body
	}

	defer resp.Body.Close()
}

func persistFile(fileUrl string, dirPath string, respBodyChannel chan []byte, wg *sync.WaitGroup) {
	defer (*wg).Done()

	// Get file data through channel
	byteBody := <-respBodyChannel
	if len(byteBody) == 0 {
		fmt.Println("Body is empty. Skip file creation.")
		return
	}
	r := bytes.NewReader(byteBody)

	fileURL, err := url.Parse(fileUrl)
	if err != nil {
		check(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(dirPath + "/" + fileName)
	if err != nil {
		check(err)
	}
	defer file.Close()

	// Write data into file
	size, err := io.Copy(file, r)

	if err != nil {
		check(err)
	}

	fmt.Printf("File %s downloaded (size: %d)\n", file.Name(), size)
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func getHttpClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	return &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
		Timeout:   10 * time.Second,
		Transport: t,
	}
}

var UrlArg string
var dryRunArg bool = false
var extensionArg FileExtension
var output string = "output"
var domain string

func main() {
	// Args parsing
	argsParsing(os.Args)
	domain := extractDomainName(UrlArg)

	// Client generation
	client := getHttpClient()

	// Http get
	resp, err := client.Get(UrlArg)
	if err != nil {
		check(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// Process data
	urls := extractUrls(string(body))
	urlsFormated := formatUrls(urls, domain)
	urlsFiltered := filterUrlByExtension(urlsFormated, extensionArg)
	urlsUnique := unique(urlsFiltered)

	if dryRunArg {
		fmt.Println("Dry Run mode ->", len(urlsUnique), "files would have been downloaded!")
		fmt.Println("Links =>\n", urlsUnique)
		os.Exit(0)
	}

	fmt.Println(len(urlsUnique), "files are going to be downloaded")

	dirHash := md5.Sum([]byte(UrlArg))
	dirHashString := hex.EncodeToString(dirHash[:])
	outputDir := output + "/" + dirHashString
	os.Mkdir(output, 0755)
	os.Mkdir(outputDir, 0755)

	respBodyChannel := make(chan []byte)
	// Wait for all Goroutines launched
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(len(urlsUnique) * 2)

	for _, url := range urlsUnique {
		go downloadFile(client, url, outputDir, respBodyChannel, &wg)
		go persistFile(url, outputDir, respBodyChannel, &wg)
	}
	wg.Wait()
	fmt.Println("Process done!")
}
