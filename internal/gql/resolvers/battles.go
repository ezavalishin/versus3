package resolvers

import (
	"context"
	"github.com/ezavalishin/versus3/internal/auth"
	gqlm "github.com/ezavalishin/versus3/internal/gql/models"
	"github.com/ezavalishin/versus3/internal/gql/resolvers/transformations"
	"github.com/ezavalishin/versus3/internal/logger"
	dbm "github.com/ezavalishin/versus3/internal/orm/models"
	"math"
)

func (r *queryResolver) GetActiveBattles(ctx context.Context) ([]*gqlm.Battle, error) {

	return activeBattles(r, ctx)
}

func (r *mutationResolver) StartBattle(ctx context.Context, battleId int) ([]*gqlm.Round, error) {

	user := auth.ForContext(ctx)

	battleUser := dbm.BattleUser{
		BattleId: battleId,
		UserId:   user.ID,
	}

	r.ORM.DB.Create(&battleUser)

	var rounds []*gqlm.Round
	var unitsCount int

	r.ORM.DB.Model(&dbm.Unit{}).Where("battle_id = ?", battleId).Count(&unitsCount)

	for i := 0; i < int(math.Log2(float64(unitsCount))); i++ {
		round := &dbm.Round{
			BattleUserId: battleUser.ID,
			Step:         i,
		}

		r.ORM.DB.Create(&round)

		gqlRound, _ := transformations.DBRoundToGQLRound(round)

		rounds = append(rounds, gqlRound)
	}

	return rounds, nil
}

func activeBattles(r *queryResolver, ctx context.Context) ([]*gqlm.Battle, error) {

	var records []*gqlm.Battle
	var dbRecords []*dbm.Battle

	db := r.ORM.DB.New()

	db = db.Find(&dbRecords)

	for _, dbRec := range dbRecords {
		if rec, err := transformations.DBBattleToGQLBattle(dbRec); err != nil {
			logger.Errorfn("error", err)
		} else {
			records = append(records, rec)
		}
	}

	return records, nil
}
