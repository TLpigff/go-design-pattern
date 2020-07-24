/*
  @Author :     lyb
  @File :       template
  @Description:
*/
package template

import "fmt"

type Downloader interface {
	Download(url string)
}

type template struct {
	implement
	url string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(url string) {
	t.url = url
	fmt.Println("prepare downloading")
	//t.implement.download()
	//t.implement.save()
	t.download()
	t.save()
	fmt.Println("finish downloading")
}

func (t *template) save() {
	fmt.Println("default save")
}

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader  {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.url)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.url)
}