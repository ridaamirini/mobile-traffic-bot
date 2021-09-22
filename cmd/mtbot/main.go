package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"math/rand"
	"time"
)

type Input struct {
	Account   string
	MinVisits int
	MaxVisits int
}

const (
	DefaultSleep = 5 * time.Second
	UserAgent    = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/93.0.4577.82 Safari/537.36"
)

func main() {
	acc := flag.String("account", "", "Account name of mobile.de")
	minv := flag.Int("min-visits", 1, "Minimum visits for random visits")
	maxv := flag.Int("max-visits", 10, "Maximum visits for random visits")

	flag.Parse()

	input := &Input{Account: *acc, MinVisits: *minv, MaxVisits: *maxv}
	validateInput(input)

	ctx, cancel := createChromeContext()
	defer cancel()

	pageUrl := fmt.Sprintf("https://home.mobile.de/%s", *acc)
	randomVisit := randomizeVisits(input)
	vcount := 0
	for ; vcount < randomVisit; vcount++ {
		time.Sleep(DefaultSleep)
		visitPage(pageUrl, ctx)
	}

	fmt.Println("Total visits:", vcount)
}

func createChromeContext() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.UserAgent(UserAgent))
	allcCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allcCtx, chromedp.WithLogf(log.Printf))

	return ctx, cancel
}

func validateInput(i *Input) {
	if i.MinVisits < 1 {
		panic("minimum visits needs to be greater then 1")
	}

	if i.MinVisits > i.MaxVisits {
		panic("minimum visits needs to be greater then maximum visits")
	}

	if len(i.Account) == 0 {
		panic("account name needs to be set")
	}
}

func visitPage(url string, ctx context.Context) {
	var buffer []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(DefaultSleep),
		chromedp.CaptureScreenshot(&buffer),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Site visited")
}

func randomizeVisits(i *Input) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(i.MaxVisits-i.MinVisits) + i.MinVisits
}
