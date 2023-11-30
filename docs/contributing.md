# Contributing
## How to contribute
First of all, thank you for your interest in contributing to salvo. 
We are always looking for ways to improve the project and we appreciate your help.

### Here are a few ways you can help out:
#### Reporting bugs
If you find a bug or a behavior that is not expected, please open an issue on the 
[Salvo issues page](http://github.com/JasonBoyett/salvo/issues).
#### Suggesting features
If you have an idea for a useful feature, please open an issue on the
[Salvo issues page](http://github.com/JasonBoyett/salvo/issues).
#### Contributing code
If you would like to contribute code, please fork the project and submit a pull request.
We will review your code and merge it if it is a good fit for the project. 
If this is your first time contributing to an open source project, check out 
[this guide.](https://docs.github.com/en/get-started/quickstart/contributing-to-projects)

## Code Standards in Salvo
### Salvo uses the Go standard library
Because Salvo is designed to be easily extended and to integrate with
other languages, we have chosen to keep dependencies to an absolute minimum.
Therefore, try to use the Go standard library whenever possible.

### Salvo uses standard Go formatting and style 
The creators of Go created a standard formatting style for Go code.
We use this [style guide](https://go.dev/doc/effective_go#formatting)
for all code in Salvo.
In order to keep the codebase consistent, 
we use the standard Go formatting tool `gofmt`.

### Salvo aims to be well tested
Because Salvo is a library and a base layer for SDKs,
it is important that it is well tested.
We do no enforce a specific test coverage percentage,
but we ask that you write tests for any code that is not trivial.

We also suggest that you use test to confirm the behavior you are trying to create as 
you are developing instead of writing your tests after you are finished.
Though we do not enforce this, we have found that it is a good practice.
We are not suggesting that you practice test driven development,
but we do suggest that you write tests as you go.

### Keep the code clean
This is not a reference to "Clean Code". It is more a directive to remove things
that make your code cluttered or harder to read.
Here are a few examples of things that make code harder to read:
- Print debugging is a great tool for debugging,
but please remove any print statements before submitting a pull request.
- Try to make sure your functions do only one thing as much as possible.
- If you find yourself repeating code in more than two places,
consider refactoring it into a function.
- Keep non documentation comments to a minimum.
- Do not use comments to explain what your code is doing.
Instead, write your code in a way that is self explanatory.
- Do not use vulgar or offensive language in your comments or code.

### Ask for help
If you are having trouble with something, please ask for help.
The Salvo core team is here to help you.
If you have any questions, feel free to open an issue or even to contact
a member of the core team directly.

#### Core team:
##### Jason Boyett
- Email: jason.boyett@gmail.com
- Github: [JasonBoyett](http://github.com/JasonBoyett)
- Discord: jabthedevboy

### Becoming a core team member
If you are interested in becoming a core team member,
please make sure you meet the following requirements:
- You have contributed to Salvo in the past.
- You have a good understanding of the Salvo codebase.
- You have a good understanding of the Salvo design philosophy.
- You have a good understanding of the Salvo roadmap.
- You have a good understanding of the Go programming language.
- You have been in contact with the core team.

If you meet these requirements and are interested in becoming a core team member,
please contact a member of the core team and let us know you are interested.
We will set up an interview with you to discuss your interest and your qualifications.

### What does it mean to be a core team member?
Being a core team member means that you have a say in the direction of the project.
You will also have the ability to merge pull requests and to create releases.

You will also be expected to help maintain the project and to be available to help
other contributors.







