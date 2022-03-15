package text

import (
	"encoding/binary"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// DataPoint TODO
type DataPoint struct {
	Year      uint64
	Month     uint64
	Day       uint64
	Hour      uint64
	Minute    uint64
	Secondx10 uint64
	Bx        float64
	By        float64
	Bz        float64
	Tex100    int16
	Tfx100    int16
	Uinx10    uint64
	GPS       GPSStatus
}

// MarshalText TODO
func (dp *DataPoint) MarshalText() (text []byte, err error) {
	year := strconv.FormatUint(dp.Year, 10)
	month := strconv.FormatUint(dp.Month, 10)
	day := strconv.FormatUint(dp.Day, 10)
	hour := strconv.FormatUint(dp.Hour, 10)
	minute := strconv.FormatUint(dp.Minute, 10)
	second := fmt.Sprintf("%d.%d", dp.Secondx10/10, dp.Secondx10%10)
	bx := strconv.FormatFloat(dp.Bx, 'f', 3, 32)
	by := strconv.FormatFloat(dp.By, 'f', 3, 32)
	bz := strconv.FormatFloat(dp.Bz, 'f', 3, 32)
	te := fmt.Sprintf("%d.%d", dp.Tex100/100, dp.Tex100%100)
	tf := fmt.Sprintf("%d.%d", dp.Tfx100/100, dp.Tfx100%100)
	uin := fmt.Sprintf("%d.%d", dp.Secondx10/10, dp.Secondx10%10)
	gps := dp.GPS.String()

	text = []byte(strings.Join([]string{year, month, day, hour, minute, second, bx, by, bz, te, tf, uin, gps}, ", "))
	err = nil

	return
}

const reDataPoint string = "^(?P<year>\\d{4}})\\s(?P<month>\\d{2}})\\s" +
	"(?P<day>\\d{2})\\s(?P<hour>\\d{2}})\\s(?P<minute>\\d{2}})\\s" +
	"(?P<second>\\d{2}}\\.\\d)\\s(?P<bx>\\d*\\.\\d{3})\\s" +
	"(?P<by>\\d*\\.\\d{3})\\s(?P<bz>\\d*\\.\\d{3})\\s" +
	"(?P<te>\\d{2}\\.\\d{2})\\s(?P<tf>\\d{2}\\.\\d{2})\\s" +
	"(?P<uin>\\d{2}\\.\\d),\\s(?P<gps>(?:65)|(?:80))$"

// UnmarshalText TODO
func (dp *DataPoint) UnmarshalText(text []byte) (err error) {
	re := regexp.MustCompile(reDataPoint)
	m, err := groupmap(text, re)
	if err != nil {
		return
	}

	year := binary.LittleEndian.Uint64(m["year"])
	month := binary.LittleEndian.Uint64(m["month"])
	day := binary.LittleEndian.Uint64(m["day"])
	hour := binary.LittleEndian.Uint64(m["hour"])
	minute := binary.LittleEndian.Uint64(m["minute"])
	second10 := binary.LittleEndian.Uint64(
		[]byte(strings.Join(strings.Split(string(m["second"]), "."), "")),
	)
	bx1000 := binary.LittleEndian.Uint64(
		[]byte(strings.Join(strings.Split(string(m["bx"]), "."), "")),
	)
	by1000 := binary.LittleEndian.Uint64(
		[]byte(strings.Join(strings.Split(string(m["by"]), "."), "")),
	)
	bz1000 := binary.LittleEndian.Uint64(
		[]byte(strings.Join(strings.Split(string(m["bz"]), "."), "")),
	)
	te100 := binary.LittleEndian.Uint64(
		[]byte(strings.Join(strings.Split(string(m["te"]), "."), "")),
	)
	tf100 := binary.LittleEndian.Uint64(
		[]byte(strings.Join(strings.Split(string(m["tf"]), "."), "")),
	)
	uin10 := binary.LittleEndian.Uint64(
		[]byte(strings.Join(strings.Split(string(m["uin"]), "."), "")),
	)
	gps := binary.LittleEndian.Uint64(m["gps"])

	dp.Year = year
	dp.Month = month
	dp.Day = day
	dp.Hour = hour
	dp.Minute = minute
	dp.Secondx10 = second10
	dp.Bx = float64(bx1000) / 1000
	dp.By = float64(by1000) / 1000
	dp.Bz = float64(bz1000) / 1000
	dp.Tex100 = int16(te100)
	dp.Tfx100 = int16(tf100)
	dp.Uinx10 = uin10
	dp.GPS = GPSStatus(gps)

	return
}

// UnmarshalText2 TODO
func (dp *DataPoint) UnmarshalText2(text []byte) (err error) {
	fields := strings.Split(string(text), ", ")

	year, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		return
	}
	month, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return
	}
	day, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return
	}
	hour, err := strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		return
	}
	minute, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return
	}
	second := strings.Split(fields[5], ".")
	second10s, err := strconv.ParseUint(second[0], 10, 64)
	if err != nil {
		return
	}
	second1s, err := strconv.ParseUint(second[1], 10, 64)
	if err != nil {
		return
	}
	second10x := second10s*10 + second1s

	bx, err := strconv.ParseFloat(fields[6], 64)
	if err != nil {
		return
	}
	by, err := strconv.ParseFloat(fields[7], 64)
	if err != nil {
		return
	}
	bz, err := strconv.ParseFloat(fields[8], 64)
	if err != nil {
		return
	}

	te, err := strconv.ParseFloat(fields[9], 64)
	if err != nil {
		return
	}
	te100x := int16(te * 100)

	tf, err := strconv.ParseFloat(fields[10], 64)
	if err != nil {
		return
	}
	tf100x := int16(tf * 100)

	uin, err := strconv.ParseFloat(fields[11], 64)
	if err != nil {
		return
	}
	uin10x := uint64(uin * 10)

	gps, err := strconv.ParseInt(fields[12], 10, 64)
	if err != nil {
		return
	}

	dp.Year = year
	dp.Month = month
	dp.Day = day
	dp.Hour = hour
	dp.Minute = minute
	dp.Secondx10 = second10x
	dp.Bx = bx
	dp.By = by
	dp.Bz = bz
	dp.Tex100 = te100x
	dp.Tfx100 = tf100x
	dp.Uinx10 = uin10x
	dp.GPS = GPSStatus(gps)

	return nil
}

// GPSStatus TODO
type GPSStatus int

const (
	// Passive TODO
	Passive GPSStatus = iota
	// Active TODO
	Active
)

// String TODO
func (gpsstatus GPSStatus) String() string {
	var s string

	switch gpsstatus {
	case Passive:
		s = "80"
	case Active:
		s = "65"
	default:
		s = "Invalid"
	}

	return s
}
