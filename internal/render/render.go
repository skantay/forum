package render

// Render will be used to parse some data into html code
type Render struct {
}

// Specify which template you need, e.g. sign_up.html, and any data that will be parsed into the html code
func (r *Render) Render(template string, data interface{}) (string, error) {
	return "", nil
}
