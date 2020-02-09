# My Team Bot

Slack bot for some activities in my team

## Owner

@tommynurwantoro

## Onboarding and Development Guide

### Prerequisite

- Git
- Go 1.13.4 or later
- MySQL 5.7

### Installation

- Install Git  
  See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Install Go (Golang)  
  See [Golang Installation](https://golang.org/doc/install)

- Install MySQL  
  See [MySQL Installation](https://www.mysql.com/downloads/)

- Clone this repo in your local

  ```sh
  git clone git@github.com:tommynurwantoro/myteambotslack.git
  ```

- Go to myteambotslack directory, then sync the vendor file

  ```sh
  cd myteambotslack
  make mod
  ```

- Copy env.sample and if necessary, modify the env value(s)

  ```sh
  cp env.sample .env
  ```

- Install Bundler
  
  ```sh
  gem install bundler
  ```

- Prepare database

  ```sh
  bundle install
  rake db:create db:migrate
  ```

- Run Bot

  ```sh
  make run
  ```
