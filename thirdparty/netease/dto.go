package netease

type PageResult struct {
	Code int
	More bool
}

// PlayList 用户歌单
type PlayList struct {
	Id          int
	Name        string
	Status      int
	SpecialType int
}

// Song 歌曲详情
type Song struct {
	IdName
	Ar []IdName
	Al IdName
}

type SongDetailResult struct {
	PageResult
	Songs []Song
}

type PlayListResult struct {
	PageResult
	PlayList []PlayList
}

type PlayListDetailsResult struct {
	PageResult
	PlayList struct {
		TrackIds []struct {
			IdName
		}
	}
}

type IdName struct {
	Id   int
	Name string
}

type Music struct {
	IdName
	Artist []string
	Album  string
}
