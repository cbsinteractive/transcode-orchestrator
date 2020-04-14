package model
import (
	"time"
)

// Only one of those properties may be set: fontSize, fontSizeExpression.
type TextFilter struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	Font TextFilterFont `json:"font,omitempty"`
	// If set to true a box is drawn around the text using the background color.
	Box *bool `json:"box,omitempty"`
	// The width of the box drawn around the text.
	BoxBorderWidth *int32 `json:"boxBorderWidth,omitempty"`
	// The background color to be used for drawing the box.
	BoxColor string `json:"boxColor,omitempty"`
	// Line spacing of the border around the box in pixels
	LineSpacing *int32 `json:"lineSpacing,omitempty"`
	// Width of the border around the text
	BorderWidth *int32 `json:"borderWidth,omitempty"`
	// If set to true, it will fix text coordinates to avoid clipping if necessary
	FixBounds *bool `json:"fixBounds,omitempty"`
	// The color to be used to draw the text
	FontColor string `json:"fontColor,omitempty"`
	// Font size to be used to draw the text
	FontSize *int32 `json:"fontSize,omitempty"`
	// An expression for the Font size. Either fontSize or fontSizeExpression can be set at the same time. The following variables are valid: main_h, h, H for input height and main_w, w, W for the input_width
	FontSizeExpression string `json:"fontSizeExpression,omitempty"`
	// If set, alpha blending for the text is applied. Values are valid between 0.0 and 1.0.
	Alpha *int32 `json:"alpha,omitempty"`
	// Color of the shadow
	ShadowColor string `json:"shadowColor,omitempty"`
	// X offset of the shadow
	ShadowX *int32 `json:"shadowX,omitempty"`
	// Y offset of the shadow
	ShadowY *int32 `json:"shadowY,omitempty"`
	// If set, the timecode representation in \"hh:mm:ss[:;.]ff\" format will be applied
	Timecode string `json:"timecode,omitempty"`
	// String to be drawn
	Text string `json:"text,omitempty"`
	// X position of the text. Also an expression can be used. The following variables are valid: line_h - height of each text line; main_h - input height; main_w - input width; n - number of input frame; text_h - Text height; text_w - Text width (required)
	X string `json:"x,omitempty"`
	// Y position of the text. Also an expression can be used. The following variables are valid: line_h - height of each text line; main_h - input height; main_w - input width; n - number of input frame; text_h - Text height; text_w - Text width (required)
	Y string `json:"y,omitempty"`
}
func (o TextFilter) FilterType() FilterType {
    return FilterType_TEXT
}

