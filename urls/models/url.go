package url

import (
	"time"
)

type Url struct {
	Id          int64
	OriginalUrl string
	Short       string
	ShortUrl    string
	Date        time.Time
	UrlId		int64
}

// URL CMD
type CreateURLCMD struct {
	Original    string `json:"original_url"`
	Short       string `json:"short"`
	ShortUrl    string `json:"short_url"`
	UrlId		int64  `json:"url_id"`
}
