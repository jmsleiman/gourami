# gourami
gourami is a shell-friendly tool for exploring jwt claims

# how to use it
gourami will respond to stdin, and will print on stdout, so it should play nice with other shell tools.

Simply provide it with a `jwt` and it will decode it for you.

# example

```
$ echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ | ./gourami
{
        "alg": "HS256",
        "typ": "JWT"
}
{
        "admin": true,
        "name": "John Doe",
        "sub": "1234567890"
}
```

You can of course choose to run it directly (`$ ./gourami`) it will continue to read the input until it's stopped.

# what it does and doesn't do

*does*:

- show you the header of a jwt, in pretty-print json
- show you the claims of a jwt, in pretty-print json

*does not*:

- validate the struct of the jwt or claims (other than confirming that the data was base64 encoded and contained json)
- validate whether the token was correctly signed (it just interprets the two first parts which aren't signed)
- work with encrpyted tokens / JOSE

If you'd like any of the features that don't work, make a PR, or open an issue, and I can add them in.

# why did i bother building this

Because often times I need to check the claims of a token to make sure they were set properly, and I hated having a fragile workflow for getting the base64 part decoded and checking the json. You can chain this up with a tool like `jq` -- it might be interesting to introduce a flag such as `--claims` to get only the claims, and pipe that directly into `jq`.

# why go

Because go is awesome :)
