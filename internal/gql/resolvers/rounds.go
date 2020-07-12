package resolvers

import (
	"context"
	"errors"
	"github.com/ezavalishin/versus3/internal/auth"
	gqlm "github.com/ezavalishin/versus3/internal/gql/models"
	"github.com/ezavalishin/versus3/internal/gql/resolvers/transformations"
	dbm "github.com/ezavalishin/versus3/internal/orm/models"
	"math/rand"
	"time"
)

func (r *queryResolver) GetRound(ctx context.Context, roundId int) ([]*gqlm.Pair, error) {

	user := auth.ForContext(ctx)

	var modelPairs []*dbm.Pair
	var pairs []*gqlm.Pair

	r.ORM.DB.Preload("UnitOne").Preload("UnitTwo").Where("round_id = ?", roundId).Find(&modelPairs)

	if len(modelPairs) > 0 {
		for _, pair := range modelPairs {
			pair.LoadPicked(r.ORM.DB)
			pairGql, _ := transformations.DBPairToGQLPair(pair)
			pairs = append(pairs, pairGql)
		}

		return pairs, nil
	}

	round := dbm.Round{}
	battleUser := dbm.BattleUser{}
	battle := dbm.Battle{}
	var units []*dbm.Unit

	r.ORM.DB.First(&round, roundId)
	r.ORM.DB.First(&battleUser, round.BattleUserId)

	if round.Step != 0 {

		prevRound := dbm.Round{}

		prevStep := round.Step - 1

		r.ORM.DB.Where("step = ? AND battle_user_id = ?", prevStep, battleUser.ID).First(&prevRound)

		if prevRound.ID == 0 {
			return nil, errors.New("invalid round, not found prev")
		}

		r.ORM.DB.Preload("UserPick").Where("round_id = ?", prevRound.ID).Find(&modelPairs)

		if len(modelPairs) == 0 {
			return nil, errors.New("invalid round, has no pairs")
		}

		for _, pair := range modelPairs {
			units = append(units, pair.UserPick)
		}

	} else {
		r.ORM.DB.First(&battle, battleUser.BattleId)
		r.ORM.DB.Where("battle_id = ?", battle.ID).Find(&units)
	}

	if len(units)%2 != 0 {
		return nil, errors.New("invalid units count")
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(units), func(i, j int) { units[i], units[j] = units[j], units[i] })

	for i := 0; i < len(units)/2; i++ {
		pair := dbm.Pair{
			RoundId:   round.ID,
			UnitOne:   *units[i*2],
			UnitTwo:   *units[i*2+1],
			UnitOneId: units[i*2].ID,
			UnitTwoId: units[i*2+1].ID,
			UserId:    user.ID,
		}

		pair.LoadPicked(r.ORM.DB)

		r.ORM.DB.Create(&pair)

		pairGql, _ := transformations.DBPairToGQLPair(&pair)

		pairs = append(pairs, pairGql)
	}

	return pairs, nil
}

func newTrue() *bool {
	b := true
	return &b
}

func (r *mutationResolver) MakePick(ctx context.Context, pairID int, unitID int) (*bool, error) {

	user := auth.ForContext(ctx)

	pair := dbm.Pair{}
	round := dbm.Round{}
	battleUser := dbm.BattleUser{}

	r.ORM.DB.First(&pair, pairID)
	r.ORM.DB.First(&round, pair.RoundId)
	r.ORM.DB.First(&battleUser, round.BattleUserId)

	if battleUser.UserId != user.ID {
		return nil, errors.New("forbidden")
	}

	if unitID != pair.UnitOneId && unitID != pair.UnitTwoId {
		return nil, errors.New("invalid unit id")
	}

	pair.UserPickId = &unitID

	r.ORM.DB.Save(pair)

	return newTrue(), nil
}
