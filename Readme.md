# http to md5
This project makes HTTP requests and prints the address of the request along with the MD5 hash of the response.
The maximum number of concurrent requests is defined by a flag `-parallel` and its default value is `10`.

## How to run
This project was tested on a machine with go version `go version go1.19.2 darwin/amd64`.
To run the project, follow the instructions below:

1. Build the project with the following command:
```shell
make build
```

2. Execute the binary passing URLs as arguments:
```
./myhttp google.com facebook.com yahoo.com
```

You can alternatively inform the limit of concurrent workers with the flag `parallel` :
```
./myhttp -parallel 3 google.com http://twitter.com facebook.com
```

  