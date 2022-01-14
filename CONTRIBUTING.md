# Contributing to go-gin-boilerplate

:+1::tada: First off, thanks for taking the time to contribute! :tada::+1:

The following is a set of guidelines for contributing to `go-gin-boilerplate`, which are hosted in the [mrinjamul/go-gin-boilerplate](#) on GitHub. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

We :heart: contributions from our community! Thank you for taking the time to review this guide and understand how we work. And please remember, all interactions in this repository should adhere to the [Code of Conduct](CODE_OF_CONDUCT.md).

| :bulb:                                                                         | :question:                                           | :bug:                                  |
| ------------------------------------------------------------------------------ | ---------------------------------------------------- | -------------------------------------- |
| **Have an idea for improvement?** Open an issue first and let's talk about it. | **Have a question on using the app?** open an issue. | **Did you find a bug?** Open an issue. |

### How you can help

There are many ways you can help us make this project better. Every repository is different and our goal is to create an app that is flexible enough to help most repositories, while still being easy for new app users to grok. If you would like to help, we use labels to organize the work that needs to be done. Look for the labels that fit your expertise:

- **help-wanted**: used when an issue or PR is up for grabs

You can also help by translating documentation and reviewing open pull requests.

## Contributing Guide

- Create pull latest code of `main` branch to your local `main` branch.

```sh
User:(any-branch)$ git checkout main
User:(main)$ git pull origin main
```

- Create a new branch from `main` branch, use proper naming convension, i.e.- `feature-xyz`, `ISSUE-ID-xyz` etc.

```sh
User:(main)$ git checkout -b feature-xyz
```

- User proper commit message, i.e- `#ISSUE-ID TYPE: your commit message`. TYPE are `chore`, `fix`, `docs`, `test`, `feature`, `refactor`.

```sh
User:(feature-xyz)$ git add -p
User:(feature-xyz)$ git commit -m "commit message"
```

- Rebase your code with letest code of `main` branch`.

```sh
User:(feature-xyz)$ git pull --all
User:(feature-xyz)$ git rebase main
```

- Push your code, and give a Pull Request.

```sh
User:(feature-xyz)$ git push origin feature-xyz
```

## Code of Conduct

This project and everyone participating in it is governed by the [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [mrinjamul@gmail.com](mailto:mrinjamul@gmail.com).

```

```
