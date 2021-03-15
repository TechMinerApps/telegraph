package telegraph_test

import (
	"log"
	"time"

	"github.com/TechMinerApps/telegraph"
)

// Content in a string format (for this example).
// Be sure to wrap every media in a <figure> tag, okay? Be easy.
const data = `
    <figure>
        <img src="/file/6a5b15e7eb4d7329ca7af.jpg"/>
    </figure>
    <p><i>Hello</i>, my name is <b>Page</b>, <u>look at me</u>!</p>
    <figure>
        <iframe src="https://youtu.be/fzQ6gRAEoy0"></iframe>
        <figcaption>
            Yes, you can embed youtube, vimeo and twitter widgets too!
        </figcaption>
    </figure>
`

var (
	account *telegraph.Account
	page    *telegraph.Page
	content []telegraph.Node
)

func errCheck(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Example_fastStart() {
	var err error
	// Create new Telegraph account.
	requisites := telegraph.Account{
		ShortName: "toby3d", // required

		// Author name/link can be epmty. So secure. Much anonymously. Wow.
		AuthorName: "Maxim Lebedev",       // optional
		AuthorURL:  "https://t.me/toby3d", // optional
	}
	client := telegraph.NewClient()
	account, err = client.CreateAccount(requisites)
	errCheck(err)

	// Make sure that you have saved acc.AuthToken for create new pages or make
	// other actions by this account in next time!

	// Format content to []telegraph.Node array. Input data can be string, []byte
	// or io.Reader.
	content, err = client.ContentFormat(data)
	errCheck(err)

	// Boom!.. And your text will be understandable for Telegraph. MAGIC.

	// Create new Telegraph page
	pageData := telegraph.Page{
		Title:   "My super-awesome page", // required
		Content: content,                 // required

		// Not necessarily, but, hey, it's just an example.
		AuthorName: account.AuthorName, // optional
		AuthorURL:  account.AuthorURL,  // optional
	}
	page, err = client.CreatePage(pageData, false)
	errCheck(err)

	// Show link from response on created page.
	log.Println("Kaboom! Page created, look what happened:", page.URL)
}

func Example_createAccount() {
	var err error
	client := telegraph.NewClient()
	account, err = client.CreateAccount(telegraph.Account{
		ShortName:  "Sandbox",
		AuthorName: "Anonymous",
	})
	errCheck(err)

	log.Println("AccessToken:", account.AccessToken)
	log.Println("AuthURL:", account.AuthorURL)
	log.Println("ShortName:", account.ShortName)
	log.Println("AuthorName:", account.AuthorName)
}

func Example_editAccountInfo() {
	var err error
	client := telegraph.NewClient()
	account, err = client.EditAccountInfo(telegraph.Account{
		ShortName:  "Sandbox",
		AuthorName: "Anonymous",
	})
	errCheck(err)

	log.Println("AuthURL:", account.AuthorURL)
	log.Println("ShortName:", account.ShortName)
	log.Println("AuthorName:", account.AuthorName)
}

func Example_getAccountInfo() {
	client := telegraph.NewClient()
	info, err := client.GetAccountInfo(
		telegraph.FieldShortName,
		telegraph.FieldPageCount,
	)
	errCheck(err)

	log.Println("ShortName:", info.ShortName)
	log.Println("PageCount:", info.PageCount, "pages")
}

func Example_revokeAccessToken() {
	var err error
	client := telegraph.NewClient()
	// You must rewrite current variable with account structure for further usage.
	account, err = client.RevokeAccessToken()
	errCheck(err)

	log.Println("AccessToken:", account.AccessToken)
}

func Example_createPage() {
	var err error
	client := telegraph.NewClient()
	page, err = client.CreatePage(telegraph.Page{
		Title:      "Sample Page",
		AuthorName: account.AuthorName,
		Content:    content,
	}, true)
	errCheck(err)

	log.Println(page.Title, "by", page.AuthorName, "has been created!")
	log.Println("PageURL:", page.URL)
}

func Example_editPage() {
	var err error

	client := telegraph.NewClient()
	page, err = client.EditPage(telegraph.Page{
		Title:      "Sample Page",
		AuthorName: account.AuthorName,
		Content:    content,
	}, true)
	errCheck(err)

	log.Println("Page on", page.Path, "path has been updated!")
	log.Println("PageURL:", page.URL)
}

func Example_getPage() {
	client := telegraph.NewClient()
	info, err := client.GetPage("Sample-Page-12-15", true)
	errCheck(err)

	log.Println("Getted info about", info.Path, "page:")
	log.Println("Author:", info.AuthorName)
	log.Println("Views:", info.Views)
	log.Println("CanEdit:", info.CanEdit)
}

func Example_getPageList() {
	client := telegraph.NewClient()
	list, err := client.GetPageList(0, 3)
	errCheck(err)

	log.Println("Getted", list.TotalCount, "pages")

	for i := range list.Pages {
		p := list.Pages[i]
		log.Printf("%s: %s\n~ %s\n\n", p.Title, p.URL, p.Description)
	}
}

func Example_getViews() {
	client := telegraph.NewClient()
	pagePath := "Sample-Page-12-15"
	dateTime := time.Date(2016, time.December, 0, 0, 0, 0, 0, time.UTC)
	views, err := client.GetViews(pagePath, dateTime)
	errCheck(err)

	log.Println(pagePath, "has been viewed", views.Views, "times")
}

func Example_contentFormat() {
	const data = `<figure>
<img src="http://telegra.ph/file/6a5b15e7eb4d7329ca7af.jpg" /></figure>
<p><i>Hello</i>, my name is <b>Page</b>, <u>look at me</u>!</p>
<figure><iframe src="https://youtu.be/fzQ6gRAEoy0"></iframe>
<figcaption>Yes, you can embed youtube, vimeo and twitter widgets too!</figcaption>
</figure>`

	var err error
	client := telegraph.NewClient()
	content, err = client.ContentFormat(data)
	errCheck(err)

	log.Printf("Content: %#v", content)
}
