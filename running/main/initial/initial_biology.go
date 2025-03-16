package initial

import (
	"math/rand"
	"time"

	"github.com/logos/ecs/components"
	"github.com/logos/ecs/components/attributes"
	"github.com/logos/ecs/entities"
	"github.com/logos/ecs/entities/unit"
	"github.com/logos/global"
)

func InitialBiology(world *entities.World) []*unit.Unit {
	var units = make([]*unit.Unit, world.TotalAura)
	for i := 0; i < int(world.TotalAura); i++ {
		unit := initUnit(world)
		units[i] = &unit
		// units[i].Display()
	}

	return units
}

func initUnit(world *entities.World) unit.Unit {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	defer global.IncrementCurrentUnitID()
	unitTypeID := components.RandomUnitType()
	characterAttributes := attributes.CharacterAttributes{
		Vitality:     global.CARSetting.VitalityRange.Min + r.Intn(global.CARSetting.VitalityRange.Max-global.CARSetting.VitalityRange.Min),
		QiEnergy:     global.CARSetting.QiEnergyRange.Min + r.Intn(global.CARSetting.QiEnergyRange.Max-global.CARSetting.QiEnergyRange.Min),
		Strength:     global.CARSetting.StrengthRange.Min + r.Intn(global.CARSetting.StrengthRange.Max-global.CARSetting.StrengthRange.Min),
		Agility:      global.CARSetting.AgilityRange.Min + r.Intn(global.CARSetting.AgilityRange.Max-global.CARSetting.AgilityRange.Min),
		Intelligence: global.CARSetting.IntelligenceRange.Min + r.Intn(global.CARSetting.IntelligenceRange.Max-global.CARSetting.IntelligenceRange.Min),
		Age:          global.CARSetting.AgeRange.Min + r.Intn(global.CARSetting.AgeRange.Max-global.CARSetting.AgeRange.Min),
	}
	practiceAttributes := attributes.PracticeAttributes{
		// TODO:
		SpiritualRoot:    "Wood",
		BodyConstitution: global.ParSetting.BodyConstitutionRange.Min + r.Intn(global.ParSetting.BodyConstitutionRange.Max-global.ParSetting.BodyConstitutionRange.Min),
		Comprehension:    global.ParSetting.ComprehensionRange.Min + r.Intn(global.ParSetting.ComprehensionRange.Max-global.ParSetting.ComprehensionRange.Min),
		LuckFortune:      global.ParSetting.LuckFortuneRange.Min + r.Intn(global.ParSetting.LuckFortuneRange.Max-global.ParSetting.LuckFortuneRange.Min),
		Willpower:        global.ParSetting.WillpowerRange.Min + r.Intn(global.ParSetting.WillpowerRange.Max-global.ParSetting.WillpowerRange.Min),
		CultivationLevel: global.ParSetting.CultivationLevelRange.Min + r.Intn(global.ParSetting.CultivationLevelRange.Max-global.ParSetting.CultivationLevelRange.Min),
	}

	newUnit := unit.NewUnit(
		global.CurrentUnitID,
		// TODO: Name generation
		"TestUnit",
		unitTypeID,
		characterAttributes,
		practiceAttributes,
		attributes.CalculateCombatAttributes(characterAttributes, practiceAttributes),
		world.RandomLocation().Aura,
	)

	return newUnit
}
