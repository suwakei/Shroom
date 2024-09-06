package repl

import (
	"Shroom/lexer"
	"Shroom/token"
	"Shroom/parser"
	"bufio"
	"fmt"
	"io"
	"os"
)

// 最初に現れる文字
const PROMPT = ">>"

// 入力と出力
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		// exitコマンドでreplの終了
		if line == "exit" {
			fmt.Print("bye bye!")
			os.Exit(0)
		}
		lex := lexer.New(line)
		parser := parser.New(lex)

		program := parser.ParseProgram()
		if len(parser.Errors()) != 0 {
			printParserErrors(out, parser.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}


func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	} 
}
