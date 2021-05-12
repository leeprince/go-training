package resolve

import (
	"github.com/leeprince/utils"
	"testing"
)

func TestMutex(t *testing.T) {
	path := "./"
	file := "TestMutex.txt"
	data := "prince"
	
	wg.Add(50000)
	for i := 1; i <= 50000; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			utils.ReadFileOfString(path + file)
			mutex.Unlock()
		}()
	}
	
	wg.Add(5000)
	for i := 1; i <= 5000; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			utils.WrtiteFile(path, file, data)
			mutex.Unlock()
		}()
	}
	
	wg.Wait()
}
