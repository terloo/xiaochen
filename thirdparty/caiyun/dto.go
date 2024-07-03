package caiyun

type BaseBody struct {
	Status     string     `json:"status"`
	APIVersion string     `json:"api_version"`
	APIStatus  string     `json:"api_status"`
	Lang       string     `json:"lang"`
	Unit       string     `json:"unit"`
	Tzshift    int        `json:"tzshift"`
	Timezone   string     `json:"timezone"`
	ServerTime int        `json:"server_time"`
	Location   []float64  `json:"location"`
	Result     BaseResult `json:"result"`
}

type BaseResult struct {
	Daily   DailyResp `json:"daily"`
	Alert   AlertResp `json:"alert"`
	Primary int       `json:"primary,omitempty"`
}

// 告警
type AlertResp struct {
	Province      string    `json:"province"`
	Status        string    `json:"status"`
	Code          string    `json:"code"`
	Description   string    `json:"description"`
	RegionID      string    `json:"regionId"`
	County        string    `json:"county"`
	Pubtimestamp  int       `json:"pubtimestamp"`
	Latlon        []float64 `json:"latlon"`
	City          string    `json:"city"`
	AlertID       string    `json:"alertId"`
	Title         string    `json:"title"`
	Adcode        string    `json:"adcode"`
	Source        string    `json:"source"`
	Location      string    `json:"location"`
	RequestStatus string    `json:"request_status"`
}

// 小时预测
type HourlyResp struct {
	Status              string         `json:"status"`
	Description         string         `json:"description"`
	Precipitation       []WeatherValue `json:"precipitation"`
	Temperature         []WeatherValue `json:"temperature"`
	ApparentTemperature []WeatherValue `json:"apparent_temperature"`
	Wind                []Wind         `json:"wind"`
	Humidity            []WeatherValue `json:"humidity"`
	Cloudrate           []WeatherValue `json:"cloudrate"`
	Skycon              []WeatherValue `json:"skycon"`
	Pressure            []WeatherValue `json:"pressure"`
	Visibility          []WeatherValue `json:"visibility"`
	Dswrf               []WeatherValue `json:"dswrf"`
	AirQuality          AirQuality     `json:"air_quality"`
}

// 天预测
type DailyResp struct {
	Status              string         `json:"status"`
	Astro               []Astro        `json:"astro"`
	Precipitation08H20H []WeatherValue `json:"precipitation_08h_20h"`
	Precipitation20H32H []WeatherValue `json:"precipitation_20h_32h"`
	Precipitation       []WeatherValue `json:"precipitation"`
	Temperature         []WeatherValue `json:"temperature"`
	Temperature08H20H   []WeatherValue `json:"temperature_08h_20h"`
	Temperature20H32H   []WeatherValue `json:"temperature_20h_32h"`
	Wind                []Wind         `json:"wind"`
	Wind08H20H          []Wind         `json:"wind_08h_20h"`
	Wind20H32H          []Wind         `json:"wind_20h_32h"`
	Humidity            []WeatherValue `json:"humidity"`
	Cloudrate           []WeatherValue `json:"cloudrate"`
	Pressure            []WeatherValue `json:"pressure"`
	Visibility          []WeatherValue `json:"visibility"`
	Dswrf               []WeatherValue `json:"dswrf"`
	AirQuality          AirQuality     `json:"air_quality"`
	Skycon              []WeatherValue `json:"skycon"`
	Skycon08H20H        []WeatherValue `json:"skycon_08h_20h"`
	Skycon20H32H        []WeatherValue `json:"skycon_20h_32h"`
	LifeIndex           LifeIndex      `json:"life_index"`
}

// 日出日落时间
type Astro struct {
	Date    string `json:"date"`
	Sunrise struct {
		Time string `json:"time"`
	} `json:"sunrise"`
	Sunset struct {
		Time string `json:"time"`
	} `json:"sunset"`
}

// 风速
type Wind struct {
	Date string `json:"date"`
	Max  struct {
		Speed     float64 `json:"speed"`
		Direction float64 `json:"direction"`
	} `json:"max"`
	Min struct {
		Speed     float64 `json:"speed"`
		Direction float64 `json:"direction"`
	} `json:"min"`
	Avg struct {
		Speed     float64 `json:"speed"`
		Direction float64 `json:"direction"`
	} `json:"avg"`
}

// 生活指数
type LifeIndex struct {
	Ultraviolet []LifeIndexDesc `json:"ultraviolet"`
	CarWashing  []LifeIndexDesc `json:"carWashing"`
	Dressing    []LifeIndexDesc `json:"dressing"`
	Comfort     []LifeIndexDesc `json:"comfort"`
	ColdRisk    []LifeIndexDesc `json:"coldRisk"`
}

type LifeIndexDesc struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}

// 空气指数
type AirQuality struct {
	Aqi []struct {
		Date string `json:"date"`
		Max  struct {
			Chn int `json:"chn"`
			Usa int `json:"usa"`
		} `json:"max"`
		Avg struct {
			Chn int `json:"chn"`
			Usa int `json:"usa"`
		} `json:"avg"`
		Min struct {
			Chn int `json:"chn"`
			Usa int `json:"usa"`
		} `json:"min"`
	} `json:"aqi"`
	Pm25 []WeatherValue `json:"pm25"`
}

type WeatherValue struct {
	Date        string  `json:"date"`
	Value       string  `json:"value"`
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Avg         float64 `json:"avg"`
	Probability float64 `json:"probability"`
}
