package main

import (
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// pubsub is a component that does a simple pubsub on ipfs. A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type autonomy struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "pubsub World!" is displayed as a heading.
func (a *autonomy) Render() app.UI {
	return app.Div().Class("l-application").Role("presentation").Body(
		app.Link().Rel("stylesheet").Href("https://assets.ubuntu.com/v1/vanilla-framework-version-3.8.0.min.css"),
		app.Link().Rel("stylesheet").Href("https://use.fontawesome.com/releases/v6.2.0/css/all.css"),
		app.Link().Rel("stylesheet").Href("/app.css"),
		app.Section().Class("p-strip--image is-light").Body(
			app.Div().Class("row u-vertically-center").Body(
				app.Div().Class("col-6").Body(
					app.H1().Text("Cyber Autonomy"),
					app.P().Text("P2P self-governance society prototype researching the intersection of moneyless economy, liquid democracy and p2p media."),
					app.Button().Text("Read the paper").OnClick(a.openModal),
				).Style("max-width", "fit-content!important"),
				app.Div().Class("col-6 u-hide--small u-align--center").Body(
					app.Img().Src("web/self-governance.png").Alt(""),
				).Style("max-width", "fit-content!important"),
			),
		),
		app.Section().Class("p-strip--image is-light").Body(
			app.Div().Class("row u-vertically-center").Body(
				app.Div().Class("col-6").Body(
					app.H1().Text("Cyber Stasis"),
					app.P().Text("Post-money economy simulator in the form of a free fictional game based on gift economy that tests the hypothesis of having a market system without any exchange be it barters or money."),
					app.A().Href("https://ipfs.io/ipfs/QmVbgUuNz28nmMn63spZRydzDQneLewwUCGLtZGBnSEx8U").Text("Try now"),
				).Style("max-width", "fit-content!important"),
				app.Div().Class("col-6 u-hide--small u-align--center").Body(
					app.Img().Src("web/cyber-stasis.png").Alt(""),
				).Style("max-width", "fit-content!important"),
			),
		),
		app.Section().Class("p-strip--image is-light").Body(
			app.Div().Class("row u-vertically-center").Body(
				app.Div().Class("col-6").Body(
					app.H1().Text("Cyber Acid"),
					app.P().Text("Liquid democracy political simulator based on the automated data feed from the moneyless economy simulator Cyber Stasis."),
					app.A().Href("https://ipfs.io/ipfs/QmUGE4LKpRboFi2wNRi35588gRKpgM22R5k96ab46jbK9g").Text("Try now"),
				).Style("max-width", "fit-content!important"),
				app.Div().Class("col-6 u-hide--small u-align--center").Body(
					app.Img().Src("web/cyber-acid.png").Alt(""),
				).Style("max-width", "fit-content!important"),
			),
		),
		app.Section().Class("p-strip--image is-light").Body(
			app.Div().Class("row u-vertically-center").Body(
				app.Div().Class("col-6").Body(
					app.H1().Text("Cyber Witness"),
					app.P().Text("Decentralized media simulator based on the reporter and witnesses concept. It aims to be an alternative to mass media, surveillance, censorship, fact checkers and centralized control of news."),
					app.A().Href("https://ipfs.io/ipfs/Qmb2sKM6j3zgBe8MfLfFHmEecrqEeT3JGZxWWzk5GFNYNs").Text("Try now"),
				).Style("max-width", "fit-content!important"),
				app.Div().Class("col-6 u-hide--small u-align--center").Body(
					app.Img().Src("web/cyber-witness.png").Alt(""),
				).Style("max-width", "fit-content!important"),
			),
		),
		app.Div().Class("p-modal").ID("paper-modal").Style("display", "none").Body(
			app.Section().Class("p-modal__dialog").Role("dialog").Aria("modal", true).Aria("labelledby", "modal-title").Aria("describedby", "modal-description").Body(
				app.Header().Class("p-modal__header").Body(
					app.H2().Class("p-modal__title").ID("modal-title").Text("Why do we need autonomy"),
					app.Button().Class("p-modal__close").Aria("label", "Close active modal").Aria("controls", "modal").OnClick(a.closeModal),
				),
				app.Div().Class("p-heading-icon--small").Body(
					app.Aside().Class("p-accordion").Body(
						app.Ul().Class("p-accordion__list").Body(
							app.Li().Class("p-accordion__group").Body(
								app.Div().Role("heading").Aria("level", "3").Class("p-accordion__heading").Body(
									app.Button().Type("button").Class("p-accordion__tab").ID("tab1").Aria("controls", "tab1-section").Aria("expanded", true).Text("Introduction").Value("tab1-section").OnClick(a.toggleAccordion),
								),
								app.Section().Class("p-accordion__panel").ID("tab1-section").Aria("hidden", false).Aria("labelledby", "tab1").Body(
									app.P().Text("Any system based on growth gets bigger and more centralized the more advanced it becomes. This happens naturally in the form of optimizations and convenience. Think of the transition from grocery stores to supermarkets and malls. Centralization makes scalability, planning and predictability easier. The downside is that individuals lose autonomy."),
								),
							),
							app.Li().Class("p-accordion__group").Body(
								app.Div().Role("heading").Aria("level", "3").Class("p-accordion__heading").Body(
									app.Button().Type("button").Class("p-accordion__tab").ID("tab2").Aria("controls", "tab2-section").Aria("expanded", true).Text("How do we get back autonomy without losing automation and scalability").Value("tab2-section").OnClick(a.toggleAccordion),
								),
								app.Section().Class("p-accordion__panel").ID("tab2-section").Aria("hidden", true).Aria("labelledby", "tab1").Body(
									app.P().Text("The most advanced technology we have to date to solve the problem is peer-to-peer technology. It allows for individuals to interact directly and maintain a horizontal structure."),
								),
							),
							app.Li().Class("p-accordion__group").Body(
								app.Div().Role("heading").Aria("level", "3").Class("p-accordion__heading").Body(
									app.Button().Type("button").Class("p-accordion__tab").ID("tab3").Aria("controls", "tab3-section").Aria("expanded", true).Text("What can autonomy do and what it can't").Value("tab3-section").OnClick(a.toggleAccordion),
								),
								app.Section().Class("p-accordion__panel").ID("tab3-section").Aria("hidden", true).Aria("labelledby", "tab3").Body(
									app.P().Text("Autonomy is basically individual freedom. The freedom to decide everything on a personal level and to also participate in global decision making equally with all other participants. Autonomy is the opposite of control and as such relies on learning from mistakes rather than preventing them from happening."),
								),
							),
							app.Li().Class("p-accordion__group").Body(
								app.Div().Role("heading").Aria("level", "3").Class("p-accordion__heading").Body(
									app.Button().Type("button").Class("p-accordion__tab").ID("tab4").Aria("controls", "tab4-section").Aria("expanded", true).Text("Where is autonomy applicable").Value("tab4-section").OnClick(a.toggleAccordion),
								),
								app.Section().Class("p-accordion__panel").ID("tab4-section").Aria("hidden", true).Aria("labelledby", "tab4").Body(
									app.P().Text("Practically everywhere you rely on an institution, organization, structure or a governing body to make decisions on your behalf."),
								),
							),
							app.Li().Class("p-accordion__group").Body(
								app.Div().Role("heading").Aria("level", "3").Class("p-accordion__heading").Body(
									app.Button().Type("button").Class("p-accordion__tab").ID("tab5").Aria("controls", "tab4-section").Aria("expanded", true).Text("What happens if we lose all our autonomy").Value("tab5-section").OnClick(a.toggleAccordion),
								),
								app.Section().Class("p-accordion__panel").ID("tab5-section").Aria("hidden", true).Aria("labelledby", "tab5").Body(
									app.P().Text("We will effectively be enslaved as individuals. We will still have a facade of rights but they will not be practiced effectively in reality."),
								),
							),
							app.Li().Class("p-accordion__group").Body(
								app.Div().Role("heading").Aria("level", "3").Class("p-accordion__heading").Body(
									app.Button().Type("button").Class("p-accordion__tab").ID("tab6").Aria("controls", "tab6-section").Aria("expanded", true).Text("Support us").Value("tab6-section").OnClick(a.toggleAccordion),
								),
								app.Section().Class("p-accordion__panel").ID("tab6-section").Aria("hidden", true).Aria("labelledby", "tab6").Body(
									app.A().Href("https://opencollective.com/stateless-minds-collective").Text("https://opencollective.com/stateless-minds-collective"),
								),
							),
							app.Li().Class("p-accordion__group").Body(
								app.Div().Role("heading").Aria("level", "3").Class("p-accordion__heading").Body(
									app.Button().Type("button").Class("p-accordion__tab").ID("tab7").Aria("controls", "tab7-section").Aria("expanded", true).Text("Terms of service").Value("tab7-section").OnClick(a.toggleAccordion),
								),
								app.Section().Class("p-accordion__panel").ID("tab7-section").Aria("hidden", true).Aria("labelledby", "tab7").Body(
									app.Div().Class("p-card").Body(
										app.H3().Text("Introduction"),
										app.P().Class("p-card__content").Text("Cyber Autonomy is a full P2P self-governance society prototype researching the intersection of moneyless economy, liquid democracy and p2p media. By using the application you are implicitly agreeing to share your peer id with the IPFS public network."),
									),
									app.Div().Class("p-card").Body(
										app.H3().Text("Application Hosting"),
										app.P().Class("p-card__content").Text("Cyber Autonomy is a decentralized website and is hosted on a public peer to peer network. By using the application you agree to host it on the public IPFS network free of charge for as long as your usage is."),
									),
									app.Div().Class("p-card").Body(
										app.H3().Text("External Links"),
										app.P().Class("p-card__content").Text("Cyber Autonomy contains links to its sister projects Cyber Stasis, Cyber Acid and Cyber Witness and depends on their availability to function properly."),
									),
								),
							),
							app.Li().Class("p-accordion__group").Body(
								app.Div().Role("heading").Aria("level", "3").Class("p-accordion__heading").Body(
									app.Button().Type("button").Class("p-accordion__tab").ID("tab8").Aria("controls", "tab7-section").Aria("expanded", true).Text("Privacy policy").Value("tab8-section").OnClick(a.toggleAccordion),
								),
								app.Section().Class("p-accordion__panel").ID("tab8-section").Aria("hidden", true).Aria("labelledby", "tab8").Body(
									app.Div().Class("p-card").Body(
										app.H3().Text("Personal data"),
										app.P().Class("p-card__content").Text("There is no personal information collected within Cyber Autonomy. We store a small portion of your peer ID encrypted as a non-unique identifier which is used for displaying the ranks interface."),
									),
									app.Div().Class("p-card").Body(
										app.H3().Text("Coookies"),
										app.P().Class("p-card__content").Text("Cyber Autonomy does not use cookies."),
									),
									app.Div().Class("p-card").Body(
										app.H3().Text("Changes to this privacy policy"),
										app.P().Class("p-card__content").Text("This Privacy Policy might be updated from time to time. Thus, it is advised to review this page periodically for any changes. You will be notified of any changes from this page. Changes are effective immediately after they are posted on this page."),
									),
								),
							),
						),
					),
				),
			),
		),
	).Style("display", "block").Style("overflow-y", "scroll")
}

func (a *autonomy) openModal(ctx app.Context, e app.Event) {
	app.Window().GetElementByID("paper-modal").Set("style", "display:flex")
}

func (a *autonomy) closeModal(ctx app.Context, e app.Event) {
	app.Window().GetElementByID("paper-modal").Set("style", "display:none")
}

func (a *autonomy) toggleAccordion(ctx app.Context, e app.Event) {
	id := ctx.JSSrc().Get("value").String()
	attr := app.Window().GetElementByID(id).Get("attributes")
	aria := attr.Get("aria-hidden").Get("value").String()
	if aria == "false" {
		app.Window().GetElementByID(id).Call("setAttribute", "aria-hidden", "true")
	} else {
		app.Window().GetElementByID(id).Call("setAttribute", "aria-hidden", "false")
	}
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &autonomy{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".

	withGz := gziphandler.GzipHandler(&app.Handler{
		Name:        "cyber-autonomy",
		Description: "Cyber Autonomy - P2P self-governance society prototype",
		Styles: []string{
			"https://assets.ubuntu.com/v1/vanilla-framework-version-3.8.0.min.css",
			"https://use.fontawesome.com/releases/v6.2.0/css/all.css",
		},
		Scripts: []string{},
	})
	http.Handle("/", withGz)

	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}
