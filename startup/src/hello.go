package main
import "fmt"
import "runtime"
func main() {
fmt.Printf("Hello, World\n")
fmt.Printf("OS: %s Architechture: %s\n",runtime.GOOS,runtime.GOARCH)
}
