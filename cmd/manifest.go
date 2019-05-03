// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wesller/wpm/domain"
)

var all bool
var folder string

// manifestCmd represents the manifest command
var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Gerar Manifesto",
	Long:  `Gera o arquivo de manifesto automaticamente`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("manifest called")
	domain.Generate(folder)
}

func init() {
	RootCmd.AddCommand(manifestCmd)
	manifestCmd.Flags().BoolVarP(&all, "all", "a", false, "Gerar um pacote para todos os arquivos da pasta")
	manifestCmd.Flags().StringVarP(&folder, "folder", "f", "", "Pasta inicial para gerar o manifest")
}
