package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	Config StructConfig
	Pickit StructPickit
)

type StructConfig struct {
	Display       int    `yaml:"display"`
	Debug         bool   `yaml:"debug"`
	LogFilePath   string `yaml:"logFilePath"`
	MaxGameLength int    `yaml:"maxGameLength"`
	Health        struct {
		HealingPotionAt     int `yaml:"healingPotionAt"`
		ManaPotionAt        int `yaml:"manaPotionAt"`
		RejuvPotionAtLife   int `yaml:"rejuvPotionAtLife"`
		RejuvPotionAtMana   int `yaml:"rejuvPotionAtMana"`
		MercHealingPotionAt int `yaml:"mercHealingPotionAt"`
		MercRejuvPotionAt   int `yaml:"mercRejuvPotionAt"`
		ChickenAt           int `yaml:"chickenAt"`
		MercChickenAt       int `yaml:"mercChickenAt"`
	} `yaml:"health"`
	Bindings struct {
		Potion1          string `yaml:"potion1"`
		Potion2          string `yaml:"potion2"`
		Potion3          string `yaml:"potion3"`
		Potion4          string `yaml:"potion4"`
		ForceMove        string `yaml:"forceMove"`
		StandStill       string `yaml:"standStill"`
		SwapWeapon       string `yaml:"swapWeapon"`
		Teleport         string `yaml:"teleport"`
		TP               string `yaml:"tp"`
		CTABattleCommand string `yaml:"CTABattleCommand"`
		CTABattleOrders  string `yaml:"CTABattleOrders"`

		// Class Specific bindings
		Sorceress struct {
			Blizzard    string `yaml:"blizzard"`
			StaticField string `yaml:"staticField"`
			FrozenArmor string `yaml:"frozenArmor"`
		} `yaml:"sorceress"`
	} `yaml:"bindings"`
	Inventory struct {
		InventoryLock [][]int `yaml:"inventoryLock"`
		BeltColumns   struct {
			Healing      int `yaml:"healing"`
			Mana         int `yaml:"mana"`
			Rejuvenation int `yaml:"rejuvenation"`
		} `yaml:"beltColumns"`
		BeltRows int `yaml:"beltRows"`
	} `yaml:"inventory"`
	Character struct {
		Class      string `yaml:"class"`
		Difficulty string `yaml:"difficulty"`
		UseMerc    bool   `yaml:"useMerc"`
		UseCTA     bool   `yaml:"useCTA"`
	} `yaml:"character"`
	Runs struct {
		Countess   bool `yaml:"countess"`
		Andariel   bool `yaml:"andariel"`
		Summoner   bool `yaml:"summoner"`
		Mephisto   bool `yaml:"mephisto"`
		Pindleskin bool `yaml:"pindleskin"`
		Nihlathak  bool `yaml:"nihlathak"`
		Council    bool `yaml:"council"`
	} `yaml:"runs"`
}

type StructPickit struct {
	PickupGold          bool `yaml:"pickupGold"`
	MinimumGoldToPickup int  `yaml:"minimumGoldToPickup"`
	Items               []ItemPickit
}

type ItemPickit struct {
	Name    string
	Quality string
}

// Load reads the config.ini file and returns a Config struct filled with data from the ini file
func Load() error {
	r, err := os.Open("config/config.yaml")
	if err != nil {
		return fmt.Errorf("error loading config.yaml: %w", err)
	}

	d := yaml.NewDecoder(r)
	if err = d.Decode(&Config); err != nil {
		return fmt.Errorf("error reading config: %w", err)
	}

	r, err = os.Open("config/pickit.yaml")
	if err != nil {
		return fmt.Errorf("error loading pickit.yaml: %w", err)
	}

	d = yaml.NewDecoder(r)
	if err = d.Decode(&Pickit); err != nil {
		return fmt.Errorf("error reading pickit: %w", err)
	}

	b, err := ioutil.ReadFile("config/pickit.yaml")
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(b, &m)
	if err != nil {
		return fmt.Errorf("error decoding pickit items: %w", err)
	}
	items := parsePickitItems(m["items"].([]interface{}))
	Pickit.Items = items

	return nil
}

func parsePickitItems(items []interface{}) []ItemPickit {
	var itemsToPickit []ItemPickit
	for _, item := range items {
		for name, props := range item.(map[interface{}]interface{}) {
			ip := ItemPickit{
				Name: name.(string),
			}

			if props != nil {
				quality, found := props.(map[interface{}]interface{})["quality"]
				if found {
					ip.Quality = quality.(string)
				}
			}
			itemsToPickit = append(itemsToPickit, ip)
		}
	}

	return itemsToPickit
}
