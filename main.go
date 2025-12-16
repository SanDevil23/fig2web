package main

type FigmaClient struct {
	accessToken string
	baseURL    	string
}

func newFigmaClient(token string) *FigmaClient {
	return &FigmaClient{
		accessToken: 	token,
		baseURL: 		"https://api.figma.com/v1",		
	} 
}