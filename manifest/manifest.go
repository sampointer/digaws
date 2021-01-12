package manifest

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/sampointer/digaws/fetcher"
)

type manifest struct {
	doc  string
	lock sync.Mutex
}

var m manifest

func init() {
	m.update()
	m.cron()
}

//GetManifest is a thread safe Getter for the manifest document
func GetManifest() io.Reader {
	m.lock.Lock()
	r := strings.NewReader(m.doc)
	m.lock.Unlock()
	return r
}

func (m *manifest) cron() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("update of AWS manifest begins")
				m.update()
				fmt.Println("update of AWS manifest ends")
			}
		}
	}()
}

func (m *manifest) update() {
	res, err := fetcher.Fetch()
	if err != nil {
		fmt.Println("failed to update AWS manifest")
		return
	}

	b, err := ioutil.ReadAll(res)
	if err != nil {
		fmt.Println("failed to read AWS manifest during update")
		return
	}

	m.lock.Lock()
	m.doc = string(b)
	m.lock.Unlock()
}
