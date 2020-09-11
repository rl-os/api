package request

// BeatmapsetLookup contains beatmap_id from query params
type BeatmapsetLookup struct {
	Id uint `query:"beatmap_id"`
}

// BeatmapsetSearch params that used for filtering beatmaps
type BeatmapsetSearch struct {
	Query    string `json:"q" query:"q"`
	Mode     int    `json:"m" query:"m"`
	Status   string `json:"s" query:"s"`
	Genre    string `json:"g" query:"g"`
	Language string `json:"l" query:"l"`
	Sort     string `json:"sort" query:"sort"`
}
