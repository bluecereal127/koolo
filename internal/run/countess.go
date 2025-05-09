package run

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/npc"
	"github.com/hectorgimenez/d2go/pkg/data/object"
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/context"
)

var ClearRange = 45

type Countess struct {
	ctx *context.Status
}

func NewCountess() *Countess {
	return &Countess{
		ctx: context.Get(),
	}
}

func (c Countess) Name() string {
	return string(config.CountessRun)
}

func (c Countess) Run() error {
	// Travel to boss level
	err := action.WayPoint(area.BlackMarsh)
	if err != nil {
		return err
	}
	action.OpenTPIfLeader()
	areas := []area.ID{
		area.ForgottenTower,
		area.TowerCellarLevel1,
		area.TowerCellarLevel2,
		area.TowerCellarLevel3,
		area.TowerCellarLevel4,
		area.TowerCellarLevel5,
	}

	for _, a := range areas {
		err = action.MoveToArea(a)
		action.OpenTPIfLeader()
		if err != nil {
			return err
		}
	}
	action.OpenTPIfLeader()
	action.Buff()
	// Try to move around Countess area
	action.MoveTo(func() (data.Position, bool) {
		if areaData, ok := context.Get().GameReader.GetData().Areas[area.TowerCellarLevel5]; ok {
			for _, o := range areaData.Objects {
				if o.Name == object.GoodChest {
					return o.Position, true
				} // Countess Chest position from 1.13c fetch
				action.ClearThroughPath(c.ctx.Data.AreaOrigin, ClearRange, c.getMonsterFilter())
			}

		}

		// Try to teleport over Countess in case we are not able to find the chest position, a bit more risky
		if countess, found := c.ctx.Data.Monsters.FindOne(npc.DarkStalker, data.MonsterTypeSuperUnique); found {
			return countess.Position, true
		}
		action.ClearThroughPath(c.ctx.Data.AreaOrigin, ClearRange, c.getMonsterFilter())
		return data.Position{}, false
	})

	// Kill Countess
	return c.ctx.Char.KillCountess()
}

func (c *Countess) getMonsterFilter() data.MonsterFilter {
	return func(monsters data.Monsters) (filteredMonsters []data.Monster) {
		for _, m := range monsters {
			if !c.ctx.Data.AreaData.IsWalkable(m.Position) {
				continue
			}

			// If FocusOnElitePacks is enabled, only return elite monsters and seal bosses
			if c.ctx.CharacterCfg.Game.Countess.ClearGhosts {
				if m.Name == 38 {
					filteredMonsters = append(filteredMonsters, m)
				}
			} else {
				continue
			}
		}

		return filteredMonsters
	}
}
