// Copyright © 2016 Salesforce.com
//

package commands

import (
	"log"
	"os"
	"time"

	"github.com/jusongchen/ora"
	"github.com/spf13/cobra"
)

var OraPingCmd = &cobra.Command{
	Use:    "pingOra",
	Short:  "Ping an Oracle DB instance",
	Long:   `Ping an Oracle DB instance`,
	Run:    PingOra,
	PreRun: verifyArgs,
}

var oraParam = ora.OraConnParams{}
var interval = time.Second

func init() {
	RootCmd.AddCommand(OraPingCmd)
	// Here you will define your flags and configuration settings.

	OraPingCmd.Flags().StringVarP(&oraParam.OraUser, "User", "u", "", "The Oracle username who is also the DB schema name of DSP repository")
	OraPingCmd.Flags().StringVarP(&oraParam.OraPasswd, "Password", "p", "tiger", "password of the Oracle username")
	OraPingCmd.Flags().StringVarP(&oraParam.OraConnID, "ConnectID", "c", "//localhost/orcl",
		`Oracle DB connection identifier in the form of Net Service Name or Easy Connect: 
		 	 [<net_service_name> | [//]Host[:Port]/<service_name>]
		`)
}

func verifyArgs(cmd *cobra.Command, args []string) {
	if oraParam.OraUser == "" {
		log.Printf("Parameter User is mandatory.")
		os.Exit(1)
	}

}

func PingOra(cmd *cobra.Command, args []string) {

	//func Import(pgConn PgConnParams, oraConn OraConnParams) {
	ora.PingOra(oraParam, interval)
	// TODO: Work your own magic here
}

//db, err := sql.Open("postgres", "postgres://pgUser:password@localhost/DBname?sslmode=verify-full")
