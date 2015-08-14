package main

import (
	"fmt"

	"github.com/zenwerk/go-wave"
)

func main() {
	fpath := "./sample.wav"
	reader, err := wave.NewReader(fpath)
	if err != nil {
		panic(err)
	}

BREAKPOINT:
	for {
		_, err := reader.ReadSample()
		if err != nil {
			fmt.Println(err.Error())
			break BREAKPOINT
		}
	}
	if reader.NumSamples != reader.ReadSampleNum {
		fmt.Printf("実際のサンプル数: %d\n読んだサンプル数: %d\n", reader.NumSamples, reader.ReadSampleNum)
		fmt.Println(reader.NumSamples, reader.ReadSampleNum)
	} else {
		fmt.Println("正常に読み込みました")
	}

	// file info
	fmt.Printf("%#v\n", reader.RiffChunk)
	fmt.Printf("%#v\n", reader.FmtChunk)
	fmt.Printf("%#v\n", reader.FmtChunk.Data)
	fmt.Printf("%#v\n", reader.DataChunk)
}
