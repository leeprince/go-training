package resolve

import (
	"fmt"
	"github.com/leeprince/utils"
	"testing"
)

func TestRWMutex(t *testing.T) {
	path := "./"
	file := "TestRWMutex.txt"
	data := "prince\r\n"
	
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func() {
			defer wg.Done()
			rwmutex.RLock()
			s, _ := utils.ReadFileOfString(path+file)
			fmt.Println("ReadFileOfString: ", s)
			rwmutex.RUnlock()
		}()
	}
	
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func() {
			defer wg.Done()
			rwmutex.Lock()
			bool, err := utils.WrtiteFile(path, file, data)
			if bool == false || err != nil {
				fmt.Println("WrtiteFile err.", bool, err)
			}
			fmt.Println("WrtiteFile successfuly")
			rwmutex.Unlock()
		}()
	}
	
	wg.Wait()
}
