package usecase

import (
	"assesment/domain"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func (uc *uc) GetTracking() (string, error) {
	var result = domain.Result{}
	stringHTMl, statusCode, err := uc.repo.GetTracking()
	if err != nil {
		return "", err
	}

	if statusCode != http.StatusOK {
		result.Status.Code = statusCode
		result.Status.Message = domain.MessageFailed

		resultJson, err := json.MarshalIndent(result, "", "  ")

		return string(resultJson), err
	}

	content := ReadHtml(stringHTMl)

	listHistory, err := ConvertToListHistory(content)
	if err != nil {
		return "", err
	}

	dat := listHistory[len(listHistory)-1].Description

	temp := strings.Split(dat, "[")
	temp = strings.Split(temp[1], "|")
	receivedBy := temp[0]

	listHistory = SortHistoryAsc(listHistory)

	var statusMessage string

	if statusCode == http.StatusOK {
		statusMessage = domain.MessageSucces
	} else {
		statusMessage = domain.MessageFailed
	}
	result = domain.Result{
		Status: domain.Status{
			Code:    statusCode,
			Message: statusMessage,
		},
		Data: domain.Data{
			ReceivedBy: receivedBy,
			Histories:  listHistory,
		},
	}

	resultJson, err := json.MarshalIndent(result, "", "  ")

	return string(resultJson), err
}

func ConvertToListHistory(content []string) ([]domain.History, error) {
	listHistoryTemp := []domain.History{}

	var createdAt string
	var createdAtFormated string
	var layoutFormat = "02-01-2006 15:04"
	for i, data := range content {
		if i%2 == 0 {
			var date, err = time.Parse(layoutFormat, data)
			if err != nil {
				return listHistoryTemp, err
			}

			createdAt = date.Format("2006-01-02T15:04:05+07:00")
			createdAtFormated = date.Format("02 January 2006, 15:04 WIB")
		} else {
			hist := domain.History{
				CreateAt:    createdAt,
				Description: data,
				Formatted: domain.Formatted{
					CreateAt: createdAtFormated,
				},
			}
			listHistoryTemp = append(listHistoryTemp, hist)

		}
	}

	return listHistoryTemp, nil
}

func ReadHtml(stringHtml string) []string {
	z := html.NewTokenizer(strings.NewReader(stringHtml))
	content := []string{}

	var i = 0
	var j = 0
	for z.Token().Data != "section" {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "table" && i < 3 {
				i++
			}
			if i == 3 {
				if t.Data == "td" {
					if j > 0 {
						inner := z.Next()

						if inner == html.TextToken {
							text := (string)(z.Text())
							t := strings.TrimSpace(text)
							content = append(content, t)
						}
					}
					j++
				}

			}

		}
	}

	return content
}

func SortHistoryAsc(listHistoryParam []domain.History) []domain.History {
	listHistory := []domain.History{}

	for itr := len(listHistoryParam) - 1; itr >= 0; itr-- {
		listHistory = append(listHistory, listHistoryParam[itr])
	}

	return listHistory
}
