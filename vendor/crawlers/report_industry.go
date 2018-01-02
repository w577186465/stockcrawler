package crawlers

import (
	"crawler"
	"crypto/md5"
	"fmt"
	"github.com/bitly/go-simplejson"
	"crawlers/models"
	"strings"
	"time"
)

type ReportIndustry struct {
	Pagesize int
}

func (m *ReportIndustry) Run() {
	set := m
	set.Pagesize = 200
	pagesize := set.Pagesize
	fmt.Println("正在抓取行业研报")
	pageNum, ok := m.pages(pagesize) // 获取分页数
	if !ok {
		fmt.Println("获取分页失败")
		fmt.Println("行业研报抓取失败")
		return
	}

	m.parsehyreport(pageNum)

}

// 获取分页数量
func (m *ReportIndustry) pages(pagesize int) (int, bool) {
	url := fmt.Sprintf(`http://datainterface.eastmoney.com//EM_DataCenter/js.aspx?type=SR&sty=HYSR&mkt=0&stat=0&cmd=4&code=&sc=&ps=%d&p=%d&js={"data":[(x)],"pages":(pc),"update":"(ud)","count":(count)}`, 1, 1)
	data, err := crawler.Download(url).Json()
	if err != nil {
		return 0, false
	}
	num, err := data.Get("count").Int() // 数据总数

	// 计算分页数量
	if num/pagesize > int(num/pagesize) {
		return int(num/pagesize) + 1, true
	} else {
		return num / pagesize, true
	}
}

func (m ReportIndustry) list(page int) (*simplejson.Json, error) {
	pagesize := m.Pagesize
	url := fmt.Sprintf(`http://datainterface.eastmoney.com//EM_DataCenter/js.aspx?type=SR&sty=HYSR&mkt=0&stat=0&cmd=4&code=&sc=&ps=%d&p=%d&js={"data":[(x)],"pages":(pc),"update":"(ud)","count":(count)}`, pagesize, page)
	return crawler.Download(url).Json()
}

func (m ReportIndustry) parsehyreport(pageNum int) bool {
	for page := 1; page <= pageNum; page++ {
		data, err := m.list(page)
		if err != nil {

		}

		arr, err := data.Get("data").Array() // 获取data
		if err != nil {
			fmt.Println(err)
			return true
		}

		for _, v := range arr {
			// 定义模块
			var m crawler.Module
			m.Name = "industry"

			item := v.(string)

			var report models.ReportIndustry

			arr := strings.Split(item, ",")
			t, _ := time.ParseInLocation("2006/1/2 15:04:05", arr[1], time.Local) // 将时间转换为时间类型
			day := t.Format("20060102")                                           // 生成详情页地址时间

			report.Pjchange = arr[0] // 评级变动
			report.CreatedAt = t
			report.Insname = arr[4] // 机构名称
			report.Indid = arr[6]   // 行业id
			report.Pjtype = arr[7]  // 评级类型
			report.Expect = arr[8]  // 看好
			report.Title = arr[9]
			report.Indname = arr[10]     // 行业名称
			report.Fluctuation = arr[11] // 涨跌幅

			report.Hash = fmt.Sprintf("%x", md5.Sum([]byte(arr[2]+report.Indname+report.Pjchange+report.Pjtype+report.Expect+day))) // 生成hash

			contenturl := fmt.Sprintf("http://data.eastmoney.com/report/%s/hy,%s.html", day, arr[2])
			report.Content = getcontent(contenturl)
			m.Addlink(contenturl, report.Hash)
		}
	}

	return true
}

func getcontent(url string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	html, err := crawler.Request(url).Retry(5).Transcoding("gbk").Download().Html()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	content, err := html.Find(".newsContent").Html()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return content
}
