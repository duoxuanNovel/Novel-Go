package model

type ZincSearchResponse struct {
	Took     int         `json:"took"`
	TimedOut bool        `json:"timed_out"`
	MaxScore float64     `json:"max_score"`
	Hits     Hits        `json:"hits"`
	Buckets  interface{} `json:"buckets"`
	Error    string      `json:"error"`
}

type Hits struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

type Total struct {
	Value int `json:"value"`
}

type Hit struct {
	Index     string      `json:"-"`
	Type      string      `json:"-"`
	ID        string      `json:"-"`
	Score     float64     `json:"-"`
	Timestamp string      `json:"-"`
	Source    interface{} `json:"_source"`
}

type Source struct {
	Athlete    string `json:"Athlete"`
	City       string `json:"City"`
	Country    string `json:"Country"`
	Discipline string `json:"Discipline"`
	Event      string `json:"Event"`
	Gender     string `json:"Gender"`
	Medal      string `json:"Medal"`
	Season     string `json:"Season"`
	Sport      string `json:"Sport"`
	Year       int    `json:"Year"`
}
