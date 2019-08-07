// +build tools

package tools

import (
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/spf13/cobra/cobra"

	_ "golang.org/x/perf/cmd/benchstat"
)
