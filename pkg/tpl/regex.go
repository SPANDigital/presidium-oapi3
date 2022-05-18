package tpl

import "regexp"

var NonSluggableRe = regexp.MustCompile(`(?mi)\W+`)
