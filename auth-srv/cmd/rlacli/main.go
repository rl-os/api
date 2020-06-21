package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string
	commitHash string
	buildDate  string
)

const asciiArt = `
                                        
            ,:loooooooooll:'            
         :oddoc:,,'''',;:loddl,         
      'lxdc,                ;lxd:       
     :xd;                     'cxd,     
    lkc                         ,dk:    
   ok:    ;dxoooooc' co'          ok;   
  :Oo     c0o''',ckkoOO;          'xx'  
  dk;     c0l     lKXXO;           lk:  
 'xk'     c0l '',lkkoOO;           :Ol  
  xk,     c0l'lOOdc,'dO;           cOc  
  lOc     c0l 'lko'  oO;           ok;  
  ,xk,    cOc   ;xx: ;xxolllc'    :kl   
   ;kx,   ';'    ';,   ;:ccc:'   :ko'   
    ,xkc                       'lkl'    
     'lkx:'                  'cxd;      
       'cdxoc,           ';cdxo;        
          ,codddoooooooooddl:'          
              ',;:cc:::;,               
          

    GitHub: github.com/deissh/osu-lazer
                      2019-2020, deissh
`

func main() {
	rootCmd := &cobra.Command{
		Use:  "rlacli",
		Long: asciiArt,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
