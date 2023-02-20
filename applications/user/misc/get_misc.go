package misc

import "math/rand"

func GetUserAvatar() string {
	link1 := "https://doutok-video.oss-cn-shanghai.aliyuncs.com/video/hear.png"
	link2 := "https://doutok-video.oss-cn-shanghai.aliyuncs.com/video/stop.png"

	r := rand.Int31n(100)
	if r%2 == 0 {
		return link1
	} else {
		return link2
	}

}
