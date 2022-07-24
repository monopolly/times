package times

//testing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)
	_ = a

	//No match
	var ts string
	var parsed *time.Time
	ts = "2023-05-05 00:00:00"
	parsed = Parse(ts)
	a.NotNil(parsed)
	a.Equal(2023, parsed.Year())

	ts = "2010-10-08T00:44:59Z"
	parsed = Parse(ts)
	a.NotNil(parsed)
	a.Equal(2010, parsed.Year())

	ts = "26-jan-2022"
	parsed = Parse(ts)
	a.NotNil(parsed)
	a.Equal(2022, parsed.Year())

	ts = "2018-01-11T05:00:00Z"
	parsed = Parse(ts)
	a.NotNil(parsed)
	a.Equal(2018, parsed.Year())

	ts = "09-nov-2017"
	parsed = Parse(ts)
	a.NotNil(parsed)
	a.Equal(2017, parsed.Year())

	ts = "20180207"
	parsed = Parse(ts)
	a.NotNil(parsed)
	a.Equal(2018, parsed.Year())

	ts = "20180319T074127Z"
	parsed = Parse(ts)
	a.NotNil(parsed)
	a.Equal(2018, parsed.Year())

}

func BenchmarkParse(b *testing.B) {
	ts := "2023-05-05 00:00:00"
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Parse(ts)

	}
}

func BenchmarkParse1(b *testing.B) {
	ts := "20180319T074127Z"
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Parse(ts)

	}
}

func BenchmarkParse2(b *testing.B) {
	ts := "20180207"
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Parse(ts)

	}
}

func BenchmarkParse3(b *testing.B) {
	ts := "09-nov-2017"
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Parse(ts)

	}
}
