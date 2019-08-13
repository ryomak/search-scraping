package search

import (
	"sync"
  "time"
  log "github.com/sirupsen/logrus"
)

func (conf *Config)AllSearch()*Result{
  words := conf.Word
  var searchResults []SearchResult
  for _,v := range words{
    searchResults = append(searchResults,*conf.GoogleWorker(v.Name))
  }
  return &Result{
    MyURL:conf.MyURL,
    SearchResults:searchResults,
  }
}

func (conf *Config) GoogleWorker(query string) *SearchResult {
	articles := Articles{}
	page := (conf.MaxNum / 100) + 1
	num := 100
	if page == 0 {
		num = conf.MaxNum
	}
	workerNum := 1
	var wg sync.WaitGroup
	q := make(chan int, 3)
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go articles.fetchWorker(&wg, conf, query, q, num)
	}
	for i := 1; i <= page; i++ {
    time.Sleep(3 * time.Second)
		q <- i
	}
	close(q)
	wg.Wait()
	return &SearchResult{
		Keyword:  query,
		Articles: articles,
	}
}

func (articles *Articles) fetchWorker(wg *sync.WaitGroup, conf *Config, query string, page chan int, num int) {
	defer wg.Done()
	for {
		s, ok := <-page
		if !ok {
			return
		}
    a ,err:= conf.GoogleFetch(query, s, num)
    if err != nil{
      log.Errorf("page:%v,Num:%v ",s,num)
      continue
    }
		*articles = append(*articles,a...)
	}
}
