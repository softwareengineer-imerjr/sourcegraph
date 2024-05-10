package gitcli

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"strconv"
	"time"

	"github.com/sourcegraph/sourcegraph/cmd/gitserver/internal/git"
	"github.com/sourcegraph/sourcegraph/internal/gitserver/gitdomain"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

func (g *gitCLIBackend) CommitLog(ctx context.Context, opt git.CommitLogOpts) ([]*git.GitCommitWithFiles, error) {
	if err := checkSpecArgSafety(opt.Range); err != nil {
		return nil, err
	}

	if opt.Range != "" && opt.AllRefs {
		return nil, errors.New("cannot specify both a Range and AllRefs")
	}

	args, err := buildCommitLogArgs(opt)
	if err != nil {
		return nil, err
	}

	r, err := g.NewCommand(ctx, WithArguments(args...))
	if err != nil {
		return nil, err
	}

	return parseCommitLogOutput(r)
}

func buildCommitLogArgs(opt git.CommitLogOpts) ([]string, error) {
	args := []string{"log", logFormatWithoutRefs}

	if opt.MaxCommits != 0 {
		args = append(args, "-n", strconv.FormatUint(uint64(opt.MaxCommits), 10))
	}
	if opt.Skip != 0 {
		args = append(args, "--skip="+strconv.FormatUint(uint64(opt.Skip), 10))
	}

	if opt.AuthorQuery != "" {
		args = append(args, "--fixed-strings", "--author="+opt.AuthorQuery)
	}

	if !opt.After.IsZero() {
		args = append(args, "--after="+opt.After.Format(time.RFC3339))
	}
	if !opt.Before.IsZero() {
		args = append(args, "--before="+opt.Before.Format(time.RFC3339))
	}
	switch opt.Order {
	case git.CommitLogOrderCommitDate:
		args = append(args, "--date-order")
	case git.CommitLogOrderTopoDate:
		args = append(args, "--topo-order")
	case git.CommitLogOrderDefault:
		// nothing to do
	default:
		return nil, errors.Newf("invalid ordering %d", opt.Order)
	}

	if opt.MessageQuery != "" {
		args = append(args, "--fixed-strings", "--regexp-ignore-case", "--grep="+opt.MessageQuery)
	}

	if opt.FollowOnlyFirstParent {
		args = append(args, "--first-parent")
	}

	if opt.Range != "" {
		args = append(args, opt.Range)
	}
	if opt.AllRefs {
		args = append(args, "--all")
	}

	if opt.IncludeModifiedFiles {
		args = append(args, "--name-only")
	}
	if opt.FollowPathRenames {
		args = append(args, "--follow")
	}
	if opt.Path != "" {
		args = append(args, "--", opt.Path)
	}

	return args, nil
}

func parseCommitLogOutput2(r io.Reader) ([]*git.GitCommitWithFiles, error) {
	commitScanner := bufio.NewScanner(r)
	// We use an increased buffer size since sub-repo permissions
	// can result in very lengthy output.
	commitScanner.Buffer(make([]byte, 0, 65536), 4294967296)
	commitScanner.Split(commitSplitFunc)

	var commits []*git.GitCommitWithFiles
	for commitScanner.Scan() {
		rawCommit := commitScanner.Bytes()
		parts := bytes.Split(rawCommit, []byte{'\x00'})
		if len(parts) != partsPerCommit {
			return nil, errors.Newf("internal error: expected %d parts, got %d", partsPerCommit, len(parts))
		}

		commit, err := parseCommitFromLog(parts)
		if err != nil {
			return nil, err
		}
		commits = append(commits, commit)
	}

	if err := commitScanner.Err(); err != nil {
		// If exit code is 128 and `fatal: bad object` is part of stderr, most likely we
		// are referencing a commit that does not exist.
		// We want to return a gitdomain.RevisionNotFoundError in that case.
		var e *CommandFailedError
		if errors.As(err, &e) && e.ExitStatus == 128 {
			if bytes.Contains(e.Stderr, []byte("not a tree object")) || bytes.Contains(e.Stderr, []byte("fatal: bad object")) {
				return nil, &gitdomain.RevisionNotFoundError{Repo: g.repoName, Spec: string(opt.Range)}
			}
		}

		return nil, err
	}

	return commits, nil
}

func commitSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		// Request more data
		return 0, nil, nil
	}

	// Safety check: ensure we are always starting with a record separator
	if data[0] != '\x1e' {
		return 0, nil, errors.New("internal error: data should always start with an ASCII record separator")
	}

	loc := bytes.IndexByte(data[1:], '\x1e')
	if loc < 0 {
		// We can't find the start of the next record
		if atEOF {
			// If we're at the end of the stream, just return the rest as the last record
			return len(data), data[1:], bufio.ErrFinalToken
		} else {
			// If we're not at the end of the stream, request more data
			return 0, nil, nil
		}
	}
	nextStart := loc + 1 // correct for searching at an offset

	return nextStart, data[1:nextStart], nil
}
