# 🦓 arcadia-lib: ___Go template for library modules___

[![A B](https://img.shields.io/badge/branching-commonflow-informational?style=flat)](https://commonflow.org)
[![A B](https://img.shields.io/badge/merge-rebase-informational?style=flat)](https://git-scm.com/book/en/v2/Git-Branching-Rebasing)
[![Go Reference](https://pkg.go.dev/badge/github.com/snivilised/arcadia-lib.svg)](https://pkg.go.dev/github.com/snivilised/arcadia-lib)
[![Go report](https://goreportcard.com/badge/github.com/snivilised/arcadia-lib)](https://goreportcard.com/report/github.com/snivilised/arcadia-lib)
[![Coverage Status](https://coveralls.io/repos/github/snivilised/arcadia-lib/badge.svg?branch=master)](https://coveralls.io/github/snivilised/arcadia-lib?branch=master&kill_cache=1)

<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->

<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->

<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<p align="left">
  <a href="https://go.dev"><img src="resources/images/go-logo-light-blue.png" width="50" /></a>
</p>

## 🔰 Introduction

This project is a template to aid in the startup of Go library modules.

## 🔨 Usage

## 🎀 Features

<p align="left">
  <a href="https://onsi.github.io/ginkgo/"><img src="https://onsi.github.io/ginkgo/images/ginkgo.png" width="100" /></a>
  <a href="https://onsi.github.io/gomega/"><img src="https://onsi.github.io/gomega/images/gomega.png" width="100" /></a>
</p>

+ unit testing with [Ginkgo](https://onsi.github.io/ginkgo/)/[Gomega](https://onsi.github.io/gomega/)
+ i18n with [go-i18n](https://github.com/nicksnyder/go-i18n)

## 🧰 Developer Info

### 📝 Checklist of required changes

The following is list of actions that must be performed before using this template. Most of the changes concern changing the name `Arcadia` to the name of the new application.

As the template is instantiated from github, the new name will automatically replace the top level directory name, that being ___arcadia___.

+ `github actions workflow`: If the client application needs to use github actions for continuous integration, then the name of the [workflow](.github/workflows/ci-workflow.yml) needs to be changed. If not, then the workflow file should be deleted
+ `code coverage publication`: by default, coverage is published to [coveralls.io](https://coveralls.io), so coveralls needs to be setup to prevent the github actions workflow from failing. Alternatively, another service can be used. The user needs to adjust the __Send coverage__ job in the workflow accordingly. 
+ `replace README content`
+ `rename lib folder` to something domain specific
+ `create .env file`: Add any appropriate secrets to a newly created .env in the root directory
+ `update message id`: This package supports i18n and as part of that defines messages that need to be translated. The user needs to update the message ids of defined messages in `messages.go`, which by default contain ___arcadia-lib___ as part of the id.
