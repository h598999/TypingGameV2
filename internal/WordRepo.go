package internal

import ("fmt")


type WordRepo struct{
  dao UserDAO
}

func (wr *WordRepo) GetTenWords(id int64) ([]string, error){
  words, err := wr.dao.getTenWords(id)
  if err != nil{
    fmt.Print(err.Error())
    return words, err
  }
  return words, err
}

func NewWordRepo() (WordRepo){
  return WordRepo{
    dao: NewUserDAO(),
  }
}
