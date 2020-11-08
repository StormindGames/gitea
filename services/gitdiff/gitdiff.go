	"regexp"
// BlobExcerptChunkSize represent max lines of excerpt
const BlobExcerptChunkSize = 20
	} else if d.SectionInfo.RightIdx-d.SectionInfo.LastRightIdx > BlobExcerptChunkSize && d.SectionInfo.RightHunkSize > 0 {
		return content
	return "\n"
	FileName string
	Name     string
	Lines    []*DiffLine
var trailingSpanRegex = regexp.MustCompile(`<span\s*[[:alpha:]="]*?[>]?$`)
var entityRegex = regexp.MustCompile(`&[#]*?[0-9[:alpha:]]*$`)

// shouldWriteInline represents combinations where we manually write inline changes
func shouldWriteInline(diff diffmatchpatch.Diff, lineType DiffLineType) bool {
	if true &&
		diff.Type == diffmatchpatch.DiffEqual ||
		diff.Type == diffmatchpatch.DiffInsert && lineType == DiffLineAdd ||
		diff.Type == diffmatchpatch.DiffDelete && lineType == DiffLineDel {
		return true
	}
	return false
}
func diffToHTML(fileName string, diffs []diffmatchpatch.Diff, lineType DiffLineType) template.HTML {
	match := ""
	for _, diff := range diffs {
		if shouldWriteInline(diff, lineType) {
			if len(match) > 0 {
				diff.Text = match + diff.Text
				match = ""
			}
			// Chroma HTML syntax highlighting is done before diffing individual lines in order to maintain consistency.
			// Since inline changes might split in the middle of a chroma span tag or HTML entity, make we manually put it back together
			// before writing so we don't try insert added/removed code spans in the middle of one of those
			// and create broken HTML. This is done by moving incomplete HTML forward until it no longer matches our pattern of
			// a line ending with an incomplete HTML entity or partial/opening <span>.

			// EX:
			// diffs[{Type: dmp.DiffDelete, Text: "language</span><span "},
			// {Type: dmp.DiffEqual, Text: "c"},
			// {Type: dmp.DiffDelete, Text: "lass="p">}]

			// After first iteration
			// diffs[{Type: dmp.DiffDelete, Text: "language</span>"}, //write out
			// {Type: dmp.DiffEqual, Text: "<span c"},
			// {Type: dmp.DiffDelete, Text: "lass="p">,</span>}]

			// After second iteration
			// {Type: dmp.DiffEqual, Text: ""}, // write out
			// {Type: dmp.DiffDelete, Text: "<span class="p">,</span>}]

			// Final
			// {Type: dmp.DiffDelete, Text: "<span class="p">,</span>}]
			// end up writing <span class="removed-code"><span class="p">,</span></span>
			// Instead of <span class="removed-code">lass="p",</span></span>

			m := trailingSpanRegex.FindStringSubmatchIndex(diff.Text)
			if m != nil {
				match = diff.Text[m[0]:m[1]]
				diff.Text = strings.TrimSuffix(diff.Text, match)
			}
			m = entityRegex.FindStringSubmatchIndex(diff.Text)
			if m != nil {
				match = diff.Text[m[0]:m[1]]
				diff.Text = strings.TrimSuffix(diff.Text, match)
			}
			// Print an existing closing span first before opening added/remove-code span so it doesn't unintentionally close it
			if strings.HasPrefix(diff.Text, "</span>") {
				buf.WriteString("</span>")
				diff.Text = strings.TrimPrefix(diff.Text, "</span>")
			}
			// If we weren't able to fix it then this should avoid broken HTML by not inserting more spans below
			// The previous/next diff section will contain the rest of the tag that is missing here
			if strings.Count(diff.Text, "<") != strings.Count(diff.Text, ">") {
				buf.WriteString(diff.Text)
				continue
			}
		}
		case diff.Type == diffmatchpatch.DiffEqual:
			buf.WriteString(diff.Text)
		case diff.Type == diffmatchpatch.DiffInsert && lineType == DiffLineAdd:
			buf.WriteString(diff.Text)
		case diff.Type == diffmatchpatch.DiffDelete && lineType == DiffLineDel:
			buf.WriteString(diff.Text)

	case DiffLineSection:
		return template.HTML(getLineContent(diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
		return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content))
	diffRecord := diffMatchPatch.DiffMain(highlight.Code(diffSection.FileName, diff1[1:]), highlight.Code(diffSection.FileName, diff2[1:]), true)
	diffRecord = diffMatchPatch.DiffCleanupEfficiency(diffRecord)

	return diffToHTML(diffSection.FileName, diffRecord, diffLine.Type)
	IsProtected        bool
	tailSection := &DiffSection{FileName: diffFile.Name, Lines: []*DiffLine{tailDiffLine}}
	NumFiles, TotalAddition, TotalDeletion int
	Files                                  []*DiffFile
	IsIncomplete                           bool
// ParsePatch builds a Diff object from a io.Reader and some parameters.
	var curFile *DiffFile
	diff := &Diff{Files: make([]*DiffFile, 0)}

	sb := strings.Builder{}

	// OK let's set a reasonable buffer size.
	// This should be let's say at least the size of maxLineCharacters or 4096 whichever is larger.
	readerSize := maxLineCharacters
	if readerSize < 4096 {
		readerSize = 4096
	}

	input := bufio.NewReaderSize(reader, readerSize)
	line, err := input.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return diff, nil
		}
		return diff, err
	}
parsingLoop:
	for {
		// 1. A patch file always begins with `diff --git ` + `a/path b/path` (possibly quoted)
		// if it does not we have bad input!
		if !strings.HasPrefix(line, cmdDiffHead) {
			return diff, fmt.Errorf("Invalid first file line: %s", line)
		// TODO: Handle skipping first n files
		if len(diff.Files) >= maxFiles {
			diff.IsIncomplete = true
			_, err := io.Copy(ioutil.Discard, reader)
			if err != nil {
				// By the definition of io.Copy this never returns io.EOF
				return diff, fmt.Errorf("Copy: %v", err)
			}
			break parsingLoop
		}
		curFile = createDiffFile(diff, line)
		diff.Files = append(diff.Files, curFile)

		// 2. It is followed by one or more extended header lines:
		//
		//     old mode <mode>
		//     new mode <mode>
		//     deleted file mode <mode>
		//     new file mode <mode>
		//     copy from <path>
		//     copy to <path>
		//     rename from <path>
		//     rename to <path>
		//     similarity index <number>
		//     dissimilarity index <number>
		//     index <hash>..<hash> <mode>
		//
		// * <mode> 6-digit octal numbers including the file type and file permission bits.
		// * <path> does not include the a/ and b/ prefixes
		// * <number> percentage of unchanged lines for similarity, percentage of changed
		//   lines dissimilarity as integer rounded down with terminal %. 100% => equal files.
		// * The index line includes the blob object names before and after the change.
		//   The <mode> is included if the file mode does not change; otherwise, separate
		//   lines indicate the old and the new mode.
		// 3. Following this header the "standard unified" diff format header may be encountered: (but not for every case...)
		//
		//     --- a/<path>
		//     +++ b/<path>
		//
		// With multiple hunks
		//
		//     @@ <hunk descriptor> @@
		//     +added line
		//     -removed line
		//      unchanged line
		//
		// 4. Binary files get:
		//
		//     Binary files a/<path> and b/<path> differ
		//
		// but one of a/<path> and b/<path> could be /dev/null.
	curFileLoop:
			line, err = input.ReadString('\n')
				if err != io.EOF {
					return diff, err
				break parsingLoop
			switch {
			case strings.HasPrefix(line, cmdDiffHead):
				break curFileLoop
			case strings.HasPrefix(line, "old mode ") ||
				strings.HasPrefix(line, "new mode "):
				if strings.HasSuffix(line, " 160000\n") {
					curFile.IsSubmodule = true
				}
			case strings.HasPrefix(line, "copy from "):
				curFile.IsRenamed = true
				curFile.Type = DiffFileCopy
			case strings.HasPrefix(line, "copy to "):
				curFile.IsRenamed = true
				curFile.Type = DiffFileCopy
			case strings.HasPrefix(line, "new file"):
				curFile.Type = DiffFileAdd
				curFile.IsCreated = true
				if strings.HasSuffix(line, " 160000\n") {
					curFile.IsSubmodule = true
				}
			case strings.HasPrefix(line, "deleted"):
				curFile.Type = DiffFileDel
				curFile.IsDeleted = true
				if strings.HasSuffix(line, " 160000\n") {
					curFile.IsSubmodule = true
				}
			case strings.HasPrefix(line, "index"):
				if strings.HasSuffix(line, " 160000\n") {
					curFile.IsSubmodule = true
				}
			case strings.HasPrefix(line, "similarity index 100%"):
				curFile.Type = DiffFileRename
			case strings.HasPrefix(line, "Binary"):
				curFile.IsBin = true
			case strings.HasPrefix(line, "--- "):
				// Do nothing with this line
			case strings.HasPrefix(line, "+++ "):
				// Do nothing with this line
				lineBytes, isFragment, err := parseHunks(curFile, maxLines, maxLineCharacters, input)
				diff.TotalAddition += curFile.Addition
				diff.TotalDeletion += curFile.Deletion
				if err != nil {
					if err != io.EOF {
						return diff, err
					}
					break parsingLoop
				}
				sb.Reset()
				_, _ = sb.Write(lineBytes)
				for isFragment {
					lineBytes, isFragment, err = input.ReadLine()
					if err != nil {
						// Now by the definition of ReadLine this cannot be io.EOF
						return diff, fmt.Errorf("Unable to ReadLine: %v", err)
					}
					_, _ = sb.Write(lineBytes)
				}
				line = sb.String()
				sb.Reset()

				break curFileLoop

	}

	// FIXME: There are numerous issues with this:
	// - we might want to consider detecting encoding while parsing but...
	// - we're likely to fail to get the correct encoding here anyway as we won't have enough information
	// - and this doesn't really account for changes in encoding
	var buf bytes.Buffer
	for _, f := range diff.Files {
		buf.Reset()
		for _, sec := range f.Sections {
			for _, l := range sec.Lines {
				if l.Type == DiffLineSection {
					continue
				buf.WriteString(l.Content[1:])
				buf.WriteString("\n")
		}
		charsetLabel, err := charset.DetectEncoding(buf.Bytes())
		if charsetLabel != "UTF-8" && err == nil {
			encoding, _ := stdcharset.Lookup(charsetLabel)
			if encoding != nil {
				d := encoding.NewDecoder()
				for _, sec := range f.Sections {
					for _, l := range sec.Lines {
						if l.Type == DiffLineSection {
							continue
						}
						if c, _, err := transform.String(d, l.Content[1:]); err == nil {
							l.Content = l.Content[0:1] + c
						}
					}
	}
	diff.NumFiles = len(diff.Files)
	return diff, nil
}
func parseHunks(curFile *DiffFile, maxLines, maxLineCharacters int, input *bufio.Reader) (lineBytes []byte, isFragment bool, err error) {
	sb := strings.Builder{}
	var (
		curSection        *DiffSection
		curFileLinesCount int
		curFileLFSPrefix  bool
	)
	leftLine, rightLine := 1, 1
	for {
		sb.Reset()
		lineBytes, isFragment, err = input.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			err = fmt.Errorf("Unable to ReadLine: %v", err)
			return
		}
		if lineBytes[0] == 'd' {
			// End of hunks
			return
		switch lineBytes[0] {
		case '@':
			if curFileLinesCount >= maxLines {
				curFile.IsIncomplete = true
				continue
			}
			_, _ = sb.Write(lineBytes)
			for isFragment {
				// This is very odd indeed - we're in a section header and the line is too long
				// This really shouldn't happen...
				lineBytes, isFragment, err = input.ReadLine()
				if err != nil {
					// Now by the definition of ReadLine this cannot be io.EOF
					err = fmt.Errorf("Unable to ReadLine: %v", err)
					return
				}
				_, _ = sb.Write(lineBytes)
			}
			line := sb.String()

			// Create a new section to represent this hunk

			curSection.FileName = curFile.Name
		case '\\':
			if curFileLinesCount >= maxLines {
				curFile.IsIncomplete = true
				continue
			}
			// This is used only to indicate that the current file does not have a terminal newline
			if !bytes.Equal(lineBytes, []byte("\\ No newline at end of file")) {
				err = fmt.Errorf("Unexpected line in hunk: %s", string(lineBytes))
				return
			}
			// Technically this should be the end the file!
			// FIXME: we should be putting a marker at the end of the file if there is no terminal new line
			continue
		case '+':
			curFileLinesCount++
			if curFileLinesCount >= maxLines {
				curFile.IsIncomplete = true
				continue
			}
			diffLine := &DiffLine{Type: DiffLineAdd, RightIdx: rightLine}
		case '-':
			curFileLinesCount++
			if curFileLinesCount >= maxLines {
				curFile.IsIncomplete = true
				continue
			}
			diffLine := &DiffLine{Type: DiffLineDel, LeftIdx: leftLine}
		case ' ':
			curFileLinesCount++
			if curFileLinesCount >= maxLines {
				curFile.IsIncomplete = true
				continue
			}
			diffLine := &DiffLine{Type: DiffLinePlain, LeftIdx: leftLine, RightIdx: rightLine}
			leftLine++
			rightLine++
			curSection.Lines = append(curSection.Lines, diffLine)
		default:
			// This is unexpected
			err = fmt.Errorf("Unexpected line in hunk: %s", string(lineBytes))
			return
		line := string(lineBytes)
		if isFragment {
			curFile.IsIncomplete = true
			for isFragment {
				lineBytes, isFragment, err = input.ReadLine()
					// Now by the definition of ReadLine this cannot be io.EOF
					err = fmt.Errorf("Unable to ReadLine: %v", err)
					return
		}
		curSection.Lines[len(curSection.Lines)-1].Content = line
		// handle LFS
		if line[1:] == models.LFSMetaFileIdentifier {
			curFileLFSPrefix = true
		} else if curFileLFSPrefix && strings.HasPrefix(line[1:], models.LFSMetaFileOidPrefix) {
			oid := strings.TrimPrefix(line[1:], models.LFSMetaFileOidPrefix)
			if len(oid) == 64 {
				m := &models.LFSMetaObject{Oid: oid}
				count, err := models.Count(m)
				if err == nil && count > 0 {
					curFile.IsBin = true
					curFile.IsLFSFile = true
					curSection.Lines = nil
}
func createDiffFile(diff *Diff, line string) *DiffFile {
	// The a/ and b/ filenames are the same unless rename/copy is involved.
	// Especially, even for a creation or a deletion, /dev/null is not used
	// in place of the a/ or b/ filenames.
	//
	// When rename/copy is involved, file1 and file2 show the name of the
	// source file of the rename/copy and the name of the file that rename/copy
	// produces, respectively.
	//
	// Path names are quoted if necessary.
	//
	// This means that you should always be able to determine the file name even when there
	// there is potential ambiguity...
	//
	// but we can be simpler with our heuristics by just forcing git to prefix things nicely
	curFile := &DiffFile{
		Index:    len(diff.Files) + 1,
		Type:     DiffFileChange,
		Sections: make([]*DiffSection, 0, 10),
	}

	rd := strings.NewReader(line[len(cmdDiffHead):] + " ")
	curFile.Type = DiffFileChange
	curFile.OldName = readFileName(rd)
	curFile.Name = readFileName(rd)
	curFile.IsRenamed = curFile.Name != curFile.OldName
	return curFile
}

func readFileName(rd *strings.Reader) string {
	var name string
	char, _ := rd.ReadByte()
	_ = rd.UnreadByte()
	if char == '"' {
		fmt.Fscanf(rd, "%q ", &name)
		if name[0] == '\\' {
			name = name[1:]
	} else {
		fmt.Fscanf(rd, "%s ", &name)
	return name[2:]
	if (len(beforeCommitID) == 0 || beforeCommitID == git.EmptySHA) && commit.ParentCount() == 0 {
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		if len(whitespaceBehavior) != 0 {
			diffArgs = append(diffArgs, whitespaceBehavior)
		}
		// append empty tree ref
		diffArgs = append(diffArgs, "4b825dc642cb6eb9a060e54bf8d69288fbee4904")
		diffArgs = append(diffArgs, afterCommitID)
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
	shortstatArgs := []string{beforeCommitID + "..." + afterCommitID}
	if len(beforeCommitID) == 0 || beforeCommitID == git.EmptySHA {
		shortstatArgs = []string{git.EmptyTreeSHA, afterCommitID}
	}
	diff.NumFiles, diff.TotalAddition, diff.TotalDeletion, err = git.GetDiffShortStat(repoPath, shortstatArgs...)
	if err != nil && strings.Contains(err.Error(), "no merge base") {
		// git >= 2.28 now returns an error if base and head have become unrelated.
		// previously it would return the results of git diff --shortstat base head so let's try that...
		shortstatArgs = []string{beforeCommitID, afterCommitID}
		diff.NumFiles, diff.TotalAddition, diff.TotalDeletion, err = git.GetDiffShortStat(repoPath, shortstatArgs...)
	}
	if err != nil {
		return nil, err
	}
