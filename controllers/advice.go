package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"fanctionary/models"
	"fanctionary/utils"

	"github.com/gorilla/mux"
	"github.com/littletwolee/commons"
)

type Advice struct {
	loger *commons.Log
}

func GetAdviceController() *Advice {
	return &Advice{
		loger: commons.GetLogger(),
	}
}

var mysqlHelper = commons.GetMysqlHelper()

func (a *Advice) GetAdvice(w http.ResponseWriter, r *http.Request) {
	openID := mux.Vars(r)["open_id"]
	if openID == "" {
		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
	}
	// adviceType := mux.Vars(r)["type"]
	// if adviceType == "" {
	// 	utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
	// }
	sqlRuler := &commons.SQLRule{
		Table: "user_advices_list",
		Where: &commons.Conditions{
			Sentence:   "open_id = ?",
			Parameters: []interface{}{openID},
		},
	}
	var exp models.UserAdvicesList
	err := mysqlHelper.FindOne(sqlRuler, &exp)
	if err != nil {
		if err.Error() == "record not found" {
			list, err := getAdviceIDList()
			if err != nil {
				utils.ServerError(w, err)
			}
			list = getRandList(list)
			err = insertUAList(openID, list[1:])
			if err != nil {
				utils.ServerError(w, err)
			}
		}
		utils.ServerError(w, err)
	}
	// var exps []models.Exp
	// err = json.Unmarshal([]byte(exp.Comment), &exps)
	// if err != nil {
	// 	utils.ServerError(w, err)
	// }
	// if err := json.NewEncoder(w).Encode(models.NewResult(err, exps)); err != nil {
	// 	utils.ServerError(w, err)
	// 	return
	// }
}

func getRandList(list []int) []int {
	for k, _ := range list {
		r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(list))
		list[k], list[r] = list[r], list[k]
	}
	return list
}

func insertUAList(openID string, list []int) error {
	insertRuler := &commons.SQLRule{
		Table: "user_advices_list",
	}
	ual := &models.UserAdvicesList{
		OpenID: openID,
		List:   getStr(list),
	}
	return mysqlHelper.Insert(insertRuler, ual)
}

func getStr(list []int) string {
	var str string
	for _, v := range list {
		str += fmt.Sprintf("%d,", v)
	}
	return str
}

func getAdviceIDList() ([]int, error) {
	sqlRuler := &commons.SQLRule{
		Table:  "advices",
		Select: []string{"id"},
	}
	var list []int
	err := mysqlHelper.FindAll(sqlRuler, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
