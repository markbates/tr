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
//
$ tt make
```

If you do want to use a `Makefile`, than T.T. asked me to let you know that you have a `test` target in the `Makefile`, otherwise, he'll just be down right confused and won't know what to do!

### Custom (`./test.sh`)

__Key file:__ *./test.sh*

It is hard for poor old T.T. to keep up with all of the latest and greatest new testing tools, programming languages, and reality TV stars out there. So, if T.T. doesn't know how to handle you're favorite way of running tests, don't worry! Just drop a `./test.sh` file in your app and you are good to go!

```
$ tt <any flags/args your ./test.sh file takes>
//
$ tt sh <any flags/args your ./test.sh file takes>
```

## History

T.T. has an incredible memory! He can remember all of the different test runs you've done! Seriously, it's a pretty cool party trick.

### Listing History

```
$ tt history
```

### Replay History

```
$ tt history <n>
```

### Replay Last Run

```
$ tt history last
```

### Clear History

```
$ tt history clear
```
