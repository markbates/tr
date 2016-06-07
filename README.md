# T.T. (The Bear)

 ![](http://images.buycostumes.com/mgen/merchandiser/grateful-dead-blue-dancing-bear-deluxe-adult-costume-bc-805288.jpg?zm=250,250,1,0,0)

__T.T. is smarter than you. I know this because he is smarter than me. He knows how to run my tests. He helps me when I'm tired. He helps me when I can't think for my self. T.T. wants to help you.__

## What is T.T?

T.T. is a tool that runs the tests for you application in a unified way. You don't have to remember if it's a `go` or a `ruby` project, or any other type of project for that matter. T.T. just knows, and all you have to do is type `tt`, and he'll do the rest! T.T. will even remember what you ran, when you ran it, what the output was! He's a genius!!

## Installation

```
$ go get github.com/markbates/tt
```

### Usage

All you should really need to do is let T.T. loose on your test suite:

```
$ tt
```

Any arguments you pass to T.T. will be thoughtfully passed along to the appropriate test runner underneath. You can also run a specific test runner, should you desire, see the "Supported Test Runners" section for more details.

## Supported Test Runners

Right now T.T. only supports a handful of test runners, but with your help we'll get more!

### Rake

__Key file:__ *Rakefile*

If you're project contains a `Rakefile`, than T.T. will assume you want to run that suite using `rake`. Oh, but wait, there's a `Gemfile` too? Oh man, T.T. knows just what to do with that! He'll run rake for you using `bundle exec rake`, he's just that nice of a bear!

```
$ tt <any flags/args the rake command takes>
// or
$ tt rake <any flags/args the rake command takes>
```

### Go

__Key file:__ **_.test.go*

If you're project contains any `_test.go` files, than T.T. will run `go test` for you. Oh, and yes, he is very aware of that pesky `vendor` folder and knows enough to stay well and truly away from it!

```
$ tt <any flags/args the go test command takes>
// or
$ tt go <any flags/args the go test command takes>
```

#### The `-run` flag

Yes, T.T. knows about the `-run` flag. If you tell him to run a particular test of tests, he will!

```
$ tt go -run Hello ./models
```

### Make

__Key file:__ *Makefile*

Are you a masochist? Do you love your old `make` files? Well, you want to know a secret? T.T. loves them too. They remind him of a simplier time.

```
$ tt <any flags/args your Makefile takes>
// or
$ tt make
```

If you do want to use a `Makefile`, than T.T. asked me to let you know that you have a `test` target in the `Makefile`, otherwise, he'll just be down right confused and won't know what to do!

### NPM

__Key file:__ *package.json*

T.T. knows that you young kids like to use NPM for everything. He doesn't quite understand why, but he has your back. If there is a `package.json` file in your app then `T.T.` will respond by running `npm test` for you.

```
$ tt <any flags/args your ./test.sh file takes>
// or
$ tt npm <any flags/args your ./test.sh file takes>
```

### Ruby (individual files)

T.T. knows that sometimes when you're running a Rails/Ruby application that you sometimes just want to run a very specific test file. That's super easy for T.T.!

```
$ tt ruby <path/to/file.rb>
```

T.T. is such a clever bear that he even knows when you're not on a test file, and will try and find the test file for you. For example, if you try to run `app/models/user.rb` T.T. will look for either `test/models/user_test.rb` or `spec/models/user_spec.rb` and try to run that for you. I told you, he's a very, very clever bear. :)

### Custom (`./test.sh`)

__Key file:__ *./test.sh*

It is hard for poor old T.T. to keep up with all of the latest and greatest new testing tools, programming languages, and reality TV stars out there. So, if T.T. doesn't know how to handle you're favorite way of running tests, don't worry! Just drop a `./test.sh` file in your app and you are good to go!

```
$ tt <any flags/args your ./test.sh file takes>
// or
$ tt sh <any flags/args your ./test.sh file takes>
```

## History

T.T. has an incredible memory! He can remember all of the different test runs you've done! Seriously, it's a pretty cool party trick.

### Listing History

```
$ tt history
// or
$ tt h
```

This will output something along the lines of the following:

```bash
1)  Tue Jun 07 17:14:35 -0400 2016  | PASS
2)  Tue Jun 07 17:15:47 -0400 2016  | PASS
3)  Tue Jun 07 17:16:26 -0400 2016  | FAIL
4)  Tue Jun 07 17:16:33 -0400 2016  | PASS
```

#### Verbose

You can also pass the `-v` flag to get more verbose details on what was ran:

```bash
$ tt h -v
1)  Tue Jun 07 17:14:35 -0400 2016  | PASS
    go test github.com/markbates/tt github.com/markbates/tt/cmd github.com/markbates/tt/cmd/models
2)  Tue Jun 07 17:15:47 -0400 2016  | PASS
    go test github.com/markbates/tt github.com/markbates/tt/cmd github.com/markbates/tt/cmd/models
3)  Tue Jun 07 17:16:26 -0400 2016  | FAIL
    go test github.com/markbates/tt github.com/markbates/tt/cmd github.com/markbates/tt/cmd/models
4)  Tue Jun 07 17:16:33 -0400 2016  | PASS
    go test github.com/markbates/tt github.com/markbates/tt/cmd github.com/markbates/tt/cmd/models
```

### Replay History

```
$ tt history <n>
// or
$ tt h <n>
```

This command will not actually run the commands again, but will rather show you the output from when the command was originally run. This is really useful for recalling what tests failed.

### Replay Last Run

```
$ tt history last
// or
$ tt h last
```
This command will not actually run the last command again, but will rather show you the output from when the command was originally run. This is really useful for recalling what tests failed.

### Clear History

```
$ tt history clear
// or
$ tt h clear
```

### History as JSON

Almost all of the History commands, with the exception of `clear` accept the `-j` flag which will output the history as JSON instead of plain text.

```text
$ tt h -j
[
  {
    "id": 3,
    "time": "2016-06-07T17:16:26.867353264-04:00",
    "cmd": [
      "go",
      "test",
      "github.com/markbates/tt",
      "github.com/markbates/tt/cmd",
      "github.com/markbates/tt/cmd/models"
    ],
    "results": "?   \tgithub.com/markbates/tt\t[no test files]\nok  \tgithub.com/markbates/tt/cmd\t0.078s\n--- FAIL: Test_History_Save (0.00s)\n\tassertions.go:225: \r                        \r\tError Trace:\thistory_test.go:39\n\t\t\r\tError:\t\tShould be false\n\t\t\r\nFAIL\nFAIL\tgithub.com/markbates/tt/cmd/models\t0.023s\n",
    "error": "exit status 1",
    "exit_code": 1
  },
  {
    "id": 4,
    "time": "2016-06-07T17:16:33.761527184-04:00",
    "cmd": [
      "go",
      "test",
      "github.com/markbates/tt",
      "github.com/markbates/tt/cmd",
      "github.com/markbates/tt/cmd/models"
    ],
    "results": "?   \tgithub.com/markbates/tt\t[no test files]\nok  \tgithub.com/markbates/tt/cmd\t0.075s\nok  \tgithub.com/markbates/tt/cmd/models\t0.015s\n",
    "error": "",
    "exit_code": 0
  }
]
```
