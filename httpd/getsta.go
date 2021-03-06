package httpd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/whyengineer/api.cryptobc.info/caculate"
	"github.com/whyengineer/api.cryptobc.info/market"
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
		hour, _ := strconv.Atoi(c.QueryParam("hour"))
		min, _ := strconv.Atoi(c.QueryParam("min"))
		var d []caculate.Min5TradeTable
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
	case "min30":
		hour, _ := strconv.Atoi(c.QueryParam("hour"))
		min, _ := strconv.Atoi(c.QueryParam("min"))
		var d []caculate.Min30TradeTable
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
	case "hour1":
		hour, _ := strconv.Atoi(c.QueryParam("hour"))
		var d []caculate.Hour1TradeTable
		tmp := fmt.Sprintf("%04d%02d%02d%02d", year, month, day, hour)
		key, _ := strconv.Atoi(tmp)
		CalRes.Db.Debug()
		if CalRes.Db.Where("time_key<=? AND prop=? AND coin_type=?",
			key, plat, coin).Order("time_key desc").
			Limit(num).Find(&d).RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, &d)
		}
	case "hour4":
		hour, _ := strconv.Atoi(c.QueryParam("hour"))
		var d []caculate.Hour4TradeTable
		tmp := fmt.Sprintf("%04d%02d%02d%02d", year, month, day, hour)
		key, _ := strconv.Atoi(tmp)
		CalRes.Db.Debug()
		if CalRes.Db.Where("time_key<=? AND prop=? AND coin_type=?",
			key, plat, coin).Order("time_key desc").
			Limit(num).Find(&d).RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, &d)
		}
	case "day":
		var d []caculate.DayTradeTable
		tmp := fmt.Sprintf("%04d%02d%02d", year, month, day)
		key, _ := strconv.Atoi(tmp)
		CalRes.Db.Debug()
		if CalRes.Db.Where("time_key<=? AND prop=? AND coin_type=?",
			key, plat, coin).Order("time_key desc").
			Limit(num).Find(&d).RecordNotFound() {
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
		data, ok := CalRes.HotData.Load(key)
		if ok {
			dataa := data.(market.CoinInfo)
			return c.JSON(http.StatusOK, &dataa)
		}
	}
	return c.NoContent(http.StatusNoContent)
}
