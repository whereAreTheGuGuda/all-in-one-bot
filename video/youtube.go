package video

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func (v *VideoDownload) YoutubeAudioDownload(url string, startAndEnd ...string) {

	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		v.MsgC <- "出现异常,请重试"
		return
	}

	filename := v.path + video.Title

	formats := video.Formats.FindByItag(140)

	stream, _, err := client.GetStream(video, formats)
	if err != nil {
		v.MsgC <- "出现异常,请重试"
		fmt.Println(err)
		return
	}


	if _, err := os.Stat(v.path); os.IsNotExist(err) { // 检查目录是否存在
		err := os.Mkdir(v.path, os.ModePerm) // 创建目录
		if err != nil {
			fmt.Println("创建本地临时文件夹失败")
			v.MsgC <- "创建本地临时文件夹失败"
			return
		}
	}
	file, err := os.Create(filename + ".m4a")
	if err != nil {
		v.MsgC <- "出现异常,请重试"
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		v.MsgC <- "出现异常,请重试"
		fmt.Println(err)
		return
	}

	if len(startAndEnd) == 2 {
		err = v.Cut(filename + ".m4a", startAndEnd[0], startAndEnd[1], filename + "_cut.m4a")
		if err != nil {
			v.MsgC <- "请检查是否安装ffmpeg"
			fmt.Println(err)
			return
		}
		
		v.C <- filename + "_cut.m4a"
		return
	}

	v.AudioC <- filename + ".m4a"

}

func (v *VideoDownload) YoutubeDownload(url string, startAndEnd ...string) {

	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		v.MsgC <- "出现异常,请重试"
		fmt.Println(err)
		return
	}

	filename := v.path + video.Title

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		v.MsgC <- "出现异常,请重试"
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(v.path); os.IsNotExist(err) { // 检查目录是否存在
		err := os.Mkdir(v.path, os.ModePerm) // 创建目录
		if err != nil {
			fmt.Println("创建本地临时文件夹失败")
			v.MsgC <- "创建本地临时文件夹失败"
			return
		}
	}
	file, err := os.Create(filename + ".mp4")
	if err != nil {
		v.MsgC <- "出现异常,请重试"
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		v.MsgC <- "出现异常,请重试"
		fmt.Println(err)
		return
	}

	if len(startAndEnd) == 2 {
		err = v.Cut(filename + ".mp4", startAndEnd[0], startAndEnd[1], filename + "_cut.mp4")
		if err != nil {
			v.MsgC <- "请检查是否安装ffmpeg"
			fmt.Println(err)
			return
		}

		v.C <- filename + "_cut.mp4"
		return
	}

	v.C <- filename + ".mp4"
}
