

### Write an application that can accept images or binary files with metadata, like file name, and store them in permanent storage. The application should allow users to list stored images/binary files and their metadata.

* Have an API with documentation defined (openapi / README / curl examples)
* Have a permanent storage
* Simple authentication mechanism
* Being easy to run from cmd line
Have basic logging
* One unit test case
Very simple UI to list files (does not need to render the images)
* Is the system scalable ? Can two instances be safely run. 
* Any caching mechanism for retrieving file (either in-memory but easily changeable to remote, or remote cache like memcache/Redis)
* One E2E test which checks if API stores data and can retrieve same data


## Install 
https://github.com/casey/just

run `cp .env.sample .env` to create .env file with exmaple values

## Run
Run docker containers with `just docker`. After that you need to go to `http://localhost:9001` and generate new access key which need to be put in .env file.
```
    just server -> runs golang server
```
and in separate tab
```
    just web -> runs web application
```

## Endpoints
* `GET localhost:8080/api/files` -> list of uploaded files
* `POST localhost:8080/api/files/upload` -> uploads file (add `file` to FormData)


## Notes
Bucket in Minio is private so there is no option to open file in browser. To do that you need to go to minio admin and change bucket policy.
There are no tests, in normal mode it would be added for golang using testify and mock packages. For web using jest.
For e2e solution I'd write tests with Playwright.
Web is not pretty, for that I'd use tailwind and shad/cn libraries. Web also would require some love regarding splitting code into separate files to make its quality better.
