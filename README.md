# sort-issue-template

`sort-issue-template` is a CLI to rename GitHub Issue Templates using a text editor to sort them.

## Install

```sh
go install github.com/suzuki-shunsuke/sort-issue-template/cmd/sort-issue-template@latest
```

## How to use

Run `sort-issue-template` in the repository root directory.

```sh
sort-issue-template
```

sort-issue-template creates a temporary file and opens it using a text editor.
Please change the order of lines in the file as you like and save and close the file.
Then sort-issue-template renames Issue templates to change the order.
Please commit the change using Git and push the change to the default branch.

## LICENSE

[MIT](LICENSE)
