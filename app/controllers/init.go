package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"html/template"
	"github.com/koreset/homefnew/app/utils"
)

func init() {
	fmt.Println("Testing intgration here.....")

	revel.TemplateFuncs["unsafeHtml"] = func(s string) template.HTML {
		return template.HTML(s)
	}

	revel.TemplateFuncs["stripSummaryTags"] = func(s string) string {
		return utils.RemoveAllTags(s)
	}

	revel.TemplateFuncs["displayDate"] = func(s int32) string {

		return utils.DisplayDate(int64(s))
	}
}
