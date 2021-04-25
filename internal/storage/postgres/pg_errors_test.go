package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsViolateFK(t *testing.T){
	fkErr := `ERROR #23503 insert or update on table "accounts" violates foreign key constraint "accounts_customer_id_fkey"`
	assert.EqualValues(t,true,IsViolateFK(fkErr))
}

func TestIsViolateFK_Error(t *testing.T){
	fkErr := `ERROR #235 insert or update on table "accounts" violates foreign key constraint "accounts_customer_id_fkey"`
	assert.EqualValues(t,false,IsViolateFK(fkErr))
}
