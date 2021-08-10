package htmlbear

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// tools for html

//GetHTML 获取网页源码
func GetHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error:status code", resp.StatusCode)
		return " "
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	res := string(all)
	return res
}

//ReGet 获取目标项
func ReGet(html, compile string) (res [][]string) {
	//re := regexp.MustCompile(`<p class="sitetext">([^<]+)<p class="sitetext">`)
	//解释规则，解析正则表达式，如果成功则返回
	re := regexp.MustCompile(compile)
	if re == nil {
		fmt.Println("error regexp")
		panic(re)
	}
	//根据规则提取信息

	result := re.FindAllStringSubmatch(html, -1)

	return result

}
