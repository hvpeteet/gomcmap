package main

import (
	"github.com/kch42/gomcmap/mcmap"
)

type rgb uint32

func (x rgb) RGBA() (r, g, b, a uint32) {
	a = 0xffff
	r = uint32((x >> 16) << 8)
	g = uint32(((x >> 8) & 0xff) << 8)
	b = uint32((x & 0xff) << 8)
	return
}

var colors = map[mcmap.BlockID]rgb{
	mcmap.BlkStone:                      0x666666,
	mcmap.BlkGrassBlock:                 0x00aa00,
	mcmap.BlkDirt:                       0x644804,
	mcmap.BlkCobblestone:                0x7a7a7a,
	mcmap.BlkWoodPlanks:                 0xa4721c,
	mcmap.BlkBedrock:                    0x111111,
	mcmap.BlkWater:                      0x0000ff,
	mcmap.BlkStationaryWater:            0x0000ff,
	mcmap.BlkLava:                       0xff4400,
	mcmap.BlkStationaryLava:             0xff4400,
	mcmap.BlkSand:                       0xf1ee85,
	mcmap.BlkGravel:                     0x9ba3a9,
	mcmap.BlkGoldOre:                    0xffa200,
	mcmap.BlkIronOre:                    0xe1e1e1,
	mcmap.BlkCoalOre:                    0x333333,
	mcmap.BlkWood:                       0xa4721c,
	mcmap.BlkLeaves:                     0x57a100,
	mcmap.BlkGlass:                      0xeeeeff,
	mcmap.BlkLapisLazuliOre:             0x3114e3,
	mcmap.BlkLapisLazuliBlock:           0x3114e3,
	mcmap.BlkDispenser:                  0x7a7a7a,
	mcmap.BlkSandstone:                  0xf1ee85,
	mcmap.BlkNoteBlock:                  0xa4721c,
	mcmap.BlkBed:                        0xa00000,
	mcmap.BlkPoweredRail:                0xff0000,
	mcmap.BlkDetectorRail:               0xff0000,
	mcmap.BlkStickyPiston:               0x91ba12,
	mcmap.BlkCobweb:                     0xdddddd,
	mcmap.BlkGrass:                      0xa0f618,
	mcmap.BlkPiston:                     0xa4721c,
	mcmap.BlkPistonExtension:            0xa4721c,
	mcmap.BlkWool:                       0xffffff,
	mcmap.BlkBlockOfGold:                0xffa200,
	mcmap.BlkBlockOfIron:                0xe1e1e1,
	mcmap.BlkTNT:                        0xa20022,
	mcmap.BlkBookshelf:                  0xa4721c,
	mcmap.BlkMossStone:                  0x589b71,
	mcmap.BlkObsidian:                   0x111144,
	mcmap.BlkTorch:                      0xffcc00,
	mcmap.BlkFire:                       0xffcc00,
	mcmap.BlkMonsterSpawner:             0x344e6a,
	mcmap.BlkOakWoodStairs:              0xa4721c,
	mcmap.BlkChest:                      0xa4721c,
	mcmap.BlkRedstoneWire:               0xff0000,
	mcmap.BlkDiamondOre:                 0x00fff6,
	mcmap.BlkBlockOfDiamond:             0x00fff6,
	mcmap.BlkCraftingTable:              0xa4721c,
	mcmap.BlkWheat:                      0xe7ae00,
	mcmap.BlkFarmland:                   0x644804,
	mcmap.BlkFurnace:                    0x7a7a7a,
	mcmap.BlkBurningFurnace:             0x7a7a7a,
	mcmap.BlkSignPost:                   0xa4721c,
	mcmap.BlkWoodenDoor:                 0xa4721c,
	mcmap.BlkLadders:                    0xa4721c,
	mcmap.BlkRail:                       0xdbdbdb,
	mcmap.BlkCobblestoneStairs:          0x7a7a7a,
	mcmap.BlkWallSign:                   0xa4721c,
	mcmap.BlkLever:                      0xa4721c,
	mcmap.BlkStonePressurePlate:         0x666666,
	mcmap.BlkIronDoor:                   0xe1e1e1,
	mcmap.BlkWoodenPressurePlate:        0xa4721c,
	mcmap.BlkRedstoneOre:                0xa00000,
	mcmap.BlkGlowingRedstoneOre:         0xff0000,
	mcmap.BlkRedstoneTorchInactive:      0xff0000,
	mcmap.BlkRedstoneTorchActive:        0xff0000,
	mcmap.BlkStoneButton:                0x666666,
	mcmap.BlkSnow:                       0xe5fffe,
	mcmap.BlkIce:                        0x9fdcff,
	mcmap.BlkSnowBlock:                  0xe5fffe,
	mcmap.BlkCactus:                     0x01bc3a,
	mcmap.BlkClay:                       0x767a82,
	mcmap.BlkSugarCane:                  0x12db50,
	mcmap.BlkJukebox:                    0xa4721c,
	mcmap.BlkFence:                      0xa4721c,
	mcmap.BlkPumpkin:                    0xff7000,
	mcmap.BlkNetherrack:                 0x851c2d,
	mcmap.BlkSoulSand:                   0x796a59,
	mcmap.BlkGlowstone:                  0xffff00,
	mcmap.BlkNetherPortal:               0xff00ff,
	mcmap.BlkJackOLantern:               0xff7000,
	mcmap.BlkRedstoneRepeaterInactive:   0xff0000,
	mcmap.BlkRedstoneRepeaterActive:     0xff0000,
	mcmap.BlkTrapdoor:                   0xa4721c,
	mcmap.BlkStoneBricks:                0x666666,
	mcmap.BlkHugeBrownMushroom:          0xb07859,
	mcmap.BlkHugeRedMushroom:            0xdd0000,
	mcmap.BlkIronBars:                   0xe1e1e1,
	mcmap.BlkGlassPane:                  0xeeeeff,
	mcmap.BlkMelon:                      0x9ac615,
	mcmap.BlkVines:                      0x50720d,
	mcmap.BlkFenceGate:                  0xa4721c,
	mcmap.BlkBrickStairs:                0xc42500,
	mcmap.BlkStoneBrickStairs:           0x666666,
	mcmap.BlkMycelium:                   0x7c668c,
	mcmap.BlkLilyPad:                    0x50720d,
	mcmap.BlkNetherBrick:                0xc42500,
	mcmap.BlkNetherBrickFence:           0xc42500,
	mcmap.BlkNetherBrickStairs:          0xc42500,
	mcmap.BlkEnchantmentTable:           0x222244,
	mcmap.BlkBrewingStand:               0x666666,
	mcmap.BlkCauldron:                   0x666666,
	mcmap.BlkEndPortal:                  0x000000,
	mcmap.BlkEndPortalBlock:             0xe0dbce,
	mcmap.BlkEndStone:                   0xe0dbce,
	mcmap.BlkRedstoneLampInactive:       0xffff00,
	mcmap.BlkRedstoneLampActive:         0xffff00,
	mcmap.BlkSandstoneStairs:            0xf1ee85,
	mcmap.BlkEmeraldOre:                 0x00c140,
	mcmap.BlkEnderChest:                 0x222244,
	mcmap.BlkBlockOfEmerald:             0x00c140,
	mcmap.BlkSpruceWoodStairs:           0xa4721c,
	mcmap.BlkBirchWoodStairs:            0xa4721c,
	mcmap.BlkJungleWoodStairs:           0xa4721c,
	mcmap.BlkCommandBlock:               0xe8ec78,
	mcmap.BlkBeacon:                     0x00fff6,
	mcmap.BlkCobblestoneWall:            0x7a7a7a,
	mcmap.BlkCarrots:                    0xff6000,
	mcmap.BlkPotatoes:                   0xc6cd0c,
	mcmap.BlkWoodenButton:               0xa4721c,
	mcmap.BlkAnvil:                      0x444444,
	mcmap.BlkTrappedChest:               0xa4721c,
	mcmap.BlkRedstoneComparatorInactive: 0xff0000,
	mcmap.BlkRedstoneComparatorActive:   0xff0000,
	mcmap.BlkBlockOfRedstone:            0xff0000,
	mcmap.BlkNetherQuartzOre:            0xe7e7e7,
	mcmap.BlkHopper:                     0x444444,
	mcmap.BlkBlockOfQuartz:              0xe7e7e7,
	mcmap.BlkQuartzStairs:               0xe7e7e7,
	mcmap.BlkActivatorRail:              0xff0000,
	mcmap.BlkDropper:                    0x444444,
	mcmap.BlkStainedClay:                0x767a82,
	mcmap.BlkHayBlock:                   0xe7ae00,
	mcmap.BlkCarpet:                     0xffffff,
	mcmap.BlkHardenedClay:               0x767a82,
	mcmap.BlkBlockOfCoal:                0x333333,
}
