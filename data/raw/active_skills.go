package raw

type ActiveSkill struct {
	AIFile                               string `json:"AIFile"`
	ActiveSkillTargetTypes               []int  `json:"ActiveSkillTargetTypes"`
	ActiveSkillTypes                     []int  `json:"ActiveSkillTypes"`
	AlternateSkillTargetingBehavioursKey *int   `json:"AlternateSkillTargetingBehavioursKey"`
	Description                          string `json:"Description"`
	DisplayedName                        string `json:"DisplayedName"`
	IconDDSFile                          string `json:"Icon_DDSFile"`
	ID                                   string `json:"Id"`
	InputStatKeys                        []int  `json:"Input_StatKeys"`
	IsManuallyCasted                     bool   `json:"IsManuallyCasted"`
	MinionActiveSkillTypes               []int  `json:"MinionActiveSkillTypes"`
	OutputStatKeys                       []int  `json:"Output_StatKeys"`
	SkillTotemID                         int    `json:"SkillTotemId"`
	WeaponRestrictionItemClassesKeys     []int  `json:"WeaponRestriction_ItemClassesKeys"`
	WebsiteDescription                   string `json:"WebsiteDescription"`
	WebsiteImage                         string `json:"WebsiteImage"`
	Key                                  int    `json:"_key"`
}

var ActiveSkills []*ActiveSkill
var ActiveSkillsMap map[int]*ActiveSkill

func InitializeActiveSkills(version string) error {
	if err := InitHelper(version, "ActiveSkills", &ActiveSkills); err != nil {
		return err
	}

	ActiveSkillsMap = make(map[int]*ActiveSkill)
	for _, i := range ActiveSkills {
		ActiveSkillsMap[i.Key] = i
	}

	return nil
}

func (g *ActiveSkill) GetActiveSkillTypes() []*ActiveSkillType {
	if g.ActiveSkillTypes == nil {
		return nil
	}

	out := make([]*ActiveSkillType, len(g.ActiveSkillTypes))
	for i, skillType := range g.ActiveSkillTypes {
		out[i] = ActiveSkillTypesMap[skillType]
	}

	return out
}

func (g *ActiveSkill) GetWeaponRestrictions() []*ItemClass {
	if g.WeaponRestrictionItemClassesKeys == nil {
		return nil
	}

	out := make([]*ItemClass, len(g.WeaponRestrictionItemClassesKeys))
	for i, skillType := range g.WeaponRestrictionItemClassesKeys {
		out[i] = ItemClassesMap[skillType]
	}

	return out
}
