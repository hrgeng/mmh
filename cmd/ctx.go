/*
 * Copyright 2018 mritd <mritd1234@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/mritd/mmh/pkg/mmh"
	"github.com/spf13/cobra"
)

var ctxCmd = &cobra.Command{
	Use:   "ctx CONTEXT",
	Short: "Change current context",
	Long: `
Change current context.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			mmh.ContextUse(args[0])
		} else {
			_ = cmd.Help()
		}
	},
}

var ctxListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List context",
	Long: `
List context`,
	Run: func(cmd *cobra.Command, args []string) {
		mmh.ContextList()
	},
}

var ctxUseCmd = &cobra.Command{
	Use:   "use",
	Short: "Use context",
	Long: `
Use context`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Help()
			return
		}
		mmh.ContextUse(args[0])
	},
}

func init() {
	ctxCmd.AddCommand(ctxListCmd, ctxUseCmd)
	RootCmd.AddCommand(ctxCmd)
}
