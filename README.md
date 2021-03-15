# Golang library for Telegraph API

[![codecov](https://codecov.io/gh/TechMinerApps/telegraph/branch/main/graph/badge.svg?token=DWO6DPMEHM)](https://codecov.io/gh/TechMinerApps/telegraph)

> This project is just to provide a wrapper around the API without any additional features.

Caution! This library is under refactoring, interface may change.

## Start using telegraph
Download and install it:

```
$ go get -u github.com/TechMinerApps/telegraph
```

Import it in your code:
```
import "github.com/TechMinerApps/telegraph"
```

## Document

See [GoDoc](https://pkg.go.dev/github.com/TechMinerApps/telegraph)

QuickStart
```

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
	if err != nil {
        // Do error check
    }

	// Make sure that you have saved acc.AuthToken for create new pages or make
	// other actions by this account in next time!

	// Format content to []telegraph.Node array. Input data can be string, []byte
	// or io.Reader.
	content, err = client.ContentFormat(data)
	if err != nil {
        // Do error check
    }

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
	if err != nil {
        // Do error check
    }

	// Show link from response on created page.
	log.Println("Kaboom! Page created, look what happened:", page.URL)
}

```

## Credit
Original developer: [@toby3d](https://gitlab.com/toby3d/telegraph)

## License

This Project is licensed under MIT License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FTechMinerApps%2Ftelegraph.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FTechMinerApps%2Ftelegraph?ref=badge_large)