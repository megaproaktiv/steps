# Stepping a file in blocks

Define a block:

```
#begin
Block
#end
```

Ignore next lines

```
#ignore
lorem ipso
```


Call app:

```
step step file1 outfile.d2 maxblocks
```

Increase Block number with "right"
Decrease Block number with "Left"
Leave with "esc"


The number will vary between 1 and maxblocks

# Usage with d2

```
1 # https://taskfile.dev
2
3 version: '3'
4
5 vars:
6   MODULE: interact
7   STEPS: 8
8
9 tasks:
10   default:
11     desc: run exercise
12     cmds:
13       - d2 -l elk  --watch app.d2 &
14       - step step {{.MODULE}}.d2 app.d2 {{.STEPS}}
15
16
17   kill:
18     desc: Kill background d2s
19     cmds:
20       - killall -HUP d2
```

1) run [d2](https://d2lang.com/) in background to watch output file
2) start steps with maximum number of blocks/steps
3) run `task`

```
task
task: [default] d2 -l elk  --watch app.d2 &
task: [default] step step interact.d2 app.d2 8
/Users/gglawe/letsdemo/aws-as-api/interact.d2
1   + - with left and right, stop with esc
success: listening on http://127.0.0.1:55372
```

4) Show chaning diagram in browser
5) Kill background process `task kill`}
