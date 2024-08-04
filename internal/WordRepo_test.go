package internal

import "testing"

var wrep WordRepo

func prepare(){
  wrep = NewWordRepo()
}

func finish(){
  wrep.dao.db.Close()
}

func TestGetTenWordsRepo(t *testing.T){
  prepare()
  defer finish()
  words, err := wrep.GetTenWords(0)
  if err != nil{
    t.Fatal("Could not retrieve ten words: ", err)
  }

  if (len(words)!=10){
    t.Fatal("Did not retrieve ten words")
  }
}
