package parser

import (
	"GoLearn/concurrentVersion/engine"
	"GoLearn/concurrentVersion/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(["离异"||"未婚"|"丧偶"])</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+座[^<]+)</div>`)
var genderRe= regexp.MustCompile(`"genderString":"([^"]+)士"`)
var educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(月收入:[^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)kg</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)

func ParseProfile(contents []byte,name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name =name
	age, _ := strconv.Atoi(extractString(contents, ageRe))
	profile.Age = age
	height, _ := strconv.Atoi(extractString(contents, heightRe))
	profile.Height = height
	weight, _ := strconv.Atoi(extractString(contents, weightRe))
	profile.Weight = weight
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
