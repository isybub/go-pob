package data

import "go-pob/utils"

type DamageType string

const (
	DamageTypeCold      = DamageType("Cold")
	DamageTypeLightning = DamageType("Lightning")
	DamageTypeFire      = DamageType("Fire")
)

type Ailment string

const (
	AilmentIgnite  = Ailment("Ignite")
	AilmentChill   = Ailment("Chill")
	AilmentFreeze  = Ailment("Freeze")
	AilmentShock   = Ailment("Shock")
	AilmentScorch  = Ailment("Scorch")
	AilmentBrittle = Ailment("Brittle")
	AilmentSap     = Ailment("Sap")
)

func (Ailment) Values() []Ailment {
	return []Ailment{
		AilmentIgnite,
		AilmentChill,
		AilmentFreeze,
		AilmentShock,
		AilmentScorch,
		AilmentBrittle,
		AilmentSap,
	}
}

type NonDamagingAilmentData struct {
	AssociatedType DamageType
	Alt            bool
	Default        *float64
	Min            float64
	Max            float64
	Precision      float64
	Duration       *float64
}

var NonDamagingAilments = map[Ailment]NonDamagingAilmentData{
	AilmentChill: {
		AssociatedType: DamageTypeCold,
		Alt:            false,
		Default:        utils.Ptr[float64](10),
		Min:            5,
		Max:            30,
		Precision:      0,
		Duration:       utils.Ptr[float64](2),
	},
	AilmentFreeze: {
		AssociatedType: DamageTypeCold,
		Alt:            false,
		Default:        nil,
		Min:            0.3,
		Max:            3,
		Precision:      2,
		Duration:       nil,
	},
	AilmentShock: {
		AssociatedType: DamageTypeLightning,
		Alt:            false,
		Default:        utils.Ptr[float64](15),
		Min:            5,
		Max:            50,
		Precision:      0,
		Duration:       utils.Ptr[float64](2),
	},
	AilmentScorch: {
		AssociatedType: DamageTypeFire,
		Alt:            true,
		Default:        utils.Ptr[float64](10),
		Min:            0,
		Max:            30,
		Precision:      0,
		Duration:       utils.Ptr[float64](4),
	},
	AilmentBrittle: {
		AssociatedType: DamageTypeCold,
		Alt:            true,
		Default:        utils.Ptr[float64](5),
		Min:            0,
		Max:            15,
		Precision:      2,
		Duration:       utils.Ptr[float64](4),
	},
	AilmentSap: {
		AssociatedType: DamageTypeLightning,
		Alt:            true,
		Default:        utils.Ptr[float64](6),
		Min:            0,
		Max:            20,
		Precision:      0,
		Duration:       utils.Ptr[float64](4),
	},
}

const (
	ServerTickTime            = 0.033
	ServerTickRate            = 1 / 0.033
	TemporalChainsEffectCap   = 75
	DamageReductionCap        = 90
	ResistFloor               = -200
	MaxResistCap              = 90
	EvadeChanceCap            = 95
	DodgeChanceCap            = 75
	SuppressionChanceCap      = 100
	SuppressionEffect         = 50
	AvoidChanceCap            = 75
	EnergyShieldRechargeBase  = 0.33
	EnergyShieldRechargeDelay = 2
	WardRechargeDelay         = 5
	Transfiguration           = 0.3
	EnemyMaxResist            = 75
	LeechRateBase             = 0.02
	BleedPercentBase          = 70
	BleedDurationBase         = 5
	PoisonPercentBase         = 0.30
	PoisonDurationBase        = 2
	IgnitePercentBase         = 0.9
	IgniteDurationBase        = 4
	IgniteMinDuration         = 0.3
	ImpaleStoredDamageBase    = 0.1
	BuffExpirationSlowCap     = 0.25
	TrapTriggerRadiusBase     = 10
	MineDetonationRadiusBase  = 60
	MineAuraRadiusBase        = 35
	MaxEnemyLevel             = 85
	LowPoolThreshold          = 0.5
	AccuracyPerDexBase        = 2
	BrandAttachmentRangeBase  = 30
	ProjectileDistanceCap     = 150

	// Expected values to calculate EHP
	stdBossDPSMult      = 4 / 4.25
	pinnacleBossDPSMult = 8 / 4.25
	pinnacleBossPen     = 25 / 5
	uberBossDPSMult     = 10 / 4.25
	uberBossPen         = 40 / 5

	// ehp helper function magic numbers
	ehpCalcSpeedUp = 8

	// depth needs to be a power of speedUp (in this case 8^3, will run 3 recursive calls deep)
	ehpCalcMaxDepth = 512

	// max hits is currently depth + speedup - 1 to give as much accuracy with as few cycles as possible, but can be increased for more accuracy
	ehpCalcMaxHitsToCalc = 519
)

var MonsterEvasionTable = []float64{67, 86, 104, 124, 144, 166, 188, 211, 234, 259, 285, 311, 339, 368, 397, 428, 460, 493, 527, 563, 600, 638, 677, 718, 760, 804, 849, 896, 944, 994, 1046, 1100, 1155, 1212, 1271, 1332, 1395, 1460, 1528, 1597, 1669, 1743, 1819, 1898, 1979, 2063, 2150, 2239, 2331, 2426, 2524, 2626, 2730, 2837, 2948, 3063, 3180, 3302, 3427, 3556, 3689, 3826, 3967, 4112, 4262, 4416, 4575, 4739, 4907, 5081, 5260, 5444, 5633, 5828, 6029, 6235, 6448, 6667, 6892, 7124, 7362, 7608, 7860, 8120, 8388, 8663, 8946, 9237, 9536, 9844, 10160, 10486, 10821, 11165, 11519, 11883, 12258, 12643, 13038, 13445}
var MonsterAccuracyTable = []float64{14, 15, 15, 16, 17, 18, 19, 20, 21, 23, 24, 25, 26, 28, 29, 31, 32, 34, 35, 37, 39, 41, 43, 45, 47, 49, 52, 54, 57, 59, 62, 65, 68, 71, 74, 77, 81, 84, 88, 92, 96, 100, 105, 109, 114, 119, 124, 129, 135, 140, 146, 152, 159, 165, 172, 179, 187, 195, 203, 211, 220, 229, 238, 247, 257, 268, 279, 290, 301, 314, 326, 339, 352, 366, 381, 396, 412, 428, 444, 462, 480, 499, 518, 538, 559, 580, 603, 626, 650, 675, 701, 728, 755, 784, 814, 845, 877, 910, 945, 980}
var MonsterLifeTable = []float64{22, 26, 31, 36, 42, 48, 55, 62, 70, 78, 87, 97, 107, 119, 131, 144, 158, 173, 190, 207, 226, 246, 267, 290, 315, 341, 370, 400, 432, 467, 504, 543, 585, 630, 678, 730, 785, 843, 905, 972, 1042, 1118, 1198, 1284, 1375, 1472, 1575, 1685, 1802, 1927, 2059, 2200, 2350, 2509, 2678, 2858, 3050, 3253, 3469, 3698, 3942, 4201, 4476, 4768, 5078, 5407, 5756, 6127, 6520, 6937, 7380, 7850, 8348, 8876, 9436, 10030, 10660, 11328, 12036, 12787, 13582, 14425, 15319, 16265, 17268, 18331, 19457, 20649, 21913, 23250, 24667, 26168, 27756, 29438, 31220, 33105, 35101, 37214, 39450, 41817}
var MonsterAllyLifeTable = []float64{15, 17, 20, 23, 26, 30, 33, 37, 41, 46, 50, 55, 60, 66, 71, 77, 84, 91, 98, 105, 113, 122, 131, 140, 150, 161, 171, 183, 195, 208, 222, 236, 251, 266, 283, 300, 318, 337, 357, 379, 401, 424, 448, 474, 501, 529, 559, 590, 622, 656, 692, 730, 769, 810, 853, 899, 946, 996, 1048, 1102, 1159, 1219, 1281, 1346, 1415, 1486, 1561, 1640, 1722, 1807, 1897, 1991, 2089, 2192, 2299, 2411, 2528, 2651, 2779, 2913, 3053, 3199, 3352, 3511, 3678, 3853, 4035, 4225, 4424, 4631, 4848, 5074, 5310, 5557, 5815, 6084, 6364, 6658, 6964, 7283}
var MonsterDamageTable = []float64{4.9899997711182, 5.5599999427795, 6.1599998474121, 6.8099999427795, 7.5, 8.2299995422363, 9, 9.8199996948242, 10.699999809265, 11.619999885559, 12.60000038147, 13.640000343323, 14.739999771118, 15.909999847412, 17.139999389648, 18.450000762939, 19.829999923706, 21.290000915527, 22.840000152588, 24.469999313354, 26.190000534058, 28.010000228882, 29.940000534058, 31.959999084473, 34.110000610352, 36.360000610352, 38.75, 41.259998321533, 43.909999847412, 46.700000762939, 49.650001525879, 52.75, 56.009998321533, 59.450000762939, 63.080001831055, 66.889999389648, 70.910003662109, 75.129997253418, 79.580001831055, 84.26000213623, 89.180000305176, 94.349998474121, 99.800003051758, 105.51999664307, 111.5299987793, 117.86000061035, 124.5, 131.49000549316, 138.83000183105, 146.5299987793, 154.63000488281, 163.13999938965, 172.07000732422, 181.44999694824, 191.30000305176, 201.63000488281, 212.47999572754, 223.86999511719, 235.83000183105, 248.36999511719, 261.5299987793, 275.32998657227, 289.82000732422, 305.01000976563, 320.94000244141, 337.64999389648, 355.17999267578, 373.54998779297, 392.80999755859, 413.01000976563, 434.17999267578, 456.36999511719, 479.61999511719, 504, 529.53997802734, 556.29998779297, 584.34997558594, 613.72998046875, 644.5, 676.75, 710.52001953125, 745.89001464844, 782.94000244141, 821.72998046875, 862.35998535156, 904.90002441406, 949.44000244141, 996.07000732422, 1044.8900146484, 1096, 1149.5, 1205.5, 1264.1099853516, 1325.4499511719, 1389.6400146484, 1456.8199462891, 1527.1199951172, 1600.6800537109, 1677.6400146484, 1758.1700439453}
var MonsterArmourTable = []float64{22, 26, 31, 36, 42, 48, 55, 62, 70, 78, 87, 97, 107, 119, 131, 144, 158, 173, 190, 207, 226, 246, 267, 290, 315, 341, 370, 400, 432, 467, 504, 543, 585, 630, 678, 730, 785, 843, 905, 972, 1042, 1118, 1198, 1284, 1375, 1472, 1575, 1685, 1802, 1927, 2059, 2200, 2350, 2509, 2678, 2858, 3050, 3253, 3469, 3698, 3942, 4201, 4476, 4768, 5078, 5407, 5756, 6127, 6520, 6937, 7380, 7850, 8348, 8876, 9436, 10030, 10660, 11328, 12036, 12787, 13582, 14425, 15319, 16265, 17268, 18331, 19457, 20649, 21913, 23250, 24667, 26168, 27756, 29438, 31220, 33105, 35101, 37214, 39450, 41817}

var UnarmedWeaponData = map[int]map[string]interface{}{
	0: {"type": "None", "AttackRate": 1.2, "CritChance": 0, "PhysicalMin": 2, "PhysicalMax": 6}, // Scion
	1: {"type": "None", "AttackRate": 1.2, "CritChance": 0, "PhysicalMin": 2, "PhysicalMax": 8}, // Marauder
	2: {"type": "None", "AttackRate": 1.2, "CritChance": 0, "PhysicalMin": 2, "PhysicalMax": 5}, // Ranger
	3: {"type": "None", "AttackRate": 1.2, "CritChance": 0, "PhysicalMin": 2, "PhysicalMax": 5}, // Witch
	4: {"type": "None", "AttackRate": 1.2, "CritChance": 0, "PhysicalMin": 2, "PhysicalMax": 6}, // Duelist
	5: {"type": "None", "AttackRate": 1.2, "CritChance": 0, "PhysicalMin": 2, "PhysicalMax": 6}, // Templar
	6: {"type": "None", "AttackRate": 1.2, "CritChance": 0, "PhysicalMin": 2, "PhysicalMax": 5}, // Shadow
}