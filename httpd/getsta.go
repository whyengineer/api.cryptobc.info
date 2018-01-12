package httpd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/whyengineer/api.cryptobc.info/caculate"
)

func GetStaStatus(c echo.Context) error {
	type Pairs struct {
		Pairs []string `json:"pairs"`
		Plat  []string `json:"plat"`
	}
	var p Pairs
	p.Pairs = append(p.Pairs, CalRes.M.Pairs...)
	p.Plat = append(p.Plat, CalRes.M.ExP...)
	return c.JSON(http.StatusOK, &p)
}
func GetStaData(c echo.Context) error {
	//ts, _ := strconv.Atoi(c.QueryParam("ts"))
	coin := c.QueryParam("coin")
	plat := c.QueryParam("plat")
	year, _ := strconv.Atoi(c.QueryParam("year"))
	month, _ := strconv.Atoi(c.QueryParam("month"))
	day, _ := strconv.Atoi(c.QueryParam("day"))
	timetype := c.QueryParam("type")
	num, _ := strconv.Atoi(c.QueryParam("num"))
	switch timetype {
	case "min1":
		hour, _ := strconv.Atoi(c.QueryParam("hour"))
		min, _ := strconv.Atoi(c.QueryParam("min"))
		var d []caculate.Min1TradeTable
		tmp := fmt.Sprintf("%04d%02d%02d%02d%02d", year, month, day, hour, min)
		key, _ := strconv.Atoi(tmp)
		CalRes.Db.Debug()
		if CalRes.Db.Where("time_key<=? AND prop=? AND coin_type=?",
			key, plat, coin).Order("time_key desc").
			Limit(num).Find(&d).RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, &d)
		}
	case "min5":
		hour := c.QueryParam("hour")
		min := c.QueryParam("min")
		var d caculate.Min5TradeTable
		if CalRes.Db.Where("year = ? AND month = ? AND day =? AND hour=? AND min=? AND prop=? AND coin_type=?",
			year, month, day, hour, min, plat, coin).
			Find(&d).RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, &d)
		}
	case "min30":
		hour := c.QueryParam("hour")
		min := c.QueryParam("min")
		var d caculate.Min30TradeTable
		if CalRes.Db.Where("year = ? AND month = ? AND day =? AND hour=? AND min=? AND prop=? AND coin_type=?",
			year, month, day, hour, min, plat, coin).
			Find(&d).RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, &d)
		}
	case "hour1":
		hour := c.QueryParam("hour")
		var d caculate.Hour1TradeTable
		if CalRes.Db.Where("year = ? AND month = ? AND day =? AND hour=? AND prop=? AND coin_type=?",
			year, month, day, hour, plat, coin).
			Find(&d).RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, &d)
		}
	case "hour4":
		hour := c.QueryParam("hour")
		var d caculate.Hour4TradeTable
		if CalRes.Db.Where("year = ? AND month = ? AND day =? AND hour=? AND prop=? AND coin_type=?",
			year, month, day, hour, plat, coin).
			Find(&d).RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, &d)
		}
	case "day":
		var d caculate.DayTradeTable
		if CalRes.Db.Where("year = ? AND month = ? AND day =? AND prop=? AND coin_type=?",
			year, month, day, plat, coin).
			Find(&d).RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, &d)
		}

	}
	return nil
}

func GetSencodeData(c echo.Context) error {
	ts, _ := strconv.Atoi(c.QueryParam("ts"))
	coin := c.QueryParam("coin")

	for i := ts; i > ts-500; i-- {
		key := coin + ":" + strconv.FormatInt(int64(i), 10)
		data, ok := CalRes.HotData[key]
		if ok {
			return c.JSON(http.StatusOK, &data)
		}
	}
	return c.NoContent(http.StatusNoContent)
}
