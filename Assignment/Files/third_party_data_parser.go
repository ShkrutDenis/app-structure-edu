package main

import "errors"

type Response any

type ThirdPartyParser struct {
	url string
}

func New3dPartyParser(url string) *ThirdPartyParser {
	return &ThirdPartyParser{url: url}
}

func (p *ThirdPartyParser) GetData(params any) (Response, error) {
	return p.doRequest()
}

func (p *ThirdPartyParser) GetAnotherData(params any) (Response, error) {
	return p.doRequest()
}

func (p *ThirdPartyParser) doRequest() (Response, error) {
	return nil, errors.New("third party isn't working")
}
