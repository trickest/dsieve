package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/publicsuffix"
)

var inputUrl *string
var inputFilePath *string
var filterLevel *string
var outputFilePath *string
var filterTLD *bool

func fail(text string) {
	fmt.Println(text)
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		fail(err.Error())
	}
}

func parseFilter(filter string) (int, int) {
	if filter == "" {
		return -1, -1
	}
	vMin := -1
	vMax := -1
	var err error
	minMax := strings.Split(filter, ":")
	if len(minMax) == 1 {
		vMin, err = strconv.Atoi(minMax[0])
		vMax = vMin + 1
		check(err)
	} else if len(minMax) == 2 {
		if minMax[0] != "" {
			vMin, err = strconv.Atoi(minMax[0])
			check(err)
		}
		if minMax[1] != "" {
			vMax, err = strconv.Atoi(minMax[1])
			check(err)
		}
	} else {
		fail("Invalid filter value: " + filter)
	}
	return vMin, vMax
}

func writeResults(domains *[]string) {
	_, err := os.Create(*outputFilePath)
	check(err)

	if len(*domains) > 0 {
		file, _ := os.OpenFile(*outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		defer file.Close()
		writer := bufio.NewWriter(file)
		for _, domain := range *domains {
			_, _ = fmt.Fprintln(writer, domain)
		}
		_ = writer.Flush()
	}
}

func parseUrl(rawUrl string, lMin, lMax int) []string {
	domains := make([]string, 0)
	if !strings.HasPrefix(rawUrl, "http") {
		rawUrl = "http://" + rawUrl
	}
	u, err := url.Parse(rawUrl)
	if err != nil {
		return domains
	}

	domainLevels := strings.Split(u.Host, ".")
	if lMin > len(domainLevels) {
		return domains
	}

	if lMax == -1 || lMax > len(domainLevels) {
		lMax = len(domainLevels)
	}
	for i := lMax - 1; i > 0 && i >= lMin; i-- {
		domain := strings.Join(domainLevels[len(domainLevels)-i:], ".")
		if domain != "" {
			domains = append(domains, domain)
		}
	}
	return domains
}

func main() {
	inputUrl = flag.String("i", "", "Input url or domain")
	inputFilePath = flag.String("if", "", "Input file path, one url/domain per line.")
	filterLevel = flag.String("f", "", "Filter domain level. "+
		"Use python slice notation to select range. \nExample input: foo.bar.baz.tld \n"+
		"  \033[3m-f 3  \033[0m    bar.baz.tld \n"+
		"  \033[3m-f 3: \033[0m    bar.baz.tld, foo.bar.baz.tld\n"+
		"  \033[3m-f 2:4\033[0m    baz.tld, bar.baz.tld\n"+
		"  \033[3m-f :3 \033[0m    tld, baz.tld")
	outputFilePath = flag.String("o", "", "Output file path, optional")
	filterTLD = flag.Bool("t", true, "Filter invalid domains according to Mozilla's publicsuffix list.")

	flag.Parse()

	inputUrls := make([]string, 0)
	if *inputUrl != "" {
		inputUrls = append(inputUrls, *inputUrl)
	}
	if *inputFilePath != "" {
		inputFile, err := os.Open(*inputFilePath)
		check(err)
		defer inputFile.Close()
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			inputUrls = append(inputUrls, scanner.Text())
		}
	}

	if len(inputUrls) == 0 {
		flag.PrintDefaults()
		fmt.Print("\nError: No input.\n")
		os.Exit(1)
	}

	lMin, lMax := parseFilter(*filterLevel)
	domainMap := make(map[string]bool)
	domains := make([]string, 0)
	for _, inputUrl := range inputUrls {
		for _, domain := range parseUrl(inputUrl, lMin, lMax) {
			if _, dup := domainMap[domain]; !dup {
				eTLD, icann := publicsuffix.PublicSuffix(domain)
				if *filterTLD {
					if icann {
						if eTLD != domain {
							fmt.Println(domain)
							domainMap[domain] = true
							domains = append(domains, domain)
						}

					}
				} else {
					fmt.Println(domain)
					domainMap[domain] = true
					domains = append(domains, domain)
				}

			}
		}
	}

	if *outputFilePath != "" && len(domains) > 0 {
		writeResults(&domains)
	}
}
