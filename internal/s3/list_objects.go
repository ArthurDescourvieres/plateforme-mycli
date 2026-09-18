package s3

import "encoding/xml"

type ObjectItem struct {
	Key          string `xml:"Key"`
	Size         int64  `xml:"Size"`
	LastModified string `xml:"LastModified"`
}

type ListObjectsResult struct {
	XMLName  xml.Name     `xml:"ListBucketResult"`
	Name     string       `xml:"Name"`
	Contents []ObjectItem `xml:"Contents"`
}

func ListObjects(bucket string) ([]ObjectItem, error) {
	return nil, nil

}
