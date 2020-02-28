package store

type Store interface {
	Beatmap() Beatmap
	BeatmapSet() BeatmapSet
}

type Beatmap interface {
	GetBeatmap(id uint) interface{}
	GetAllBeatmap(page int, limit int) []interface{}
	CreateBeatmap(from interface{}) interface{}
	UpdateBeatmap(id uint, from interface{}) interface{}
	DeleteBeatmap(id uint)
}

type BeatmapSet interface {
	Get(id uint) interface{}
	GetAll(page int, limit int) []interface{}
	Create(from interface{}) interface{}
	Update(id uint, from interface{}) interface{}
	Delete(id uint)
}
