package main

type URLRequest struct {
	LongUrl    string `json:"long_url"`
	Expiration int    `json:"expiration"`
}

func (urlRequest *URLRequest) Init(longUrl string, expiration int) {
	urlRequest.LongUrl = longUrl
	urlRequest.Expiration = expiration
}
