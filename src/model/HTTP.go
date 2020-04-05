package model

// HTTPConfig.go

type HTTPConfig struct {
	URL         *string     `json:"url,omitempty"`
	Header      *Header     `json:"header,omitempty"`
	ContentType *string     `json:"ContentType,omitempty"`
	Username    interface{} `json:"username"`
	Password    interface{} `json:"password"`
}


// http.go

type HTTP struct {
	Path             *string `json:"path,omitempty"`
	Type             *string `json:"type,omitempty"`
	ValidStatusCode  *int64  `json:"ValidStatusCode,omitempty"`
	ValidStringRegex *string `json:"ValidStringRegex,omitempty"`
	Header           *Header `json:"header,omitempty"`
	Params           *Params `json:"params,omitempty"`
	Body             *Body   `json:"body,omitempty"`
	Items            []Item  `json:"items"`
}

// header.go

type Header struct { }

// body.go

type Body struct {
	Type    *string `json:"type,omitempty"`
	Raw     *string `json:"raw,omitempty"`
	Encoded *bool   `json:"encoded,omitempty"`
	Form    *Form   `json:"form,omitempty"`
}


// form.go

type Form struct { }


// params.go

type Params struct { }

