// Filename: chanPrinter.go
// Usage: $ go run chanPrinter.go dptpFileName.go
// GPLv3 2024, MobilityReadingGroup, University of Oxford

package chanprinter

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func init() {

	rand.Seed(time.Now().UnixNano())

}

// Map from channels to printable format
var ChanTbl struct {
	mu  sync.Mutex
	Tbl map[any]any
}

// Get elem associated to channel
func getC(k any) any {
	ChanTbl.mu.Lock()
	defer ChanTbl.mu.Unlock()

	if val, ok := ChanTbl.Tbl[k]; ok {
		return val
	}

	return nil
}

// Set pair (channel, elem)
func SetC(k any, v any) {
	ChanTbl.mu.Lock()
	defer ChanTbl.mu.Unlock()

	ChanTbl.Tbl[k] = v
}

// Make base value and register it
func MakeB(v any) any {
	SetC(v, v)
	return any(v)
}

// Apply function @f to all elements of list @ls
func sliceApp(f func(any) any, ls []any) []any {
	res := make([]any, len(ls))
	for i, a := range ls {
		res[i] = f(a)
	}
	return res
}

// Printf @str with verbs %s referring to channels @k
func Print(str string, k ...any) {
	fmt.Printf(str, sliceApp(getC, k)...)
}

// Printf @str with verbs referring to values @k
// Use %s for string base values
// Use %d for int base values, etc
func PrintB(str string, k ...any) {
	fmt.Printf(str, k...)
}

// Generate fresh channel names from seed @s
func genName(s string) string {
	return s + strconv.Itoa(rand.Intn(999))
}

// Generate short channel names from seed @s
func GenNameS(s string) string {
	return s + strconv.Itoa(rand.Intn(99))
}
