package main

import (
	"os"

	"math"

	"github.com/zenwerk/go-wave"
)

func main() {
	f, err := os.Create("./write_test.wav")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	param := wave.WriterParam{
		Out:            f,
		WaveFormatType: 1,
		Channel:        1,
		SampleRate:     44100,
		BitsPerSample:  16,
	}

	w, err := wave.NewWriter(param)

	amplitude := 0.1
	hz := 440.0
	length := param.SampleRate * 1

	for i := 0; i < length; i++ {
		_data := amplitude * math.Sin(2.0*math.Pi*hz*float64(i)/float64(param.SampleRate))
		_data = (_data + 1.0) / 2.0 * 65536.0
		if _data > 65535.0 {
			_data = 65535.0
		} else if _data < 0.0 {
			_data = 0.0
		}
		data := uint16(_data+0.5) - 32768 // 四捨五入とデータのオフセット調整
		_, err = w.WriteSample16([]uint16{data})
		if err != nil {
			panic(err)
		}
	}

	defer w.Close()
	if err != nil {
		panic(err)
	}
}
