# testscript-symlink
Example repo to show https://github.com/rogpeppe/go-internal/pull/79

```bash
$ go test . -v
=== RUN   TestMain
=== RUN   TestMain/foo
=== PAUSE TestMain/foo
=== CONT  TestMain/foo
--- FAIL: TestMain (0.00s)
    --- FAIL: TestMain/foo (0.01s)
        testscript.go:308: 
            WORK=$WORK
            PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/Users/pontus/go/bin
            HOME=/no-home
            TMPDIR=$WORK/tmp
            devnull=/dev/null
            :=:
            exe=
            
            # This works fine since it returns the symlinked path (0.003s)
            > exec pwd
            [stdout]
            $WORK
            > stdout ^$WORK
            # But any program that works with the evaluated path (like vim or pwd -P) will fail in macOS (0.003s)
            > exec pwd -P
            [stdout]
            /private$WORK
            > stdout ^$WORK
            FAIL: scripts/foo.txt:7: no match for `^$WORK` found in stdout
            
FAIL
FAIL	github.com/leitzler/testscript-symlink	0.015s
```
