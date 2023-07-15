package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/yyt030/myfaker/internal/getters"
	"github.com/yyt030/myfaker/tableparser"
	tu "github.com/yyt030/myfaker/testutils"
)

func TestZHStr(t *testing.T) {
	fake.SetLang("zh")
	for i := 0; i < 10; i++ {
		t.Logf("%s", fake.FullName())
	}
}

func TestGetSamples(t *testing.T) {
	conn := tu.GetMySQLConnection(t)
	var wantRows int64 = 100
	samples, err := getSamples(conn, "sakila", "inventory", "inventory_id", wantRows, "int")
	tu.Ok(t, err, "error getting samples")
	_, ok := samples[0].(int64)
	tu.Assert(t, ok, "Wrong data type.")
	tu.Assert(t, int64(len(samples)) == wantRows,
		"Wrong number of samples. Have %d, want 100.", len(samples))
}

func TestGenerateInsertData(t *testing.T) {
	wantRows := 3

	values := []getter{
		getters.NewRandomInt("f1", 100, false),
		getters.NewRandomString("f2", 10, false),
		getters.NewRandomDate("f3", false),
	}

	rowsChan := make(chan []getter, 100)
	count := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			select {
			case <-time.After(10 * time.Millisecond):
				wg.Done()
				return
			case row := <-rowsChan:
				if reflect.TypeOf(row[0]).String() != "*getters.RandomInt" {
					fmt.Printf("Expected '*getters.RandomInt' for field [0], got %q\n", reflect.TypeOf(row[0]).String())
					t.Fail()
				}
				if reflect.TypeOf(row[1]).String() != "*getters.RandomString" {
					fmt.Printf("Expected '*getters.RandomString' for field [1], got %q\n", reflect.TypeOf(row[1]).String())
					t.Fail()
				}
				if reflect.TypeOf(row[2]).String() != "*getters.RandomDate" {
					fmt.Printf("Expected '*getters.RandomDate' for field [2], got %q\n", reflect.TypeOf(row[2]).String())
					t.Fail()
				}
				count++
			}
		}
	}()

	generateInsertData(wantRows, values, rowsChan)

	wg.Wait()
	tu.Assert(t, count == 3, "Invalid number of rows")
}

func TestGenerateInsertStmt(t *testing.T) {
	var table *tableparser.Table
	tu.LoadJson(t, "sakila.film.json", &table)
	want := "INSERT IGNORE INTO `sakila`.`film` " +
		"(`title`,`description`,`release_year`,`language_id`," +
		"`original_language_id`,`rental_duration`,`rental_rate`," +
		"`length`,`replacement_cost`,`rating`,`special_features`," +
		"`last_update`) VALUES "

	query := generateInsertStmt(table)
	tu.Equals(t, want, query)
}

func TestCNStr(t *testing.T) {
	err := fake.SetLang("zh")
	if err != nil {
		panic(err)
	}

	t.Logf(">>>%s", fake.Sentence())
}

func TestSplitUser(t *testing.T) {
	users := []string{"u@t#g:gi", "u@t#g:gi1:gi2", "u@t#g"}
	for _, v := range users {
		ss := strings.SplitN(v, ":", 2)
		t.Logf("%s=>%s", v, ss[0])
	}
}
