# u2imgur
> Uploads images to imgur

*u2imgur* is a silly utility that uploads an URL, a path or a stream to [imgur]().

## Installing
```bash
$ go get github.com/victorgama/u2imgur
```

## Configuring
Sadly it requires a little configuration. Bear with me.

1. Create a new [OAuth App]()
2. Add it to your environment variables using the keys `U2IMGUR_CLIENT_ID` and `U2IMGUR_CLIENT_SECRET`.
3. Run `u2imgur config` and follow the instructions provided by the application. This will generate a new `.u2imgur` on your `$HOME`, containing authentication data.
4. You're good to go!

## Usage
You can upload a file, a stream or an URL. Check it out:

1. Uploading a file
```bash
$ u2imgur /Users/victorgama/dog.gif
Upload complete. Access it through the following URL: http://i.imgur.com/dbrlb5W.gif
$ # yay
```

2. Uploading file through a stream
```bash
$ u2imgur < /Users/victorgama/dog.gif
Upload complete. Access it through the following URL: http://i.imgur.com/dbrlb5W.gif
$ # yay
```

3. `curl`ing an url into it
```bash
$ curl http://barkpost.com/wp-content/uploads/2013/03/oie_5181838bU3HJXJp.gif | u2imgur
Upload complete. Access it through the following URL: http://i.imgur.com/H5G6zBx.gif
```

4. Providing an URL directly to it
```bash
$ u2imgur http://66.media.tumblr.com/tumblr_mbfie8INpM1ro1jt0o1_500.gif
Upload complete. Access it through the following URL: http://i.imgur.com/CGWrSnl.gif
```

5. It also plays nice when piping data to other utilities, like [goom](https://github.com/victorgama/goom):
```bash
$ u2imgur http://66.media.tumblr.com/tumblr_mbfie8INpM1ro1jt0o1_500.gif | xargs goom gif confusedaussie
Goom! confusedaussie in gif is http://i.imgur.com/3QMcOiB.gif. Move along!
```

## License
```
MIT License

Copyright (c) 2016 Victor Gama

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
