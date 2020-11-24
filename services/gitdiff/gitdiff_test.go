	"strconv"
	"gopkg.in/ini.v1"
	setting.Cfg = ini.Empty()
	assertEqual(t, "foo <span class=\"added-code\">bar</span> biz", diffToHTML("", []dmp.Diff{
	assertEqual(t, "foo <span class=\"removed-code\">bar</span> biz", diffToHTML("", []dmp.Diff{

	assertEqual(t, "<span class=\"k\">if</span> <span class=\"p\">!</span><span class=\"nx\">nohl</span> <span class=\"o\">&amp;&amp;</span> <span class=\"added-code\"><span class=\"p\">(</span></span><span class=\"nx\">lexer</span> <span class=\"o\">!=</span> <span class=\"kc\">nil</span><span class=\"added-code\"> <span class=\"o\">||</span> <span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nx\">GuessLanguage</span><span class=\"p\">)</span></span> <span class=\"p\">{</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"k\">if</span> <span class=\"p\">!</span><span class=\"nx\">nohl</span> <span class=\"o\">&amp;&amp;</span> <span class=\""},
		{Type: dmp.DiffInsert, Text: "p\">(</span><span class=\""},
		{Type: dmp.DiffEqual, Text: "nx\">lexer</span> <span class=\"o\">!=</span> <span class=\"kc\">nil"},
		{Type: dmp.DiffInsert, Text: "</span> <span class=\"o\">||</span> <span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nx\">GuessLanguage</span><span class=\"p\">)"},
		{Type: dmp.DiffEqual, Text: "</span> <span class=\"p\">{</span>"},
	}, DiffLineAdd))

	assertEqual(t, "<span class=\"nx\">tagURL</span> <span class=\"o\">:=</span> <span class=\"removed-code\"><span class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nf\">Sprintf</span><span class=\"p\">(</span><span class=\"s\">&#34;## [%s](%s/%s/%s/%s?q=&amp;type=all&amp;state=closed&amp;milestone=%d) - %s&#34;</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Milestone\"</span></span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">BaseURL</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Owner</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Repo</span><span class=\"p\">,</span> <span class=\"removed-code\"><span class=\"nx\">from</span><span class=\"p\">,</span> <span class=\"nx\">milestoneID</span><span class=\"p\">,</span> <span class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nf\">Now</span><span class=\"p\">(</span><span class=\"p\">)</span><span class=\"p\">.</span><span class=\"nf\">Format</span><span class=\"p\">(</span><span class=\"s\">&#34;2006-01-02&#34;</span><span class=\"p\">)</span></span><span class=\"p\">)</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"nx\">tagURL</span> <span class=\"o\">:=</span> <span class=\"n"},
		{Type: dmp.DiffDelete, Text: "x\">fmt</span><span class=\"p\">.</span><span class=\"nf\">Sprintf</span><span class=\"p\">(</span><span class=\"s\">&#34;## [%s](%s/%s/%s/%s?q=&amp;type=all&amp;state=closed&amp;milestone=%d) - %s&#34;</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Milestone\""},
		{Type: dmp.DiffInsert, Text: "f\">getGiteaTagURL</span><span class=\"p\">(</span><span class=\"nx\">client"},
		{Type: dmp.DiffEqual, Text: "</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">BaseURL</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Owner</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Repo</span><span class=\"p\">,</span> <span class=\"nx\">"},
		{Type: dmp.DiffDelete, Text: "from</span><span class=\"p\">,</span> <span class=\"nx\">milestoneID</span><span class=\"p\">,</span> <span class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nf\">Now</span><span class=\"p\">(</span><span class=\"p\">)</span><span class=\"p\">.</span><span class=\"nf\">Format</span><span class=\"p\">(</span><span class=\"s\">&#34;2006-01-02&#34;</span><span class=\"p\">)"},
		{Type: dmp.DiffInsert, Text: "ge</span><span class=\"p\">.</span><span class=\"nx\">Milestone</span><span class=\"p\">,</span> <span class=\"nx\">from</span><span class=\"p\">,</span> <span class=\"nx\">milestoneID"},
		{Type: dmp.DiffEqual, Text: "</span><span class=\"p\">)</span>"},
	}, DiffLineDel))

	assertEqual(t, "<span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nf\">WrapperRenderer</span><span class=\"p\">(</span><span class=\"nx\">w</span><span class=\"p\">,</span> <span class=\"removed-code\"><span class=\"nx\">language</span></span><span class=\"removed-code\"><span class=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs</span></span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nf\">WrapperRenderer</span><span class=\"p\">(</span><span class=\"nx\">w</span><span class=\"p\">,</span> <span class=\"nx\">"},
		{Type: dmp.DiffDelete, Text: "language</span><span "},
		{Type: dmp.DiffEqual, Text: "c"},
		{Type: dmp.DiffDelete, Text: "lass=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs"},
		{Type: dmp.DiffEqual, Text: "</span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>"},
	}, DiffLineDel))

	assertEqual(t, "<span class=\"added-code\">language</span></span><span class=\"added-code\"><span class=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs</span></span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffInsert, Text: "language</span><span "},
		{Type: dmp.DiffEqual, Text: "c"},
		{Type: dmp.DiffInsert, Text: "lass=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs"},
		{Type: dmp.DiffEqual, Text: "</span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>"},
	}, DiffLineAdd))

	assertEqual(t, "<span class=\"k\">print</span><span class=\"added-code\"></span><span class=\"added-code\"><span class=\"p\">(</span></span><span class=\"sa\"></span><span class=\"s2\">&#34;</span><span class=\"s2\">// </span><span class=\"s2\">&#34;</span><span class=\"p\">,</span> <span class=\"n\">sys</span><span class=\"o\">.</span><span class=\"n\">argv</span><span class=\"added-code\"><span class=\"p\">)</span></span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"k\">print</span>"},
		{Type: dmp.DiffInsert, Text: "<span"},
		{Type: dmp.DiffEqual, Text: " "},
		{Type: dmp.DiffInsert, Text: "class=\"p\">(</span>"},
		{Type: dmp.DiffEqual, Text: "<span class=\"sa\"></span><span class=\"s2\">&#34;</span><span class=\"s2\">// </span><span class=\"s2\">&#34;</span><span class=\"p\">,</span> <span class=\"n\">sys</span><span class=\"o\">.</span><span class=\"n\">argv</span>"},
		{Type: dmp.DiffInsert, Text: "<span class=\"p\">)</span>"},
	}, DiffLineAdd))

	assertEqual(t, "sh <span class=\"added-code\">&#39;useradd -u $(stat -c &#34;%u&#34; .gitignore) jenkins</span>&#39;", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "sh &#3"},
		{Type: dmp.DiffDelete, Text: "4;useradd -u 111 jenkins&#34"},
		{Type: dmp.DiffInsert, Text: "9;useradd -u $(stat -c &#34;%u&#34; .gitignore) jenkins&#39"},
		{Type: dmp.DiffEqual, Text: ";"},
	}, DiffLineAdd))

	assertEqual(t, "<span class=\"x\">							&lt;h<span class=\"added-code\">4 class=</span><span class=\"added-code\">&#34;release-list-title df ac&#34;</span>&gt;</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"x\">							&lt;h"},
		{Type: dmp.DiffInsert, Text: "4 class=&#"},
		{Type: dmp.DiffEqual, Text: "3"},
		{Type: dmp.DiffInsert, Text: "4;release-list-title df ac&#34;"},
		{Type: dmp.DiffEqual, Text: "&gt;</span>"},
	}, DiffLineAdd))
--- "\\a/a b/file b/a a/file"	` + `
+++ "\\b/a b/file b/a a/file"	` + `
 ` + `
--- "\\a/file with blanks" ` + `
			if got.NumFiles != 1 {
	// Test max lines
	diffBuilder := &strings.Builder{}

	var diff = `diff --git a/newfile2 b/newfile2
new file mode 100644
index 0000000..6bb8f39
--- /dev/null
+++ b/newfile2
@@ -0,0 +1,35 @@
`
	diffBuilder.WriteString(diff)

	for i := 0; i < 35; i++ {
		diffBuilder.WriteString("+line" + strconv.Itoa(i) + "\n")
	}
	diff = diffBuilder.String()
	result, err := ParsePatch(20, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if result.Files[0].IsIncomplete {
		t.Errorf("Files should not be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, 5, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}

	// Test max characters
	diff = `diff --git a/newfile2 b/newfile2
new file mode 100644
index 0000000..6bb8f39
--- /dev/null
+++ b/newfile2
@@ -0,0 +1,35 @@
`
	diffBuilder.Reset()
	diffBuilder.WriteString(diff)

	for i := 0; i < 33; i++ {
		diffBuilder.WriteString("+line" + strconv.Itoa(i) + "\n")
	}
	diffBuilder.WriteString("+line33")
	for i := 0; i < 512; i++ {
		diffBuilder.WriteString("0123456789ABCDEF")
	}
	diffBuilder.WriteByte('\n')
	diffBuilder.WriteString("+line" + strconv.Itoa(34) + "\n")
	diffBuilder.WriteString("+line" + strconv.Itoa(35) + "\n")
	diff = diffBuilder.String()

	result, err = ParsePatch(20, 4096, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, 4096, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}

	diff = `diff --git "a/README.md" "b/README.md"
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))