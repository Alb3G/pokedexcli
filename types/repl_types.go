package types

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *Config) error
}

type Config struct {
	PreviousUrl string
	NextUrl     string
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
