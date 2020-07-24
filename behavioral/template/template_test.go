/*
  @Author :     lyb
  @File :       template_test
  @Description:
*/
package template

import "testing"

func TestTemplate(t *testing.T) {
	var downloader1 Downloader = NewHTTPDownloader()

	downloader1.Download("http://example.com/abc.zip")

	var downloader2 Downloader = NewFTPDownloader()

	downloader2.Download("ftp://example.com/abc.zip")
}