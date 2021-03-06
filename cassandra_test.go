package main

import (
	"fmt"
	"testing"
	"github.com/carloscm/gossie/src/gossie"
)

func TestGossie(t *testing.T) {
	var keyspace = "hecunatest"
	var colspace = "snps"
	pool, err :=
gossie.NewConnectionPool([]string{"127.0.0.1:9160"}, keyspace, gossie.PoolOptions{Size: 1, Timeout: 30000})
	if err != nil {
		ExitMsg(fmt.Sprint("Connecting: ", err))
	}
	fmt.Println("Connected to keyspace ", keyspace)

	var mapping, mappingErr = gossie.NewMapping(&SNP{})
	if mappingErr != nil {
		ExitMsg(fmt.Sprint("mapping: ", mappingErr))
	}

	var snpObj = &SNP{"testkey", "does nothing", "C", "CAT"}
	var testSNP, mapErr = mapping.Map(snpObj)
	if mapErr != nil {
		ExitMsg("Can't map value.")
	}
	err = pool.Writer().Insert(colspace, testSNP).Run()
	if err != nil {
		ExitMsg(fmt.Sprint("Couldn't write: ", err))
	}
	fmt.Println("Connected.")
	var query = pool.Query(mapping)
	var ret, readErr = query.Get(12)
	if readErr != nil {
		ExitMsg(fmt.Sprint("Couldn't read: ", readErr))
	}
	for {
		res_snp := &SNP{}
		err := ret.Next(res_snp)
		if err != nil {
			break
		}
		fmt.Println(res_snp.Alleles)
	}
	// Output: CAT
}

