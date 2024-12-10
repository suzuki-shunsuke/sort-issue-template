# sort-issue-template

[MIT LICENSE](LICENSE) | [Install](INSTALL.md)

`sort-issue-template` is a CLI to change the order of GitHub Issue Templates using a text editor by renaming template files.

e.g. https://github.com/aquaproj/aqua/issues/new/choose

<img width="1334" alt="image" src="https://github.com/user-attachments/assets/15f2eb61-2a6f-4b3e-83f9-7d283777a401">

Please run this tool in the repository root directory and change the order of templates as you like!

## Motivation

Sometimes you would like to change the order of GitHub Issue Templates.
[You can change the order by renaming issue templates](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/configuring-issue-templates-for-your-repository#changing-the-order-of-templates).
But it's a bit bothersome to rename files.
Using `sort-issue-template`, you can rename them using a text editor.

## How to use

Run `sort-issue-template` in the repository root directory.

```sh
sort-issue-template
```

`sort-issue-template` creates a temporary file and opens it using a text editor.
Please change the order of lines in the file as you like and save and close the file.
Then `sort-issue-template` renames Issue templates to change the order.
Please commit the change using Git and push the change to the default branch.

### Cancel the change

After a text editor is launched, you may want to cancel the change.
In that case, please delete all lines in the temporary file and save and close it.
Then issue templates aren't changed.

### Change the editor

By default, the editor is `vi`.
You can change the editor via the command line option `-editor` or the environment variable `EDITOR`.

```sh
sort-issue-template -editor nvim
```

```sh
export EDITOR=nvim
```

### Change the separator

Change the separator from `-` (default) to `_`:

```sh
sort-issue-template -separator _
```

Separator must match with the regular expression `[-_.]*`.
sort-issue-template detects the separator in current template file names automatically.

## Example

```console
$ ls .github/ISSUE_TEMPLATE | sort
bug-report.yml
config.yml
feature-request.yml
general.yml
question.yml
support-request.yml
```

Run `sort-issue-template`:

```sh
sort-issue-template
```

Then an editor is launched.

```
bug-report.yml
feature-request.yml
general.yml
question.yml
support-request.yml
```

Please change the order:

```
bug-report.yml
feature-request.yml
support-request.yml
question.yml
general.yml
```

Close the editor. Then template files are renamed.

```console
$ sort-issue-template
01-bug-report.yml
02-feature-request.yml
03-support-request.yml
04-question.yml
05-general.yml
```

Let's add a template between `02-feature-request.yml` and `03-support-request.yml`.

```sh
touch .github/ISSUE_TEMPLATE/foo.md
sort-issue-template
```

Please change the order:

```
01-bug-report.yml
02-feature-request.yml
foo.md
03-support-request.yml
04-question.yml
05-general.yml
```

Close the editor. Then template files are renamed again.

```console
$ sort-issue-template
01-bug-report.yml
02-feature-request.yml
03-foo.md
04-support-request.yml
05-question.yml
06-general.yml
```

## LICENSE

[MIT](LICENSE)
