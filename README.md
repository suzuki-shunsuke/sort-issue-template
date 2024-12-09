# sort-issue-template

`sort-issue-template` is a CLI to rename GitHub Issue Templates using a text editor to sort them.

e.g. https://github.com/aquaproj/aqua/issues/new/choose

<img width="1334" alt="image" src="https://github.com/user-attachments/assets/15f2eb61-2a6f-4b3e-83f9-7d283777a401">

## Motivation

Sometimes you would like to change the order of GitHub Issue Templates.
[You can change the order by renaming issue templates](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/configuring-issue-templates-for-your-repository#changing-the-order-of-templates).
But it's a bit bothersome to rename files.
Using `sort-issue-template`, you can rename them using a text editor.

## Install

```sh
go install github.com/suzuki-shunsuke/sort-issue-template/cmd/sort-issue-template@latest
```

## How to use

Run `sort-issue-template` in the repository root directory.

```sh
sort-issue-template
```

`sort-issue-template` creates a temporary file and opens it using a text editor.
Please change the order of lines in the file as you like and save and close the file.
Then `sort-issue-template` renames Issue templates to change the order.
Please commit the change using Git and push the change to the default branch.
If the temporary file isn't changed or all lines are deleted, templates aren't renamed.

By default, the editor is `vi`.
You can change the editor via the command line option `-editro` or the environment variable `EDITOR`.

```sh
sort-issue-template -editor nvim
```

```sh
export EDITOR=nvim
```

## LICENSE

[MIT](LICENSE)
