package datadog

import (
	"fmt"
	"regexp"
	"sort"
)

// DataDog allows '.', '_' and alphas only.
// If we don't validate this here then the datadog error logs can fill up disk really quickly
// noinspection RegExpRedundantEscape
var nameRegex = regexp.MustCompile(`[^\._a-zA-Z0-9]+`)

func formatName(name string) string {
	return nameRegex.ReplaceAllString(fmt.Sprint(prefix, name), "_")
}

func (tags Tags) StringSlice() []string {
	var stringSlice []string
	for k, v := range tags {
		if k != "" && v != "" {
			stringSlice = append(stringSlice, formatName(k)+":"+v)
		}
	}
	sort.Strings(stringSlice)
	return stringSlice
}

func mergeTags(tagsSlice []Tags) Tags {
	merged := Tags{}
	for k, v := range globalTags {
		merged[formatName(k)] = formatName(v)
	}
	for _, tags := range tagsSlice {
		for k, v := range tags {
			merged[formatName(k)] = formatName(v)
		}
	}
	return merged
}
