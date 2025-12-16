package main

type FigmaNode struct {
	ID               		string                 	`json:"id"`
	Name             		string                 	`json:"name"`
	Type             		string                 	`json:"type"`
	Children         		[]FigmaNode            	`json:"children,omitempty"`
	AbsoluteBoundingBox 	*BoundingBox        	`json:"absoluteBoundingBox,omitempty"`
	BackgroundColor  		*Color                 	`json:"backgroundColor,omitempty"`
	Fills            		[]Fill                 	`json:"fills,omitempty"`
	Strokes          		[]Stroke               	`json:"strokes,omitempty"`
	StrokeWeight     		float64                	`json:"strokeWeight,omitempty"`
	CornerRadius     		float64                	`json:"cornerRadius,omitempty"`
	Characters       		string                 	`json:"characters,omitempty"`
	Style            		*TextStyle             	`json:"style,omitempty"`
	LayoutMode       		string            		`json:"layoutMode,omitempty"`
	PrimaryAxisAlignItems 	string            		`json:"primaryAxisAlignItems,omitempty"`
	CounterAxisAlignItems 	string            		`json:"counterAxisAlignItems,omitempty"`
	PaddingLeft      		float64                	`json:"paddingLeft,omitempty"`
	PaddingRight     		float64                	`json:"paddingRight,omitempty"`
	PaddingTop       		float64                	`json:"paddingTop,omitempty"`
	PaddingBottom    		float64                	`json:"paddingBottom,omitempty"`
	ItemSpacing      		float64                	`json:"itemSpacing,omitempty"`
}

type FigmaFileResponse struct {
	Document	FigmaNode `json:"document"`
	Name    	string    `json:"name"`	
}
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