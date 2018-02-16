package model
import (
"time"
"crypto/rand"
"encoding/base64"
"sync"
)
type News struct{
	ID string
	Title string
	Image string
	Detail string
	CreatedAt time.Time
	UpdateAt time.Time
}

var (
	newsStorage []News
	mutexNews sync.Mutex
	//lastNewsID int
)
func genarateID() string{
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}
func CreateNews(news News) {
		news.ID  =genarateID()
		news.CreatedAt = time.Now()
		news.UpdateAt = news.CreatedAt
		mutexNews.Lock()
		defer mutexNews.Unlock()
 		newsStorage = append(newsStorage, news)

}
func ListNews() []*News{
		mutexNews.Lock()
		defer mutexNews.Unlock()
		r := make([]*News, len(newsStorage))
		for i := range newsStorage{
			n := newsStorage[i]
			r[i] = &n
		}
	return r
}
func GetNews(id string) *News {
	mutexNews.Lock()
		defer mutexNews.Unlock()
	for _, news := range newsStorage{
		if news.ID == id{
			n := news
			return &n
		}
	}
	return nil
}
func DeleteNews(id string) {
		mutexNews.Lock()
		defer mutexNews.Unlock()
	for i, news := range newsStorage{
		if news.ID == id{
			newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
			return
		}
	}
}