package alert

import (
	"AutoReservationSys/model"
	"AutoReservationSys/util"
	"encoding/json"
	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

const (
	fieldRoom       = "室場"
	fieldDate       = "日付"
	fieldTime       = "時間帯"
	fieldStats      = "状態"
	AlertText       = "<!channel>空いている・開放待ちの体育館を見つけました！"
	attachmentColor = "#3eb991"
	attachmentTitle = "施設："
)

type AlertRepository struct {
	hostname string
	url      string
}

func NewAlertRepository() AlertRepository {
	return AlertRepository{
		"SAKURA",
		"https://hooks.slack.com/services/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",
	}
}

func (n *AlertRepository) SendAlert(m model.BookableGymList) error {

	// リストのデータをすべて処理する
	for _, v := range m {

		roomAlert := Field{
			Title: fieldRoom,
			Value: v.Room,
			Short: true,
		}

		dateAlert := Field{
			Title: fieldDate,
			Value: v.Date,
			Short: true,
		}

		timeAlert := Field{
			Title: fieldTime,
			Value: v.Time,
			Short: true,
		}

		st := ""
		if v.Status == model.BookableStatusBookable {
			st = "予約可能"
		} else {
			st = "開放待ち"
		}
		statsAlert := Field{
			Title: fieldStats,
			Value: st,
			Short: true,
		}

		attachment := Attachment{
			Color:     attachmentColor,
			Field:     []Field{roomAlert, dateAlert, timeAlert, statsAlert},
			TimeStamp: time.Now().Unix(),
			Title:     attachmentTitle,
			Text:      v.Name,
		}

		alert := Alert{
			Text:       AlertText,
			Attachment: []Attachment{attachment},
		}

		// jsonテキストへ変換
		bytes, err := json.Marshal(&alert)
		if err != nil {
			return xerrors.Errorf("error converting to json: %w", err)
		}

		req, err := retryablehttp.NewRequest(http.MethodPost, n.url, bytes)
		if err != nil {
			return xerrors.Errorf("error requesting to slack: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")

		cli := util.Httpclient()

		_, err = cli.Do(req)

		if err != nil {
			return xerrors.Errorf("failed to send alert to slack: %w", err)
		}
	}

	return nil
}
