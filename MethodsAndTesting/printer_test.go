package asciiart

import "testing"

func TestFormatPrinter(t *testing.T) {
	checks := []struct {
		input          string
		expectedOutput string
	}{
		{
			`[\]^_ 'a\n`,
			`
 ___  __       ___   /\                  _          
|  _| \ \     |_  | |/\|          "      ( )         
| |    \ \      | |                     |/    __ _  
| |     \ \     | |                          / _` + "`" + ` | 
| |      \ \    | |                         | (_| | 
| |_      \_\  _| |                          \__,_| 
|___|         |___|       ______                    
                         |______|                   
`,
		},
		{
			`hello\n`,
			`
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
							   $
`,
		},
	}

	for _, word := range checks {
		funcOutput := FormatPrinter(word.input)
		if word.expectedOutput != funcOutput {
			t.Errorf("Expected: %v, Got: %v", word.expectedOutput, funcOutput)
		}

	}
}
