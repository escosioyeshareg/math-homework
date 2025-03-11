
package main
import(
"math/rand"
"time"
)
func main(){
rand.Seed(time.Now().UnixNano())
// random number between 1 and 100
fmt.Println(rand.Intn(100))
}