# Thingiverse + Go Web App Sample

This sample demonstrates how to add authentication to a Go web app using Thingiverse.
## Running the App

To run the app, make sure you have **go** and **go get** installed.

Rename the `.env.example` file to `.env` and provide your Thingiverse credentials.

```bash
# .env

TV_CLIENT_ID={CLIENT_ID}
TV_DOMAIN={DOMAIN}
TV_CLIENT_SECRET={CLIENT_SECRET}
TV_CALLBACK_URL=http://localhost:3000/callback
TV_AUDIENCE=
```

__Note:__ If you are not implementing any API, leave the `TV_AUDIENCE` variable empty, will be set with `https://TV_DOMAIN/userinfo`.

Once you've set your TV credentials in the `.env` file, run `go get -d` to install the Go dependencies.

Run `go run main.go server.go` to start the app and navigate to [http://localhost:3000/](http://localhost:3000/).


## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE.txt) file for more info.
