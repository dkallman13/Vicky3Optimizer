package model

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/bzick/tokenizer"
	"github.com/dkallman13/Vicky3Optimizer/types"
)

const (
	TokenCurlyOpen = iota + 1
	TokenCurlyClose
	TokenEquals
	TokenDoubleQuoted
)

var DecodedSaveFile types.SaveFile
var SaveFiles []string
var SaveFileLoc string
var saveFileLocBuilder strings.Builder
var isInMetadata bool

func SaveFileLocSetter() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	saveFileLocBuilder.WriteString(homedir)
	saveFileLocBuilder.WriteString("\\Documents\\Paradox Interactive\\Victoria 3\\save games\\")
	SaveFileLoc = saveFileLocBuilder.String()
}

func SaveFileLister() {
	files, err := os.ReadDir(SaveFileLoc)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		SaveFiles = append(SaveFiles, file.Name())
	}
}

func TokenizeSave(saveFileName string) string {
	for _, file := range SaveFiles {
		switch file == saveFileName {
		case true:
			DecodedSaveFile := new(types.SaveFile)
			var saveFileBuilder strings.Builder
			var rawFileTextBuilder strings.Builder
			saveFileBuilder.WriteString(SaveFileLoc)
			saveFileBuilder.WriteString(saveFileName)
			file, err := os.Open(saveFileBuilder.String())
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
			isInMetadata = false

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
							switch bracketCount {
							case 0:
								isInMetadata = false
							}
						case false:
							switch stream.CurrentToken().Is(tokenizer.TokenKeyword) {
							case true:
								switch isInMetadata {
								case true:
									switch stream.CurrentToken().ValueString() == "name" {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(TokenEquals) {
										case true:
											stream.GoNext()
											DecodedSaveFile.PlayerCountryName = stream.CurrentToken().ValueString()
											rawFileTextBuilder.WriteString(stream.CurrentToken().ValueString())
										}
									}
								}
								switch stream.CurrentToken().ValueString() {
								case "meta_data":
									isInMetadata = true
								}
							}
						}
					}
					stream.GoNext()
				}
				stream.Close()
			}
			return rawFileTextBuilder.String()
		}
	}
	return "save file" + saveFileName + " not found"
}
