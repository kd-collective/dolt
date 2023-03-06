#!/usr/bin/env bats
load $BATS_TEST_DIRNAME/helper/common.bash

setup() {
    setup_common


    dolt sql -q "CREATE TABLE test(pk BIGINT PRIMARY KEY, v varchar(10))"
    dolt add .
    dolt commit -am "Created table"
}

teardown() {
    assert_feature_version
    teardown_common
}

@test "stash: stash is not supported for old format" {
    if [ "$DOLT_DEFAULT_BIN_FORMAT" = "__LD_1__" ]; then
        dolt sql -q "INSERT INTO test VALUES (1, 'a')"

        run dolt stash
        [ "$status" -eq 1 ]
        [[ "$output" =~ "stash is not supported for old storage format" ]] || false
    fi
}

@test "stash: stashing on clean working set" {
    skip_nbf_ld_1
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "No local changes to save" ]] || false
}

@test "stash: simple stashing and popping stash" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    dolt sql -q "SELECT * FROM test"
    run dolt sql -q "SELECT * FROM test"
    [ "$status" -eq 0 ]
    result=$output

    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Saved working directory and index state" ]] || false

    dolt sql -q "SELECT * FROM test"
    run dolt sql -q "SELECT * FROM test"
    [ "$status" -eq 0 ]
    [ "$output" = "" ]

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 1 ]
    [[ "$output" =~ "stash@{0}" ]] || false

    run dolt stash pop
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Dropped refs/stash@{0}" ]] || false

    run dolt sql -q "SELECT * FROM test"
    [ "$status" -eq 0 ]
    [ "$output" = "$result" ]
}

@test "stash: clearing stash when stash list is empty" {
    skip_nbf_ld_1
    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 0 ]

    run dolt stash clear
    [ "$status" -eq 0 ]
    [ "$output" = "" ]
}

@test "stash: clearing stash removes all entries in stash list" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    dolt stash

    dolt sql -q "INSERT INTO test VALUES (2, 'b')"
    dolt stash

    dolt sql -q "INSERT INTO test VALUES (3, 'c')"
    dolt stash

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 3 ]

    run dolt stash clear
    [ "$status" -eq 0 ]

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 0 ]
}

@test "stash: clearing stash and stashing again" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    dolt stash

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 1 ]

    run dolt stash clear
    [ "$status" -eq 0 ]

    dolt sql -q "INSERT INTO test VALUES (2, 'b')"
    dolt stash
    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 1 ]
}

@test "stash: clearing stash and popping returns error of no entries found" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    dolt stash
    dolt sql -q "INSERT INTO test VALUES (2, 'b')"
    dolt stash

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 2 ]

    run dolt stash clear
    [ "$status" -eq 0 ]

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 0 ]

    run dolt stash pop
    [ "$status" -eq 1 ]
    [[ "$output" =~ "No stash entries found." ]] || false

    dolt sql -q "INSERT INTO test VALUES (2, 'b')"
    dolt stash
    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 1 ]
}

@test "stash: popping oldest stash" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Saved working directory and index state" ]] || false
    [[ "$output" =~ "Created table" ]] || false

    dolt sql -q "INSERT INTO test VALUES (2, 'b')"
    dolt commit -am "Added row 2 b"

    dolt sql -q "INSERT INTO test VALUES (3, 'c')"
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Saved working directory and index state" ]] || false
    [[ "$output" =~ "Added row 2 b" ]] || false

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 2 ]
    [[ "${lines[0]}" =~ "stash@{0}" ]] || false
    [[ "${lines[1]}" =~ "stash@{1}" ]] || false

    # stash@{1} is older stash than stash@{0}, which is the latest
    run dolt stash pop stash@{1}
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Dropped refs/stash@{1}" ]] || false

    run dolt sql -q "SELECT * FROM test" -r csv
    [ "$status" -eq 0 ]
    [[ "$output" =~ "1,a" ]] || false
}

@test "stash: popping neither latest nor oldest stash" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Created table" ]] || false

    dolt sql -q "INSERT INTO test VALUES (2, 'b')"
    dolt commit -am "Added row 2 b"

    dolt sql -q "INSERT INTO test VALUES (3, 'c')"
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Added row 2 b" ]] || false

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 2 ]

    dolt sql -q "INSERT INTO test VALUES (4, 'd')"
    dolt commit -am "Added row 4 d"

    dolt sql -q "INSERT INTO test VALUES (5, 'e')"
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Added row 4 d" ]] || false

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 3 ]

    run dolt stash pop stash@{1}
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Dropped refs/stash@{1}" ]] || false

    run dolt sql -q "SELECT * FROM test" -r csv
    [ "$status" -eq 0 ]
    [[ "$output" =~ "3,c" ]] || false

    dolt checkout test
    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 2 ]

    run dolt stash pop stash@{1}
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Dropped refs/stash@{1}" ]] || false

    run dolt sql -q "SELECT * FROM test" -r csv
    [ "$status" -eq 0 ]
    [[ "$output" =~ "1,a" ]] || false
}

@test "stash: stashing multiple entries on different branches" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    run dolt stash
    [ "$status" -eq 0 ]

    dolt checkout -b newbranch
    dolt sql -q "INSERT INTO test VALUES (1, 'b')"
    run dolt stash
    [ "$status" -eq 0 ]

    run dolt stash list
    [ "$status" -eq 0 ]
    [[ "$output" =~ "stash@{0}: WIP on refs/heads/newbranch:" ]] || false
    [[ "$output" =~ "stash@{1}: WIP on refs/heads/main:" ]] || false
}

@test "stash: popping stash on different branch" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    run dolt stash
    [ "$status" -eq 0 ]

    run dolt stash list
    [ "$status" -eq 0 ]
    [[ "$output" =~ "stash@{0}: WIP on refs/heads/main:" ]] || false

    dolt checkout -b newbranch
    run dolt stash pop
    [ "$status" -eq 0 ]

    run dolt sql -q "SELECT * FROM test" -r csv
    [[ "$output" =~ "1,a" ]] || false
}

@test "stash: dropping stash removes an entry at given index in stash list" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Created table" ]] || false

    dolt sql -q "INSERT INTO test VALUES (2, 'b')"
    dolt commit -am "Added row 2 b"

    dolt sql -q "INSERT INTO test VALUES (3, 'c')"
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Added row 2 b" ]] || false

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 2 ]

    dolt sql -q "INSERT INTO test VALUES (4, 'd')"
    dolt commit -am "Added row 4 d"

    dolt sql -q "INSERT INTO test VALUES (5, 'e')"
    run dolt stash
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Added row 4 d" ]] || false

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 3 ]
    [[ "${lines[0]}" =~ "Added row 4 d" ]] || false
    [[ "${lines[1]}" =~ "Added row 2 b" ]] || false
    [[ "${lines[2]}" =~ "Created table" ]] || false

    run dolt stash drop stash@{1}
    [ "$status" -eq 0 ]

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 2 ]
    [[ "${lines[0]}" =~ "Added row 4 d" ]] || false
    [[ "${lines[1]}" =~ "Created table" ]] || false
    [[ ! "$output" =~ "Added row 2 b" ]] || false
}

@test "stash: popping stash on dirty working set with no conflict" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    run dolt stash
    [ "$status" -eq 0 ]

    dolt sql -q "INSERT INTO test VALUES (2, 'b')"
    run dolt sql -q "SELECT * FROM test" -r csv
    [ "$status" -eq 0 ]
    [[ ! "$output" = "1,a" ]] || false

    run dolt stash pop 0
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Dropped refs/stash@{0}" ]] || false

    run dolt sql -q "SELECT * FROM test" -r csv
    [ "$status" -eq 0 ]
    [[ "$output" =~ "1,a" ]] || false
    [[ "$output" =~ "2,b" ]] || false
}

@test "stash: popping stash on dirty working set with conflict" {
    skip_nbf_ld_1
    dolt sql -q "INSERT INTO test VALUES (1, 'a')"
    run dolt stash
    [ "$status" -eq 0 ]

    dolt sql -q "INSERT INTO test VALUES (1, 'b')"
    run dolt stash pop
    [ "$status" -eq 1 ]
    [[ "$output" =~ "error: Your local changes to the following tables would be overwritten by applying stash" ]] || false
    [[ "$output" =~ "The stash entry is kept in case you need it again." ]] || false

    run dolt sql -q "SELECT * FROM test" -r csv
    [[ "$output" =~ "1,b" ]] || false
    [[ ! "$output" =~ "1,a" ]] || false

    run dolt stash list
    [ "$status" -eq 0 ]
    [ "${#lines[@]}" -eq 1 ]
}