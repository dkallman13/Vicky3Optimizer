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
func PopLookup() []types.Pop {
	var pops []types.Pop
	DB.Find(&pops)
	return pops
}
func StateLookup() []types.State {
	var states []types.State
	DB.Table("States").Select("*")
	return states
}
func TokenizeSave(saveFileName string) string {
	// Fast existence check
	found := false
	for _, file := range SaveFiles {
		if file == saveFileName {
			found = true
			break
		}
	}
	if !found {
		return "save file" + saveFileName + " not found"
	}

	DecodedSaveFile := new(types.SaveFile)
	var rawFileTextBuilder strings.Builder
	var saveFileBuilder strings.Builder
	saveFileBuilder.WriteString(SaveFileLoc)
	saveFileBuilder.WriteString(saveFileName)
	file, err := os.Open(saveFileBuilder.String())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the entire file at once for faster processing
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	parser := tokenizer.New()
	parser.
		AllowKeywordSymbols(tokenizer.Underscore, tokenizer.Numbers).
		DefineTokens(TokenCurlyOpen, []string{"{"}).
		DefineTokens(TokenCurlyClose, []string{"}"}).
		DefineTokens(TokenEquals, []string{"="}).
		DefineStringToken(TokenDoubleQuoted, `"`, `"`).
		SetEscapeSymbol(tokenizer.BackSlash).AddSpecialStrings(tokenizer.DefaultSpecialString)

	bracketCount := 0
	isInMetadata = false
	var isInPops = false
	var isInQualifications = false
	var popId int
	var jobType string
	var workforceSize int
	var dependantsSize int
	var workplaceId int
	var qualifications []types.Qualifications
	var jobSatisfaction float64
	var soL int
	var wealth int
	var newPop types.Pop

	stream := parser.ParseBytes(fileBytes)
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
					isInPops = false
					isInQualifications = false
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
					switch isInPops {
					case true:
						switch bracketCount {
						case 2:
							isInQualifications = false
							switch stream.CurrentToken().Is(tokenizer.TokenInteger) {
							case true:
								if popId != 0 {
									newPop = types.NewPop(popId, jobType, workforceSize, dependantsSize, workplaceId, qualifications, jobSatisfaction, soL, wealth)
									var existingPop types.Pop
									DB.Table("Pops").Select("Id", "WorkforceSize", "DependantsSize", "WorkplaceId", "qualifications", "jobSatisfaction", "soL", "wealth").Where(&types.Pop{Id: popId}).Find(&existingPop)
									if newPop.Id != existingPop.Id {
    									if err := DB.Save(&newPop).Error; err != nil {
											log.Printf("Failed to save pop: %v", err)
										}
										newPop = types.Pop{}}
								}
								popId, err = strconv.Atoi(stream.CurrentToken().ValueString())
								if err != nil {
									log.Fatal(err)
								}
							}
						case 3:
							switch stream.CurrentToken().Is(tokenizer.TokenKeyword) {
							case true:
								switch stream.CurrentToken().ValueString() {
								case "type":
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(tokenizer.TokenKeyword) {
										case true:
											jobType = stream.CurrentToken().ValueString()
										}
									}
								case "workforce":
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(tokenizer.TokenInteger) {
										case true:
											workforceSize, err = strconv.Atoi(stream.CurrentToken().ValueString())
											if err != nil {
												log.Fatal(err)
											}
										}
									}
								case "dependants":
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(tokenizer.TokenInteger) {
										case true:
											dependantsSize, err = strconv.Atoi(stream.CurrentToken().ValueString())
											if err != nil {
												log.Fatal(err)
											}
										}
									}
								case "workplace":
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(tokenizer.TokenInteger) {
										case true:
											workplaceId, err = strconv.Atoi(stream.CurrentToken().ValueString())
											if err != nil {
												log.Fatal(err)
											}
										}
									}
								case "qualifications":
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(TokenCurlyOpen) {
										case true:
											bracketCount++
											isInQualifications = true
											qualifications = make([]types.Qualifications, 0)
										}
									}
								case "job_satisfaction":
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(tokenizer.TokenFloat) {
										case true:
											jobSatisfaction, err = strconv.ParseFloat(stream.CurrentToken().ValueString(), 64)
											if err != nil {
												log.Fatal(err)
											}
										}
									}
								case "previous_quality_of_life":
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(tokenizer.TokenInteger) {
										case true:
											soL, err = strconv.Atoi(stream.CurrentToken().ValueString())
											if err != nil {
												log.Fatal(err)
											}
										}
									}
								case "wealth":
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(tokenizer.TokenInteger) {
										case true:
											wealth, err = strconv.Atoi(stream.CurrentToken().ValueString())
											if err != nil {
												log.Fatal(err)
											}
										}
									}
								}
							}
						case 4:
							switch isInQualifications {
							case true:
								switch stream.CurrentToken().IsInteger() {
								case true:
									switch stream.CurrentToken().ValueInt64() {
									case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
										qualjob := stream.CurrentToken().ValueInt64()
										if err != nil {
											log.Fatal(err)
										}
										stream.GoNext()
										switch stream.CurrentToken().Is(TokenEquals) {
										case true:
											stream.GoNext()
											switch stream.CurrentToken().Is(tokenizer.TokenFloat) {
											case true:
												qualpop, err := strconv.ParseFloat(stream.CurrentToken().ValueString(), 64)
												if err != nil {
													log.Fatal(err)
												}
												qualifications = append(qualifications, types.NewQualifications(popId, qualjob, qualpop))
											}
										}
									}
								}
							}
						}
					}
					switch stream.CurrentToken().ValueString() {
					case "meta_data":
						isInMetadata = true
					case "pops":
						stream.GoNext()
						switch stream.CurrentToken().Is(TokenEquals) {
						case true:
							stream.GoNext()
							switch stream.CurrentToken().Is(TokenCurlyOpen) {
							case true:
								bracketCount++
								stream.GoNext()
								switch stream.CurrentToken().Is(tokenizer.TokenKeyword) {
								case true:
									switch stream.CurrentToken().ValueString() {
									case "database":
										isInPops = true
									}
								}
							}
						}
					}
				}
			}
		}
		stream.GoNext()
	}
	stream.Close()
	return rawFileTextBuilder.String()
}
