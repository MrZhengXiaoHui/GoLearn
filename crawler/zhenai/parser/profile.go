package parser

import (
	"GoLearn/crawler/engine"
	"GoLearn/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(["离异"||"未婚"|"丧偶"])</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+座[^<]+)</div>`)
var genderRe = regexp.MustCompile(`(["男"||"女"])`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(["高中及以下"||"中专"|"大学本科"||"大专"])</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(月收入:[^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)kg</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)

func ParseProfile(contents []byte,name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name =name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}
	profile.Marriage = extractString(contents, marriageRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Education = extractString(contents, educationRe)
	profile.Income = extractString(contents, incomeRe)

	profile.Hokou = extractString(contents, hokouRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
