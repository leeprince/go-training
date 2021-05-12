package resolve

import (
	"github.com/leeprince/utils"
	"testing"
)

func TestRWMutex(t *testing.T) {
	path := "./"
	file := "TestRWMutex.txt"
	data := "prince"
	
	wg.Add(50000)
	for i := 1; i <= 50000; i++ {
		go func() {
			defer wg.Done()
			rwmutex.RLock()
			utils.ReadFileOfString(path+file)
			rwmutex.RUnlock()
		}()
	}
	
	wg.Add(5000)
	for i := 1; i <= 5000; i++ {
		go func() {
			defer wg.Done()
			rwmutex.Lock()
			utils.WrtiteFile(path, file, data)
			rwmutex.Unlock()
		}()
	}
	
	wg.Wait()
}
