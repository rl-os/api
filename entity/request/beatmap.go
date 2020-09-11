package request

type BeatmapLookup struct {
	Id       uint   `query:"id"`
	CheckSum string `query:"checksum"`
	Filename string `query:"filename"`
}

type GetBeatmapScores struct {
	Type string `query:"type"`
	Mode string `query:"mode"`
}
