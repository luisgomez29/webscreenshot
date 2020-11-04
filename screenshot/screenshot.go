package screenshot

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"
)

// GenerateScreenshot get the screenshot and stores it in the images folder
func GenerateScreenshot(url *url.URL) (string, error) {
	// chromedp configuration
	var options []chromedp.ExecAllocatorOption

	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)
	options = append(options, chromedp.WindowSize(1440, 900))

	actx, acancel := chromedp.NewExecAllocator(context.Background(), options...)
	ctx, cancel := chromedp.NewContext(actx)
	defer acancel()
	defer cancel()

	// capture screen of an element
	var buf []byte

	// capture the entire browser screen returning the PNG image
	if err := chromedp.Run(ctx, takeScreenshot(url.String(), &buf)); err != nil {
		return "", err
	}

	// storage path and image name
	createFolderIFNoExist("images")
	rand.Seed(time.Now().Unix())
	imgPath := "images/screenshot" + strconv.Itoa(int(rand.Int31())) + ".png"
	if err := ioutil.WriteFile(imgPath, buf, 0644); err != nil {
		return "", err
	}
	return imgPath, nil
}

// takeScreenshot take a screenshot of the entire browser window
func takeScreenshot(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(res),
	}
}

// createFolderIFNoExist create folder where images are stored if it doesn't exist
func createFolderIFNoExist(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.Mkdir(folderPath, 0755)
		if err != nil {
			fmt.Printf("error while creating folder")
		}
	}
}
