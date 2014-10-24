package controllers

import "github.com/revel/revel"

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    return c.Render()
}

func (c App) Report(email string, description string) revel.Result {
    c.Validation.Required(email).Message("Your name is required!")
    c.Validation.MinSize(email, 3).Message("Your name is not long enough!")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
    }
    // Here we'd dispatch a specially formulated GET request to the
    // perl script that powers the Maine DOT form.
    // We may also want to save the request to a database and send
    // a confirmation email to folks about their successful report.
    return c.Render(email, description)
}
