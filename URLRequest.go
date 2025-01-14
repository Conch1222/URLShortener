package main

type URLRequest struct {
	LongUrl    string `json:"long_url"`
	Expiration string `json:"expiration"`
}

func (urlRequest *URLRequest) Init(longUrl string, expiration string) {
	urlRequest.LongUrl = longUrl
	urlRequest.Expiration = expiration
}
