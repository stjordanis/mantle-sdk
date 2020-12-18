package generate

import (
	"github.com/terra-project/mantle-sdk/utils"
	"strings"
)

var ReservedArgKeys = []string{
	"Limit",
	"Order",
	"offset",
}

func FilterArgs(args map[string]interface{}, skip []string) map[string]interface{} {
	next := make(map[string]interface{})
	for argKey, argValue := range args {
		if utils.SliceContainsString(skip, argKey) {
			continue
		}

		next[argKey] = argValue
	}

	return next
}

func FilterRangedArgs(args map[string]interface{}) map[string]interface{} {
	next := make(map[string]interface{})
	for argKey, argValue := range args {
		if strings.HasSuffix(argKey, "_range") {
			next[argKey] = argValue
		}
	}

	return next
}

func FilterNonRangeArgs(args map[string]interface{}) map[string]interface{} {
	next := make(map[string]interface{})
	for argKey, argValue := range args {
		if !strings.HasSuffix(argKey, "_range") {
			next[argKey] = argValue
		}
	}

	return next
}
