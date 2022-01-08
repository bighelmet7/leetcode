# Leetcode WIKI

Wiki for all my leetcode solutions

# TODO
- [ ] Add golden files for larger test cases
- [x] Add code coverage
- [x] Get docs for each exercice

# Requirements
- Go version 1.17.5

# Doc
```bash
# Use the PKG_NAME variable to get the docs for one specific package.
make doc PKG_NAME=add_two_numbers
make doc PKG_NAME=longest_substring
```

# Tests
```bash
# all tests
make tests
```
```bash
# one specific test
make test TEST_REGEXP=TestMyAtoi
```

# Code coverage
```bash
make coverage
```
```bash
# Cleaner
make clean
```
