package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func UploadApp(loomHost string, apikey string, filename string, slug string) {
	targetUrl := fmt.Sprintf("%s/upload", loomHost)

	fmt.Printf("Deploying %s to Loom Network... \n", filename)
	fmt.Printf("DApp deployed to ")
	color.Blue("https://%s.loomapps.io\n", slug)

	postFile(filename, targetUrl, apikey, slug)
}

func postFile(filename, targetUrl, apikey, slug string) error {
	client := &http.Client{}

	var err error
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.Boundary()
	bodyWriter.FormDataContentType()

	if err = bodyWriter.WriteField("application_slug", slug); err != nil {
		return err
	}
	//Lets the backend know to create a new application if it doesn't already exist
	if err = bodyWriter.WriteField("auto_create", "true"); err != nil {
		return err
	}

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	bodyWriter.Close()

	req, err := http.NewRequest("POST", targetUrl, bodyBuf)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	//	requestDump, err := httputil.DumpRequest(req, true)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(string(requestDump))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(resp_body))
	return nil
}
