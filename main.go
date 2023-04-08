package main

import (
	"fmt"
	"os"

	"github.com/yannh/kubeconform/cmd/kubeconform"
	"github.com/yannh/kubeconform/pkg/config"
)

var version = "development"

func main() {
	cfg, out, err := config.FromFlags(os.Args[0], os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parsing command line: %s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(kubeconform.Validate(cfg, out))
}
