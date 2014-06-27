package main

import (
	"fmt"
	"math"
)

type Cookieverse struct {

	// items
	Cursors              int32
	Grandmas             int32
	Farms                int32
	Factories            int32
	Mines                int32
	Shipments            int32
	AlchemyLabs          int32
	Portals              int32
	TimeMachines         int32
	AntimatterCondensers int32
	Prisms               int32

	// cursor upgrades
	ReinforcedIndexFinger       bool
	CarpalTunnelPreventionCream bool
	Ambidextrous                bool
	ThousandFingers             bool
	MillionFingers              bool
	BillionFingers              bool
	TrillionFingers             bool
	QuadrillionFingers          bool
	QuintillionFingers          bool
	SextillionFingers           bool
	SeptillionFingers           bool
	OctillionFingers            bool

	// grandma upgrades
	ForwardsFromGrandma    bool
	SteelPlatedRollingPins bool
	LubricatedDentures     bool
	PruneJuice             bool
	DoubleThickGlasses     bool
	AgingAgents            bool
	FarmerGrandmas         bool
	WorkerGrandmas         bool
	MinerGrandmas          bool
	CosmicGrandmas         bool
	TransmutedGrandmas     bool
	AlteredGrandmas        bool
	GrandmasGrandmas       bool
	AntiGrandmas           bool
	RainbowGrandmas        bool

	// farm upgrades
	CheapHoes                  bool
	Fertilizer                 bool
	CookieTrees                bool
	GeneticallyModifiedCookies bool
	GingerbreadScarecrows      bool
	PulsarSprinklers           bool

	// factory upgrades
	SturdierConveyorBelts bool
	ChildLabor            bool
	SweatShop             bool
	RadiumReactors        bool
	Recombobulators       bool
	DeepBakeProcess       bool

	// mine upgrades
	SugarGas    bool
	Megadrill   bool
	Ultradrill  bool
	Ultimadrill bool
	HBombMining bool
	Coreforge   bool

	// shipment upgrades
	VanillaNebulae     bool
	Wormholes          bool
	FrequentFlyer      bool
	WarpDrive          bool
	ChocolateMonoliths bool
	GenerationShip     bool

	// alchemy lab upgrades
	Antimony       bool
	EssenceOfDough bool
	TrueChocolate  bool
	Ambrosia       bool
	AcquaCrustulae bool
	OriginCrucible bool

	// portals upgrades
	AncientTablet        bool
	InsaneOatlingWorkers bool
	SouldBond            bool
	SanityDance          bool
	BraneTransplant      bool
	DeitySizedPortals    bool

	// time machines upgrades
	FluxCapacitors          bool
	TimeParadoxResolver     bool
	QuantumConundrum        bool
	CausalityEnforcer       bool
	YestermorrowComparators bool
	FarFutureEnactment      bool

	// antimatter condenser upgrades
	SugarBosons          bool
	StringTheory         bool
	LargeMacaronCollider bool
	BigBangCake          bool
	ReverseCyclotrons    bool
	Nanocosmics          bool

	// prisms upgrades
	GemPolish       bool
	NinthColor      bool
	ChocolateLight  bool
	Grainbow        bool
	PureCosmicLight bool
	GlowInTheDark   bool

	// research upgrades
	DisableResearch             bool
	BingoCenterResearchFacility bool
	SpecializedChocolateChips   bool
	DesignerCocoaBeans          bool
	RitualRollingPins           bool
	UnderworldOvens             bool
	OneMind                     bool
	ExoticNuts                  bool
	CommunalBrainsweep          bool
	ArcaneSugar                 bool
	ElderPact                   bool

	// cookies upgrades
	OutmealRaisinCookies                           bool
	PeanutButterCookies                            bool
	PlainCookies                                   bool
	CoconutCookies                                 bool
	WhiteChocolateCookies                          bool
	MacadamiaNutCookies                            bool
	SugarCookies                                   bool
	DoubleChipCookies                              bool
	WhiteChocolateMacadamiaNutCookies              bool
	AllChocolateCookies                            bool
	DarkCholocateCoatedCookies                     bool
	WhiteCholocateCoatedCookies                    bool
	EclipseCookies                                 bool
	ZebraCookies                                   bool
	Snickerdoodles                                 bool
	Stroopwafels                                   bool
	Macaroons                                      bool
	EmpireBiscuits                                 bool
	BritishTeaBiscuits                             bool
	ChocolateBritishTeaBiscuits                    bool
	RoundBritishTeaBiscuits                        bool
	RoundChocolateBritishTeaBiscuits               bool
	RoundBritishTeaBiscuitsWithHeartMotif          bool
	RoundChocolateBritishTeaBiscuitsWithHeartMotif bool
	Madeleines                                     bool
	Palmiers                                       bool
	Palets                                         bool
	Sables                                         bool
	Caramoas                                       bool
	Sagalongs                                      bool
	Shortfoils                                     bool
	WinMints                                       bool
	FigGluttons                                    bool
	Loreols                                        bool
	JaffaCakes                                     bool
	GreasesCups                                    bool
	RoseMacarons                                   bool
	LemonMacarons                                  bool
	ChocolateMacarons                              bool
	PistachioMacarons                              bool
	HazelnutMacarons                               bool
	VioletMacarons                                 bool
	CaramelMacarons                                bool
	LicoriceMacarons                               bool
	GingerbreadMen                                 bool
	GingerbreadTrees                               bool

	// wrinklers
	DisableSkullCookies   bool
	SkullCookies          bool
	DisableSlimeCookies   bool
	SlimeCookies          bool
	DisableEyeballCookies bool
	EyeballCookies        bool
	DisableSpiderCookies  bool
	SpiderCookies         bool
	DisableGhostCookies   bool
	GhostCookies          bool
	DisableBatCookies     bool
	BatCookies            bool
	DisablePumpkinCookies bool
	PumpkinCookies        bool

	// easter
	DisableChickenEgg   bool
	ChickenEgg          bool
	DisableDuckEgg      bool
	DuckEgg             bool
	DisableTurkeyEgg    bool
	TurkeyEgg           bool
	DisableQuailEgg     bool
	QuailEgg            bool
	DisableRobinEgg     bool
	RobinEgg            bool
	DisableOstrichEgg   bool
	OstrichEgg          bool
	DisableCassowaryEgg bool
	CassowaryEgg        bool
	DisableSalmonRoe    bool
	SalmonRoe           bool
	DisableFrogspawn    bool
	Frogspawn           bool
	DisableSharkEgg     bool
	SharkEgg            bool
	DisableTurtleEgg    bool
	TurtleEgg           bool
	DisableAntLarva     bool
	AntLarva            bool
	GoldenGooseEgg      bool
	FabergeEgg          bool
	Wrinklerspawn       bool
	CookieEgg           bool
	Omelette            bool
	ChocolateEgg        bool
	CenturyDays         Production
	CenturyEgg          bool
	Egg                 bool

	// heavenly upgrades
	HeavenlyChipSecret    bool
	HeavenlyCookieStand   bool
	HeavenlyBakery        bool
	HeavenlyConfectionery bool
	HeavenlyKey           bool

	// milk upgrades
	KittenHelpers   bool
	KittenWorkers   bool
	KittenEngineers bool
	KittenOverseers bool
	KittenManagers  bool

	// prestige
	Milk          int32
	HeavenlyChips int32
}

func (c *Cookieverse) CursorCpS() Production {
	baseCursors := Production(0.1)
	if c.ReinforcedIndexFinger {
		baseCursors += 0.1
	}
	mulCursors := Production(1)
	if c.CarpalTunnelPreventionCream {
		mulCursors *= 2
	}
	if c.Ambidextrous {
		mulCursors *= 2
	}
	baseNonCursor := Production(0)
	sumNonCursor := Production(c.Grandmas + c.Farms + c.Factories + c.Mines + c.Shipments + c.AlchemyLabs + c.Portals + c.TimeMachines + c.AntimatterCondensers + c.Prisms)
	if c.ThousandFingers {
		baseNonCursor += 0.1
	}
	if c.MillionFingers {
		baseNonCursor += 0.5
	}
	if c.BillionFingers {
		baseNonCursor += 2
	}
	if c.TrillionFingers {
		baseNonCursor += 10
	}
	if c.QuadrillionFingers {
		baseNonCursor += 20
	}
	if c.QuintillionFingers {
		baseNonCursor += 100
	}
	if c.SextillionFingers {
		baseNonCursor += 200
	}
	if c.SeptillionFingers {
		baseNonCursor += 400
	}
	if c.OctillionFingers {
		baseNonCursor += 800
	}
	return Production(c.Cursors) * (baseCursors*mulCursors + baseNonCursor*sumNonCursor)
}

func (c *Cookieverse) GrandmaCpS() Production {
	baseGrandmas := Production(0.5)
	if c.ForwardsFromGrandma {
		baseGrandmas += 0.3
	}
	mulGrandmas := Production(1)
	if c.SteelPlatedRollingPins {
		mulGrandmas *= 2
	}
	if c.LubricatedDentures {
		mulGrandmas *= 2
	}
	if c.PruneJuice {
		mulGrandmas *= 2
	}
	if c.DoubleThickGlasses {
		mulGrandmas *= 2
	}
	if c.AgingAgents {
		mulGrandmas *= 2
	}
	if c.FarmerGrandmas {
		mulGrandmas *= 2
	}
	if c.WorkerGrandmas {
		mulGrandmas *= 2
	}
	if c.MinerGrandmas {
		mulGrandmas *= 2
	}
	if c.CosmicGrandmas {
		mulGrandmas *= 2
	}
	if c.TransmutedGrandmas {
		mulGrandmas *= 2
	}
	if c.AlteredGrandmas {
		mulGrandmas *= 2
	}
	if c.GrandmasGrandmas {
		mulGrandmas *= 2
	}
	if c.AntiGrandmas {
		mulGrandmas *= 2
	}
	if c.RainbowGrandmas {
		mulGrandmas *= 2
	}
	if c.BingoCenterResearchFacility {
		mulGrandmas *= 4
	}
	if c.RitualRollingPins {
		mulGrandmas *= 2
	}
	baseModGrandmas := Production(0)
	if c.OneMind {
		baseModGrandmas += 1
	}
	if c.CommunalBrainsweep {
		baseModGrandmas += 1
	}
	baseModPortals := Production(0)
	if c.ElderPact {
		baseModPortals += 1
	}
	return Production(c.Grandmas) * (baseGrandmas + baseModGrandmas*Production(c.Grandmas)/50 + baseModPortals*Production(c.Portals)/20) * mulGrandmas
}

func (c *Cookieverse) FarmCpS() Production {
	baseFarms := Production(4)
	if c.CheapHoes {
		baseFarms += 1
	}
	mulFarms := Production(1)
	if c.Fertilizer {
		mulFarms *= 2
	}
	if c.CookieTrees {
		mulFarms *= 2
	}
	if c.GeneticallyModifiedCookies {
		mulFarms *= 2
	}
	if c.GingerbreadScarecrows {
		mulFarms *= 2
	}
	if c.PulsarSprinklers {
		mulFarms *= 2
	}
	return Production(c.Farms) * baseFarms * mulFarms
}

func (c *Cookieverse) FactoryCpS() Production {
	baseFactories := Production(10)
	if c.SturdierConveyorBelts {
		baseFactories += 4
	}
	mulFactories := Production(1)
	if c.ChildLabor {
		mulFactories *= 2
	}
	if c.SweatShop {
		mulFactories *= 2
	}
	if c.RadiumReactors {
		mulFactories *= 2
	}
	if c.Recombobulators {
		mulFactories *= 2
	}
	if c.DeepBakeProcess {
		mulFactories *= 2
	}
	return Production(c.Factories) * baseFactories * mulFactories
}

func (c *Cookieverse) MineCpS() Production {
	baseMines := Production(40)
	if c.SugarGas {
		baseMines += 10
	}
	mulMines := Production(1)
	if c.Megadrill {
		mulMines *= 2
	}
	if c.Ultradrill {
		mulMines *= 2
	}
	if c.Ultimadrill {
		mulMines *= 2
	}
	if c.HBombMining {
		mulMines *= 2
	}
	if c.Coreforge {
		mulMines *= 2
	}
	return Production(c.Mines) * baseMines * mulMines
}

func (c *Cookieverse) ShipmentCpS() Production {
	baseShipments := Production(100)
	if c.VanillaNebulae {
		baseShipments += 30
	}
	mulShipments := Production(1)
	if c.Wormholes {
		mulShipments *= 2
	}
	if c.FrequentFlyer {
		mulShipments *= 2
	}
	if c.WarpDrive {
		mulShipments *= 2
	}
	if c.ChocolateMonoliths {
		mulShipments *= 2
	}
	if c.GenerationShip {
		mulShipments *= 2
	}
	return Production(c.Shipments) * baseShipments * mulShipments
}

func (c *Cookieverse) AlchemyLabCpS() Production {
	baseAlchemyLabs := Production(400)
	if c.Antimony {
		baseAlchemyLabs += 100
	}
	mulAlchemyLabs := Production(1)
	if c.EssenceOfDough {
		mulAlchemyLabs *= 2
	}
	if c.TrueChocolate {
		mulAlchemyLabs *= 2
	}
	if c.Ambrosia {
		mulAlchemyLabs *= 2
	}
	if c.AcquaCrustulae {
		mulAlchemyLabs *= 2
	}
	if c.OriginCrucible {
		mulAlchemyLabs *= 2
	}
	return Production(c.AlchemyLabs) * baseAlchemyLabs * mulAlchemyLabs
}

func (c *Cookieverse) PortalCpS() Production {
	basePortals := Production(6666)
	if c.AncientTablet {
		basePortals += 1666
	}
	mulPortals := Production(1)
	if c.InsaneOatlingWorkers {
		mulPortals *= 2
	}
	if c.SouldBond {
		mulPortals *= 2
	}
	if c.SanityDance {
		mulPortals *= 2
	}
	if c.BraneTransplant {
		mulPortals *= 2
	}
	if c.DeitySizedPortals {
		mulPortals *= 2
	}
	return Production(c.Portals) * basePortals * mulPortals
}

func (c *Cookieverse) TimeMachineCpS() Production {
	baseTimeMachines := Production(98765)
	if c.FluxCapacitors {
		baseTimeMachines += 9876
	}
	mulTimeMachines := Production(1)
	if c.TimeParadoxResolver {
		mulTimeMachines *= 2
	}
	if c.QuantumConundrum {
		mulTimeMachines *= 2
	}
	if c.CausalityEnforcer {
		mulTimeMachines *= 2
	}
	if c.YestermorrowComparators {
		mulTimeMachines *= 2
	}
	if c.FarFutureEnactment {
		mulTimeMachines *= 2
	}
	return Production(c.TimeMachines) * baseTimeMachines * mulTimeMachines
}

func (c *Cookieverse) AntimatterCondenserCpS() Production {
	baseAntimatterCondensers := Production(999999)
	if c.SugarBosons {
		baseAntimatterCondensers += 99999
	}
	mulAntimatterCondensers := Production(1)
	if c.StringTheory {
		mulAntimatterCondensers *= 2
	}
	if c.LargeMacaronCollider {
		mulAntimatterCondensers *= 2
	}
	if c.BigBangCake {
		mulAntimatterCondensers *= 2
	}
	if c.ReverseCyclotrons {
		mulAntimatterCondensers *= 2
	}
	if c.Nanocosmics {
		mulAntimatterCondensers *= 2
	}
	return Production(c.AntimatterCondensers) * baseAntimatterCondensers * mulAntimatterCondensers
}

func (c *Cookieverse) PrismsCpS() Production {
	basePrisms := Production(10e6)
	if c.GemPolish {
		basePrisms += 1e6
	}
	mulPrisms := Production(1)
	if c.NinthColor {
		mulPrisms *= 2
	}
	if c.ChocolateLight {
		mulPrisms *= 2
	}
	if c.Grainbow {
		mulPrisms *= 2
	}
	if c.PureCosmicLight {
		mulPrisms *= 2
	}
	if c.GlowInTheDark {
		mulPrisms *= 2
	}
	return Production(c.Prisms) * basePrisms * mulPrisms
}

func (c *Cookieverse) BuildingCpS() (cps Production) {
	cps += c.CursorCpS()
	cps += c.GrandmaCpS()
	cps += c.FarmCpS()
	cps += c.FactoryCpS()
	cps += c.MineCpS()
	cps += c.ShipmentCpS()
	cps += c.AlchemyLabCpS()
	cps += c.PortalCpS()
	cps += c.TimeMachineCpS()
	cps += c.AntimatterCondenserCpS()
	cps += c.PrismsCpS()
	return
}

func (c *Cookieverse) NormalMultiplier() (mul Production) {
	if c.SpecializedChocolateChips {
		mul += 1
	}
	if c.DesignerCocoaBeans {
		mul += 2
	}
	if c.UnderworldOvens {
		mul += 3
	}
	if c.ExoticNuts {
		mul += 4
	}
	if c.ArcaneSugar {
		mul += 5
	}
	if c.OutmealRaisinCookies {
		mul += 5
	}
	if c.PeanutButterCookies {
		mul += 5
	}
	if c.PlainCookies {
		mul += 5
	}
	if c.CoconutCookies {
		mul += 5
	}
	if c.WhiteChocolateCookies {
		mul += 5
	}
	if c.MacadamiaNutCookies {
		mul += 5
	}
	if c.SugarCookies {
		mul += 5
	}
	if c.DoubleChipCookies {
		mul += 10
	}
	if c.WhiteChocolateMacadamiaNutCookies {
		mul += 10
	}
	if c.AllChocolateCookies {
		mul += 10
	}
	if c.DarkCholocateCoatedCookies {
		mul += 15
	}
	if c.WhiteCholocateCoatedCookies {
		mul += 15
	}
	if c.EclipseCookies {
		mul += 15
	}
	if c.ZebraCookies {
		mul += 15
	}
	if c.Snickerdoodles {
		mul += 15
	}
	if c.Stroopwafels {
		mul += 15
	}
	if c.Macaroons {
		mul += 15
	}
	if c.EmpireBiscuits {
		mul += 15
	}
	if c.BritishTeaBiscuits {
		mul += 15
	}
	if c.ChocolateBritishTeaBiscuits {
		mul += 15
	}
	if c.RoundBritishTeaBiscuits {
		mul += 15
	}
	if c.RoundChocolateBritishTeaBiscuits {
		mul += 15
	}
	if c.RoundBritishTeaBiscuitsWithHeartMotif {
		mul += 15
	}
	if c.RoundChocolateBritishTeaBiscuitsWithHeartMotif {
		mul += 15
	}
	if c.Madeleines {
		mul += 20
	}
	if c.Palmiers {
		mul += 20
	}
	if c.Palets {
		mul += 20
	}
	if c.Sables {
		mul += 20
	}
	if c.Caramoas {
		mul += 25
	}
	if c.Sagalongs {
		mul += 25
	}
	if c.Shortfoils {
		mul += 25
	}
	if c.WinMints {
		mul += 25
	}
	if c.FigGluttons {
		mul += 25
	}
	if c.Loreols {
		mul += 25
	}
	if c.JaffaCakes {
		mul += 25
	}
	if c.GreasesCups {
		mul += 25
	}
	if c.RoseMacarons {
		mul += 30
	}
	if c.LemonMacarons {
		mul += 30
	}
	if c.ChocolateMacarons {
		mul += 30
	}
	if c.PistachioMacarons {
		mul += 30
	}
	if c.HazelnutMacarons {
		mul += 30
	}
	if c.VioletMacarons {
		mul += 30
	}
	if c.CaramelMacarons {
		mul += 30
	}
	if c.LicoriceMacarons {
		mul += 30
	}
	if c.GingerbreadMen {
		mul += 25
	}
	if c.GingerbreadTrees {
		mul += 25
	}
	if c.SkullCookies {
		mul += 20
	}
	if c.SlimeCookies {
		mul += 20
	}
	if c.EyeballCookies {
		mul += 20
	}
	if c.SpiderCookies {
		mul += 20
	}
	if c.GhostCookies {
		mul += 20
	}
	if c.BatCookies {
		mul += 20
	}
	if c.PumpkinCookies {
		mul += 20
	}
	return
}

func (c *Cookieverse) HeavenlyMultiplier() Production {
	var mul Production
	if c.HeavenlyKey {
		mul = 1
	} else if c.HeavenlyConfectionery {
		mul = .75
	} else if c.HeavenlyBakery {
		mul = .5
	} else if c.HeavenlyCookieStand {
		mul = .25
	} else if c.HeavenlyChipSecret {
		mul = .05
	}
	return 2 * Production(c.HeavenlyChips) * mul
}

func (c *Cookieverse) MultiplierCpS() Production {
	mul := c.NormalMultiplier() + c.HeavenlyMultiplier()
	return c.BuildingCpS() * (1 + mul/100)
}

func (c *Cookieverse) KittenCpS() Production {
	cps := c.MultiplierCpS()
	if c.KittenHelpers {
		cps *= (1 + .0005*Production(c.Milk))
	}
	if c.KittenWorkers {
		cps *= (1 + .0010*Production(c.Milk))
	}
	if c.KittenEngineers {
		cps *= (1 + .0020*Production(c.Milk))
	}
	if c.KittenOverseers {
		cps *= (1 + .0020*Production(c.Milk))
	}
	if c.KittenManagers {
		cps *= (1 + .0020*Production(c.Milk))
	}
	return cps
}

func (c *Cookieverse) EggsCps() Production {
	var mul Production = 0
	if c.ChickenEgg {
		mul += 1
	}
	if c.DuckEgg {
		mul += 1
	}
	if c.TurkeyEgg {
		mul += 1
	}
	if c.QuailEgg {
		mul += 1
	}
	if c.RobinEgg {
		mul += 1
	}
	if c.OstrichEgg {
		mul += 1
	}
	if c.CassowaryEgg {
		mul += 1
	}
	if c.SalmonRoe {
		mul += 1
	}
	if c.Frogspawn {
		mul += 1
	}
	if c.SharkEgg {
		mul += 1
	}
	if c.TurtleEgg {
		mul += 1
	}
	if c.AntLarva {
		mul += 1
	}
	if c.GoldenGooseEgg {
		// nothing
	}
	if c.FabergeEgg {
		// nothing
	}
	if c.Wrinklerspawn {
		// nothing
	}
	if c.CookieEgg {
		// nothing
	}
	if c.Omelette {
		// nothing
	}
	if c.ChocolateEgg {
		// nothing
	}
	if c.CenturyEgg {
		f := c.CenturyDays / 100
		if f > 1 {
			f = 1
		}
		mul += 10 * (1 - (1-f)*(1-f)*(1-f))
	}
	if c.Egg {
		// TODO +9 CpS
	}
	return c.KittenCpS() * (1 + Production(mul)/100)
}

func (c *Cookieverse) Production() Production {
	cps := c.EggsCps()
	if cps < 0.1 {
		cps = 0.1
	}
	return cps
}

func (c *Cookieverse) buildingCost(baseCost Stock, owned int32) Stock {
	cost := baseCost * Stock(math.Pow(1.15, float64(owned)))
	if c.FabergeEgg {
		cost *= 0.99
	}
	return cost
}

func (c *Cookieverse) Updates() (updates []Update) {
	{
		cc := *c
		cc.Cursors++
		u := Update{&cc, "cu", UpdateDescription(fmt.Sprintf("%d Cursors", cc.Cursors)), c.buildingCost(15, c.Cursors)}
		updates = append(updates, u)
	}
	if c.Cursors > 0 {
		cc := *c
		cc.Grandmas++
		u := Update{&cc, "gr", UpdateDescription(fmt.Sprintf("%d Grandmas", cc.Grandmas)), c.buildingCost(100, c.Grandmas)}
		updates = append(updates, u)
	}
	if c.Grandmas > 0 {
		cc := *c
		cc.Farms++
		u := Update{&cc, "fr", UpdateDescription(fmt.Sprintf("%d Farms", cc.Farms)), c.buildingCost(500, c.Farms)}
		updates = append(updates, u)
	}
	if c.Farms > 0 {
		cc := *c
		cc.Factories++
		u := Update{&cc, "fc", UpdateDescription(fmt.Sprintf("%d Factories", cc.Factories)), c.buildingCost(3e3, c.Factories)}
		updates = append(updates, u)
	}
	if c.Factories > 0 {
		cc := *c
		cc.Mines++
		u := Update{&cc, "mi", UpdateDescription(fmt.Sprintf("%d Mines", cc.Mines)), c.buildingCost(1e4, c.Mines)}
		updates = append(updates, u)
	}
	if c.Mines > 0 {
		cc := *c
		cc.Shipments++
		u := Update{&cc, "sh", UpdateDescription(fmt.Sprintf("%d Shipments", cc.Shipments)), c.buildingCost(4e4, c.Shipments)}
		updates = append(updates, u)
	}
	if c.Shipments > 0 {
		cc := *c
		cc.AlchemyLabs++
		u := Update{&cc, "al", UpdateDescription(fmt.Sprintf("%d AlchemyLabs", cc.AlchemyLabs)), c.buildingCost(2e5, c.AlchemyLabs)}
		updates = append(updates, u)
	}
	if c.AlchemyLabs > 0 {
		cc := *c
		cc.Portals++
		u := Update{&cc, "po", UpdateDescription(fmt.Sprintf("%d Portals", cc.Portals)), c.buildingCost(1666666, c.Portals)}
		updates = append(updates, u)
	}
	if c.Portals > 0 {
		cc := *c
		cc.TimeMachines++
		u := Update{&cc, "tm", UpdateDescription(fmt.Sprintf("%d TimeMachines", cc.TimeMachines)), c.buildingCost(123456789, c.TimeMachines)}
		updates = append(updates, u)
	}
	if c.TimeMachines > 0 {
		cc := *c
		cc.AntimatterCondensers++
		u := Update{&cc, "ac", UpdateDescription(fmt.Sprintf("%d AntimatterCondensers", cc.AntimatterCondensers)), c.buildingCost(4e9-1, c.AntimatterCondensers)}
		updates = append(updates, u)
	}
	if c.AntimatterCondensers > 0 {
		cc := *c
		cc.Prisms++
		u := Update{&cc, "pr", UpdateDescription(fmt.Sprintf("%d Prisms", cc.Prisms)), c.buildingCost(75e9, c.Prisms)}
		updates = append(updates, u)
	}
	if !c.ReinforcedIndexFinger && c.Cursors >= 1 {
		cc := *c
		cc.ReinforcedIndexFinger = true
		u := Update{&cc, "ReinforcedIndexFinger", "ReinforcedIndexFinger", 100}
		updates = append(updates, u)
	}
	if !c.CarpalTunnelPreventionCream && c.Cursors >= 1 {
		cc := *c
		cc.CarpalTunnelPreventionCream = true
		u := Update{&cc, "CarpalTunnelPreventionCream", "CarpalTunnelPreventionCream", 400}
		updates = append(updates, u)
	}
	if !c.Ambidextrous && c.Cursors >= 10 && c.CarpalTunnelPreventionCream {
		cc := *c
		cc.Ambidextrous = true
		u := Update{&cc, "Ambidextrous", "Ambidextrous", 1e4}
		updates = append(updates, u)
	}
	if !c.ThousandFingers && c.Cursors >= 20 {
		cc := *c
		cc.ThousandFingers = true
		u := Update{&cc, "ThousandFingers", "ThousandFingers", 5e5}
		updates = append(updates, u)
	}
	if !c.MillionFingers && c.Cursors >= 40 {
		cc := *c
		cc.MillionFingers = true
		u := Update{&cc, "MillionFingers", "MillionFingers", 5e7}
		updates = append(updates, u)
	}
	if !c.BillionFingers && c.Cursors >= 80 {
		cc := *c
		cc.BillionFingers = true
		u := Update{&cc, "BillionFingers", "BillionFingers", 5e8}
		updates = append(updates, u)
	}
	if !c.TrillionFingers && c.Cursors >= 120 {
		cc := *c
		cc.TrillionFingers = true
		u := Update{&cc, "TrillionFingers", "TrillionFingers", 5e9}
		updates = append(updates, u)
	}
	if !c.QuadrillionFingers && c.Cursors >= 160 {
		cc := *c
		cc.QuadrillionFingers = true
		u := Update{&cc, "QuadrillionFingers", "QuadrillionFingers", 5e10}
		updates = append(updates, u)
	}
	if !c.QuintillionFingers && c.Cursors >= 200 {
		cc := *c
		cc.QuintillionFingers = true
		u := Update{&cc, "QuintillionFingers", "QuintillionFingers", 5e13}
		updates = append(updates, u)
	}
	if !c.SextillionFingers && c.Cursors >= 240 {
		cc := *c
		cc.SextillionFingers = true
		u := Update{&cc, "SextillionFingers", "SextillionFingers", 5e14}
		updates = append(updates, u)
	}
	if !c.SeptillionFingers && c.Cursors >= 280 {
		cc := *c
		cc.SeptillionFingers = true
		u := Update{&cc, "SeptillionFingers", "SeptillionFingers", 5e15}
		updates = append(updates, u)
	}
	if !c.OctillionFingers && c.Cursors >= 320 {
		cc := *c
		cc.OctillionFingers = true
		u := Update{&cc, "OctillionFingers", "OctillionFingers", 50e15}
		updates = append(updates, u)
	}
	if !c.ForwardsFromGrandma && c.Grandmas >= 1 {
		cc := *c
		cc.ForwardsFromGrandma = true
		u := Update{&cc, "ForwardsFromGrandma", "ForwardsFromGrandma", 1e3}
		updates = append(updates, u)
	}
	if !c.SteelPlatedRollingPins && c.Grandmas >= 1 {
		cc := *c
		cc.SteelPlatedRollingPins = true
		u := Update{&cc, "SteelPlatedRollingPins", "SteelPlatedRollingPins", 1e4}
		updates = append(updates, u)
	}
	if !c.LubricatedDentures && c.Grandmas >= 10 {
		cc := *c
		cc.LubricatedDentures = true
		u := Update{&cc, "LubricatedDentures", "LubricatedDentures", 1e5}
		updates = append(updates, u)
	}
	if !c.PruneJuice && c.Grandmas >= 50 {
		cc := *c
		cc.PruneJuice = true
		u := Update{&cc, "PruneJuice", "PruneJuice", 5e6}
		updates = append(updates, u)
	}
	if !c.DoubleThickGlasses && c.Grandmas >= 100 {
		cc := *c
		cc.DoubleThickGlasses = true
		u := Update{&cc, "DoubleThickGlasses", "DoubleThickGlasses", 100e6}
		updates = append(updates, u)
	}
	if !c.AgingAgents && c.Grandmas >= 200 {
		cc := *c
		cc.AgingAgents = true
		u := Update{&cc, "AgingAgents", "AgingAgents", 800e9}
		updates = append(updates, u)
	}
	if !c.FarmerGrandmas && c.Grandmas >= 15 {
		cc := *c
		cc.FarmerGrandmas = true
		u := Update{&cc, "FarmerGrandmas", "FarmerGrandmas", 5e4}
		updates = append(updates, u)
	}
	if !c.WorkerGrandmas && c.Factories >= 15 {
		cc := *c
		cc.WorkerGrandmas = true
		u := Update{&cc, "WorkerGrandmas", "WorkerGrandmas", 3e5}
		updates = append(updates, u)
	}
	if !c.MinerGrandmas && c.Mines >= 15 {
		cc := *c
		cc.MinerGrandmas = true
		u := Update{&cc, "MinerGrandmas", "MinerGrandmas", 1e6}
		updates = append(updates, u)
	}
	if !c.CosmicGrandmas && c.Shipments >= 15 {
		cc := *c
		cc.CosmicGrandmas = true
		u := Update{&cc, "CosmicGrandmas", "CosmicGrandmas", 4e6}
		updates = append(updates, u)
	}
	if !c.TransmutedGrandmas && c.AlchemyLabs >= 15 {
		cc := *c
		cc.TransmutedGrandmas = true
		u := Update{&cc, "TransmutedGrandmas", "TransmutedGrandmas", 2e7}
		updates = append(updates, u)
	}
	if !c.AlteredGrandmas && c.Portals >= 15 {
		cc := *c
		cc.AlteredGrandmas = true
		u := Update{&cc, "AlteredGrandmas", "AlteredGrandmas", 166666600}
		updates = append(updates, u)
	}
	if !c.GrandmasGrandmas && c.TimeMachines >= 15 {
		cc := *c
		cc.GrandmasGrandmas = true
		u := Update{&cc, "GrandmasGrandmas", "GrandmasGrandmas", 12345678900}
		updates = append(updates, u)
	}
	if !c.AntiGrandmas && c.AntimatterCondensers >= 15 {
		cc := *c
		cc.AntiGrandmas = true
		u := Update{&cc, "AntiGrandmas", "AntiGrandmas", 399999999000}
		updates = append(updates, u)
	}
	if !c.RainbowGrandmas && c.Prisms >= 15 {
		cc := *c
		cc.RainbowGrandmas = true
		u := Update{&cc, "RainbowGrandmas", "RainbowGrandmas", 7500e12}
		updates = append(updates, u)
	}
	if !c.CheapHoes && c.Farms >= 1 {
		cc := *c
		cc.CheapHoes = true
		u := Update{&cc, "CheapHoes", "CheapHoes", 5e3}
		updates = append(updates, u)
	}
	if !c.Fertilizer && c.Farms >= 1 {
		cc := *c
		cc.Fertilizer = true
		u := Update{&cc, "Fertilizer", "Fertilizer", 5e4}
		updates = append(updates, u)
	}
	if !c.CookieTrees && c.Farms >= 10 {
		cc := *c
		cc.CookieTrees = true
		u := Update{&cc, "CookieTrees", "CookieTrees", 5e5}
		updates = append(updates, u)
	}
	if !c.GeneticallyModifiedCookies && c.Farms >= 50 {
		cc := *c
		cc.GeneticallyModifiedCookies = true
		u := Update{&cc, "GeneticallyModifiedCookies", "GeneticallyModifiedCookies", 5e6}
		updates = append(updates, u)
	}
	if !c.GingerbreadScarecrows && c.Farms >= 100 {
		cc := *c
		cc.GingerbreadScarecrows = true
		u := Update{&cc, "GingerbreadScarecrows", "GingerbreadScarecrows", 25e7}
		updates = append(updates, u)
	}
	if !c.PulsarSprinklers && c.Farms >= 200 {
		cc := *c
		cc.PulsarSprinklers = true
		u := Update{&cc, "PulsarSprinklers", "PulsarSprinklers", 4e12}
		updates = append(updates, u)
	}
	if !c.SturdierConveyorBelts && c.Factories >= 1 {
		cc := *c
		cc.SturdierConveyorBelts = true
		u := Update{&cc, "SturdierConveyorBelts", "SturdierConveyorBelts", 3e4}
		updates = append(updates, u)
	}
	if !c.ChildLabor && c.Factories >= 1 {
		cc := *c
		cc.ChildLabor = true
		u := Update{&cc, "ChildLabor", "ChildLabor", 3e5}
		updates = append(updates, u)
	}
	if !c.SweatShop && c.Factories >= 10 {
		cc := *c
		cc.SweatShop = true
		u := Update{&cc, "SweatShop", "SweatShop", 3e6}
		updates = append(updates, u)
	}
	if !c.RadiumReactors && c.Factories >= 50 {
		cc := *c
		cc.RadiumReactors = true
		u := Update{&cc, "RadiumReactors", "RadiumReactors", 150e6}
		updates = append(updates, u)
	}
	if !c.Recombobulators && c.Factories >= 100 {
		cc := *c
		cc.Recombobulators = true
		u := Update{&cc, "Recombobulators", "Recombobulators", 3e9}
		updates = append(updates, u)
	}
	if !c.DeepBakeProcess && c.Factories >= 200 {
		cc := *c
		cc.DeepBakeProcess = true
		u := Update{&cc, "DeepBakeProcess", "DeepBakeProcess", 24e12}
		updates = append(updates, u)
	}
	if !c.SugarGas && c.Mines >= 1 {
		cc := *c
		cc.SugarGas = true
		u := Update{&cc, "SugarGas", "SugarGas", 1e5}
		updates = append(updates, u)
	}
	if !c.Megadrill && c.Mines >= 1 {
		cc := *c
		cc.Megadrill = true
		u := Update{&cc, "Megadrill", "Megadrill", 1e6}
		updates = append(updates, u)
	}
	if !c.Ultradrill && c.Mines >= 10 {
		cc := *c
		cc.Ultradrill = true
		u := Update{&cc, "Ultradrill", "Ultradrill", 1e7}
		updates = append(updates, u)
	}
	if !c.Ultimadrill && c.Mines >= 50 {
		cc := *c
		cc.Ultimadrill = true
		u := Update{&cc, "Ultimadrill", "Ultimadrill", 1e8}
		updates = append(updates, u)
	}
	if !c.HBombMining && c.Mines >= 100 {
		cc := *c
		cc.HBombMining = true
		u := Update{&cc, "HBombMining", "HBombMining", 5e9}
		updates = append(updates, u)
	}
	if !c.Coreforge && c.Mines >= 200 {
		cc := *c
		cc.Coreforge = true
		u := Update{&cc, "Coreforge", "Coreforge", 80e12}
		updates = append(updates, u)
	}
	if !c.VanillaNebulae && c.Shipments >= 1 {
		cc := *c
		cc.VanillaNebulae = true
		u := Update{&cc, "VanillaNebulae", "VanillaNebulae", 4e5}
		updates = append(updates, u)
	}
	if !c.Wormholes && c.Shipments >= 1 {
		cc := *c
		cc.Wormholes = true
		u := Update{&cc, "Wormholes", "Wormholes", 4e6}
		updates = append(updates, u)
	}
	if !c.FrequentFlyer && c.Shipments >= 10 {
		cc := *c
		cc.FrequentFlyer = true
		u := Update{&cc, "FrequentFlyer", "FrequentFlyer", 40e6}
		updates = append(updates, u)
	}
	if !c.WarpDrive && c.Shipments >= 50 {
		cc := *c
		cc.WarpDrive = true
		u := Update{&cc, "WarpDrive", "WarpDrive", 2e9}
		updates = append(updates, u)
	}
	if !c.ChocolateMonoliths && c.Shipments >= 100 {
		cc := *c
		cc.ChocolateMonoliths = true
		u := Update{&cc, "ChocolateMonoliths", "ChocolateMonoliths", 40e9}
		updates = append(updates, u)
	}
	if !c.GenerationShip && c.Shipments >= 200 {
		cc := *c
		cc.GenerationShip = true
		u := Update{&cc, "GenerationShip", "GenerationShip", 32e12}
		updates = append(updates, u)
	}
	if !c.Antimony && c.AlchemyLabs >= 1 {
		cc := *c
		cc.Antimony = true
		u := Update{&cc, "Antimony", "Antimony", 2e6}
		updates = append(updates, u)
	}
	if !c.EssenceOfDough && c.AlchemyLabs >= 1 {
		cc := *c
		cc.EssenceOfDough = true
		u := Update{&cc, "EssenceOfDough", "EssenceOfDough", 2e7}
		updates = append(updates, u)
	}
	if !c.TrueChocolate && c.AlchemyLabs >= 10 {
		cc := *c
		cc.TrueChocolate = true
		u := Update{&cc, "TrueChocolate", "TrueChocolate", 2e8}
		updates = append(updates, u)
	}
	if !c.Ambrosia && c.AlchemyLabs >= 50 {
		cc := *c
		cc.Ambrosia = true
		u := Update{&cc, "Ambrosia", "Ambrosia", 10e9}
		updates = append(updates, u)
	}
	if !c.AcquaCrustulae && c.AlchemyLabs >= 100 {
		cc := *c
		cc.AcquaCrustulae = true
		u := Update{&cc, "AcquaCrustulae", "AcquaCrustulae", 200e9}
		updates = append(updates, u)
	}
	if !c.OriginCrucible && c.AlchemyLabs >= 200 {
		cc := *c
		cc.OriginCrucible = true
		u := Update{&cc, "OriginCrucible", "OriginCrucible", 1600e12}
		updates = append(updates, u)
	}
	if !c.AncientTablet && c.Portals >= 1 {
		cc := *c
		cc.AncientTablet = true
		u := Update{&cc, "AncientTablet", "AncientTablet", 16666660}
		updates = append(updates, u)
	}
	if !c.InsaneOatlingWorkers && c.Portals >= 1 {
		cc := *c
		cc.InsaneOatlingWorkers = true
		u := Update{&cc, "InsaneOatlingWorkers", "InsaneOatlingWorkers", 166666600}
		updates = append(updates, u)
	}
	if !c.SouldBond && c.Portals >= 10 {
		cc := *c
		cc.SouldBond = true
		u := Update{&cc, "SouldBond", "SouldBond", 1666666000}
		updates = append(updates, u)
	}
	if !c.SanityDance && c.Portals >= 50 {
		cc := *c
		cc.SanityDance = true
		u := Update{&cc, "SanityDance", "SanityDance", 83333300000}
		updates = append(updates, u)
	}
	if !c.BraneTransplant && c.Portals >= 100 {
		cc := *c
		cc.BraneTransplant = true
		u := Update{&cc, "BraneTransplant", "BraneTransplant", 1666666e6}
		updates = append(updates, u)
	}
	if !c.DeitySizedPortals && c.Portals >= 200 {
		cc := *c
		cc.DeitySizedPortals = true
		u := Update{&cc, "DeitySizedPortals", "DeitySizedPortals", 13333328e9}
		updates = append(updates, u)
	}
	if !c.FluxCapacitors && c.TimeMachines >= 1 {
		cc := *c
		cc.FluxCapacitors = true
		u := Update{&cc, "FluxCapacitors", "FluxCapacitors", 1234567890}
		updates = append(updates, u)
	}
	if !c.TimeParadoxResolver && c.TimeMachines >= 1 {
		cc := *c
		cc.TimeParadoxResolver = true
		u := Update{&cc, "TimeParadoxResolver", "TimeParadoxResolver", 9876543210}
		updates = append(updates, u)
	}
	if !c.QuantumConundrum && c.TimeMachines >= 10 {
		cc := *c
		cc.QuantumConundrum = true
		u := Update{&cc, "QuantumConundrum", "QuantumConundrum", 98765456789}
		updates = append(updates, u)
	}
	if !c.CausalityEnforcer && c.TimeMachines >= 50 {
		cc := *c
		cc.CausalityEnforcer = true
		u := Update{&cc, "CausalityEnforcer", "CausalityEnforcer", 123456789e4}
		updates = append(updates, u)
	}
	if !c.YestermorrowComparators && c.TimeMachines >= 100 {
		cc := *c
		cc.YestermorrowComparators = true
		u := Update{&cc, "YestermorrowComparators", "YestermorrowComparators", 123456789e6}
		updates = append(updates, u)
	}
	if !c.FarFutureEnactment && c.TimeMachines >= 200 {
		cc := *c
		cc.FarFutureEnactment = true
		u := Update{&cc, "FarFutureEnactment", "FarFutureEnactment", 987654321e12}
		updates = append(updates, u)
	}
	if !c.SugarBosons && c.AntimatterCondensers >= 1 {
		cc := *c
		cc.SugarBosons = true
		u := Update{&cc, "SugarBosons", "SugarBosons", 39999999990}
		updates = append(updates, u)
	}
	if !c.StringTheory && c.AntimatterCondensers >= 1 {
		cc := *c
		cc.StringTheory = true
		u := Update{&cc, "StringTheory", "StringTheory", 399999999900}
		updates = append(updates, u)
	}
	if !c.LargeMacaronCollider && c.AntimatterCondensers >= 10 {
		cc := *c
		cc.LargeMacaronCollider = true
		u := Update{&cc, "LargeMacaronCollider", "LargeMacaronCollider", 3999999999e3}
		updates = append(updates, u)
	}
	if !c.BigBangCake && c.AntimatterCondensers >= 50 {
		cc := *c
		cc.BigBangCake = true
		u := Update{&cc, "BigBangCake", "BigBangCake", 199999999950e3}
		updates = append(updates, u)
	}
	if !c.ReverseCyclotrons && c.AntimatterCondensers >= 100 {
		cc := *c
		cc.ReverseCyclotrons = true
		u := Update{&cc, "ReverseCyclotrons", "ReverseCyclotrons", 3999999999e6}
		updates = append(updates, u)
	}
	if !c.Nanocosmics && c.AntimatterCondensers >= 200 {
		cc := *c
		cc.Nanocosmics = true
		u := Update{&cc, "Nanocosmics", "Nanocosmics", 31999999992e12}
		updates = append(updates, u)
	}
	if !c.GemPolish && c.Prisms >= 1 {
		cc := *c
		cc.GemPolish = true
		u := Update{&cc, "GemPolish", "GemPolish", 750e9}
		updates = append(updates, u)
	}
	if !c.NinthColor && c.Prisms >= 1 {
		cc := *c
		cc.NinthColor = true
		u := Update{&cc, "NinthColor", "NinthColor", 7500e9}
		updates = append(updates, u)
	}
	if !c.ChocolateLight && c.Prisms >= 10 {
		cc := *c
		cc.ChocolateLight = true
		u := Update{&cc, "ChocolateLight", "ChocolateLight", 75e12}
		updates = append(updates, u)
	}
	if !c.Grainbow && c.Prisms >= 50 {
		cc := *c
		cc.Grainbow = true
		u := Update{&cc, "Grainbow", "Grainbow", 3750e12}
		updates = append(updates, u)
	}
	if !c.PureCosmicLight && c.Prisms >= 100 {
		cc := *c
		cc.PureCosmicLight = true
		u := Update{&cc, "PureCosmicLight", "PureCosmicLight", 75e15}
		updates = append(updates, u)
	}
	if !c.GlowInTheDark && c.Prisms >= 200 {
		cc := *c
		cc.GlowInTheDark = true
		u := Update{&cc, "GlowInTheDark", "GlowInTheDark", 600e18}
		updates = append(updates, u)
	}
	if !c.BingoCenterResearchFacility {
		cc := *c
		cc.BingoCenterResearchFacility = true
		u := Update{&cc, "BingoCenterResearchFacility", "BingoCenterResearchFacility", 1e11}
		updates = append(updates, u)
	}
	if !c.SpecializedChocolateChips && c.BingoCenterResearchFacility && !c.DisableResearch {
		cc := *c
		cc.SpecializedChocolateChips = true
		u := Update{&cc, "SpecializedChocolateChips", "SpecializedChocolateChips", 1e10}
		updates = append(updates, u)
	}
	if !c.DesignerCocoaBeans && c.SpecializedChocolateChips && !c.DisableResearch {
		cc := *c
		cc.DesignerCocoaBeans = true
		u := Update{&cc, "DesignerCocoaBeans", "DesignerCocoaBeans", 2e10}
		updates = append(updates, u)
	}
	if !c.RitualRollingPins && c.DesignerCocoaBeans && !c.DisableResearch {
		cc := *c
		cc.RitualRollingPins = true
		u := Update{&cc, "RitualRollingPins", "RitualRollingPins", 4e10}
		updates = append(updates, u)
	}
	if !c.UnderworldOvens && c.RitualRollingPins && !c.DisableResearch {
		cc := *c
		cc.UnderworldOvens = true
		u := Update{&cc, "UnderworldOvens", "UnderworldOvens", 8e10}
		updates = append(updates, u)
	}
	if !c.OneMind && c.UnderworldOvens && !c.DisableResearch {
		cc := *c
		cc.OneMind = true
		u := Update{&cc, "OneMind", "OneMind", 16e10}
		updates = append(updates, u)
	}
	if !c.ExoticNuts && c.OneMind && !c.DisableResearch {
		cc := *c
		cc.ExoticNuts = true
		u := Update{&cc, "ExoticNuts", "ExoticNuts", 32e10}
		updates = append(updates, u)
	}
	if !c.CommunalBrainsweep && c.ExoticNuts && !c.DisableResearch {
		cc := *c
		cc.CommunalBrainsweep = true
		u := Update{&cc, "CommunalBrainsweep", "CommunalBrainsweep", 64e10}
		updates = append(updates, u)
	}
	if !c.ArcaneSugar && c.CommunalBrainsweep && !c.DisableResearch {
		cc := *c
		cc.ArcaneSugar = true
		u := Update{&cc, "ArcaneSugar", "ArcaneSugar", 128e10}
		updates = append(updates, u)
	}
	if !c.ElderPact && c.ArcaneSugar && !c.DisableResearch {
		cc := *c
		cc.ElderPact = true
		u := Update{&cc, "ElderPact", "ElderPact", 256e10}
		updates = append(updates, u)
	}
	if !c.HeavenlyChipSecret {
		cc := *c
		cc.HeavenlyChipSecret = true
		u := Update{&cc, "HeavenlyChipSecret", "HeavenlyChipSecret", 11}
		updates = append(updates, u)
	}
	if !c.HeavenlyCookieStand && c.HeavenlyChipSecret {
		cc := *c
		cc.HeavenlyCookieStand = true
		u := Update{&cc, "HeavenlyCookieStand", "HeavenlyCookieStand", 1111}
		updates = append(updates, u)
	}
	if !c.HeavenlyBakery && c.HeavenlyCookieStand {
		cc := *c
		cc.HeavenlyBakery = true
		u := Update{&cc, "HeavenlyBakery", "HeavenlyBakery", 111111}
		updates = append(updates, u)
	}
	if !c.HeavenlyConfectionery && c.HeavenlyBakery {
		cc := *c
		cc.HeavenlyConfectionery = true
		u := Update{&cc, "HeavenlyConfectionery", "HeavenlyConfectionery", 11111111}
		updates = append(updates, u)
	}
	if !c.HeavenlyKey && c.HeavenlyConfectionery {
		cc := *c
		cc.HeavenlyKey = true
		u := Update{&cc, "HeavenlyKey", "HeavenlyKey", 1111111111}
		updates = append(updates, u)
	}
	if !c.KittenHelpers && c.Milk >= 52 {
		cc := *c
		cc.KittenHelpers = true
		u := Update{&cc, "KittenHelpers", "KittenHelpers", 9e6}
		updates = append(updates, u)
	}
	if !c.KittenWorkers && c.Milk >= 100 {
		cc := *c
		cc.KittenWorkers = true
		u := Update{&cc, "KittenWorkers", "KittenWorkers", 9e9}
		updates = append(updates, u)
	}
	if !c.KittenEngineers && c.Milk >= 200 {
		cc := *c
		cc.KittenEngineers = true
		u := Update{&cc, "KittenEngineers", "KittenEngineers", 90e12}
		updates = append(updates, u)
	}
	if !c.KittenOverseers && c.Milk >= 300 {
		cc := *c
		cc.KittenOverseers = true
		u := Update{&cc, "KittenOverseers", "KittenOverseers", 90e15}
		updates = append(updates, u)
	}
	if !c.KittenManagers && c.Milk >= 400 {
		cc := *c
		cc.KittenManagers = true
		u := Update{&cc, "KittenManagers", "KittenManagers", 900e18}
		updates = append(updates, u)
	}
	if !c.OutmealRaisinCookies {
		cc := *c
		cc.OutmealRaisinCookies = true
		u := Update{&cc, "OutmealRaisinCookies", "OutmealRaisinCookies", 1e8 - 1}
		updates = append(updates, u)
	}
	if !c.PeanutButterCookies && c.OutmealRaisinCookies {
		cc := *c
		cc.PeanutButterCookies = true
		u := Update{&cc, "PeanutButterCookies", "PeanutButterCookies", 1e8 - 1}
		updates = append(updates, u)
	}
	if !c.PlainCookies && c.PeanutButterCookies {
		cc := *c
		cc.PlainCookies = true
		u := Update{&cc, "PlainCookies", "PlainCookies", 1e8 - 1}
		updates = append(updates, u)
	}
	if !c.CoconutCookies {
		cc := *c
		cc.CoconutCookies = true
		u := Update{&cc, "CoconutCookies", "CoconutCookies", 1e9 - 1}
		updates = append(updates, u)
	}
	if !c.WhiteChocolateCookies && c.CoconutCookies {
		cc := *c
		cc.WhiteChocolateCookies = true
		u := Update{&cc, "WhiteChocolateCookies", "WhiteChocolateCookies", 1e9 - 1}
		updates = append(updates, u)
	}
	if !c.MacadamiaNutCookies && c.WhiteChocolateCookies {
		cc := *c
		cc.MacadamiaNutCookies = true
		u := Update{&cc, "MacadamiaNutCookies", "MacadamiaNutCookies", 1e9 - 1}
		updates = append(updates, u)
	}
	if !c.SugarCookies {
		cc := *c
		cc.SugarCookies = true
		u := Update{&cc, "SugarCookies", "SugarCookies", 1e8 - 11}
		updates = append(updates, u)
	}
	if !c.DoubleChipCookies {
		cc := *c
		cc.DoubleChipCookies = true
		u := Update{&cc, "DoubleChipCookies", "DoubleChipCookies", 1e11 - 1}
		updates = append(updates, u)
	}
	if !c.WhiteChocolateMacadamiaNutCookies && c.DoubleChipCookies {
		cc := *c
		cc.WhiteChocolateMacadamiaNutCookies = true
		u := Update{&cc, "WhiteChocolateMacadamiaNutCookies", "WhiteChocolateMacadamiaNutCookies", 1e11 - 1}
		updates = append(updates, u)
	}
	if !c.AllChocolateCookies && c.WhiteChocolateMacadamiaNutCookies {
		cc := *c
		cc.AllChocolateCookies = true
		u := Update{&cc, "AllChocolateCookies", "AllChocolateCookies", 1e11 - 1}
		updates = append(updates, u)
	}
	if !c.DarkCholocateCoatedCookies {
		cc := *c
		cc.DarkCholocateCoatedCookies = true
		u := Update{&cc, "DarkCholocateCoatedCookies", "DarkCholocateCoatedCookies", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.WhiteCholocateCoatedCookies && c.DarkCholocateCoatedCookies {
		cc := *c
		cc.WhiteCholocateCoatedCookies = true
		u := Update{&cc, "WhiteCholocateCoatedCookies", "WhiteCholocateCoatedCookies", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.EclipseCookies {
		cc := *c
		cc.EclipseCookies = true
		u := Update{&cc, "EclipseCookies", "EclipseCookies", 1e13 - 1}
		updates = append(updates, u)
	}
	if !c.ZebraCookies && c.EclipseCookies {
		cc := *c
		cc.ZebraCookies = true
		u := Update{&cc, "ZebraCookies", "ZebraCookies", 1e13 - 1}
		updates = append(updates, u)
	}
	if !c.Snickerdoodles {
		cc := *c
		cc.Snickerdoodles = true
		u := Update{&cc, "Snickerdoodles", "Snickerdoodles", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.Stroopwafels && c.Snickerdoodles {
		cc := *c
		cc.Stroopwafels = true
		u := Update{&cc, "Stroopwafels", "Stroopwafels", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.Macaroons && c.Stroopwafels {
		cc := *c
		cc.Macaroons = true
		u := Update{&cc, "Macaroons", "Macaroons", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.EmpireBiscuits && c.Macaroons {
		cc := *c
		cc.EmpireBiscuits = true
		u := Update{&cc, "EmpireBiscuits", "EmpireBiscuits", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.BritishTeaBiscuits && c.EmpireBiscuits {
		cc := *c
		cc.BritishTeaBiscuits = true
		u := Update{&cc, "BritishTeaBiscuits", "BritishTeaBiscuits", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.ChocolateBritishTeaBiscuits && c.BritishTeaBiscuits {
		cc := *c
		cc.ChocolateBritishTeaBiscuits = true
		u := Update{&cc, "ChocolateBritishTeaBiscuits", "ChocolateBritishTeaBiscuits", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.RoundBritishTeaBiscuits && c.ChocolateBritishTeaBiscuits {
		cc := *c
		cc.RoundBritishTeaBiscuits = true
		u := Update{&cc, "RoundBritishTeaBiscuits", "RoundBritishTeaBiscuits", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.RoundChocolateBritishTeaBiscuits && c.RoundBritishTeaBiscuits {
		cc := *c
		cc.RoundChocolateBritishTeaBiscuits = true
		u := Update{&cc, "RoundChocolateBritishTeaBiscuits", "RoundChocolateBritishTeaBiscuits", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.RoundBritishTeaBiscuitsWithHeartMotif && c.RoundChocolateBritishTeaBiscuits {
		cc := *c
		cc.RoundBritishTeaBiscuitsWithHeartMotif = true
		u := Update{&cc, "RoundBritishTeaBiscuitsWithHeartMotif", "RoundBritishTeaBiscuitsWithHeartMotif", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.RoundChocolateBritishTeaBiscuitsWithHeartMotif && c.RoundBritishTeaBiscuitsWithHeartMotif {
		cc := *c
		cc.RoundChocolateBritishTeaBiscuitsWithHeartMotif = true
		u := Update{&cc, "RoundChocolateBritishTeaBiscuitsWithHeartMotif", "RoundChocolateBritishTeaBiscuitsWithHeartMotif", 1e14 - 1}
		updates = append(updates, u)
	}
	if !c.Madeleines {
		cc := *c
		cc.Madeleines = true
		u := Update{&cc, "Madeleines", "Madeleines", 2e14 - 1}
		updates = append(updates, u)
	}
	if !c.Palmiers && c.Madeleines {
		cc := *c
		cc.Palmiers = true
		u := Update{&cc, "Palmiers", "Palmiers", 2e14 - 1}
		updates = append(updates, u)
	}
	if !c.Palets && c.Palmiers {
		cc := *c
		cc.Palets = true
		u := Update{&cc, "Palets", "Palets", 2e14 - 1}
		updates = append(updates, u)
	}
	if !c.Sables && c.Palets {
		cc := *c
		cc.Sables = true
		u := Update{&cc, "Sables", "Sables", 2e14 - 1}
		updates = append(updates, u)
	}
	if !c.Caramoas && c.HeavenlyChips >= 1 {
		cc := *c
		cc.Caramoas = true
		u := Update{&cc, "Caramoas", "Caramoas", 1e15 - 1}
		updates = append(updates, u)
	}
	if !c.Sagalongs && c.Caramoas && c.HeavenlyChips >= 2 {
		cc := *c
		cc.Sagalongs = true
		u := Update{&cc, "Sagalongs", "Sagalongs", 1e15 - 1}
		updates = append(updates, u)
	}
	if !c.Shortfoils && c.Sagalongs && c.HeavenlyChips >= 3 {
		cc := *c
		cc.Shortfoils = true
		u := Update{&cc, "Shortfoils", "Shortfoils", 1e15 - 1}
		updates = append(updates, u)
	}
	if !c.WinMints && c.Shortfoils && c.HeavenlyChips >= 4 {
		cc := *c
		cc.WinMints = true
		u := Update{&cc, "WinMints", "WinMints", 1e15 - 1}
		updates = append(updates, u)
	}
	if !c.FigGluttons && c.WinMints && c.HeavenlyChips >= 10 {
		cc := *c
		cc.FigGluttons = true
		u := Update{&cc, "FigGluttons", "FigGluttons", 1e15 - 1}
		updates = append(updates, u)
	}
	if !c.Loreols && c.FigGluttons && c.HeavenlyChips >= 100 {
		cc := *c
		cc.Loreols = true
		u := Update{&cc, "Loreols", "Loreols", 1e15 - 1}
		updates = append(updates, u)
	}
	if !c.JaffaCakes && c.Loreols && c.HeavenlyChips >= 500 {
		cc := *c
		cc.JaffaCakes = true
		u := Update{&cc, "JaffaCakes", "JaffaCakes", 1e15 - 1}
		updates = append(updates, u)
	}
	if !c.GreasesCups && c.JaffaCakes && c.HeavenlyChips >= 2e3 {
		cc := *c
		cc.GreasesCups = true
		u := Update{&cc, "GreasesCups", "GreasesCups", 1e15 - 1}
		updates = append(updates, u)
	}
	if !c.RoseMacarons && c.GreasesCups && c.HeavenlyChips >= 10e3 {
		cc := *c
		cc.RoseMacarons = true
		u := Update{&cc, "RoseMacarons", "RoseMacarons", 1e18}
		updates = append(updates, u)
	}
	if !c.LemonMacarons && c.RoseMacarons {
		cc := *c
		cc.LemonMacarons = true
		u := Update{&cc, "LemonMacarons", "LemonMacarons", 1e18}
		updates = append(updates, u)
	}
	if !c.ChocolateMacarons && c.LemonMacarons {
		cc := *c
		cc.ChocolateMacarons = true
		u := Update{&cc, "ChocolateMacarons", "ChocolateMacarons", 1e18}
		updates = append(updates, u)
	}
	if !c.PistachioMacarons && c.ChocolateMacarons {
		cc := *c
		cc.PistachioMacarons = true
		u := Update{&cc, "PistachioMacarons", "PistachioMacarons", 1e18}
		updates = append(updates, u)
	}
	if !c.HazelnutMacarons && c.PistachioMacarons {
		cc := *c
		cc.HazelnutMacarons = true
		u := Update{&cc, "HazelnutMacarons", "HazelnutMacarons", 1e18}
		updates = append(updates, u)
	}
	if !c.VioletMacarons && c.HazelnutMacarons {
		cc := *c
		cc.VioletMacarons = true
		u := Update{&cc, "VioletMacarons", "VioletMacarons", 1e18}
		updates = append(updates, u)
	}
	if !c.CaramelMacarons && c.VioletMacarons {
		cc := *c
		cc.CaramelMacarons = true
		u := Update{&cc, "CaramelMacarons", "CaramelMacarons", 1e18}
		updates = append(updates, u)
	}
	if !c.LicoriceMacarons && c.CaramelMacarons {
		cc := *c
		cc.LicoriceMacarons = true
		u := Update{&cc, "LicoriceMacarons", "LicoriceMacarons", 1e18}
		updates = append(updates, u)
	}
	if !c.GingerbreadMen {
		cc := *c
		cc.GingerbreadMen = true
		u := Update{&cc, "GingerbreadMen", "GingerbreadMen", 10e15}
		updates = append(updates, u)
	}
	if !c.GingerbreadTrees && c.GingerbreadMen {
		cc := *c
		cc.GingerbreadTrees = true
		u := Update{&cc, "GingerbreadTrees", "GingerbreadTrees", 10e15}
		updates = append(updates, u)
	}
	if !c.SkullCookies && !c.DisableSkullCookies {
		cc := *c
		cc.SkullCookies = true
		u := Update{&cc, "SkullCookies", "SkullCookies", 444444444444}
		updates = append(updates, u)
	}
	if !c.SlimeCookies && !c.DisableSlimeCookies {
		cc := *c
		cc.SlimeCookies = true
		u := Update{&cc, "SlimeCookies", "SlimeCookies", 444444444444}
		updates = append(updates, u)
	}
	if !c.EyeballCookies && !c.DisableEyeballCookies {
		cc := *c
		cc.EyeballCookies = true
		u := Update{&cc, "EyeballCookies", "EyeballCookies", 444444444444}
		updates = append(updates, u)
	}
	if !c.SpiderCookies && !c.DisableSpiderCookies {
		cc := *c
		cc.SpiderCookies = true
		u := Update{&cc, "SpiderCookies", "SpiderCookies", 444444444444}
		updates = append(updates, u)
	}
	if !c.GhostCookies && !c.DisableGhostCookies {
		cc := *c
		cc.GhostCookies = true
		u := Update{&cc, "GhostCookies", "GhostCookies", 444444444444}
		updates = append(updates, u)
	}
	if !c.BatCookies && !c.DisableBatCookies {
		cc := *c
		cc.BatCookies = true
		u := Update{&cc, "BatCookies", "BatCookies", 444444444444}
		updates = append(updates, u)
	}
	if !c.PumpkinCookies && !c.DisablePumpkinCookies {
		cc := *c
		cc.PumpkinCookies = true
		u := Update{&cc, "PumpkinCookies", "PumpkinCookies", 444444444444}
		updates = append(updates, u)
	}
	if !c.ChickenEgg && !c.DisableChickenEgg {
		cc := *c
		cc.ChickenEgg = true
		u := Update{&cc, "ChickenEgg", "ChickenEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.DuckEgg && !c.DisableDuckEgg {
		cc := *c
		cc.DuckEgg = true
		u := Update{&cc, "DuckEgg", "DuckEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.TurkeyEgg && !c.DisableTurkeyEgg {
		cc := *c
		cc.TurkeyEgg = true
		u := Update{&cc, "TurkeyEgg", "TurkeyEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.QuailEgg && !c.DisableQuailEgg {
		cc := *c
		cc.QuailEgg = true
		u := Update{&cc, "QuailEgg", "QuailEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.RobinEgg && !c.DisableRobinEgg {
		cc := *c
		cc.RobinEgg = true
		u := Update{&cc, "RobinEgg", "RobinEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.OstrichEgg && !c.DisableOstrichEgg {
		cc := *c
		cc.OstrichEgg = true
		u := Update{&cc, "OstrichEgg", "OstrichEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.CassowaryEgg && !c.DisableCassowaryEgg {
		cc := *c
		cc.CassowaryEgg = true
		u := Update{&cc, "CassowaryEgg", "CassowaryEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.SalmonRoe && !c.DisableSalmonRoe {
		cc := *c
		cc.SalmonRoe = true
		u := Update{&cc, "SalmonRoe", "SalmonRoe", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.Frogspawn && !c.DisableFrogspawn {
		cc := *c
		cc.Frogspawn = true
		u := Update{&cc, "Frogspawn", "Frogspawn", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.SharkEgg && !c.DisableSharkEgg {
		cc := *c
		cc.SharkEgg = true
		u := Update{&cc, "SharkEgg", "SharkEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.TurtleEgg && !c.DisableTurtleEgg {
		cc := *c
		cc.TurtleEgg = true
		u := Update{&cc, "TurtleEgg", "TurtleEgg", 1e12 - 1}
		updates = append(updates, u)
	}
	if !c.AntLarva && !c.DisableAntLarva {
		cc := *c
		cc.AntLarva = true
		u := Update{&cc, "AntLarva", "AntLarva", 1e12 - 1}
		updates = append(updates, u)
	}
	return
}
