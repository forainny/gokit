package s3

import "testing"

func TestPresignPutObject(t *testing.T) {
	s3Client := NewS3Client("bucket")
	req, err := s3Client.PresignPutObject("path/to/filename.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(req.URL)
}
