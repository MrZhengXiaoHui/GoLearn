package parser

import (
	"TestGo/crawler/engine"
	"TestGo/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(["离异"||"未婚"|"丧偶"])</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+座[^<]+)</div>`)
var genderRe = regexp.MustCompile(`(["男"||"女"])`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(["高中及以下"||"中专"|"大学本科"||"大专"])</div>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age,err:=strconv.Atoi(extractString(contents,ageRe))
	if err != nil {
		profile.Age=age
	}
	profile.Marriage=extractString(contents,marriageRe)
	profile.Xinzuo=extractString(contents,xinzuoRe)
	profile.Gender=extractString(contents,genderRe)
	profile.Education=extractString(contents,educationRe)

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
