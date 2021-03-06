package main

import (
	"fmt"
	"mgrep/worker"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"sync"

	arg "github.com/alexflint/go-arg"
)

func discoverDirs(wl *worklist.Worklist, path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Readdir error:", err)

		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			nextpath := filepath.Join(path, entry.Name())
			discoverDirs(wl, nextpath)
		} else {
			wl.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

var args struct {
	SearchTerm string `arg:"positional,required"`
	SearchDir  string `arg:"positional"`
}

func main() {
	arg.MustParse(&args)

	var workersWg sync.WaitGroup

	wl := worklist.New(100)

	results := make(chan worker.Result, 100)

	numWorkers := 10

	workersWg.Add(1)

	go func() {
		defer workersWg.Done()

		discoverDirs(&wl, args.SearchDir)

		wl.Finalize(numWorkers)
	}()

	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)

		go func() {
			defer workersWg.Done()
			for {
				workEntry := wl.Next()

				if workEntry.Path != "" {

				}
			}
		}()
	}

	fmt.Println("Results", results)
}
