package s3 

import "fmt"

type S3Error struct {
	Code       string
	HTTPStatus int
	Message    string
}

func (e *S3Error) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("[%s] %s (http %d)", e.Code, e.Message, e.HTTPStatus)
	}
	return e.Message

}
