// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"github.com/q191201771/lal/pkg/hls"
)

func main() {
	hlsOutPath := "hls"
	fragmentDurationMs := 3000
	fragmentNum := 6
	hlsMuxerConfig := hls.MuxerConfig{
		OutPath:            hlsOutPath,
		FragmentDurationMs: fragmentDurationMs,
		FragmentNum:        fragmentNum,
	}

	streamName := "ddd"
	hlsMuexer := hls.NewMuxer(streamName, &hlsMuxerConfig, nil)
	hlsMuexer.Start()

}
