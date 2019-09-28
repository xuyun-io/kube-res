package utils

import (
	"fmt"
)

const (
	template = "%-20s%-22s%-16s%-20s\n"
)

func PrintHeader() {
	fmt.Printf(template, "APIVERSION", "KIND", "NAMESPACE", "NAME")

}

func Print(apiversion, resourcetype, namespace, name string) {
	fmt.Printf(template, apiversion, resourcetype, namespace, name)
}
