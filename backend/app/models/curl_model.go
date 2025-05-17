package models

// CurlRequestBase struct to describe CurlRequest object.
type CurlRequestBase struct {
	Name string `gorm:"type:varchar(255)" json:"name"`
	URL  string `gorm:"type:varchar(2048)" json:"url"`
}

type CurlRequestSecret struct {
	Method  string            `gorm:"type:varchar(10)" json:"method"`
	Headers map[string]string `gorm:"type:mediumblob;serializer:json" json:"headers"`
	Body    string            `gorm:"type:mediumblob" json:"body"`
	Timeout int64             `gorm:"type:int" json:"timeout"`
	Raw     string            `gorm:"type:mediumblob" json:"raw"`
}

type CurlRequestListItem struct {
	Base
	CurlRequestBase
}

// TableName переопределяет название таблицы для CurlRequestListItem на `curl_requests`
func (CurlRequestListItem) TableName() string {
	return "curl_requests"
}

type CurlRequest struct {
	Base
	CurlRequestBase
	CurlRequestSecret
}
