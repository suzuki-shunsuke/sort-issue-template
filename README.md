# sort-issue-template

[![DeepWiki](https://img.shields.io/badge/DeepWiki-suzuki--shunsuke%2Fsort--issue--template-blue.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAyCAYAAAAnWDnqAAAAAXNSR0IArs4c6QAAA05JREFUaEPtmUtyEzEQhtWTQyQLHNak2AB7ZnyXZMEjXMGeK/AIi+QuHrMnbChYY7MIh8g01fJoopFb0uhhEqqcbWTp06/uv1saEDv4O3n3dV60RfP947Mm9/SQc0ICFQgzfc4CYZoTPAswgSJCCUJUnAAoRHOAUOcATwbmVLWdGoH//PB8mnKqScAhsD0kYP3j/Yt5LPQe2KvcXmGvRHcDnpxfL2zOYJ1mFwrryWTz0advv1Ut4CJgf5uhDuDj5eUcAUoahrdY/56ebRWeraTjMt/00Sh3UDtjgHtQNHwcRGOC98BJEAEymycmYcWwOprTgcB6VZ5JK5TAJ+fXGLBm3FDAmn6oPPjR4rKCAoJCal2eAiQp2x0vxTPB3ALO2CRkwmDy5WohzBDwSEFKRwPbknEggCPB/imwrycgxX2NzoMCHhPkDwqYMr9tRcP5qNrMZHkVnOjRMWwLCcr8ohBVb1OMjxLwGCvjTikrsBOiA6fNyCrm8V1rP93iVPpwaE+gO0SsWmPiXB+jikdf6SizrT5qKasx5j8ABbHpFTx+vFXp9EnYQmLx02h1QTTrl6eDqxLnGjporxl3NL3agEvXdT0WmEost648sQOYAeJS9Q7bfUVoMGnjo4AZdUMQku50McDcMWcBPvr0SzbTAFDfvJqwLzgxwATnCgnp4wDl6Aa+Ax283gghmj+vj7feE2KBBRMW3FzOpLOADl0Isb5587h/U4gGvkt5v60Z1VLG8BhYjbzRwyQZemwAd6cCR5/XFWLYZRIMpX39AR0tjaGGiGzLVyhse5C9RKC6ai42ppWPKiBagOvaYk8lO7DajerabOZP46Lby5wKjw1HCRx7p9sVMOWGzb/vA1hwiWc6jm3MvQDTogQkiqIhJV0nBQBTU+3okKCFDy9WwferkHjtxib7t3xIUQtHxnIwtx4mpg26/HfwVNVDb4oI9RHmx5WGelRVlrtiw43zboCLaxv46AZeB3IlTkwouebTr1y2NjSpHz68WNFjHvupy3q8TFn3Hos2IAk4Ju5dCo8B3wP7VPr/FGaKiG+T+v+TQqIrOqMTL1VdWV1DdmcbO8KXBz6esmYWYKPwDL5b5FA1a0hwapHiom0r/cKaoqr+27/XcrS5UwSMbQAAAABJRU5ErkJggg==)](https://deepwiki.com/suzuki-shunsuke/sort-issue-template)

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
