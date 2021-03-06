package blogic

import (
	"github.com/abondar24/SocialTournamentService/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogic_AddPlayer(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, pId > 0)
	assert.Nil(t, err)

}

func TestLogic_Take(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	err = l.Take(pId, 50)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := ds.BeginTx()
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	p, err := ds.GetPlayerById(pId, tx)
	if err != nil {
		t.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 50, p.Points)
}

func TestLogic_Fund(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	err = l.Fund(pId, 50)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := ds.BeginTx()
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	p, err := ds.GetPlayerById(pId, tx)
	if err != nil {
		t.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 150, p.Points)
}

func TestLogic_AnnounceTournament(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	tName := "blacjack"
	tDeposit := 50

	tId, err := l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, tId > 0)
	assert.Nil(t, err)

}

func Test_Logic_Join_Tournament(t *testing.T) {

	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	tName := "blacjack"
	tDeposit := 50

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{}

	err = l.JoinTournament(tId, pId, backerIds)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nil, err)
}

func Test_Logic_Join_Tournament_InsufficientFunds(t *testing.T) {

	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	tName := "blacjack"
	tDeposit := 200

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{}

	err = l.JoinTournament(tId, pId, backerIds)
	assert.Equal(t, "not enough points", err.Error())
}

func Test_Logic_Join_Tournament_With_Backers(t *testing.T) {

	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	tName := "blacjack"
	tDeposit := 200

	pbName := "Rudi"
	pbPoints := 1000

	pb1Name := "Hans"
	pb1Points := 1000

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	pbId, err := l.AddPlayer(pbName, pbPoints)
	if err != nil {
		t.Fatal(err)
	}

	pb1Id, err := l.AddPlayer(pb1Name, pb1Points)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	err = l.JoinTournament(tId, pbId, &[]int64{})
	if err != nil {
		t.Fatal(err)
	}

	err = l.JoinTournament(tId, pb1Id, &[]int64{})
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{pbId, pb1Id}

	err = l.JoinTournament(tId, pId, backerIds)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nil, err)
}

func Test_Logic_Join_Tournament_With_Backers_Not_In_Tournament(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	tName := "blacjack"
	tDeposit := 200

	pbName := "Rudi"
	pbPoints := 1000

	pb1Name := "Hans"
	pb1Points := 1000

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	pbId, err := l.AddPlayer(pbName, pbPoints)
	if err != nil {
		t.Fatal(err)
	}

	pb1Id, err := l.AddPlayer(pb1Name, pb1Points)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	err = l.JoinTournament(tId, pbId, &[]int64{})
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{pbId, pb1Id}

	err = l.JoinTournament(tId, pId, backerIds)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nil, err)
}

func TestLogic_Balance(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	b, err := l.Balance(pId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, pId, b.PlayerId)
	assert.Equal(t, pPoints, b.Balance)
}

func TestLogic_ResultTournament(t *testing.T) {

	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	tName := "blacjack"
	tDeposit := 200

	pbName := "Rudi"
	pbPoints := 1000

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	pbId, err := l.AddPlayer(pbName, pbPoints)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{pbId}

	err = l.JoinTournament(tId, pbId, &[]int64{})
	if err != nil {
		t.Fatal(err)
	}

	err = l.JoinTournament(tId, pId, backerIds)
	if err != nil {
		t.Fatal(err)
	}

	err = l.UpdatePrize(tId, pId, 50)
	if err != nil {
		t.Fatal(err)
	}

	err = l.UpdatePrize(tId, pbId, 50)
	if err != nil {
		t.Fatal(err)
	}

	tr, err := l.ResultTournament(tId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tId, tr.TournamentId)
	assert.Equal(t, 2, len(tr.Winners))
}

func TestLogic_UpdatePrize(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	tName := "blacjack"
	tDeposit := 200

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	prize := 100

	err = l.UpdatePrize(tId, pId, prize)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nil, err)
}

func TestLogic_GetAllTournaments(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	tName := "blacjack"
	tDeposit := 200

	tName1 := "blacjack"
	tDeposit1 := 200

	_, err = l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	_, err = l.AnnounceTournament(tName1, tDeposit1)
	if err != nil {
		t.Fatal(err)
	}

	tournaments, err := l.GetAllTournaments()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(*tournaments))

}

func TestLogic_GetAllPlayers(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	pName1 := "Arsen"
	pPoints1 := 100

	_, err = l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	_, err = l.AddPlayer(pName1, pPoints1)
	if err != nil {
		t.Fatal(err)
	}

	players, err := l.GetAllPlayers()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(*players))

}

func TestLogic_GetPlayersTournament(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	pName := "Ahmed"
	pPoints := 100

	tName := "blacjack"
	tDeposit := 200

	pbName := "Rudi"
	pbPoints := 1000

	pb1Name := "Hans"
	pb1Points := 1000

	pId, err := l.AddPlayer(pName, pPoints)
	if err != nil {
		t.Fatal(err)
	}

	pbId, err := l.AddPlayer(pbName, pbPoints)
	if err != nil {
		t.Fatal(err)
	}

	pb1Id, err := l.AddPlayer(pb1Name, pb1Points)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := l.AnnounceTournament(tName, tDeposit)
	if err != nil {
		t.Fatal(err)
	}

	err = l.JoinTournament(tId, pbId, &[]int64{})
	if err != nil {
		t.Fatal(err)
	}

	err = l.JoinTournament(tId, pb1Id, &[]int64{})
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{pbId, pb1Id}

	err = l.JoinTournament(tId, pId, backerIds)
	if err != nil {
		t.Fatal(err)
	}

	players, err := l.GetPlayersTournament(tId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(*players))
	assert.Equal(t, "Ahmed", (*players)[0].Name)
}

func Test_Clear_DataBase(t *testing.T) {
	ds, err := data.ConnectToTestBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}
	err = ds.ClearDB()

}
