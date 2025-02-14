package model

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bzick/tokenizer"
	"github.com/dkallman13/Vicky3Optimizer/types"
)

func TokenizeStateFile(stateFileName string) {
	var stateFileBuilder strings.Builder
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	stateFileBuilder.WriteString(workingDir)
	stateFileBuilder.WriteString(stateFileName)
	file, err := os.Open(stateFileBuilder.String())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parser := tokenizer.New()
	parser.
		AllowKeywordSymbols(tokenizer.Underscore, tokenizer.Numbers).
		DefineTokens(TokenCurlyOpen, []string{"{"}).
		DefineTokens(TokenCurlyClose, []string{"}"}).
		DefineTokens(TokenEquals, []string{"="}).
		DefineStringToken(TokenDoubleQuoted, `"`, `"`).
		SetEscapeSymbol(tokenizer.BackSlash).AddSpecialStrings(tokenizer.DefaultSpecialString)
	const chunkSize = 4096
	buffer := make([]byte, chunkSize)
	bracketCount := 0
	var newstate types.State
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}

		stream := parser.ParseBytes(buffer[:n])
		for stream.IsValid() {
			switch stream.CurrentToken().Is(TokenCurlyOpen) {
			case true:
				bracketCount++
			case false:
				switch stream.CurrentToken().Is(TokenCurlyClose) {
				case true:
					bracketCount--
				}
			}
			switch strings.Contains(stream.CurrentToken().ValueString(), "STATE_") {
			case true:
				newstate = types.NewStateNameOnly(stream.CurrentToken().ValueString())
			case false:
				switch strings.Contains(stream.CurrentToken().ValueString(), "id") {
				case true:
					stream.GoNext()
					switch stream.CurrentToken().Is(TokenEquals) {
					case true:
						stream.GoNext()
						newstate.Id, err = strconv.Atoi(stream.CurrentToken().ValueString())
						if err != nil {
							log.Fatal(err)
						}
					}
				case false:
					switch strings.Contains(stream.CurrentToken().ValueString(), "arable_land") {
					case true:
						stream.GoNext()
						switch stream.CurrentToken().Is(TokenEquals) {
						case true:
							stream.GoNext()
							newstate.AirableLand, err = strconv.Atoi(stream.CurrentToken().ValueString())
							if err != nil {
								log.Fatal(err)
							}
						}
					}
				}
			}
			stream.GoNext()
		}
		stream.Close()
	}
}
