package transformations

import (
	gql "github.com/ezavalishin/versus3/internal/gql/models"
	dbm "github.com/ezavalishin/versus3/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBPairToGQLPair(i *dbm.Pair) (o *gql.Pair, err error) {

	unitOneGl, _ := DBUnitToGQLUnit(&i.UnitOne)
	unitTwoGl, _ := DBUnitToGQLUnit(&i.UnitTwo)
	unitOnePickedGl, _ := DBPickedToGQLPicked(&i.UnitOnePicked)
	unitTwoPickedGl, _ := DBPickedToGQLPicked(&i.UnitTwoPicked)

	o = &gql.Pair{
		ID:            i.ID,
		UnitOne:       unitOneGl,
		UnitTwo:       unitTwoGl,
		UnitOnePicked: unitOnePickedGl,
		UnitTwoPicked: unitTwoPickedGl,
	}
	return o, err
}
