#!/usr/bin/env bats

load "../support/bats-support/load"
load "../support/bats-assert/load"

setup() {
  source /tmp/.env-*
  export PATH="/home/semaphore/.rbenv/bin:$PATH"
  eval "$(rbenv init -)"
  set -u
  source ~/.toolbox/toolbox
}

#  Ruby
@test "change ruby to 2.3.8" {

  run sem-version ruby 2.3.8
  assert_success    
  run ruby --version
  assert_line --partial "ruby 2.3.8"
}

@test "change ruby to 2.5.3" {

  run sem-version ruby 2.5.3
  assert_success
  run ruby --version
  assert_line --partial "ruby 2.5.3"
}

@test "change ruby to 2.6.10" {

  run sem-version ruby 2.6.10
  assert_success    
  run ruby --version
  assert_line --partial "ruby 2.6.10"
}

@test "change ruby to 2.7.7" {

  run sem-version ruby 2.7.7
  assert_success
  run ruby --version
  assert_line --partial "ruby 2.7.7"
}

@test "change ruby to 3.0.5" {

  run sem-version ruby 3.0.5
  assert_success
  run ruby --version
  assert_line --partial "ruby 3.0.5"
}

@test "change ruby to 3.1.3" {

  run sem-version ruby 3.1.3
  assert_success
  run ruby --version
  assert_line --partial "ruby 3.1.3"
}

@test "change ruby to 3.2.1" {

  run sem-version ruby 3.2.1
  assert_success
  run ruby --version
  assert_line --partial "ruby 3.2.1"
}

@test "ruby minor versions test" {

  run sem-version ruby 2.5
  assert_success
  run ruby --version
  assert_line --partial "ruby 2.5.9"

  run sem-version ruby 2.6
  assert_success
  run ruby --version
  assert_line --partial "ruby 2.6.10"

  run sem-version ruby 2.7
  assert_success
  run ruby --version
  assert_line --partial "ruby 2.7.7"

  run sem-version ruby 3.0
  assert_success
  run ruby --version
  assert_line --partial "ruby 3.0.5"

  run sem-version ruby 3.1
  assert_success
  run ruby --version
  assert_line --partial "ruby 3.1.3"

  run sem-version ruby 3.2
  assert_success
  run ruby --version
  assert_line --partial "ruby 3.2.1"
}

@test "change ruby to 4.0.1" {

  run sem-version ruby 4.0.1
  assert_failure
}
