package gorm_gen

import (
	"log"
	"testing"
)

func TestGen(t *testing.T) {
	log.SetFlags(log.LstdFlags)
	Gen("root:P@ssw0rd@tcp(192.168.4.32:3306)/ark_mall?parseTime=true", "item", "ark_mall", "po")
}
