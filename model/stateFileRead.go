package model

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bzick/tokenizer"
	"github.com/dkallman13/Vicky3Optimizer/types"
)

func TokenizeStateFile(stateFile *os.File) {
	parser := tokenizer.New()
	parser.
		AllowKeywordSymbols(tokenizer.Underscore, tokenizer.Numbers).
		DefineTokens(TokenCurlyOpen, []string{"{"}).
		DefineTokens(TokenCurlyClose, []string{"}"}).
		DefineTokens(TokenEquals, []string{"="}).
		DefineStringToken(TokenDoubleQuoted, `"`, `"`).
		SetEscapeSymbol(tokenizer.BackSlash).AddSpecialStrings(tokenizer.DefaultSpecialString)
	const chunkSize uint = 4096
	bracketCount := 0
	var newstate types.State
	var stateName string
	var stateId int
	var err error
	stream := parser.ParseStream(stateFile, chunkSize)
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
					DB.Save(&newstate)
					newstate = types.State{}
				}
			}
		}
		switch strings.Contains(stream.CurrentToken().ValueString(), "STATE_") && !strings.Contains(stream.CurrentToken().ValueString(), "state_trait") {
		case true:
			stateName = stream.CurrentToken().ValueString()
		case false:
			switch strings.Contains(stream.CurrentToken().ValueString(), "id") {
			case true:
				stream.GoNext()
				switch stream.CurrentToken().Is(TokenEquals) {
				case true:
					stream.GoNext()
					stateId, err = strconv.Atoi(stream.CurrentToken().ValueString())
					if err != nil {
						log.Println("error here4")
						log.Fatal(err)
					}

					DB.Table("states").Select("Id", "Name").Where(&types.State{Id: stateId}).Find(&newstate)
					if newstate.Name != stateName {
						newstate.Name = stateName
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
							log.Println("error here5")
							log.Fatal(err)
						}
					}
				case false:
					switch strings.Contains(stream.CurrentToken().ValueString(), "bg_logging") {
					case true:
						stream.GoNext()
						switch stream.CurrentToken().Is(TokenEquals) {
						case true:
							stream.GoNext()
							newstate.WoodCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
							if err != nil {
								log.Println("error here6")
								log.Fatal(err)
							}
						}

					case false:
						switch strings.Contains(stream.CurrentToken().ValueString(), "bg_fishing") {
						case true:
							stream.GoNext()
							switch stream.CurrentToken().Is(TokenEquals) {
							case true:
								stream.GoNext()
								newstate.FishCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
								if err != nil {
									log.Println("error here7")
									log.Fatal(err)
								}
							}
						case false:
							switch strings.Contains(stream.CurrentToken().ValueString(), "bg_iron_mining") {
							case true:
								stream.GoNext()
								switch stream.CurrentToken().Is(TokenEquals) {
								case true:
									stream.GoNext()
									newstate.IronCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
									if err != nil {
										log.Println("error here8")
										log.Fatal(err)
									}
								}
							case false:
								switch strings.Contains(stream.CurrentToken().ValueString(), "bg_coal_mining") {
								case true:
									stream.GoNext()
									switch stream.CurrentToken().Is(TokenEquals) {
									case true:
										stream.GoNext()
										newstate.CoalCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
										if err != nil {
											log.Println("error here9")
											log.Fatal(err)
										}
									}
								case false:
									switch strings.Contains(stream.CurrentToken().ValueString(), "bg_sulfur_mining") {
									case true:
										stream.GoNext()
										switch stream.CurrentToken().Is(TokenEquals) {
										case true:
											stream.GoNext()
											newstate.SulfurCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
											if err != nil {
												log.Println("error here10")
												log.Fatal(err)
											}
										}
									case false:
										switch strings.Contains(stream.CurrentToken().ValueString(), "bg_lead_mining") {
										case true:
											stream.GoNext()
											switch stream.CurrentToken().Is(TokenEquals) {
											case true:
												stream.GoNext()
												newstate.LeadCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
												if err != nil {
													log.Println("error here11")
													log.Fatal(err)
												}
											}
										case false:
											switch strings.Contains(stream.CurrentToken().ValueString(), "bg_whaling") {
											case true:
												stream.GoNext()
												switch stream.CurrentToken().Is(TokenEquals) {
												case true:
													stream.GoNext()
													newstate.WhaleCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
													if err != nil {
														log.Println("error here12")
														log.Fatal(err)
													}
												}
											case false:
												switch strings.Contains(stream.CurrentToken().ValueString(), "bg_oil_extraction") {
												case true:
													stream.GoNext()
													switch strings.Contains(stream.CurrentToken().ValueString(), "undiscovered_amount") {
													case true:
														stream.GoNext()
														switch stream.CurrentToken().Is(TokenEquals) {
														case true:
															stream.GoNext()
															newstate.OilCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
															if err != nil {
																log.Println("error here13")
																log.Fatal(err)
															}
														}
													}
												case false:
													switch strings.Contains(stream.CurrentToken().ValueString(), "bg_rubber") {
													case true:
														stream.GoNext()
														switch strings.Contains(stream.CurrentToken().ValueString(), "undiscovered_amount") {
														case true:
															stream.GoNext()
															switch stream.CurrentToken().Is(TokenEquals) {
															case true:
																stream.GoNext()
																newstate.RubberCap, err = strconv.Atoi(stream.CurrentToken().ValueString())
																if err != nil {
																	log.Println("error here14")
																	log.Fatal(err)
																}
															}
														}
													case false:
														switch strings.Contains(stream.CurrentToken().ValueString(), "bg_cotton_plantations") {
														case true:
															newstate.HasCotton = true
														case false:
															switch strings.Contains(stream.CurrentToken().ValueString(), "bg_opium_plantations") {
															case true:
																newstate.HasOpium = true
															case false:
																switch strings.Contains(stream.CurrentToken().ValueString(), "bg_dye_plantations") {
																case true:
																	newstate.HasDye = true
																case false:
																	switch strings.Contains(stream.CurrentToken().ValueString(), "bg_vineyard_plantations") {
																	case true:
																		newstate.HasWine = true
																	case false:
																		switch strings.Contains(stream.CurrentToken().ValueString(), "bg_silk_plantations") {
																		case true:
																			newstate.HasSilk = true
																		case false:
																			switch strings.Contains(stream.CurrentToken().ValueString(), "bg_sugar_plantations") {
																			case true:
																				newstate.HasSugar = true
																			case false:
																				switch strings.Contains(stream.CurrentToken().ValueString(), "bg_tea_plantations") {
																				case true:
																					newstate.HasTea = true
																				case false:
																					switch strings.Contains(stream.CurrentToken().ValueString(), "bg_banana_plantations") {
																					case true:
																						newstate.HasBananas = true
																					case false:
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
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
}
