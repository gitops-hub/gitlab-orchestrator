package main

import (
	"fmt"

	// "gitlab.com/HnBI/shared-projects/gitops-bot/cmd"
	"github.com/gitops-hub/gitlab-orchestrator/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
