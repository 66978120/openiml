package server

import (
	"k8s.io/klog/v2"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
}

func ParseFlags() Options {
	for index, arg := range os.Args {
		if strings.HasPrefix(arg, "-mode") {
			os.Args[index] = strings.Replace(arg, "-mode", "--mode", 1)
			break
		}
	}

	options := Options{}
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}
	klog.Infof("Parsed options: %v\n", options)
	return options
}
