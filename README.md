# EnvGit
[![Release](https://img.shields.io/github/release/unravela/envgit.svg?style=flat-square)](https://github.com/goreleaser/goreleaser/releases/latest)
[![Software License](https://img.shields.io/github/license/unravela/envgit?style=flat-square)](/LICENSE.md)
[![Build](https://img.shields.io/github/workflow/status/unravela/envgit/build/main?style=flat-square)](/actions?query=workflow%3Abuild)

EnvGit runs another command with env. variables loaded directly from file in git repository.

    envgit --url https://github.com/my/repo --file devel.env ./run_my_app.sh

## Motivation

I fall into situations, where I wanted to have centralized and versioned configuration but I 
wanted to avoid installation of systems like Consul. I needed something very simple and lightweight. 

I wrote this tool as response for my need - having easy to set up, easy to manage centralized and 
versioned configuration.


## Installation

### Install from source code

    go get https://github.com/unravela/envgit/cmd/envgit
    envgit --version

### Install on Linux - Snap

    snap install envgit

### Install on Linux - DEB package 
   
    wget https://github.com/unravela/envgit/releases/download/v0.1.0/envgit_0.1.0_linux_64-bit.deb
    dpkg -i ./envgit_0.1.0_linux_64-bit.deb

### Install on Linux - RPM package

    wget https://github.com/unravela/envgit/releases/download/v0.1.0/envgit_0.1.0_linux_64-bit.rpm
    sudo rpm –i envgit_0.1.0_linux_64-bit.rpm

### Install on MacOS (Homebrew)

If you have Homebrew present in your environment, run command:

    brew install unravela/tap/envgit

### Install on Windows (Scoop)

First, ensure the [scoop](https://scoop.sh/) is present in your environment. If not install it.

Run commands:

    scoop bucket add unravela https://github.com/unravela/scoop-bucket
    scoop install envgit
    envgit --version
    

## How to use

The best way how to explore how you can use envgit is by help:

    envgit --help

Let's imagine we have `customapp` which is using Postgres and this can be configured via env. variables.
The configuration for QA environment placed in `customapp/qa.env` file, in our private repository `http://github.com/company/config-repo`. 
We will run the application as follows:

    envgit --url https://github.com/company/config-repo --username usr --password token --file customapp/qa.env ./customapp
    
The command will download the `customapp/qa.env` file from `main` branch, set env. variables and execute the `./customapp` 
with these env. variables. 

     