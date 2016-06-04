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

If you're project contains a `Rakefile`, than T.T. will assume you want to run that suite using `rake`. Oh, but wait, there's a `Gemfile` too? Oh man, T.T. knows just what to do with that! He'll run rake for you using `bundle exec rake`, he's just that nice of a bear!

```
$ tt rake
```
