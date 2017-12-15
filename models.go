package main

// AjaxModel -
type AjaxModel struct {
	Type string    `json:"type" binding:"required"`
	Data DataModel `json:"data" binding:"required"`
}

// DataModel -
type DataModel struct {
	Class string     `json:"class" binding:"required"`
	ID    int        `json:"id" binding:"required"`
	Model PhoneModel `json:"model" binding:"required"`
}

// PhoneModel -
type PhoneModel struct {
	Phone string `json:"phone" binding:"required"`
}

// Car -
type Car struct {
	Brand  string
	Name   string
	Year   string
	City   string
	Volume string
}

func (c *Car) getHeader() string {
	return c.Brand + " " + c.Name + " " + c.Year
}
