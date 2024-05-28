package main

import (
	"fmt"
	"os/exec"
	"testing"
)

// First, we initialize a testCases [] struct to store all our input and want cases
// We then loop through testCases and run subtests for each case
// We create a cmd using os/exec.Command to construct a command to run the program
// We then capture the output of the command (got) and compare it to want
// If they're not equal, we throw an error.
func TestMain(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},
		{
			input: "\n",
			want:  "$\n",
		},
		{
			input: "Hello\n",
			want: ` _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
`,
		},
		{
			input: "Hello",
			want: ` _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
`,
		},
		{
			input: "HeLlO",
			want: ` _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $
`,
		},
		{
			input: "Hello There",
			want: ` _    _          _   _                 _______   _                           $
| |  | |        | | | |               |__   __| | |                          $
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| $
                                                                             $
                                                                             $
`,
		},
		{
			input: "1Hello 2There",
			want: `     _    _          _   _                         _______   _                           $
 _  | |  | |        | | | |                ____   |__   __| | |                          $
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  $
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ $
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ $
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| $
                                                                                         $
                                                                                         $
`,
		},
		{
			input: "{Hello There}",
			want: `   __  _    _          _   _                 _______   _                           __    $
  / / | |  | |        | | | |               |__   __| | |                          \ \   $
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  $
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / $
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  $
  \_\                                                                              /_/   $
                                                                                         $
`,
		},
		{
			input: "Hello\nThere",
			want: ` _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
`,
		},
		{
			input: "Hello\n\nThere",
			want: ` _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			cmd := exec.Command("sh", "-c", fmt.Sprintf("go run . \"%s\" | cat -e", tc.input))
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Error running program: %v", err)
			}
			got := string(output)
			if got != tc.want {
				t.Errorf("\ngot:\n%v\nwant:\n%v\n", got, tc.want)
			}
		})
	}
}
