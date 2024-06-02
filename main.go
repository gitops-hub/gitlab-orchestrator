package main

import (
	"fmt"

	"github.com/gitops-hub/gitlab-orchestrator/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
