// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  myjson

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( AccountBalance ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* AccountBalance ) UnmarshalJSON([]byte) error { return nil }
func ( AccountBalance ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* AccountBalance ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_AccountBalance *AccountBalance

func ( CurrencyAmount ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* CurrencyAmount ) UnmarshalJSON([]byte) error { return nil }
func ( CurrencyAmount ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* CurrencyAmount ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_CurrencyAmount *CurrencyAmount