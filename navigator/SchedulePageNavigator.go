package navigator

import (
	"AutoReservationSys/alert"
	"AutoReservationSys/model"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"golang.org/x/xerrors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type SchedulePageNavigator struct {
	Driver *agouti.WebDriver
	Page   *agouti.Page
}

func NewSchedulePageNavigator(wd *agouti.WebDriver, p *agouti.Page) SchedulePageNavigator {
	return SchedulePageNavigator{wd, p}
}

func (n *SchedulePageNavigator) Navigate() error {

	err := n.navigateToSearch()

	if err != nil {
		return xerrors.Errorf("failed to navigate: %w", err)
	}

	err = n.chooseFacilityCategory()

	if err != nil {
		return xerrors.Errorf("failed to navigate: %w", err)
	}

	err = n.choosePurpose()

	if err != nil {
		return xerrors.Errorf("failed to navigate: %w", err)
	}

	err = n.chooseTime()

	if err != nil {
		return xerrors.Errorf("failed to navigate: %w", err)
	}

	err = n.enterDateRange()

	if err != nil {
		return xerrors.Errorf("failed to navigate: %w", err)
	}

	err = n.selectWeekDate()

	if err != nil {
		return xerrors.Errorf("failed to navigate: %w", err)
	}

	n.startSearch()

	list, err := n.crawlInformation()

	if err != nil {
		return xerrors.Errorf("failed to navigate: %w", err)
	}

	err = n.sendAlert(list)

	if err != nil {
		return xerrors.Errorf("failed to navigate: %w", err)
	}

	return nil

}

func (n *SchedulePageNavigator) navigateToSearch() error {
	bt := n.Page.FindByID("RSGK001_99")
	if err := bt.Click(); err != nil {
		return xerrors.Errorf("Failed to click `空施設検索`:  %w", err)
	}

	return nil
}

func (n *SchedulePageNavigator) chooseFacilityCategory() error {
	bt := n.Page.FirstByXPath("//*[@id=\"tbl_kensaku\"]/tbody/tr[2]/td[1]/button")
	if err := bt.Click(); err != nil {
		return xerrors.Errorf("Failed to click `室場分類`:  %w", err)
	}

	bt = n.Page.FindByID("fbox_03")
	if err := bt.Click(); err != nil {
		return xerrors.Errorf("Failed to click `体育室`:  %w", err)
	}
	return nil
}

func (n *SchedulePageNavigator) choosePurpose() error {
	bt := n.Page.FirstByXPath("//*[@id=\"tbl_kensaku\"]/tbody/tr[3]/td[1]/button")
	if err := bt.Click(); err != nil {
		return xerrors.Errorf("Failed to click `利用目的`:  %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"FRM_RSGK301\"]/div[3]/div/div/p/select")
	if err := bt.Select("スポーツ"); err != nil {
		return xerrors.Errorf("Failed to click `利用目的 Selector`:  %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"fbox_0009\"]")
	if err := bt.Click(); err != nil {
		return xerrors.Errorf("Failed to click `バスケットボール`: %w", err)
	}
	return nil
}

func (n *SchedulePageNavigator) chooseTime() error {
	bt := n.Page.FirstByXPath("//*[@id=\"CHK_JIKANTAI_KBN12\"]")
	if err := bt.Check(); err != nil {
		return xerrors.Errorf("Failed to click `9:00`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_JIKANTAI_KBN14\"]")
	if err := bt.Check(); err != nil {
		return xerrors.Errorf("Failed to click `12:00`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_JIKANTAI_KBN16\"]")
	if err := bt.Check(); err != nil {
		return xerrors.Errorf("Failed to click `15:00`:: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_JIKANTAI_KBN18\"]")
	if err := bt.Check(); err != nil {
		return xerrors.Errorf("Failed to click `17:00`:: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_JIKANTAI_KBN20\"]")
	if err := bt.Check(); err != nil {
		return xerrors.Errorf("Failed to click `19:00`: %w", err)
	}

	return nil
}

func (n *SchedulePageNavigator) enterDateRange() error {

	t := time.Now()
	fd := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	ld := fd.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	bt := n.Page.FirstByXPath("//*[@id=\"TXT_TO_DAY\"]")
	if err := bt.Fill(strconv.Itoa(ld.Day())); err != nil {
		return xerrors.Errorf("Failed to enter `DAY`: %w", err)
	}
	return nil
}

func (n *SchedulePageNavigator) selectWeekDate() error {
	bt := n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN9\"]")
	if err := bt.Check(); err != nil {
		return xerrors.Errorf("Failed to click `土日祝`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN0\"]")
	if err := bt.Uncheck(); err != nil {
		return xerrors.Errorf("Failed to click `平日`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN1\"]")
	if err := bt.Uncheck(); err != nil {
		return xerrors.Errorf("Failed to click `日曜`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN2\"]")
	if err := bt.Uncheck(); err != nil {
		return xerrors.Errorf("Failed to click `月曜`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN3\"]")
	if err := bt.Uncheck(); err != nil {
		return xerrors.Errorf("Failed to click `火曜`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN4\"]")
	if err := bt.Uncheck(); err != nil {
		return xerrors.Errorf("Failed to click `水曜`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN5\"]")
	if err := bt.Uncheck(); err != nil {
		return xerrors.Errorf("Failed to click `木曜`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN6\"]")
	if err := bt.Uncheck(); err != nil {
		return xerrors.Errorf("Failed to click `金曜`: %w", err)
	}

	bt = n.Page.FirstByXPath("//*[@id=\"CHK_YOUBI_KBN7\"]")
	if err := bt.Uncheck(); err != nil {
		return xerrors.Errorf("Failed to click `土曜`: %w", err)
	}
	return nil
}

func (n *SchedulePageNavigator) startSearch() {
	bt := n.Page.FirstByXPath("//*[@id=\"footer\"]/div/button[2]")
	if err := bt.Click(); err != nil {
		log.Fatalf("Failed to click `検索`: %v", err)
	}
}

func (n *SchedulePageNavigator) crawlInformation() (model.BookableGymList, error) {

	curContentsDom, err := n.Page.HTML()
	if err != nil {
		log.Printf("Failed to get html: %v", err)
	}

	readerCurContents := strings.NewReader(curContentsDom)
	contentsDom, _ := goquery.NewDocumentFromReader(readerCurContents)

	// ページ数を確認する。必要であれば回数分ループする
	pagingInfo := contentsDom.Find("#FRM_RSGK351 > div:nth-child(17) > div > div > table.tbl_page > tbody > tr > td:nth-child(2) > strong")
	pagingText := pagingInfo.Text()
	prevPagingText := ""

	// 今までの予約状態を確認する
	prevList, err := n.readPrevList()
	if err != nil {
		log.Printf("could not read file: %v", err)
	}

	// 解析開始
	list := model.BookableGymList{}
	for prevPagingText != pagingText {

		// テーブルをクロールする
		tb := contentsDom.Find("#tbl_setsubi")
		tb.Find("tr").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			m := model.BookableGymModel{}
			s.Find("td").Each(func(i int, d *goquery.Selection) {
				switch i % model.BookableGymColumnNum {
				case model.BookableGymColumnName:
					m.Name = d.Text()
				case model.BookableGymColumnRoom:
					m.Room = d.Text()
				case model.BookableGymColumnDate:
					m.Date = d.Text()
				case model.BookableGymColumnTime:
					m.Time = d.Text()
				case model.BookableGymColumnButton:
					if strings.Contains(d.Text(), "開放待ち") {
						m.Status = model.BookableStatusNotReleased
					} else {
						m.Status = model.BookableStatusBookable
					}
				}
			})

			list = append(list, m)
		})

		// 次のページに移動してページ情報が変わらなければ処理を終了させる
		prevPagingText = pagingText

		bt := n.Page.FirstByXPath("//*[@id=\"FRM_RSGK351\"]/div[3]/div/div/table[4]/tbody/tr/td[3]/button[1]")
		if err := bt.Click(); err != nil {
			log.Fatalf("Failed to click `次のページ`: %v", err)
		}

		curContentsDom, err := n.Page.HTML()
		if err != nil {
			log.Printf("Failed to get html: %v", err)
		}
		readerCurContents := strings.NewReader(curContentsDom)
		contentsDom, _ = goquery.NewDocumentFromReader(readerCurContents)

		pagingInfo := contentsDom.Find("#FRM_RSGK351 > div:nth-child(17) > div > div > table.tbl_page > tbody > tr > td:nth-child(2) > strong")
		pagingText = pagingInfo.Text()
	}

	// 新規のデータを検索
	alertList := n.checkNewStateGym(prevList, list)

	bytes, err := json.Marshal(&list)
	if err != nil {
		return nil, xerrors.Errorf("failed to encode to json: %w", err)
	}

	err = ioutil.WriteFile("/var/tmp/schedule.json", bytes, 0644)
	if err != nil {
		return nil, xerrors.Errorf("failed to output json: %w", err)
	}
	return alertList, nil
}

func (n *SchedulePageNavigator) readPrevList() (model.BookableGymList, error) {
	raw, err := ioutil.ReadFile("/var/tmp/schedule.json")
	if err != nil {
		return nil, xerrors.Errorf("failed to read prev list: %w", err)
	}

	var prevList = model.BookableGymList{}
	err = json.Unmarshal(raw, &prevList)
	if err != nil {
		return nil, xerrors.Errorf("failed to unmarshal prev list: %w", err)
	}

	return prevList, nil
}

func (n *SchedulePageNavigator) checkNewStateGym(prevList model.BookableGymList, newList model.BookableGymList) model.BookableGymList {

	var list = model.BookableGymList{}

	for _, v := range newList {
		isExist := false
		for _, w := range prevList {

			if v.Name == w.Name &&
				v.Room == w.Room &&
				v.Time == w.Time &&
				v.Status == w.Status &&
				v.Date == w.Date {
				isExist = true
				break
			}
		}

		if !isExist {
			list = append(list, v)
		}
	}

	return list
}

func (n *SchedulePageNavigator) sendAlert(m model.BookableGymList) error {

	na := alert.NewAlertRepository()

	err := na.SendAlert(m)
	if err != nil {
		return xerrors.Errorf("failed to send schedule alert: %w", err)
	}

	return nil
}
