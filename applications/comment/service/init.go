package service

import "github.com/TremblingV5/DouTok/pkg/safeMap"

func InitMemoryMap() {
	comCount := safeMap.New()
	comContent := safeMap.New()

	ComCount = comCount
	ComContent = comContent
}
