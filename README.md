# go-learning
Go lang learning repository. 
Premise: This document focuses on the development go application with VS code.

## Emvironment
1. Install go extention for VS code.
- Search "Go" and install fundamental Go extention as follows:      
    - https://github.com/Microsoft/vscode-go
2. Install required package including debugger.
- Open command palette(Ctrl+Shift+P) and search "Go: Install/Update tools"
- Select all tools.
    - "All tools successfully installed. You're ready to Go :)."
    - It takes several minutes to complete install tools.

## Operation Comfirmation
- Make hello.go file and run (go run hello.go).

```
[hello.go]

package main 

import "fmt"

func main(){
	fmt.Printf("Hello World\n");
}
```