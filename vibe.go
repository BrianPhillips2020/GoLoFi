package main

import(
	"fmt"
	"runtime"
	"os/exec"
)



func main() {
	fmt.Println("gotchu bud")
	status := openBrowser()
	if status == false {
		fmt.Println("that's not very chill")
	}
	
}

// openBrowser tries to open the URL in a browser,
// and returns whether it succeed in doing so.
func openBrowser() bool {
	url := "https://www.youtube.com/watch?v=jfKfPfyJRdk"
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}