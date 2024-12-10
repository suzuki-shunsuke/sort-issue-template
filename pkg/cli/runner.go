package cli

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Runner struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	LDFlags *LDFlags
	FS      fs.FS
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

type OSFS struct{}

func (fs *OSFS) Open(name string) (fs.File, error) {
	f, err := os.Open(name)
	return f, err //nolint:wrapcheck
}

const waitDelay = 1000 * time.Hour

type Flag struct {
	Editor    string
	Separator string
	Help      bool
	Version   bool
}

const help = `sort-issue-template - Sort GitHub Issue Templates using a text editor
https://github.com/suzuki-shunsuke/sort-issue-template

Usage:
	sort-issue-template [-help] [-version] [-editor editor] [-separator separator]

Options:
	-help		Show help
	-version	Show sort-issue-template version
	-editor		Editor to sort issue templates
	-separator 	Separator in file names`

const header = `# https://github.com/suzuki-shunsuke/sort-issue-template
# sort-issue-template created this temporary file and opened it using the editor.
# Please change the order of lines to change the order of issue templates.
# Save and close the editor, then issue templates will be renamed.
# Lines starting with # are comments and are ignored.
# If you want to cancel the change, please remove all lines and save and close the file.
`

func (r *Runner) Run(ctx context.Context, _ ...string) error { //nolint:funlen,cyclop
	flg := &Flag{}
	parseFlags(flg)
	if flg.Version {
		fmt.Fprintln(r.Stdout, r.LDFlags.Version)
		return nil
	}
	allowedSeparator := regexp.MustCompile(`[-_.]+`)
	if !allowedSeparator.MatchString(flg.Separator) {
		return errors.New(`the separator must match with the regular expression [-_.]+`)
	}
	if flg.Help {
		fmt.Fprintln(r.Stdout, help)
		return nil
	}
	files, err := r.findTemplates()
	if err != nil {
		return err
	}

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "sort-issue-template-*.txt")
	if err != nil {
		return fmt.Errorf("create a temporary file: %w", err)
	}
	tempFileName := tempFile.Name()
	defer tempFile.Close()
	defer os.Remove(tempFileName)
	// Write issue template file names to the temporary file
	if err := os.WriteFile(tempFileName, []byte(header+strings.Join(files, "\n")), 0o644); err != nil { //nolint:mnd,gosec
		return fmt.Errorf("write the issue template file names to the temporary file: %w", err)
	}
	// Open the temporary file using the editor
	if err := r.openEditor(ctx, flg.Editor, tempFileName); err != nil {
		return err
	}

	// Read the temporary file
	f, err := r.FS.Open(tempFileName)
	if err != nil {
		return fmt.Errorf("open the temporary file: %w", err)
	}
	defer f.Close()

	fileNames, err := r.readTempFile(f)
	if err != nil {
		return err
	}

	size := len(fileNames)
	if size == 0 {
		return nil
	}

	// padding
	pd := len(strconv.Itoa(size))
	if pd == 1 {
		pd = 2
	}

	afterFilenames := make([]string, size)
	filenamePattern, err := regexp.Compile(`^(\d+)[-_.]+(.+)$`)
	if err != nil {
		return fmt.Errorf("compile the regular expression: %w", err)
	}
	for i, fileName := range fileNames {
		afterFilename := getNewFilename(fileName, filenamePattern, i+1, pd, flg.Separator)
		afterFilenames[i] = afterFilename
		// Rename issue templates
		if fileName == afterFilename {
			continue
		}
		if err := os.Rename(filepath.Join(".github", "ISSUE_TEMPLATE", fileName), filepath.Join(".github", "ISSUE_TEMPLATE", afterFilename)); err != nil {
			return fmt.Errorf("rename issue templates: %w", err)
		}
	}
	fmt.Fprintln(r.Stdout, strings.Join(afterFilenames, "\n"))
	return nil
}

func getNewFilename(fileName string, pattern *regexp.Regexp, index, padding int, separator string) string {
	matches := pattern.FindStringSubmatch(fileName)
	if matches == nil {
		return fmt.Sprintf("%0"+strconv.Itoa(padding)+"d%s%s", index, separator, fileName)
	}
	return fmt.Sprintf("%0"+strconv.Itoa(padding)+"d%s%s", index, separator, matches[2])
}

func (r *Runner) readTempFile(f io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(f)
	var fileNames []string
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" || strings.HasPrefix(txt, "#") {
			continue
		}
		fileNames = append(fileNames, txt)
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("read the temporary file: %w", err)
		}
	}
	return fileNames, nil
}

func parseFlags(f *Flag) {
	flag.StringVar(&f.Editor, "editor", "", "editor")
	flag.BoolVar(&f.Help, "help", false, "Show help")
	flag.BoolVar(&f.Version, "version", false, "Show version")
	flag.StringVar(&f.Separator, "separator", "-", "separator")
	flag.Parse()
	if f.Editor == "" {
		f.Editor = os.Getenv("EDITOR")
		if f.Editor == "" {
			f.Editor = "vi"
		}
	}
}

func (r *Runner) findTemplates() ([]string, error) {
	// Find issue templates
	dirEntries, err := os.ReadDir(filepath.Join(".github", "ISSUE_TEMPLATE"))
	if err != nil {
		return nil, fmt.Errorf("read the directory .github/ISSUE_TEMPLATE: %w", err)
	}
	files := make([]string, 0, len(dirEntries))
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}
		name := dirEntry.Name()
		suffix := filepath.Ext(name)
		if suffix != ".md" && suffix != ".yml" && suffix != ".yaml" {
			continue
		}
		if name == "config.yml" || name == "config.yaml" {
			continue
		}
		files = append(files, name)
	}

	sort.Strings(files)
	return files, nil
}

func (r *Runner) openEditor(ctx context.Context, editor, filepath string) error {
	cmd := exec.CommandContext(ctx, editor, filepath)
	cmd.Stdin = r.Stdin
	cmd.Stdout = r.Stdout
	cmd.Stderr = r.Stderr
	cmd.Cancel = func() error {
		return cmd.Process.Signal(os.Interrupt)
	}
	cmd.WaitDelay = waitDelay
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("open the temporary file using the editor: %w", err)
	}
	return nil
}
