package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("home.html"))
}

// ResumeHandler will direct the user to the
// resume page
func ResumeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("resume.html"))
}

// ContactHandler will direct the user to the
// contact page
func ContactHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("contact.html"))
}
