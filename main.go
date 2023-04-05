package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const mainGoTemplate = `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is your Go microservice!")
	})

	http.ListenAndServe(":8080", nil)
}
`

const gitignoreTemplate = `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output files
*.out
*.o

# Dependency directories
vendor/

`

func main() {
	goVersion := flag.String("go-version", "1.19", "Specify the Go version to use in the go.mod file")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: newgoapp [--go-version <version>] <module_path>")
		os.Exit(1)
	}

	modulePath := args[0]
	projectDir := filepath.Base(modulePath)

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	currentDirName := filepath.Base(currentDir)

	if currentDirName != projectDir {
		if err := os.Mkdir(projectDir, 0755); err != nil {
			fmt.Println("Error creating project directory:", err)
			os.Exit(1)
		}
		if err := os.Chdir(projectDir); err != nil {
			fmt.Println("Error changing to project directory:", err)
			os.Exit(1)
		}
	}

	goModContent := fmt.Sprintf("module %s\n\ngo %s\n", modulePath, *goVersion)

	if err := os.WriteFile("go.mod", []byte(goModContent), 0644); err != nil {
		fmt.Println("Error creating go.mod file:", err)
		os.Exit(1)
	}

	if err := os.WriteFile("main.go", []byte(mainGoTemplate), 0644); err != nil {
		fmt.Println("Error creating main.go file:", err)
		os.Exit(1)
	}

	gitignoreContent := fmt.Sprintf("%s\n%s\n", gitignoreTemplate, projectDir)
	if err := os.WriteFile(".gitignore", []byte(gitignoreContent), 0644); err != nil {
		fmt.Println("Error creating .gitignore file:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully created Go project skeleton at", filepath.Join(currentDir, projectDir))
}
